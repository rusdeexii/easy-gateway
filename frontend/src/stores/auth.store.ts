// stores/auth.store.ts
import { defineStore } from 'pinia'
import api from '@/utils/axios'
import type { User, CreateUserRequest, UpdateUserRequest  } from '@/models/user.model';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as User | null,
    isCheckingAuth: true,
    users: [] as User[],
  }),

  actions: {
    async init() {
      try {
        const response = await api.get('/auth/me')
        this.user = response.data.user
      } catch {
        this.user = null
      } finally {
        this.isCheckingAuth = false
      }
    },

    async login(username: string, password: string) {
      const res = await api.post('/auth/login', { username, password })
      this.user = res.data.user
    },

    async refreshToken() {
      await api.post('/auth/refresh')
      await this.init()
    },

    async logout() {
      await api.post('/auth/logout')
      this.user = null
    },

    async fetchUsers() {
      const res = await api.get<User[]>('/auth/users')
      this.users = res.data
      return res.data
    },

    async createUser(data: CreateUserRequest) {
      const res = await api.post<User>('/auth/users', data)
      return res.data
    },

    async updateUser(id: string, data: UpdateUserRequest) {
      const res = await api.put<User>(`/auth/users/${id}`, data)
      return res.data
    },


    async deleteUser(userId: string) {
      await api.delete(`/auth/users/${userId}`)
    },

    isAuthenticated(): boolean {
      return !!this.user
    },

    hasRole(role: string): boolean {
      return this.user?.role === role
    },
  },
})
