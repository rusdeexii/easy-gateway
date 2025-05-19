<template>
  <component :is="layout" v-if="isReady && isAuthenticated">
    <router-view />
  </component>
  <router-view v-else-if="isReady" />
</template>

<script setup lang="ts">
import { computed, ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import DefaultLayout from '@/layouts/DefaultLayout.vue';
import AuthLayout from '@/layouts/AuthLayout.vue';
import { useAuthStore } from '@/stores/auth.store';

const route = useRoute();
const authStore = useAuthStore();
const isReady = ref(false);

onMounted(async () => {
  if (!authStore.user) {
    await authStore.init();
  }
  isReady.value = true;
});

const layout = computed(() => {
  const layoutName = route.meta.layout || 'default';
  return layoutName === 'auth' ? AuthLayout : DefaultLayout;
});

const isAuthenticated = computed(() => authStore.isAuthenticated());
</script>
