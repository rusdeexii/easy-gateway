// services/auth.service.ts
import api from '@/utils/axios'
import type { User, CreateUserRequest, UpdateUserRequest } from '@/models/user.model';

export const login = async (username: string, password: string) => {
  const response = await api.post('/auth/login', { username, password });
  return response.data; // { user, expires_at }
};

export const logout = async () => {
  await api.post('/auth/logout');
};

export const fetchUsers = async (): Promise<User[]> => {
  const res = await api.get('/auth/users')
  return res.data
}

export const createUser = async (user: CreateUserRequest): Promise<User> => {
  const res = await api.post('/auth/users', user)
  return res.data
}

export const updateUser = async (userId: string, user: UpdateUserRequest): Promise<User> => {
  const res = await api.put(`/auth/users/${userId}`, user)
  return res.data
}


export const deleteUser = async (userId: string): Promise<void> => {
  await api.delete(`/auth/users/${userId}`)
}
