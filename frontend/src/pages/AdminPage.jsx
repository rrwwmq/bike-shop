import { useState } from 'react';
import { getBikes, createBike, updateBike, deleteBike } from '../api/bikes';
import { getOrders, updateOrderStatus } from '../api/orders';
import { login } from '../api/auth';

export default function AdminPage({ showToast }) {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [token, setToken] = useState('');
  const [authed, setAuthed] = useState(false);
  const [tab, setTab] = useState('orders');
  const [orders, setOrders] = useState([]);
  const [bikes, setBikes] = useState([]);
  const [modal, setModal] = useState(false);
  const [editingBike, setEditingBike] = useState(null);
  const [form, setForm] = useState({ brand: '', model: '', type: '', price: '', stock: '', description: '' });

  const formatPrice = p => new Intl.NumberFormat('ru-RU').format(Math.round(p)) + ' ₽';
  const formatDate = s => new Date(s).toLocaleDateString('ru-RU');

  async function handleLogin() {
    if (!email || !password) { showToast('Введите email и пароль', 'error'); return; }
    try {
      const data = await login(email, password);
      setToken(data.token);
      const orders = await getOrders(data.token);
      setOrders(Array.isArray(orders) ? orders : []);
      setAuthed(true);
      loadBikes();
    } catch(e) {
      showToast(e.message || 'Неверные данные', 'error');
    }
  }

  async function loadBikes() {
    const data = await getBikes();
    setBikes(Array.isArray(data) ? data : []);
  }

  async function loadOrders() {
    const data = await getOrders(token);
    setOrders(Array.isArray(data) ? data : []);
  }

  async function handleStatusChange(id, status) {
    try {
      await updateOrderStatus(id, status, token);
      showToast('Статус обновлён', 'success');
      loadOrders();
    } catch { showToast('Ошибка', 'error'); }
  }

  async function handleDeleteBike(id) {
    if (!confirm('Удалить велосипед?')) return;
    try {
      await deleteBike(id, token);
      showToast('Удалено', 'success');
      loadBikes();
    } catch { showToast('Ошибка удаления', 'error'); }
  }

  async function handleSaveBike() {
    const data = { brand: form.brand, model: form.model, type: form.type, price: parseFloat(form.price), stock: parseInt(form.stock), description: form.description };
    if (!data.brand || !data.model || !data.type || !data.price) { showToast('Заполните обязательные поля', 'error'); return; }
    try {
      if (editingBike) await updateBike(editingBike.id, data, token);
      else await createBike(data, token);
      showToast(editingBike ? 'Обновлено' : 'Добавлено', 'success');
      setModal(false);
      loadBikes();
    } catch { showToast('Ошибка сохранения', 'error'); }
  }

  function openModal(bike = null) {
    setEditingBike(bike);
    setForm(bike ? { brand: bike.brand, model: bike.model, type: bike.type, price: bike.price, stock: bike.stock, description: bike.description || '' } : { brand: '', model: '', type: '', price: '', stock: '', description: '' });
    setModal(true);
  }

  const input = (key, label, placeholder, type = 'text') => (
    <div style={{ marginBottom: '14px' }}>
      <label style={{ display: 'block', fontSize: '12px', fontWeight: 500, color: 'var(--muted)', textTransform: 'uppercase', letterSpacing: '0.5px', marginBottom: '6px' }}>{label}</label>
      <input type={type} value={form[key]} placeholder={placeholder}
        onChange={e => setForm(p => ({ ...p, [key]: e.target.value }))}
        style={{ width: '100%', padding: '9px 12px', border: '1px solid var(--border)', borderRadius: 'var(--radius-sm)', fontFamily: 'Manrope', fontSize: '14px', background: 'var(--bg)', outline: 'none' }}
      />
    </div>
  );

  return (
    <div style={{ padding: '2rem', maxWidth: '1100px', margin: '0 auto' }}>
      <h1 style={{ fontFamily: 'Syne', fontSize: '26px', fontWeight: 600, marginBottom: '1.5rem' }}>Панель управления</h1>

      {!authed ? (
        <div style={{ background: 'var(--surface)', border: '1px solid var(--border)', borderRadius: 'var(--radius)', padding: '1.5rem', maxWidth: '400px' }}>
          <h3 style={{ fontFamily: 'Syne', fontSize: '16px', marginBottom: '1rem' }}>Вход для администратора</h3>
          <div style={{ marginBottom: '14px' }}>
            <label style={{ display: 'block', fontSize: '12px', fontWeight: 500, color: 'var(--muted)', textTransform: 'uppercase', letterSpacing: '0.5px', marginBottom: '6px' }}>Email</label>
            <input type="email" value={email} onChange={e => setEmail(e.target.value)}
              onKeyDown={e => e.key === 'Enter' && handleLogin()}
              placeholder="admin@bikeshop.com"
              style={{ width: '100%', padding: '9px 12px', border: '1px solid var(--border)', borderRadius: 'var(--radius-sm)', fontFamily: 'Manrope', fontSize: '14px', background: 'var(--bg)', outline: 'none' }}
            />
          </div>
          <div style={{ marginBottom: '14px' }}>
            <label style={{ display: 'block', fontSize: '12px', fontWeight: 500, color: 'var(--muted)', textTransform: 'uppercase', letterSpacing: '0.5px', marginBottom: '6px' }}>Пароль</label>
            <input type="password" value={password} onChange={e => setPassword(e.target.value)}
              onKeyDown={e => e.key === 'Enter' && handleLogin()}
              placeholder="••••••"
              style={{ width: '100%', padding: '9px 12px', border: '1px solid var(--border)', borderRadius: 'var(--radius-sm)', fontFamily: 'Manrope', fontSize: '14px', background: 'var(--bg)', outline: 'none' }}
            />
          </div>
          <button onClick={handleLogin} style={{ background: 'var(--text)', color: '#fff', border: 'none', padding: '10px 20px', borderRadius: 'var(--radius-sm)', fontFamily: 'Manrope', fontSize: '14px', fontWeight: 500, cursor: 'pointer' }}>Войти</button>
        </div>
      ) : (
        <>
          <div style={{ display: 'flex', gap: '4px', marginBottom: '1.5rem', borderBottom: '1px solid var(--border)' }}>
            {['orders', 'bikes'].map(t => (
              <button key={t} onClick={() => setTab(t)} style={{ background: 'none', border: 'none', padding: '9px 18px', fontFamily: 'Manrope', fontSize: '14px', color: tab === t ? 'var(--text)' : 'var(--muted)', borderBottom: tab === t ? '2px solid var(--text)' : '2px solid transparent', marginBottom: '-1px', cursor: 'pointer', fontWeight: tab === t ? 500 : 400 }}>
                {t === 'orders' ? 'Заказы' : 'Каталог'}
              </button>
            ))}
          </div>

          {tab === 'orders' && (
            <div style={{ background: 'var(--surface)', border: '1px solid var(--border)', borderRadius: 'var(--radius)', overflow: 'hidden' }}>
              <table style={{ width: '100%', borderCollapse: 'collapse' }}>
                <thead>
                  <tr style={{ background: 'var(--bg)' }}>
                    {['#', 'Клиент', 'Email', 'Сумма', 'Дата', 'Статус'].map(h => (
                      <th key={h} style={{ padding: '10px 16px', textAlign: 'left', fontSize: '11px', fontWeight: 500, textTransform: 'uppercase', letterSpacing: '0.5px', color: 'var(--muted)', borderBottom: '1px solid var(--border)' }}>{h}</th>
                    ))}
                  </tr>
                </thead>
                <tbody>
                  {orders.length === 0
                    ? <tr><td colSpan={6} style={{ padding: '2rem', textAlign: 'center', color: 'var(--muted)' }}>Заказов пока нет</td></tr>
                    : orders.map(o => (
                      <tr key={o.id} style={{ borderBottom: '1px solid var(--border)' }}>
                        <td style={{ padding: '12px 16px', fontWeight: 500 }}>#{o.id}</td>
                        <td style={{ padding: '12px 16px' }}>{o.full_name}</td>
                        <td style={{ padding: '12px 16px', color: 'var(--muted)' }}>{o.email}</td>
                        <td style={{ padding: '12px 16px', fontWeight: 500 }}>{formatPrice(o.total_price)}</td>
                        <td style={{ padding: '12px 16px', color: 'var(--muted)' }}>{formatDate(o.created_at)}</td>
                        <td style={{ padding: '12px 16px' }}>
                          <select value={o.status} onChange={e => handleStatusChange(o.id, e.target.value)}
                            style={{ padding: '5px 8px', border: '1px solid var(--border)', borderRadius: 'var(--radius-sm)', fontFamily: 'Manrope', fontSize: '13px', background: 'var(--bg)', cursor: 'pointer' }}>
                            <option value="pending">Ожидает</option>
                            <option value="completed">Завершён</option>
                            <option value="cancelled">Отменён</option>
                          </select>
                        </td>
                      </tr>
                    ))
                  }
                </tbody>
              </table>
            </div>
          )}

          {tab === 'bikes' && (
            <>
              <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '1rem' }}>
                <span style={{ color: 'var(--muted)', fontSize: '14px' }}>{bikes.length} велосипедов</span>
                <button onClick={() => openModal()} style={{ background: 'var(--text)', color: '#fff', border: 'none', padding: '8px 16px', borderRadius: 'var(--radius-sm)', fontFamily: 'Manrope', fontSize: '13px', fontWeight: 500, cursor: 'pointer' }}>+ Добавить</button>
              </div>
              <div style={{ background: 'var(--surface)', border: '1px solid var(--border)', borderRadius: 'var(--radius)', overflow: 'hidden' }}>
                <table style={{ width: '100%', borderCollapse: 'collapse' }}>
                  <thead>
                    <tr style={{ background: 'var(--bg)' }}>
                      {['Модель', 'Бренд', 'Тип', 'Цена', 'Склад', 'Действия'].map(h => (
                        <th key={h} style={{ padding: '10px 16px', textAlign: 'left', fontSize: '11px', fontWeight: 500, textTransform: 'uppercase', letterSpacing: '0.5px', color: 'var(--muted)', borderBottom: '1px solid var(--border)' }}>{h}</th>
                      ))}
                    </tr>
                  </thead>
                  <tbody>
                    {bikes.map(b => (
                      <tr key={b.id} style={{ borderBottom: '1px solid var(--border)' }}>
                        <td style={{ padding: '12px 16px', fontWeight: 500 }}>{b.model}</td>
                        <td style={{ padding: '12px 16px' }}>{b.brand}</td>
                        <td style={{ padding: '12px 16px' }}>
                          <span style={{ background: 'var(--accent-light)', color: 'var(--muted)', fontSize: '11px', padding: '3px 9px', borderRadius: '999px', textTransform: 'uppercase' }}>{b.type}</span>
                        </td>
                        <td style={{ padding: '12px 16px' }}>{formatPrice(b.price)}</td>
                        <td style={{ padding: '12px 16px' }}>{b.stock} шт.</td>
                        <td style={{ padding: '12px 16px' }}>
                          <div style={{ display: 'flex', gap: '6px' }}>
                            <button onClick={() => openModal(b)} style={{ background: 'none', border: '1px solid var(--border)', padding: '6px 12px', borderRadius: 'var(--radius-sm)', fontFamily: 'Manrope', fontSize: '13px', cursor: 'pointer' }}>Изменить</button>
                            <button onClick={() => handleDeleteBike(b.id)} style={{ background: 'none', border: '1px solid #fbb', padding: '6px 12px', borderRadius: 'var(--radius-sm)', fontFamily: 'Manrope', fontSize: '13px', cursor: 'pointer', color: 'var(--danger)' }}>Удалить</button>
                          </div>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            </>
          )}
        </>
      )}

      {modal && (
        <div onClick={e => e.target === e.currentTarget && setModal(false)} style={{ position: 'fixed', inset: 0, background: 'rgba(0,0,0,0.35)', zIndex: 200, display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
          <div style={{ background: 'var(--surface)', borderRadius: 'var(--radius)', padding: '1.75rem', width: '440px', maxWidth: 'calc(100vw - 2rem)' }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '1.25rem' }}>
              <h3 style={{ fontFamily: 'Syne', fontSize: '18px', fontWeight: 600 }}>{editingBike ? 'Редактировать' : 'Добавить велосипед'}</h3>
              <button onClick={() => setModal(false)} style={{ background: 'none', border: 'none', fontSize: '20px', cursor: 'pointer', color: 'var(--muted)' }}>×</button>
            </div>
            {input('brand', 'Бренд', 'Trek')}
            {input('model', 'Модель', 'Marlin 5')}
            {input('type', 'Тип', 'mountain')}
            <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '12px' }}>
              {input('price', 'Цена (₽)', '45000', 'number')}
              {input('stock', 'Количество', '5', 'number')}
            </div>
            {input('description', 'Описание', 'Краткое описание')}
            <div style={{ display: 'flex', gap: '8px', justifyContent: 'flex-end', marginTop: '1.25rem' }}>
              <button onClick={() => setModal(false)} style={{ background: 'none', border: '1px solid var(--border)', padding: '8px 16px', borderRadius: 'var(--radius-sm)', fontFamily: 'Manrope', fontSize: '13px', cursor: 'pointer' }}>Отмена</button>
              <button onClick={handleSaveBike} style={{ background: 'var(--text)', color: '#fff', border: 'none', padding: '8px 16px', borderRadius: 'var(--radius-sm)', fontFamily: 'Manrope', fontSize: '13px', fontWeight: 500, cursor: 'pointer' }}>Сохранить</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}