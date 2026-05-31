import { BrowserRouter, Routes, Route, Navigate, useNavigate, useParams } from 'react-router-dom';
import { useCart } from './store/cart';
import { useToast, Toast } from './components/Toast';
import Navbar from './components/Navbar';
import CatalogPage from './pages/CatalogPage';
import CartPage from './pages/CartPage';
import AdminPage from './pages/AdminPage';

function SuccessPage() {
  const { id } = useParams();
  const navigate = useNavigate();

  return (
    <div style={{ padding: '2rem', display: 'flex', justifyContent: 'center' }}>
      <div style={{ background: 'var(--surface)', border: '1px solid var(--border)', borderRadius: 'var(--radius)', padding: '3rem', textAlign: 'center', maxWidth: '480px', width: '100%' }}>
        <div style={{ width: '56px', height: '56px', background: '#e8f5e9', borderRadius: '50%', display: 'flex', alignItems: 'center', justifyContent: 'center', margin: '0 auto 1.5rem', fontSize: '24px' }}>✓</div>
        <h2 style={{ fontFamily: 'Syne', fontSize: '22px', fontWeight: 600, marginBottom: '8px' }}>Заказ оформлен!</h2>
        <p style={{ color: 'var(--muted)', marginBottom: '1rem' }}>Ваш номер заказа</p>
        <div style={{ fontFamily: 'Syne', fontSize: '28px', fontWeight: 600, marginBottom: '1.5rem' }}>#{id}</div>
        <p style={{ color: 'var(--muted)', marginBottom: '1.5rem' }}>Мы свяжемся с вами по указанной почте</p>
        <button onClick={() => navigate('/')} style={{ background: 'var(--text)', color: '#fff', border: 'none', padding: '11px 24px', borderRadius: 'var(--radius-sm)', fontFamily: 'Manrope', fontSize: '15px', fontWeight: 500, cursor: 'pointer' }}>Вернуться в каталог</button>
      </div>
    </div>
  );
}

export default function App() {
  const { cart, items, totalCount, totalPrice, addToCart, changeQty, removeFromCart, clearCart } = useCart();
  const { toast, showToast } = useToast();

  return (
    <BrowserRouter>
      <Navbar totalCount={totalCount} />
      <Routes>
        <Route path="/" element={<CatalogPage cart={cart} addToCart={addToCart} changeQty={changeQty} />} />
        <Route path="/cart" element={<CartPage items={items} totalPrice={totalPrice} removeFromCart={removeFromCart} clearCart={clearCart} showToast={showToast} />} />
        <Route path="/admin" element={<AdminPage showToast={showToast} />} />
        <Route path="/success/:id" element={<SuccessPage />} />
        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
      <Toast toast={toast} />
    </BrowserRouter>
  );
}