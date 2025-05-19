<template>
    <Transition name="modal">
        <div v-if="show" class="fixed inset-0 bg-opacity-10 backdrop-blur-xs flex justify-center items-start z-50">
            <div class="bg-white rounded-lg shadow-lg w-full max-w-2xl mt-16">
          <!-- Modal Header -->
          <div class="px-6 py-4 border-b border-gray-200 flex items-center">
            <div class="bg-blue-100 p-3 rounded-full mr-4">
              <i 
                :class="[
                  'fas text-blue-600 text-xl', 
                  user ? 'fa-user-edit' : 'fa-user-plus'
                ]"
              ></i>
            </div>
            <h2 class="text-xl font-bold text-gray-800">
              {{ user ? 'แก้ไขข้อมูลผู้ใช้' : 'เพิ่มผู้ใช้ใหม่' }}
            </h2>
            <button 
              @click="$emit('close')" 
              class="ml-auto text-gray-500 hover:text-gray-700 transition-colors"
            >
              <i class="fas fa-times text-xl"></i>
            </button>
          </div>
  
          <!-- Modal Body -->
          <form @submit.prevent="onSubmit" class="p-6 space-y-6">
            <!-- Username Input -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                ชื่อผู้ใช้
              </label>
              <div class="relative">
                <span class="absolute inset-y-0 left-0 pl-3 flex items-center text-gray-400">
                  <i class="fas fa-user"></i>
                </span>
                <input 
                  v-model="formData.username"
                  type="text" 
                  required
                  class="w-full pl-10 pr-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all"
                  placeholder="กรอกชื่อผู้ใช้"
                />
              </div>
            </div>
  
           
  
            <!-- Password Input -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                รหัสผ่าน
              </label>
              <div class="relative">
                <span class="absolute inset-y-0 left-0 pl-3 flex items-center text-gray-400">
                  <i class="fas fa-lock"></i>
                </span>
                <input 
                  v-model="formData.password"
                  type="password" 
                  :required="!user"
                  class="w-full pl-10 pr-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all"
                  :placeholder="user ? 'เว้นว่างถ้าไม่ต้องการเปลี่ยน' : 'กรอกรหัสผ่าน'"
                />
              </div>
              <p v-if="user" class="text-xs text-gray-500 mt-1">
                <i class="fas fa-info-circle mr-1"></i> 
                ปล่อยว่างถ้าไม่ต้องการเปลี่ยนรหัสผ่าน
              </p>
            </div>
  
            <!-- Role Select -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                บทบาท
              </label>
              <div class="relative">
                <span class="absolute inset-y-0 left-0 pl-3 flex items-center text-gray-400">
                  <i class="fas fa-user-tag"></i>
                </span>
                <select 
                  v-model="formData.role"
                  class="w-full pl-10 pr-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all"
                >
                  <option value="admin">Admin</option>
                  <option value="superadmin">Super Admin</option>
                </select>
              </div>
            </div>
  
            <!-- Active Toggle -->
            <div class="flex items-center">
              <div 
                @click="formData.is_active = !formData.is_active"
                :class="[
                  'relative inline-flex flex-shrink-0 h-6 w-11 border-2 rounded-full cursor-pointer transition-colors ease-in-out duration-200',
                  formData.is_active ? 'bg-green-500 border-green-500' : 'bg-gray-200 border-gray-300'
                ]"
              >
                <span 
                  :class="[
                    'inline-block h-5 w-5 rounded-full bg-white shadow transform transition-transform ease-in-out duration-200',
                    formData.is_active ? 'translate-x-5' : 'translate-x-0'
                  ]"
                ></span>
              </div>
              <span 
                :class="[
                  'ml-3 text-sm font-medium',
                  formData.is_active ? 'text-green-700' : 'text-gray-500'
                ]"
              >
                {{ formData.is_active ? 'เปิดใช้งาน' : 'ปิดการใช้งาน' }}
              </span>
            </div>
  
            <!-- Action Buttons -->
            <div class="flex justify-end space-x-3 pt-4">
              <button 
                type="button"
                @click="$emit('close')"
                class="px-4 py-2 bg-gray-100 text-gray-800 rounded-lg hover:bg-gray-200 transition-colors"
              >
                ยกเลิก
              </button>
              <button 
                type="submit"
                class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
              >
                {{ user ? 'บันทึกการแก้ไข' : 'เพิ่มผู้ใช้' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </Transition>
  </template>
  
  <script setup lang="ts">
  import { ref, watch } from 'vue'
  import type { User, CreateUserRequest } from '@/models/user.model'
  
  interface Props {
    show: boolean
    user?: User | null
  }
  
  const props = defineProps<Props>()
  
  const emit = defineEmits(['close', 'submit'])
  
  const formData = ref<CreateUserRequest>({
    username: '',
    password: '',
    role: 'admin',
    is_active: true,
    email: ''
  })
  
  // Reset form when modal opens/closes or user changes
  watch(() => props.show, (newShow) => {
    if (newShow) {
      // If editing existing user
      if (props.user) {
        formData.value = {
          username: props.user.username,
          password: '',
          role: props.user.role as 'admin' | 'superadmin',
          is_active: props.user.is_active,
          email: props.user.email || ''
        }
      } else {
        // Reset for new user
        formData.value = {
          username: '',
          password: '',
          role: 'admin',
          is_active: true,
          email: ''
        }
      }
    }
  }, { immediate: true })
  
  const onSubmit = () => {
    emit('submit', formData.value)
  }
  </script>
  
  <style scoped>
  .modal-enter-active,
  .modal-leave-active {
    transition: all 0.3s ease;
  }
  
  .modal-enter-from,
  .modal-leave-to {
    opacity: 0;
    transform: scale(0.9);
  }
  
  .modal-enter-to,
  .modal-leave-from {
    opacity: 1;
    transform: scale(1);
  }
  </style>

