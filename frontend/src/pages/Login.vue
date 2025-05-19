<template>
  <div class="min-h-screen bg-gradient-to-br from-emerald-50 to-green-100 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
    <!-- Background decoration elements -->
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
          <img src="/logo.png" alt="Logo" class="w-full h-full object-contain">
        </div>
      </div>
      
      <!-- Title section -->
      <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
        Dashboard System
      </h2>
      <p class="mt-2 text-center text-sm text-gray-600">
        เข้าสู่ระบบเพื่อใช้งาน
      </p>
    </div>

    <!-- Login card -->
    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md z-10 relative">
      <div class="bg-white py-8 px-4 shadow-xl sm:rounded-lg sm:px-10 border-t-4 border-emerald-500 transition-all duration-300 hover:shadow-2xl">
        <!-- Login form -->
        <form class="space-y-6" @submit.prevent="handleLogin">
          <!-- Username Field -->
          <div>
            <label for="username" class="block text-sm font-medium text-gray-700">
              ชื่อผู้ใช้งานระบบ
            </label>
            <div class="mt-1 relative rounded-md shadow-sm">
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <i class="fas fa-user text-emerald-500"></i>
              </div>
              <input 
                id="username" 
                v-model="username" 
                name="username" 
                type="text" 
                required 
                class="block w-full pl-10 pr-3 py-3 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-500 focus:outline-none focus:ring-emerald-500 focus:border-emerald-500 transition duration-150 ease-in-out sm:text-sm"
                placeholder="กรอกชื่อผู้ใช้งาน"
                :class="{'border-red-500': usernameError}"
                
              >
            </div>
            <p v-if="usernameError" class="mt-1 text-xs text-red-600">{{ usernameError }}</p>
          </div>

          <!-- Password Field -->
          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">
              รหัสผู้ใช้งาน
            </label>
            <div class="mt-1 relative rounded-md shadow-sm">
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <i class="fas fa-lock text-emerald-500"></i>
              </div>
              <input 
                id="password" 
                v-model="password" 
                name="password" 
                :type="showPassword ? 'text' : 'password'" 
                required 
                class="block w-full pl-10 pr-10 py-3 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-500 focus:outline-none focus:ring-emerald-500 focus:border-emerald-500 transition duration-150 ease-in-out sm:text-sm"
                placeholder="กรอกรหัสผ่าน"
                :class="{'border-red-500': passwordError}"
                
              >
              <div class="absolute inset-y-0 right-0 pr-3 flex items-center">
                <button 
                  type="button" 
                  @click="togglePasswordVisibility" 
                  class="text-gray-400 hover:text-gray-600 focus:outline-none"
                >
                  <i class="fas" :class="showPassword ? 'fa-eye-slash' : 'fa-eye'"></i>
                </button>
              </div>
            </div>
            <p v-if="passwordError" class="mt-1 text-xs text-red-600">{{ passwordError }}</p>
          </div>

          <!-- Remember me checkbox -->
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <input 
                id="remember_me" 
                v-model="rememberMe" 
                name="remember_me" 
                type="checkbox" 
                class="h-4 w-4 text-emerald-600 focus:ring-emerald-500 border-gray-300 rounded cursor-pointer"
              >
              <label for="remember_me" class="ml-2 block text-sm text-gray-700 cursor-pointer">
                จดจำฉัน
              </label>
            </div>

            <div class="text-sm">
              <a href="#" class="font-medium text-emerald-600 hover:text-emerald-500 transition duration-150 ease-in-out">
                ลืมรหัสผ่าน?
              </a>
            </div>
          </div>

          <!-- Submit button -->
          <div>
            <button 
              type="submit" 
              :disabled="isLoading" 
              class="group relative w-full flex justify-center py-3 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-gradient-to-r from-emerald-600 to-green-500 hover:from-emerald-700 hover:to-green-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-emerald-500 transition-all duration-300 transform hover:scale-[1.02] disabled:opacity-70 disabled:cursor-not-allowed disabled:hover:scale-100"
            >
              <span class="absolute left-0 inset-y-0 flex items-center pl-3">
                <i 
                  class="fas" 
                  :class="isLoading ? 'fa-spinner fa-spin' : 'fa-sign-in-alt'"
                ></i>
              </span>
              <span v-if="!isLoading">เข้าสู่ระบบ</span>
              <span v-else>กำลังเข้าสู่ระบบ...</span>
            </button>
          </div>
        </form>
        
        <!-- Login Status Alert -->
        <transition name="fade">
          <div v-if="loginStatus.show" :class="`mt-6 p-4 rounded-md ${loginStatus.type === 'error' ? 'bg-red-50 text-red-700' : 'bg-green-50 text-green-700'}`">
            <div class="flex">
              <div class="flex-shrink-0">
                <i :class="`fas ${loginStatus.type === 'error' ? 'fa-exclamation-circle text-red-400' : 'fa-check-circle text-green-400'}`"></i>
              </div>
              <div class="ml-3">
                <p class="text-sm font-medium">{{ loginStatus.message }}</p>
              </div>
              <div class="ml-auto pl-3">
                <div class="-mx-1.5 -my-1.5">
                  <button
                    @click="loginStatus.show = false"
                    class="inline-flex rounded-md p-1.5 focus:outline-none focus:ring-2 focus:ring-offset-2"
                    :class="loginStatus.type === 'error' ? 'text-red-500 hover:bg-red-100 focus:ring-red-600' : 'text-green-500 hover:bg-green-100 focus:ring-green-600'"
                  >
                    <span class="sr-only">ปิด</span>
                    <i class="fas fa-times"></i>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </transition>
      </div>
      </div>

    <!-- Footer -->
    <div class="mt-8 text-center text-sm text-gray-500 z-10 relative">
      <p>© 2025 Developed by Devmour. All rights reserved.</p>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth.store'

const router = useRouter()
const authStore = useAuthStore()

// ฟอร์มและสถานะ
const username = ref('')
const password = ref('')
const rememberMe = ref(false)

const usernameError = ref('')
const passwordError = ref('')
const showPassword = ref(false)

const loginStatus = reactive({
  show: false,
  type: 'error' as 'success' | 'error',
  message: ''
})

const isLoading = ref(false)

const togglePasswordVisibility = () => {
  showPassword.value = !showPassword.value
}

const handleLogin = async () => {
  if (usernameError.value || passwordError.value) return;

  isLoading.value = true;
  loginStatus.show = false;

  try {
    await authStore.login(username.value, password.value);

    if (authStore.user) {
      loginStatus.type = 'success';
      loginStatus.message = 'เข้าสู่ระบบสำเร็จ กำลังนำคุณไปยังหน้าแรก...';
      loginStatus.show = true;

      setTimeout(() => {
        router.push('/system/dashboard');
      }, 1000);
    } else {
      throw new Error('Authentication failed');
    }
  } catch (error) {
    loginStatus.type = 'error';
    loginStatus.message = 'ชื่อผู้ใช้ หรือรหัสผ่านไม่ถูกต้อง';
    loginStatus.show = true;
    console.error('Login error:', error);
  } finally {
    isLoading.value = false;
  }
};


const loadRememberedUser = () => {
  const rememberedUser = localStorage.getItem('rememberedUser')
  if (rememberedUser) {
    const userData = JSON.parse(rememberedUser)
    username.value = userData.username || ''
    rememberMe.value = true
  }
}

loadRememberedUser()
</script>
