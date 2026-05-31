import { useEffect, useState } from 'react';
import { getBikes } from '../api/bikes';

export default function CatalogPage({ cart, addToCart, changeQty }) {
  const [bikes, setBikes] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    getBikes().then(data => { setBikes(Array.isArray(data) ? data : []); setLoading(false); })
      .catch(() => setLoading(false));
  }, []);

  const formatPrice = p => new Intl.NumberFormat('ru-RU').format(Math.round(p)) + ' ₽';

  if (loading) return <div style={{ padding: '4rem', textAlign: 'center', color: 'var(--muted)' }}>Загрузка...</div>;

  return (
    <div style={{ padding: '2rem', maxWidth: '1100px', margin: '0 auto' }}>
      <div style={{ marginBottom: '1.5rem' }}>
        <h1 style={{ fontFamily: 'Syne', fontSize: '26px', fontWeight: 600 }}>Каталог</h1>
        <p style={{ color: 'var(--muted)', fontSize: '14px' }}>{bikes.length} велосипедов</p>
      </div>
      <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(240px, 1fr))', gap: '16px' }}>
        {bikes.map(b => {
          const inCart = cart[b.id]?.qty || 0;
          return (
            <div key={b.id} style={{ background: 'var(--surface)', border: '1px solid var(--border)', borderRadius: 'var(--radius)', padding: '1.25rem', display: 'flex', flexDirection: 'column', gap: '8px' }}>
              <span style={{ background: 'var(--accent-light)', color: 'var(--muted)', fontSize: '11px', fontWeight: 500, padding: '3px 9px', borderRadius: '999px', textTransform: 'uppercase', letterSpacing: '0.5px', width: 'fit-content' }}>{b.type}</span>
              <div style={{ fontFamily: 'Syne', fontSize: '16px', fontWeight: 600 }}>{b.model}</div>
              <div style={{ color: 'var(--muted)', fontSize: '13px' }}>{b.brand}</div>
              {b.description && <div style={{ color: 'var(--muted)', fontSize: '13px', flex: 1 }}>{b.description}</div>}
              <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', paddingTop: '12px', borderTop: '1px solid var(--border)' }}>
                <div>
                  <div style={{ fontFamily: 'Syne', fontSize: '18px', fontWeight: 600 }}>{formatPrice(b.price)}</div>
                  <div style={{ fontSize: '12px', color: b.stock === 0 ? 'var(--danger)' : b.stock <= 2 ? 'var(--warning)' : 'var(--muted)' }}>
                    {b.stock === 0 ? 'Нет в наличии' : `В наличии: ${b.stock} шт.`}
                  </div>
                </div>
                {b.stock === 0
                  ? <button disabled style={{ opacity: 0.3, background: 'var(--text)', color: '#fff', border: 'none', padding: '7px 14px', borderRadius: 'var(--radius-sm)', fontSize: '13px' }}>Нет</button>
                  : inCart === 0
                    ? <button onClick={() => addToCart(b)} style={{ background: 'var(--text)', color: '#fff', border: 'none', padding: '7px 14px', borderRadius: 'var(--radius-sm)', fontSize: '13px', fontWeight: 500, cursor: 'pointer' }}>В корзину</button>
                    : <div style={{ display: 'flex', alignItems: 'center', gap: '8px' }}>
                        <button onClick={() => changeQty(b.id, -1, b.stock)} style={{ width: '28px', height: '28px', border: '1px solid var(--border)', background: 'var(--surface)', borderRadius: '6px', cursor: 'pointer', fontSize: '16px' }}>−</button>
                        <span style={{ fontSize: '14px', fontWeight: 500, minWidth: '20px', textAlign: 'center' }}>{inCart}</span>
                        <button onClick={() => changeQty(b.id, 1, b.stock)} style={{ width: '28px', height: '28px', border: '1px solid var(--border)', background: 'var(--surface)', borderRadius: '6px', cursor: 'pointer', fontSize: '16px' }}>+</button>
                      </div>
                }
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
}