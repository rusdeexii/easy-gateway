<template>
    <div>
      <div v-if="loading" class="fixed inset-0 flex items-center justify-center bg-cyan-100 z-50">
        <div class="text-center">
          <div class="loader mx-auto mb-2 border-4 border-t-4 border-cyan-600 rounded-full w-12 h-12 animate-spin"></div>
          <h2 class="text-lg font-semibold text-gray-800">Authorizing</h2>
          <p class="text-gray-600">We are processing your request. Please wait a moment.</p>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, onMounted } from 'vue';
  import { router } from '@/router';
  
  const loading = ref(false);
  
  router.beforeEach((to, from, next) => {
    loading.value = true;
    next();
  });
  
  router.afterEach(() => {
    setTimeout(() => loading.value = false, 1000); // simulate delay
  });
  </script>
  
  <style>
  .loader {
    border-color: #c0e8f9 transparent transparent transparent;
  }
  </style>
  