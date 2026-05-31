import { useNavigate, useLocation } from 'react-router-dom';

export default function Navbar({ totalCount }) {
  const navigate = useNavigate();
  const location = useLocation();

  const navBtn = (label, path) => (
    <button
      onClick={() => navigate(path)}
      style={{
        background: location.pathname === path ? 'var(--accent-light)' : 'none',
        border: 'none', padding: '7px 14px',
        borderRadius: 'var(--radius-sm)',
        fontFamily: 'Manrope, sans-serif', fontSize: '14px',
        color: location.pathname === path ? 'var(--text)' : 'var(--muted)',
        fontWeight: location.pathname === path ? 500 : 400,
        cursor: 'pointer'
      }}
    >
      {label}
    </button>
  );

  return (
    <nav style={{
      background: 'var(--surface)', borderBottom: '1px solid var(--border)',
      padding: '0 2rem', display: 'flex', alignItems: 'center',
      justifyContent: 'space-between', height: '60px',
      position: 'sticky', top: 0, zIndex: 100
    }}>
      <div
        onClick={() => navigate('/')}
        style={{ fontFamily: 'Syne, sans-serif', fontSize: '18px', fontWeight: 600, cursor: 'pointer' }}
      >
        Вело<span style={{ color: 'var(--muted)', fontWeight: 400 }}>шоп</span>
      </div>
      <div style={{ display: 'flex', gap: '4px', alignItems: 'center' }}>
        {navBtn('Каталог', '/')}
        {navBtn('Админ', '/admin')}
        <button
          onClick={() => navigate('/cart')}
          style={{
            background: 'var(--text)', color: '#fff', border: 'none',
            padding: '7px 16px', borderRadius: 'var(--radius-sm)',
            fontFamily: 'Manrope, sans-serif', fontSize: '14px',
            fontWeight: 500, cursor: 'pointer', display: 'flex',
            alignItems: 'center', gap: '6px'
          }}
        >
          Корзина
          <span style={{
            background: '#fff', color: 'var(--text)',
            borderRadius: '999px', padding: '1px 7px',
            fontSize: '12px', fontWeight: 600
          }}>
            {totalCount}
          </span>
        </button>
      </div>
    </nav>
  );
}