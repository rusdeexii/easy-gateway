// router/index.ts
import { createRouter, createWebHistory } from 'vue-router';
import { useAuthStore } from '../stores/auth.store';

const routes = [
  {
    path: '/',
    redirect: '/system/dashboard'
  },
  {
    path: '/auth/signin',
    name: 'Login',
    component: () => import('@/pages/Login.vue'),
    meta: { layout: 'auth', guest: true }
  },
  {
    path: '/system/dashboard',
    name: 'Dashboard',
    component: () => import('@/pages/admin/Dashboard.vue'),
    meta: { layout: 'default', requiresAuth: true }
  },
  {
    path: '/system/users',
    component: () => import('@/pages/admin/Users.vue'),
    meta: { layout: 'default', requiresAuth: true }
  },
  {
    path: '/system/products',
    component: () => import('@/pages/admin/Products.vue'),
    meta: { layout: 'default', requiresAuth: true }
  },
  {
    path: '/setting/profile',
    component: () => import('@/pages/admin/Profile-Setting.vue'),
    meta: { layout: 'default', requiresAuth: true }
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/pages/NotFound.vue'),
    meta: { layout: 'default' }
  }
];

export const router = createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore();

  // รอให้เช็ค auth เสร็จก่อน (ตอน reload ใหม่)
  if (authStore.isCheckingAuth) {
    await new Promise(resolve => {
      const unwatch = authStore.$subscribe(() => {
        if (!authStore.isCheckingAuth) {
          unwatch();
          resolve(true);
        }
      });
    });
  }

  const isAuthenticated = authStore.isAuthenticated();

  if (to.meta.requiresAuth && !isAuthenticated) {
    // ถ้า route ต้อง login แต่ไม่มี user → ไปหน้า Login
    return next({ name: 'Login' });
  }

  // ถ้า login แล้ว แต่พยายามเข้า /login อีก → redirect ไป /system/dashboard
  if (to.name === 'Login' && isAuthenticated) {
    return next({ name: 'Dashboard' });
  }

  next();
});

export default router;