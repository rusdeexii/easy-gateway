import axios from 'axios';
import { useAuthStore } from '@/stores/auth.store';
import router from '@/router';

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
  withCredentials: true, // ส่งคุกกี้ไปกับทุกคำขอ
});

// Interceptor สำหรับคำขอ
api.interceptors.request.use(
  (config) => config,
  (error) => Promise.reject(error)
);

// Interceptor สำหรับการตอบสนอง
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const authStore = useAuthStore();
    const originalRequest = error.config;

    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;
      try {
        await authStore.refreshToken();
        return api(originalRequest);
      } catch (refreshError) {
        authStore.clearAuth();
        router.push({ name: 'Login' });
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
);

export default api;
