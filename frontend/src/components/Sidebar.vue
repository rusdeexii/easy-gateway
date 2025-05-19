<template>
  <div class="flex h-screen bg-gray-50">
    <!-- Mobile menu button with animation -->
    <div class="lg:hidden fixed top-0 left-0 z-20 m-4">
      <button 
        @click="toggleSidebar" 
        class="p-2 rounded-md bg-emerald-600 text-white hover:bg-emerald-700 focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2 transition-all duration-200 shadow-lg hover:shadow-emerald-500/50"
      >
        <span class="sr-only">Open sidebar</span>
        <div class="w-6 h-6 relative">
          <span 
            class="block absolute h-0.5 w-6 bg-white transform transition-all duration-300 ease-in-out"
            :class="isSidebarOpen ? 'rotate-45 top-3' : 'top-1'"
          ></span>
          <span 
            class="block absolute h-0.5 w-6 bg-white transform transition-all duration-300 ease-in-out top-3"
            :class="isSidebarOpen ? 'opacity-0' : 'opacity-100'"
          ></span>
          <span 
            class="block absolute h-0.5 w-6 bg-white transform transition-all duration-300 ease-in-out"
            :class="isSidebarOpen ? '-rotate-45 top-3' : 'top-5'"
          ></span>
        </div>
      </button>
    </div>

    <!-- Sidebar with green theme -->
    <aside 
      :class="`${isSidebarOpen ? 'translate-x-0' : '-translate-x-full'} lg:translate-x-0 fixed lg:static inset-y-0 left-0 z-10 w-64 bg-gradient-to-b from-emerald-800 to-emerald-900 text-gray-100 flex flex-col transition-all duration-300 ease-in-out shadow-xl`"
    >
      <!-- Header with gradient effect -->
      <div class="flex items-center p-4 border-b border-emerald-700 bg-gradient-to-r from-emerald-800 to-emerald-700">
        <div class="w-10 h-10 rounded-full bg-white flex items-center justify-center mr-3 shadow-lg transform transition-transform duration-200 hover:scale-110">
          <img src="/logo.png" alt="Logo" class="w-8 h-8 object-cover rounded-full">
        </div>
        <h1 class="text-xl font-bold text-white">BackOffice</h1>
      </div>
      
      <!-- Navigation Sections with improved styling -->
      <div class="flex-1 overflow-y-auto py-2 scrollbar-thin scrollbar-thumb-emerald-600 scrollbar-track-emerald-900">
        <!-- Manage Inventory Section -->
        <div class="px-4 pt-4 pb-2">
          <p class="text-xs font-medium text-emerald-300 uppercase tracking-wider mb-2 pl-2">จัดการระบบ</p>
          <nav class="space-y-1">
            <router-link 
              to="/system/dashboard" 
              custom
              v-slot="{ navigate, isActive }"
            >
              <a 
                @click="navigate(); closeSidebarOnMobile()"
                :class="`flex items-center px-3 py-2.5 text-sm rounded-lg transition-all duration-200 ${
                  isActive 
                    ? 'bg-emerald-600 text-white font-medium shadow-md' 
                    : 'text-gray-200 hover:bg-emerald-700/70'
                }`"
              >
                <div :class="`mr-3 w-6 h-6 flex items-center justify-center rounded-md ${isActive ? 'bg-white text-emerald-600' : 'text-emerald-300'}`">
                  <i class="fas fa-tachometer-alt"></i>
                </div>
                <span>หน้าหลัก</span>
                <span v-if="isActive" class="ml-auto w-1.5 h-8 bg-white rounded-sm"></span>
              </a>
            </router-link>
            
            <router-link 
              to="/system/products" 
              custom
              v-slot="{ navigate, isActive }"
            >
              <a 
                @click="navigate(); closeSidebarOnMobile()"
                :class="`flex items-center px-3 py-2.5 text-sm rounded-lg transition-all duration-200 ${
                  isActive 
                    ? 'bg-emerald-600 text-white font-medium shadow-md' 
                    : 'text-gray-200 hover:bg-emerald-700/70'
                }`"
              >
                <div :class="`mr-3 w-6 h-6 flex items-center justify-center rounded-md ${isActive ? 'bg-white text-emerald-600' : 'text-emerald-300'}`">
                  <i class="fas fa-file-alt"></i>
                </div>
                <span>รายงานธุรกรรม</span>
                <span v-if="isActive" class="ml-auto w-1.5 h-8 bg-white rounded-sm"></span>
              </a>
            </router-link>
            
            <router-link 
              to="/system/quotations" 
              custom
              v-slot="{ navigate, isActive }"
            >
              <a 
                @click="navigate(); closeSidebarOnMobile()"
                :class="`flex items-center px-3 py-2.5 text-sm rounded-lg transition-all duration-200 ${
                  isActive 
                    ? 'bg-emerald-600 text-white font-medium shadow-md' 
                    : 'text-gray-200 hover:bg-emerald-700/70'
                }`"
              >
                <div :class="`mr-3 w-6 h-6 flex items-center justify-center rounded-md ${isActive ? 'bg-white text-emerald-600' : 'text-emerald-300'}`">
                  <i class="fas fa-box"></i>
                </div>
                <span>จัดการร้านค้า</span>
                <span v-if="isActive" class="ml-auto w-1.5 h-8 bg-white rounded-sm"></span>
              </a>
            </router-link>

          </nav>
        </div>
        
        <!-- Account Settings Section -->
        <div class="px-4 pt-6 pb-2">
          <p class="text-xs font-medium text-emerald-300 uppercase tracking-wider mb-2 pl-2">เฉพาะเจ้าหน้าที่</p>
          <nav class="space-y-1">
            <router-link 
              to="/system/users" 
              custom
              v-slot="{ navigate, isActive }"
            >
              <a 
                @click="navigate(); closeSidebarOnMobile()"
                :class="`flex items-center px-3 py-2.5 text-sm rounded-lg transition-all duration-200 ${
                  isActive 
                    ? 'bg-emerald-600 text-white font-medium shadow-md' 
                    : 'text-gray-200 hover:bg-emerald-700/70'
                }`"
              >
                <div :class="`mr-3 w-6 h-6 flex items-center justify-center rounded-md ${isActive ? 'bg-white text-emerald-600' : 'text-emerald-300'}`">
                  <i class="fas fa-user"></i>
                </div>
                <span>จัดการผู้ใช้งาน</span>
                <span v-if="isActive" class="ml-auto w-1.5 h-8 bg-white rounded-sm"></span>
              </a>
            </router-link>
            <a 
              href="#"
              @click.prevent="logout"
              class="flex items-center px-3 py-2.5 text-sm rounded-lg transition-all duration-200 text-gray-200 hover:bg-emerald-700/70"
            >
              <div class="mr-3 w-6 h-6 flex items-center justify-center rounded-md text-emerald-300">
                <i class="fas fa-sign-out-alt"></i>
              </div>
              <span>ออกจากระบบ</span>
            </a>



          </nav>
        </div>
      </div>
      
     
      
      <!-- Footer -->
      <footer class="border-t border-emerald-700 p-3 bg-emerald-900/50">
        <div class="text-sm text-emerald-200 text-center font-semibold">
          &copy; 2025  Devmour.com
        </div>
      </footer>
    </aside>

    <!-- Mini Sidebar (for collapsed state) -->
    

    <!-- Main Content with padding adjustment -->
   <!-- Main Content - FIXED: removed lg:ml-64 -->
<div class="flex-1 flex flex-col">
  <!-- Top bar for mobile -->
  <div class="lg:hidden h-16 bg-white shadow-sm flex items-center px-4">
    <h1 class="text-xl font-semibold text-gray-800">Quotation System</h1>
  </div>
  
  <!-- Page Content - FIXED: added lg:pl-64 to maintain proper spacing from sidebar -->
  <main class="flex-1 overflow-y-auto px-4 py-6">
    <router-view></router-view>
  </main>
</div>
    <!-- Overlay when sidebar is open on mobile -->
    <div 
      v-if="isSidebarOpen" 
      class="lg:hidden fixed inset-0 bg-opacity-50 z-0 transition-all duration-300"
      @click="isSidebarOpen = false"
    ></div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, watch } from 'vue';
import { useRouter } from 'vue-router';
import Swal from 'sweetalert2';
import { useAuthStore } from '@/stores/auth.store';


const router = useRouter();
const authStore = useAuthStore();
const isSidebarOpen = ref(false);

const windowWidth = ref(window.innerWidth);

// Toggle sidebar for mobile
const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value;
};

// Close sidebar on mobile when navigation occurs
const closeSidebarOnMobile = () => {
  if (windowWidth.value < 1024) {
    isSidebarOpen.value = false;
  }
};

// Toggle collapse state for desktop


// Handle window resize
const handleResize = () => {
  windowWidth.value = window.innerWidth;
  
  // Auto-close sidebar on small screens
  if (windowWidth.value < 1024) {
    isSidebarOpen.value = false;
  }
};

const logout = async () => {
  const result = await Swal.fire({
    title: 'ออกจากระบบ',
    text: 'คุณต้องการออกจากระบบใช่หรือไม่?',
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#cb4040',
    cancelButtonColor: '#3085d6',
    confirmButtonText: 'ออกจากระบบ',
    cancelButtonText: 'ยกเลิก',
  });

  if (result.isConfirmed) {
    try {
      await authStore.logout();

      Swal.fire({
        icon: 'success',
        title: 'ออกจากระบบแล้ว',
        text: 'ขอบคุณที่ใช้งาน',
        timer: 1500,
        showConfirmButton: false,
      });

      // เปลี่ยนหน้าไปยัง login หลังจาก logout เสร็จ
      setTimeout(() => {
        router.push('/auth/signin');
      }, 1500);
    } catch (err) {
      Swal.fire('เกิดข้อผิดพลาด', 'ไม่สามารถออกจากระบบได้', 'error');
    }
  }
};



onMounted(() => {
  window.addEventListener('resize', handleResize);
  window.addEventListener('beforeunload', closeSidebarOnMobile);
});

onBeforeUnmount(() => {
  window.removeEventListener('resize', handleResize);
  window.removeEventListener('beforeunload', closeSidebarOnMobile);
});
</script>

<style>
@import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css');

/* Custom scrollbar */
.scrollbar-thin::-webkit-scrollbar {
  width: 4px;
}

.scrollbar-thin::-webkit-scrollbar-track {
  background: rgba(16, 185, 129, 0.1);
}

.scrollbar-thin::-webkit-scrollbar-thumb {
  background: rgba(16, 185, 129, 0.4);
  border-radius: 4px;
}

.scrollbar-thin::-webkit-scrollbar-thumb:hover {
  background: rgba(16, 185, 129, 0.6);
}

/* Smooth transitions */
.router-link-active {
  position: relative;
}

.router-link-active::after {
  content: '';
  position: absolute;
  right: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 3px;
  height: 70%;
  background: white;
  border-radius: 2px;
}
</style>