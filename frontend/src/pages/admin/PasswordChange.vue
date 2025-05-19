<template>
    <div class="min-h-screen flex bg-gray-100">
      <div class="w-full md:w-1/2 lg:w-1/3 bg-white p-8 rounded-2xl shadow-xl">
        <h2 class="text-2xl font-bold mb-6 text-gray-700">เปลี่ยนรหัสผ่าน</h2>
        <form @submit.prevent="handleSubmit" class="space-y-4">
          <!-- รหัสผ่านเดิม -->
          <div>
            <label class="block text-sm font-semibold text-gray-600 mb-1">รหัสผ่านเดิม</label>
            <div class="relative">
              <i class="fas fa-lock absolute top-2.5 left-3 text-gray-400"></i>
              <input
                type="password"
                v-model="oldPassword"
                class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-400"
                required
              />
            </div>
          </div>
  
          <!-- รหัสผ่านใหม่ -->
          <div>
            <label class="block text-sm font-semibold text-gray-600 mb-1">รหัสผ่านใหม่</label>
            <div class="relative">
              <i class="fas fa-key absolute top-2.5 left-3 text-gray-400"></i>
              <input
                type="password"
                v-model="newPassword"
                class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-400"
                required
              />
            </div>
          </div>
  
          <!-- ยืนยันรหัสผ่านใหม่ -->
          <div>
            <label class="block text-sm font-semibold text-gray-600 mb-1">ยืนยันรหัสผ่านใหม่</label>
            <div class="relative">
              <i class="fas fa-check-circle absolute top-2.5 left-3 text-gray-400"></i>
              <input
                type="password"
                v-model="confirmPassword"
                class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-400"
                required
              />
            </div>
            <p v-if="passwordMismatch" class="text-sm text-red-500 mt-1">รหัสผ่านไม่ตรงกัน</p>
          </div>
  
          <!-- ปุ่มยืนยัน -->
          <button
            type="submit"
            :disabled="passwordMismatch"
            class="w-full bg-blue-500 hover:bg-blue-600 text-white py-2 rounded-xl transition disabled:opacity-50"
          >
            <i class="fas fa-save mr-2"></i> ยืนยันการเปลี่ยนรหัสผ่าน
          </button>
        </form>
       
      </div>
    </div>
  
  </template>
  
  <script lang="ts" setup>
  import { ref, computed } from 'vue'
  import Swal from 'sweetalert2'


  
  const oldPassword = ref('')
  const newPassword = ref('')
  const confirmPassword = ref('')
  
  const passwordMismatch = computed(() => {
  return newPassword.value !== confirmPassword.value
})

  const handleSubmit = () => {
  if (passwordMismatch.value) return

  // แสดง SweetAlert2
  Swal.fire({
    icon: 'success',
    title: 'สำเร็จ!',
    text: 'เปลี่ยนรหัสผ่านเรียบร้อยแล้ว',
    confirmButtonText: 'ตกลง',
    confirmButtonColor: '#1f9f4c',
  })
  }
  </script>
  