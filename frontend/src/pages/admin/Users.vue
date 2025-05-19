<template>
  <section class="p-6 bg-white shadow-lg rounded-lg min-h-screen border border-gray-100">
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-center mb-8 gap-4">
      <div class="flex items-center">
        <span class="bg-blue-500 text-white p-3 rounded-lg mr-4 shadow-md">
          <i class="fas fa-users text-xl"></i>
        </span>
        <h1 class="text-2xl font-bold text-gray-800 tracking-tight">จัดการผู้ใช้ระบบ</h1>
      </div>
          
          <div class="flex flex-col md:flex-row items-center gap-4 w-full md:w-auto">           
            <!-- Add User Button -->
            <button 
              @click="openModal()" 
              class="flex items-center justify-center w-full md:w-auto px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-all duration-200 shadow-md hover:shadow-lg transform hover:scale-105"
            >
              <i class="fas fa-plus-circle mr-2"></i>
              <span>เพิ่มผู้ใช้</span>
            </button>
          </div>
        </div>

        <!-- Status Cards -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
          <div v-for="(stat, index) in userStats" :key="index" 
            :class="[
              'border-l-4 rounded-lg p-5 shadow-sm transform transition-all duration-300 hover:scale-105 hover:shadow-md',
              stat.borderColor, 
              stat.bgColor
            ]"
          >
            <div class="flex items-center">
              <i :class="['text-2xl mr-4', stat.icon]"></i>
              <div>
                <h3 class="font-semibold text-gray-700">{{ stat.title }}</h3>
                <p class="text-2xl font-bold" :class="stat.textColor">
                  {{ stat.count }}
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- User Table -->
        <div class="bg-white rounded-lg border border-gray-200 overflow-hidden shadow-sm">
          <div v-if="loading" class="flex justify-center items-center py-16">
            <div class="animate-spin rounded-full h-12 w-12 border-t-4 border-b-4 border-blue-500"></div>
          </div>
          
          <div v-else-if="users.length === 0" class="text-center py-16">
            <i class="fas fa-user-slash text-6xl text-gray-300 mb-4"></i>
            <p class="text-gray-500 text-lg mb-4">ไม่พบข้อมูลผู้ใช้</p>
            <button 
              @click="openModal()" 
              class="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            >
              <i class="fas fa-plus-circle mr-2"></i>
              เพิ่มผู้ใช้ใหม่
            </button>
          </div>

          <table v-else class="w-full">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  ชื่อผู้ใช้
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  บทบาท
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  สถานะ
                </th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  จัดการ
                </th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr 
                v-for="user in paginatedUsers" 
                :key="user.user_id"
                class="hover:bg-gray-50 transition-colors"
              >
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="flex-shrink-0 h-10 w-10 bg-blue-100 rounded-full flex items-center justify-center">
                      <i :class="[
                        'fas text-blue-600', 
                        user.role === 'superadmin' ? 'fa-user-shield' : 'fa-user'
                      ]"></i>
                    </div>
                    <div class="ml-4">
                      <div class="text-sm font-medium text-gray-900">
                        {{ user.username }}
                      </div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="[
                    'px-3 py-1 rounded-full text-xs capitalize font-medium',
                    user.role === 'superadmin' 
                      ? 'bg-purple-100 text-purple-800' 
                      : 'bg-blue-100 text-blue-800'
                  ]">
                    {{ user.role }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="[
                    'px-3 py-1 rounded-full text-xs font-medium',
                    user.is_active 
                      ? 'bg-green-100 text-green-800' 
                      : 'bg-red-100 text-red-800'
                  ]">
                    {{ user.is_active ? 'ใช้งาน' : 'ปิดใช้งาน' }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-right">
                  <div class="flex justify-end space-x-2">
                    <button 
                      @click="openModal(user)" 
                      class="text-yellow-600 hover:text-yellow-800 bg-yellow-100 hover:bg-yellow-200 p-2 rounded-full transition-colors"
                    >
                      <i class="fas fa-edit"></i>
                    </button>
                    <button 
                      @click="confirmDelete(user.user_id)" 
                      class="text-red-600 hover:text-red-800 bg-red-100 hover:bg-red-200 p-2 rounded-full transition-colors"
                    >
                      <i class="fas fa-trash-alt"></i>
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>

         <!-- Enhanced Pagination -->
<div v-if="users.length > 0" class="bg-gray-50 px-4 py-3 flex flex-col md:flex-row items-center justify-between border-t">
  <div class="hidden md:block w-full md:w-auto mb-2 md:mb-0">
    <p class="text-sm text-gray-700 flex items-center">
      <span class="mr-2">
        <i class="fas fa-list-ol text-gray-500 mr-1"></i>
        แสดง
      </span>
      <select 
        v-model="itemsPerPage" 
        class="form-select mx-2 py-1 px-2 border rounded-md text-sm focus:ring-2 focus:ring-blue-500"
      >
        <option :value="5">5 รายการ</option>
        <option :value="10">10 รายการ</option>
        <option :value="25">25 รายการ</option>
        <option :value="50">50 รายการ</option>
      </select>
      <span>
        จาก 
        <span class="font-medium mx-1">{{ users.length }}</span> 
        รายการทั้งหมด
      </span>
    </p>
  </div>
  
  <div class="flex items-center space-x-2">
    <nav class="inline-flex rounded-md shadow-sm -space-x-px">
      <!-- First Page Button -->
      <button 
        @click="goToPage(1)"
        :disabled="currentPage === 1"
        class="px-2 py-2 border bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 rounded-l-md"
        title="หน้าแรก"
      >
        <i class="fas fa-angle-double-left"></i>
      </button>
      
      <!-- Previous Page Button -->
      <button 
        @click="goToPage(currentPage - 1)"
        :disabled="currentPage === 1"
        class="px-3 py-2 border bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50"
        title="หน้าก่อนหน้า"
      >
        <i class="fas fa-chevron-left"></i>
      </button>
      
      <!-- Page Number Buttons -->
      <template v-if="totalPages <= 7">
        <button 
          v-for="page in totalPages" 
          :key="page"
          @click="goToPage(page)"
          :class="[
            'px-4 py-2 border text-sm font-medium transition-all duration-200',
            page === currentPage 
              ? 'bg-blue-500 text-white' 
              : 'bg-white text-gray-700 hover:bg-gray-100'
          ]"
        >
          {{ page }}
        </button>
      </template>
      
      <!-- Condensed Pagination for Many Pages -->
      <template v-else>
        <!-- First page -->
        <button 
          v-if="currentPage > 3"
          @click="goToPage(1)"
          class="px-4 py-2 border bg-white text-sm font-medium text-gray-700 hover:bg-gray-100"
        >
          1
        </button>
        
        <!-- Ellipsis before current page group -->
        <span 
          v-if="currentPage > 3" 
          class="px-4 py-2 border bg-white text-sm font-medium text-gray-700"
        >
          ...
        </span>
        
        <!-- Dynamic page range -->
        <button 
          v-for="page in pageRange" 
          :key="page"
          @click="goToPage(page)"
          :class="[
            'px-4 py-2 border text-sm font-medium transition-all duration-200',
            page === currentPage 
              ? 'bg-blue-500 text-white' 
              : 'bg-white text-gray-700 hover:bg-gray-100'
          ]"
        >
          {{ page }}
        </button>
        
        <!-- Ellipsis after current page group -->
        <span 
          v-if="currentPage < totalPages - 2" 
          class="px-4 py-2 border bg-white text-sm font-medium text-gray-700"
        >
          ...
        </span>
        
        <!-- Last page -->
        <button 
          v-if="currentPage < totalPages - 2"
          @click="goToPage(totalPages)"
          class="px-4 py-2 border bg-white text-sm font-medium text-gray-700 hover:bg-gray-100"
        >
          {{ totalPages }}
        </button>
      </template>
      
      <!-- Next Page Button -->
      <button 
        @click="goToPage(currentPage + 1)"
        :disabled="currentPage === totalPages"
        class="px-3 py-2 border bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50"
        title="หน้าถัดไป"
      >
        <i class="fas fa-chevron-right"></i>
      </button>
      
      <!-- Last Page Button -->
      <button 
        @click="goToPage(totalPages)"
        :disabled="currentPage === totalPages"
        class="px-2 py-2 border bg-white text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:opacity-50 rounded-r-md"
        title="หน้าสุดท้าย"
      >
        <i class="fas fa-angle-double-right"></i>
      </button>
    </nav>
  </div>
</div>
        </div>
        <!-- Pagination -->
      </section>
  

    <!-- Modal for Add/Edit User -->
    <UserModal 
      :show="showModal" 
      :user="selectedUser" 
      @close="closeModal" 
      @submit="submitForm"
    />

</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue'
import { useAuthStore } from '@/stores/auth.store'
import UserModal from '@/components/Modal.vue'
import Swal from 'sweetalert2'
import type { User, CreateUserRequest } from '@/models/user.model'

const auth = useAuthStore()
const users = ref<User[]>([])
const loading = ref(false)
const showModal = ref(false)
const selectedUser = ref<User | null>(null)
const currentPage = ref(1)
const itemsPerPage = ref(10)
// Pagination-related computed properties and methods
const pageRange = computed(() => {
  const range = []
  const start = Math.max(1, currentPage.value - 1)
  const end = Math.min(totalPages.value, currentPage.value + 1)
  
  for (let i = start; i <= end; i++) {
    range.push(i)
  }
  
  return range
})

const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
  }
}

// Watch for changes in items per page to reset current page
watch(itemsPerPage, () => {
  currentPage.value = 1
})
const userStats = computed(() => [
  {
    title: 'ผู้ใช้ทั้งหมด',
    count: users.value.length,
    icon: 'fas fa-users text-blue-500',
    borderColor: 'border-blue-500',
    bgColor: 'bg-blue-50',
    textColor: 'text-blue-600'
  },
  {
    title: 'ผู้ใช้ที่ใช้งานอยู่',
    count: users.value.filter(u => u.is_active).length,
    icon: 'fas fa-toggle-on text-green-500',
    borderColor: 'border-green-500',
    bgColor: 'bg-green-50',
    textColor: 'text-green-600'
  },
  {
    title: 'ผู้ใช้ที่ปิดใช้งาน',
    count: users.value.filter(u => !u.is_active).length,
    icon: 'fas fa-toggle-off text-red-500',
    borderColor: 'border-red-500',
    bgColor: 'bg-red-50',
    textColor: 'text-red-600'
  }
])

const paginatedUsers = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return users.value.slice(start, end)
})

const totalPages = computed(() => 
  Math.ceil(users.value.length / itemsPerPage.value)
)

const fetchUsers = async () => {
  loading.value = true
  try {
    users.value = await auth.fetchUsers()
  } catch (e) {
    console.error('Failed to fetch users', e)
    Swal.fire('เกิดข้อผิดพลาด', 'ไม่สามารถโหลดข้อมูลผู้ใช้', 'error')
  } finally {
    loading.value = false
  }
}



const openModal = (user?: User) => {
  selectedUser.value = user || null
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  selectedUser.value = null
}

const submitForm = async (formData: CreateUserRequest) => {
  try {
    if (selectedUser.value) {
      await auth.updateUser(selectedUser.value.user_id, formData)
      Swal.fire('สำเร็จ', 'แก้ไขผู้ใช้เรียบร้อย', 'success')
    } else {
      await auth.createUser(formData)
      Swal.fire('สำเร็จ', 'เพิ่มผู้ใช้เรียบร้อย', 'success')
    }

    closeModal()
    await fetchUsers()
  } catch (e) {
    Swal.fire('เกิดข้อผิดพลาด', 'ไม่สามารถดำเนินการได้', 'error')
  }
}

const confirmDelete = async (userId: string) => {
  const result = await Swal.fire({
    title: 'แน่ใจหรือไม่?',
    text: 'คุณกำลังจะลบผู้ใช้นี้',
    icon: 'warning',
    showCancelButton: true,
    confirmButtonText: 'ลบเลย',
    cancelButtonText: 'ยกเลิก',
  });

  if (result.isConfirmed) {
    try {
      await auth.deleteUser(userId)
      await fetchUsers()
      Swal.fire('สำเร็จ', 'ลบผู้ใช้เรียบร้อย', 'success')
    } catch (e) {
      Swal.fire('เกิดข้อผิดพลาด', 'ไม่สามารถดำเนินการได้', 'error')
    }
  }
} 
onMounted(() => {
  fetchUsers()
})
</script>
<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: all 0.1s ease-in-out;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
  transform: scale(0.95);
}

.modal-enter-to,
.modal-leave-from {
  opacity: 1;
  transform: scale(1);
}
</style>
