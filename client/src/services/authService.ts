// src/services/authService.ts

import api from './api';

export const register = async (username: string, email: string, password: string) => {
  const response = await api.post('/users/register', { username, email, password });
  return response.data;
};

export const login = async (username: string, password: string) => {
  const response = await api.post('/users/login', { username, password });
  const token = response.data.token;
  localStorage.setItem('token', token);
  return token;
};

export const logout = () => {
  localStorage.removeItem('token');
};