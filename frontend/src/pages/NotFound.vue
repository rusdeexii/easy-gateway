<template>
  <div class="min-h-screen bg-gradient-to-br from-emerald-50 to-green-100 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
    <!-- Background decoration elements - matching login page style -->
    <div class="absolute top-0 left-0 w-full h-64 bg-gradient-to-r from-emerald-600 to-green-500 transform -skew-y-3 origin-top-left z-0 opacity-20"></div>
    <div class="absolute inset-0 overflow-hidden z-0">
      <div class="absolute -top-10 -right-10 w-64 h-64 bg-emerald-300 rounded-full opacity-10"></div>
      <div class="absolute top-1/4 -left-20 w-72 h-72 bg-green-300 rounded-full opacity-10"></div>
      <div class="absolute bottom-1/3 right-10 w-80 h-80 bg-emerald-400 rounded-full opacity-10"></div>
    </div>
    
    <div class="sm:mx-auto sm:w-full sm:max-w-md z-10 relative">
      <!-- Logo section -->
      <div class="flex justify-center">
        <div class="w-20 h-20 rounded-full bg-white shadow-lg flex items-center justify-center p-4">
          <div class="w-12 h-12 text-amber-500 animate-pulse-slow">
            <svg class="w-full h-full" focusable="false" viewBox="0 0 24 24" aria-hidden="true">
              <path d="M4.47 21h15.06c1.54 0 2.5-1.67 1.73-3L13.73 4.99c-.77-1.33-2.69-1.33-3.46 0L2.74 18c-.77 1.33.19 3 1.73 3zM12 14c-.55 0-1-.45-1-1v-2c0-.55.45-1 1-1s1 .45 1 1v2c0 .55-.45 1-1 1zm1 4h-2v-2h2v2z" fill="currentColor"></path>
            </svg>
          </div>
        </div>
      </div>
      
      <!-- Title section -->
      <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
        <span class="text-red-600">503</span> Access Denied
      </h2>
      <p class="mt-2 text-center text-sm text-gray-600">
        ขออภัย ไม่สามารถดำเนินการคำขอของคุณได้ในขณะนี้
      </p>
    </div>
  

    <!-- Footer -->
    <div class="mt-8 text-center text-sm text-gray-500 z-10 relative">
      <p>© 2025 Developed by Devmour. All rights reserved.</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const isShaking = ref(false);


</script>

