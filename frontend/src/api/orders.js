const API = 'http://localhost:5050/api/v1';

export async function getOrders(token) {
  const r = await fetch(`${API}/orders`, {
    headers: { 'Authorization': `Bearer ${token}` }
  });
  if (!r.ok) throw new Error('Ошибка загрузки заказов');
  return r.json();
}

export async function createOrder(data) {
  const r = await fetch(`${API}/orders`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data)
  });
  if (!r.ok) throw new Error('Ошибка создания заказа');
  return r.json();
}

export async function updateOrderStatus(id, status, token) {
  const r = await fetch(`${API}/orders/${id}/status`, {
    method: 'PATCH',
    headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${token}` },
    body: JSON.stringify({ status })
  });
  if (!r.ok) throw new Error('Ошибка обновления статуса');
  return r.json();
}