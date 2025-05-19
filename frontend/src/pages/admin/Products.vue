<template>
  <section class="min-h-screen bg-gray-50">
    <!-- Header with Product Management title -->
    <div class="flex justify-between items-center p-4">
      <div class="flex items-center space-x-2">
        <i class="fas fa-box text-gray-600"></i>
        <h1 class="text-lg font-medium">จัดการสินค้า</h1>
      </div>
      <button v-if="!isAddingProduct" @click="startAddProduct" class="bg-blue-600 text-white px-3 py-1.5 rounded-md flex items-center space-x-1">
        <i class="fas fa-plus-circle text-sm"></i>
        <span>เพิ่มสินค้าใหม่</span>
      </button>
      <button v-else @click="cancelAddProduct" class="bg-gray-500 text-white px-3 py-1.5 rounded-md flex items-center space-x-1">
        <i class="fas fa-arrow-left text-sm"></i>
        <span>กลับไปหน้าสินค้า</span>
      </button>
    </div>

    <!-- Main Content Area -->
    <div class="mt-2 mx-4">
      <!-- Empty State - shown when no products -->
      <div v-if="products.length === 0 && !isAddingProduct" class="bg-white rounded-lg shadow overflow-hidden flex flex-col items-center justify-center py-16">
        <div class="text-gray-500 mb-2">
          <i class="fas fa-box text-5xl"></i>
        </div>
        <p class="text-gray-600 mb-1">ยังไม่มีสินค้า</p>
        <p class="text-gray-500 text-sm mb-4">เริ่มเพิ่มสินค้าให้เมนูเลย</p>
        <button @click="startAddProduct" class="bg-blue-600 text-white px-4 py-2 rounded-md flex items-center space-x-1">
          <i class="fas fa-plus-circle"></i>
          <span>เพิ่มสินค้า</span>
        </button>
      </div>

      <!-- Product Table -->
      <div v-if="!isAddingProduct && products.length > 0" class="bg-white rounded-lg shadow overflow-hidden">
        <!-- Blue header bar -->
        <div class="bg-blue-600 text-white p-3 flex items-center space-x-2">
          <i class="fas fa-list-ul"></i>
          <span>รายการสินค้าทั้งหมด</span>
        </div>

        <div class="px-4 py-3">
          <!-- Filter and Display Options -->
          <div class="flex justify-between items-center mb-3">
            <div class="flex items-center space-x-2">
              <span class="text-sm text-gray-600">แสดง</span>
              <select v-model="itemsPerPage" class="border rounded px-2 py-1 text-sm">
                <option value="10">10</option>
                <option value="20">20</option>
                <option value="50">50</option>
              </select>
              <span class="text-sm text-gray-600">รายการ</span>
            </div>
            
            <div class="flex items-center relative">
              <span class="text-sm text-gray-600 mr-2">ค้นหา:</span>
              <input type="text" v-model="searchTerm" class="border rounded pl-2 pr-2 py-1 text-sm" />
            </div>
          </div>

          <!-- Table -->
          <table class="w-full text-sm">
            <thead>
              <tr class="text-left border-b">
                <th class="py-2 px-2">รหัสสินค้า</th>
                <th class="py-2 px-2">ชื่อสินค้า</th>
                <th class="py-2 px-2">รายละเอียด</th>
                <th class="py-2 px-2">หน่วย</th>
                <th class="py-2 px-2">ราคา</th>
                <th class="py-2 px-2">จัดการ</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="product in paginatedProducts" :key="product.id" class="border-b hover:bg-gray-50">
                <td class="py-2 px-2">{{ product.id }}</td>
                <td class="py-2 px-2">{{ product.name }}</td>
                <td class="py-2 px-2">{{ product.description }}</td>
                <td class="py-2 px-2">{{ product.unit }}</td>
                <td class="py-2 px-2">{{ product.price }} บาท</td>
                <td class="py-2 px-2">
                  <div class="flex space-x-1">
                    <button @click="editProduct(product)" class="bg-cyan-500 text-white px-2 py-1 rounded text-xs">
                      <i class="fas fa-edit"></i> แก้ไข
                    </button>
                    <button @click="deleteProduct(product.id)" class="bg-red-500 text-white px-2 py-1 rounded text-xs">
                      <i class="fas fa-trash"></i> ลบ
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>

          <!-- Pagination -->
          <div class="mt-3 flex justify-between items-center text-sm">
            <div>
              แสดง {{ startItem }} ถึง {{ endItem }} จาก {{ filteredProducts.length }} รายการ
            </div>
            <div class="flex items-center space-x-1">
              <button @click="previousPage" :disabled="currentPage === 1" 
                :class="currentPage === 1 ? 'border px-3 py-1 rounded bg-gray-100 text-gray-400' : 'border px-3 py-1 rounded bg-gray-100 hover:bg-gray-200 text-gray-700'">
                <i class="fas fa-chevron-left"></i> ก่อนหน้า
              </button>
              <span class="border px-3 py-1 rounded bg-blue-600 text-white">{{ currentPage }}</span>
              <button @click="nextPage" :disabled="currentPage >= totalPages" 
                :class="currentPage >= totalPages ? 'border px-3 py-1 rounded bg-gray-100 text-gray-400' : 'border px-3 py-1 rounded bg-gray-100 hover:bg-gray-200 text-gray-700'">
                ถัดไป <i class="fas fa-chevron-right"></i>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Add Product Modal -->
    <div v-if="showAddModal" class="fixed inset-0 bg-opacity-10 backdrop-blur-xs flex justify-center items-start z-50">
      <div class="bg-white rounded-lg shadow-lg w-full max-w-2xl mt-16">
        <!-- Modal Header -->
        <div class="bg-blue-600 text-white p-3">
          <h3 class="font-medium">เพิ่มสินค้าใหม่</h3>
        </div>

        <!-- Modal Form -->
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">ชื่อสินค้า <span class="text-red-500">*</span></label>
              <input v-model="newProduct.name" type="text" class="w-full border border-gray-300 rounded p-2" />
            </div>
            <div class="flex gap-4">
              <div class="flex-grow">
                <label class="block text-sm font-medium text-gray-700 mb-1">ราคา (บาท) <span class="text-red-500">*</span></label>
                <input v-model="newProduct.price" type="number" class="w-full border border-gray-300 rounded p-2" />
              </div>
              <div class="w-32">
                <label class="block text-sm font-medium text-gray-700 mb-1">หน่วย</label>
                <input v-model="newProduct.unit" type="text" class="w-full border border-gray-300 rounded p-2" placeholder="เช่น ชิ้น, อัน, กล่อง" />
              </div>
            </div>
          </div>
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-1">รายละเอียดสินค้า</label>
            <textarea v-model="newProduct.description" rows="4" class="w-full border border-gray-300 rounded p-2"></textarea>
          </div>
          <div class="flex">
            <button @click="saveProduct" class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded mr-2 flex items-center">
              <i class="fas fa-save mr-1"></i>
              บันทึกข้อมูล
            </button>
            <button @click="closeAddModal" class="border border-gray-300 text-gray-700 px-4 py-2 rounded hover:bg-gray-50 flex items-center">
              <i class="fas fa-times mr-1"></i>
              ยกเลิก
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Edit Product Modal -->
    <div v-if="showEditModal" class="fixed inset-0  bg-opacity-10 backdrop-blur-xs flex justify-center items-start z-50 pt-16">
      <div class="bg-white rounded-lg shadow-lg w-full max-w-2xl">
        <!-- Modal Header -->
        <div class="bg-blue-600 text-white p-3">
          <h3 class="font-medium">แก้ไขข้อมูลสินค้า</h3>
        </div>

        <!-- Modal Form -->
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">ชื่อสินค้า <span class="text-red-500">*</span></label>
              <input v-model="editingProduct.name" type="text" class="w-full border border-gray-300 rounded p-2" />
            </div>
            <div class="flex gap-4">
              <div class="flex-grow">
                <label class="block text-sm font-medium text-gray-700 mb-1">ราคา (บาท) <span class="text-red-500">*</span></label>
                <input v-model="editingProduct.price" type="number" class="w-full border border-gray-300 rounded p-2" />
              </div>
              <div class="w-32">
                <label class="block text-sm font-medium text-gray-700 mb-1">หน่วย</label>
                <input v-model="editingProduct.unit" type="text" class="w-full border border-gray-300 rounded p-2" placeholder="เช่น ชิ้น, อัน, กล่อง" />
              </div>
            </div>
          </div>
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-1">รายละเอียดสินค้า</label>
            <textarea v-model="editingProduct.description" rows="4" class="w-full border border-gray-300 rounded p-2"></textarea>
          </div>
          <div class="flex">
            <button @click="updateProduct" class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded mr-2 flex items-center">
              <i class="fas fa-save mr-1"></i>
              บันทึกข้อมูล
            </button>
            <button @click="closeEditModal" class="border border-gray-300 text-gray-700 px-4 py-2 rounded hover:bg-gray-50 flex items-center">
              <i class="fas fa-times mr-1"></i>
              ยกเลิก
            </button>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import Swal from 'sweetalert2'
// Define product interface
interface Product {
  id: number;
  name: string;
  description: string;
  unit: string;
  price: number;
}

// Sample product data
const products = ref<Product[]>([
  {
    id: 1,
    name: 'กาแฟอเมริกาโน่',
    description: 'กาแฟดำรสชาติเข้มข้น',
    unit: 'แก้ว',
    price: 55
  },
  {
    id: 2,
    name: 'ลาเต้',
    description: 'กาแฟผสมนมสดรสนุ่มละมุน',
    unit: 'แก้ว',
    price: 65
  }
]);

// UI state
const isAddingProduct = ref(false);
const showAddModal = ref(false);
const showEditModal = ref(false);


// Pagination state
const currentPage = ref(1);
const itemsPerPage = ref(10);

// Search
const searchTerm = ref('');

// New product form
const newProduct = ref<Omit<Product, 'id'>>({
  name: '',
  description: '',
  unit: '',
  price: 0
});

// Editing product form
const editingProduct = ref<Product>({
  id: 0,
  name: '',
  description: '',
  unit: '',
  price: 0
});

// Computed properties
const filteredProducts = computed(() => {
  if (!searchTerm.value.trim()) {
    return products.value;
  }

  const term = searchTerm.value.toLowerCase();
  return products.value.filter(product => 
    product.name.toLowerCase().includes(term) ||
    product.description.toLowerCase().includes(term) ||
    product.unit.toLowerCase().includes(term)
  );
});

const totalPages = computed(() => {
  return Math.ceil(filteredProducts.value.length / itemsPerPage.value);
});

const paginatedProducts = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value;
  const end = start + itemsPerPage.value;
  return filteredProducts.value.slice(start, end);
});

const startItem = computed(() => {
  return filteredProducts.value.length === 0 
    ? 0 
    : (currentPage.value - 1) * itemsPerPage.value + 1;
});

const endItem = computed(() => {
  return Math.min(currentPage.value * itemsPerPage.value, filteredProducts.value.length);
});

// Methods
function startAddProduct() {
  newProduct.value = {
    name: '',
    description: '',
    unit: '',
    price: 0
  };
  showAddModal.value = true;
}

function closeAddModal() {
  showAddModal.value = false;
}

function cancelAddProduct() {
  isAddingProduct.value = false;
}

function saveProduct() {
  if (!newProduct.value.name.trim()) {
    alert('กรุณากรอกชื่อสินค้า');
    return;
  }

  if (newProduct.value.price <= 0) {
    alert('กรุณากรอกราคาสินค้า');
    return;
  }

  // Generate a new ID
  const newId = products.value.length > 0 
    ? Math.max(...products.value.map(p => p.id)) + 1 
    : 1;

  // Add to list
  products.value.push({
    id: newId,
    ...newProduct.value
  });

  Swal.fire({
    icon: 'success',
    title: 'สำเร็จ!',
    text: 'เพิ่มสินค้าเรียบร้อยแล้ว',
    confirmButtonText: 'ตกลง',
    confirmButtonColor: '#1f9f4c',
  })
  closeAddModal()
}

function editProduct(product: Product) {
  editingProduct.value = { ...product };
  showEditModal.value = true;
}

function closeEditModal() {
  showEditModal.value = false;
}

function updateProduct() {
  if (!editingProduct.value.name.trim()) {
    alert('กรุณากรอกชื่อสินค้า');
    return;
  }

  if (editingProduct.value.price <= 0) {
    alert('กรุณากรอกราคาสินค้า');
    return;
  }

  // Find and update product
  const index = products.value.findIndex(p => p.id === editingProduct.value.id);
  if (index !== -1) {
    products.value[index] = { ...editingProduct.value };
  }

  // Close modal and show success message
  Swal.fire({
    icon: 'success',
    title: 'สำเร็จ!',
    text: 'แก้ไขสินค้าเรียบร้อยแล้ว',
    confirmButtonText: 'ตกลง',
    confirmButtonColor: '#1f9f4c',
  })
  closeEditModal()
}

function deleteProduct(id: number) {
  Swal.fire({
    title: 'คุณแน่ใจหรือไม่?',
    text: 'คุณต้องการลบสินค้านี้ใช่หรือไม่?',
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#cb4040',
    cancelButtonColor: '#3085d6',
    confirmButtonText: 'ลบ',
    cancelButtonText: 'ยกเลิก',
  }).then((result) => {
    if (result.isConfirmed) {
      // ลบสินค้า
      products.value = products.value.filter(p => p.id !== id)

      // แสดงข้อความสำเร็จ
      Swal.fire({
        icon: 'success',
        title: 'สำเร็จ!',
        text: 'ลบสินค้าเรียบร้อยแล้ว',
        confirmButtonText: 'ตกลง',
        confirmButtonColor: '#1f9f4c',
      })
    }
  })
}

function nextPage() {
  if (currentPage.value < totalPages.value) {
    currentPage.value++;
  }
}

function previousPage() {
  if (currentPage.value > 1) {
    currentPage.value--;
  }
}

// Lifecycle hooks
onMounted(() => {
  // Here you would fetch products from your API
});
</script>