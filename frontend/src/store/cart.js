import { useState, useCallback } from 'react';

export function useCart() {
  const [cart, setCart] = useState({});

  const addToCart = useCallback((bike) => {
    setCart(prev => ({ ...prev, [bike.id]: { bike, qty: (prev[bike.id]?.qty || 0) + 1 } }));
}, []);

const changeQty = useCallback((id, delta, maxStock) => {
    setCart(prev => {
      const cur = prev[id]?.qty || 0;
      const next = cur + delta;
      if (next <= 0) { const { [id]: _, ...rest } = prev; return rest; }
      if (next > maxStock) return prev;
      return { ...prev, [id]: { ...prev[id], qty: next } };
    });
}, []);

  const removeFromCart = useCallback((id) => {
    setCart(prev => { const { [id]: _, ...rest } = prev; return rest; });
  }, []);

  const clearCart = useCallback(() => setCart({}), []);

  const totalCount = Object.values(cart).reduce((s, i) => s + i.qty, 0);
  const totalPrice = Object.values(cart).reduce((s, i) => s + i.bike.price * i.qty, 0);
  const items = Object.values(cart);

  return { cart, items, totalCount, totalPrice, addToCart, changeQty, removeFromCart, clearCart };
}