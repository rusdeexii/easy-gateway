<template>
    <div class="min-h-screen bg-gray-50">
      <!-- Header -->
      <header class="bg-blue-600 text-white p-4 flex items-center">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
          <path fill-rule="evenodd" d="M10 2a2 2 0 00-2 2v1H5a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-3V4a2 2 0 00-2-2zM5 7h10v10H5V7z" clip-rule="evenodd" />
        </svg>
        <h1 class="font-medium">ตั้งค่าระบบ</h1>
      </header>
  
      <!-- Main Content -->
      <div class="p-4">
        <div class="bg-white rounded-lg shadow">
          <!-- Tabs -->
          <div class="flex border-b">
            <button 
              @click="activeTab = 'company'"
              :class="[
                'px-4 py-3 flex items-center',
                activeTab === 'company' ? 'border-b-2 border-blue-600 text-blue-600' : 'text-gray-600'
              ]"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 14v3m4-3v3m4-3v3M3 21h18M3 10h18M3 7l9-4 9 4M4 10h16v11H4V10z" />
              </svg>
              ข้อมูลบริษัท
            </button>
            <button 
              @click="activeTab = 'email'"
              :class="[
                'px-4 py-3 flex items-center',
                activeTab === 'email' ? 'border-b-2 border-blue-600 text-blue-600' : 'text-gray-600'
              ]"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
              </svg>
              การตั้งค่าอีเมล
            </button>
          </div>
          
          <!-- Company Information Tab -->
          <div v-if="activeTab === 'company'" class="p-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">ชื่อบริษัท <span class="text-red-500">*</span></label>
                <input v-model="companyData.name" type="text" class="w-full border border-gray-300 rounded p-2" />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">เลขประจำตัวผู้เสียภาษี</label>
                <input v-model="companyData.taxId" type="text" class="w-full border border-gray-300 rounded p-2" />
              </div>
              <div class="md:col-span-2">
                <label class="block text-sm font-medium text-gray-700 mb-1">ที่อยู่</label>
                <input v-model="companyData.address" type="text" class="w-full border border-gray-300 rounded p-2" />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">โทรศัพท์</label>
                <input v-model="companyData.phone" type="text" class="w-full border border-gray-300 rounded p-2" />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">อีเมล</label>
                <input v-model="companyData.email" type="email" class="w-full border border-gray-300 rounded p-2" />
              </div>
            </div>
            <div class="mt-6">
              <button class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4" />
                </svg>
                บันทึกการตั้งค่า
              </button>
            </div>
          </div>
          
          <!-- Email Settings Tab -->
          <div v-if="activeTab === 'email'" class="p-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">SMTP Host</label>
                <input v-model="emailConfig.smtpHost" type="text" class="w-full border border-gray-300 rounded p-2" />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">SMTP Port</label>
                <input v-model="emailConfig.smtpPort" type="text" class="w-full border border-gray-300 rounded p-2" />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">SMTP Secure</label>
                <div class="relative">
                  <select v-model="emailConfig.smtpSecure" class="w-full border border-gray-300 rounded p-2 appearance-none pr-8">
                    <option value="TLS">TLS</option>
                    <option value="SSL">SSL</option>
                    <option value="none">None</option>
                  </select>
                  <div class="absolute inset-y-0 right-0 flex items-center px-2 pointer-events-none">
                    <svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path d="M19 9l-7 7-7-7" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" />
                    </svg>
                  </div>
                </div>
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">SMTP Username</label>
                <input v-model="emailConfig.smtpUsername" type="text" class="w-full border border-gray-300 rounded p-2" />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">SMTP Password</label>
                <input v-model="emailConfig.smtpPassword" type="password" class="w-full border border-gray-300 rounded p-2" />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">อีเมลผู้ส่ง</label>
                <input v-model="emailConfig.senderEmail" type="email" class="w-full border border-gray-300 rounded p-2" />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1">ชื่อผู้ส่ง</label>
                <input v-model="emailConfig.senderName" type="text" class="w-full border border-gray-300 rounded p-2" />
              </div>
            </div>
            
            <!-- Info box -->
            <div class="mt-6 bg-blue-50 border border-blue-100 rounded p-4 text-blue-800 flex items-start">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2 mt-0.5 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <p class="text-sm">การตั้งค่า SMTP ใช้สำหรับการส่งอีเมลในระบบจากที่นี่ถึงลูกค้า คุณสามารถใช้บริการอีเมลเซิร์ฟเวอร์ เช่น Gmail, Outlook หรือเซิร์ฟเวอร์อีเมลของบริษัทได้</p>
            </div>
            
            <!-- Buttons -->
            <div class="mt-6 flex items-center space-x-3">
              <button class="bg-cyan-500 hover:bg-cyan-600 text-white px-4 py-2 rounded flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                </svg>
                ทดสอบการส่งอีเมล
              </button>
              <button class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4" />
                </svg>
                บันทึกการตั้งค่า
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue';
  
  // Active tab state
  const activeTab = ref<'company' | 'email'>('company');
  
  // Company information data
  const companyData = ref({
    name: 'บริษัท ตัวอย่าง จำกัด',
    taxId: '0123456789012',
    address: '123 ถนนตัวอย่าง แขวงตัวอย่าง เขตตัวอย่าง กรุงเทพฯ 10000',
    phone: '02-123-4567',
    email: 'info@example.com'
  });
  
  // Email configuration data
  const emailConfig = ref({
    smtpHost: 'mail.rayongall.com',
    smtpPort: '587',
    smtpSecure: 'TLS',
    smtpUsername: 'contact@rayongall.com',
    smtpPassword: 'รหัสผ่านนี้ไม่ปลอดภัยควรใช้รหัสผ่านที่ซับซ้อน',
    senderEmail: 'contact@rayongall.com',
    senderName: ''
  });
  </script>