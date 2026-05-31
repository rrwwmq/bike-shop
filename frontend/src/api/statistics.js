const API = 'http://localhost:5050/api/v1';

export async function getStatistics() {
  const r = await fetch(`${API}/statistics`);
  if (!r.ok) throw new Error('Ошибка загрузки статистики');
  return r.json();
}