const API = 'http://localhost:5050/api/v1';

export async function login(email, password) {
  const r = await fetch(`${API}/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  });
  if (!r.ok) throw new Error('Неверный email или пароль');
  return r.json();
}

export async function register(email, password) {
  const r = await fetch(`${API}/auth/register`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  });
  if (!r.ok) throw new Error('Ошибка регистрации');
  return r.json();
}