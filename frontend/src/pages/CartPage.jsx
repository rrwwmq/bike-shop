import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { createOrder } from '../api/orders';

export default function CartPage({ items, totalPrice, removeFromCart, clearCart, showToast }) {
  const navigate = useNavigate();
  const [form, setForm] = useState({ full_name: '', email: '', address: '' });
  const [loading, setLoading] = useState(false);

  const formatPrice = p => new Intl.NumberFormat('ru-RU').format(Math.round(p)) + ' ₽';

  if (!items.length) return (
    <div style={{ padding: '4rem', textAlign: 'center' }}>
      <h2 style={{ fontFamily: 'Syne', fontSize: '18px', marginBottom: '8px' }}>Корзина пуста</h2>
      <p style={{ color: 'var(--muted)', marginBottom: '1rem' }}>Добавьте велосипеды из каталога</p>
      <button onClick={() => navigate('/')} style={{ background: 'var(--text)', color: '#fff', border: 'none', padding: '10px 20px', borderRadius: 'var(--radius-sm)', cursor: 'pointer', fontFamily: 'Manrope' }}>Перейти в каталог</button>
    </div>
  );

  async function handleSubmit() {
    if (!form.full_name || !form.email || !form.address) { showToast('Заполните все поля', 'error'); return; }
    if (!form.email.includes('@')) { showToast('Некорректный email', 'error'); return; }
    setLoading(true);
    try {
      const order = await createOrder({
        ...form,
        items: items.map(i => ({ bike_id: i.bike.id, quantity: i.qty }))
      });
      clearCart();
      navigate('/success/' + order.id);
    } catch(e) {
      showToast(e.message, 'error');
    } finally {
      setLoading(false);
    }
  }

  return (
    <div style={{ padding: '2rem', maxWidth: '1100px', margin: '0 auto' }}>
      <h1 style={{ fontFamily: 'Syne', fontSize: '26px', fontWeight: 600, marginBottom: '1.5rem' }}>Корзина</h1>
      <div style={{ display: 'grid', gridTemplateColumns: '1fr 340px', gap: '24px', alignItems: 'start' }}>
        <div style={{ display: 'flex', flexDirection: 'column', gap: '12px' }}>
          {items.map(({ bike, qty }) => (
            <div key={bike.id} style={{ background: 'var(--surface)', border: '1px solid var(--border)', borderRadius: 'var(--radius)', padding: '1rem 1.25rem', display: 'flex', alignItems: 'center', gap: '16px' }}>
              <div style={{ flex: 1 }}>
                <div style={{ fontWeight: 500 }}>{bike.brand} {bike.model}</div>
                <div style={{ color: 'var(--muted)', fontSize: '13px' }}>{bike.type} · {qty} шт.</div>
              </div>
              <div style={{ fontFamily: 'Syne', fontWeight: 600, fontSize: '16px' }}>{formatPrice(bike.price * qty)}</div>
              <button onClick={() => removeFromCart(bike.id)} style={{ background: 'none', border: 'none', color: 'var(--muted)', cursor: 'pointer', fontSize: '20px', lineHeight: 1 }}>×</button>
            </div>
          ))}
        </div>
        <div style={{ background: 'var(--surface)', border: '1px solid var(--border)', borderRadius: 'var(--radius)', padding: '1.5rem', position: 'sticky', top: '80px' }}>
          <h3 style={{ fontFamily: 'Syne', fontSize: '16px', fontWeight: 600, marginBottom: '1.25rem' }}>Оформление заказа</h3>
          {[['full_name', 'Имя', 'Иван Иванов'], ['email', 'Email', 'ivan@example.com'], ['address', 'Адрес доставки', 'Москва, ул. Ленина 1']].map(([key, label, placeholder]) => (
            <div key={key} style={{ marginBottom: '14px' }}>
              <label style={{ display: 'block', fontSize: '12px', fontWeight: 500, color: 'var(--muted)', textTransform: 'uppercase', letterSpacing: '0.5px', marginBottom: '6px' }}>{label}</label>
              <input
                value={form[key]} placeholder={placeholder}
                onChange={e => setForm(p => ({ ...p, [key]: e.target.value }))}
                style={{ width: '100%', padding: '9px 12px', border: '1px solid var(--border)', borderRadius: 'var(--radius-sm)', fontFamily: 'Manrope', fontSize: '14px', background: 'var(--bg)', outline: 'none' }}
              />
            </div>
          ))}
          <div style={{ borderTop: '1px solid var(--border)', paddingTop: '14px', marginTop: '14px', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '14px' }}>
            <span style={{ color: 'var(--muted)', fontSize: '13px' }}>Итого</span>
            <span style={{ fontFamily: 'Syne', fontSize: '20px', fontWeight: 600 }}>{formatPrice(totalPrice)}</span>
          </div>
          <button onClick={handleSubmit} disabled={loading} style={{ width: '100%', background: 'var(--text)', color: '#fff', border: 'none', padding: '11px', borderRadius: 'var(--radius-sm)', fontFamily: 'Manrope', fontSize: '15px', fontWeight: 500, cursor: loading ? 'not-allowed' : 'pointer', opacity: loading ? 0.6 : 1 }}>
            {loading ? 'Оформляем...' : 'Оформить заказ'}
          </button>
        </div>
      </div>
    </div>
  );
}