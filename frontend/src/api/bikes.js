const API = 'http://localhost:5050/api/v1';

export async function getBikes() {
  const r = await fetch(`${API}/bikes`);
  if (!r.ok) throw new Error('Ошибка загрузки каталога');
  return r.json();
}

export async function createBike(data, token) {
  const r = await fetch(`${API}/bikes`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${token}` },
    body: JSON.stringify(data)
  });
  if (!r.ok) throw new Error('Ошибка создания велосипеда');
  return r.json();
}

export async function updateBike(id, data, token) {
  const r = await fetch(`${API}/bikes/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${token}` },
    body: JSON.stringify(data)
  });
  if (!r.ok) throw new Error('Ошибка обновления велосипеда');
  return r.json();
}

export async function deleteBike(id, token) {
  const r = await fetch(`${API}/bikes/${id}`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${token}` }
  });
  if (!r.ok) throw new Error('Ошибка удаления велосипеда');
}