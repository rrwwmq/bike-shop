import { useState, useCallback } from 'react';

export function useToast() {
  const [toast, setToast] = useState(null);

  const showToast = useCallback((msg, type = '') => {
    setToast({ msg, type });
    setTimeout(() => setToast(null), 3000);
  }, []);

  return { toast, showToast };
}

export function Toast({ toast }) {
  if (!toast) return null;
  const colors = {
    success: '#27ae60',
    error: '#c0392b',
    '': '#1a1917'
  };
  return (
    <div style={{
      position: 'fixed', bottom: '2rem', right: '2rem',
      background: colors[toast.type] || colors[''],
      color: '#fff', padding: '12px 20px',
      borderRadius: '8px', fontSize: '14px', zIndex: 300
    }}>
      {toast.msg}
    </div>
  );
}