<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { BarChart, LineChart, PieChart } from 'echarts/charts';
import { RadarChart } from 'echarts/charts';
import { 
  TitleComponent, 
  TooltipComponent, 
  LegendComponent, 
  GridComponent,
  DatasetComponent,
  TransformComponent
} from 'echarts/components';
import { LabelLayout, UniversalTransition } from 'echarts/features';
import { CanvasRenderer } from 'echarts/renderers';
import { use } from 'echarts/core';
import VChart from 'vue-echarts';

// Register ECharts components
use([
  CanvasRenderer,
  BarChart,
  LineChart,
  PieChart,
  RadarChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent,
  DatasetComponent,
  TransformComponent,
  LabelLayout,
  UniversalTransition
]);

// Types
interface Transaction {
  bank: string;
  transaction_id: string;
  amount: number;
  timestamp: string;
}

interface BankSummary {
  bank: string;
  count: number;
  totalAmount: number;
  color: string;
}

// Bank colors
const BANK_COLORS = {
  SCB: '#4F2D7F',
  KBANK: '#138F2D',
  TTB: '#0A5EAA',
};

// Reactive state
const transactions = ref<Transaction[]>([]);
const loading = ref(true);
const dateRange = ref('today');
const searchTerm = ref('');
const selectedBank = ref('all');
const bankSummaries = ref<BankSummary[]>([]);
const dailyData = ref<any[]>([]);
const stats = ref({
  totalTransactions: 0,
  totalAmount: 0,
  avgAmount: 0,
  growth: 12.5
});
const darkMode = ref(false);
const showTransactionDetails = ref(false);
const selectedTransaction = ref<Transaction | null>(null);
const chartViewMode = ref('bar');

// Computed properties
const filteredTransactions = computed(() => {
  return transactions.value.filter(txn => {
    const matchesSearch = searchTerm.value === '' || 
      txn.transaction_id.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
      txn.amount.toString().includes(searchTerm.value);
    
    const matchesBank = selectedBank.value === 'all' || txn.bank === selectedBank.value;
    
    return matchesSearch && matchesBank;
  });
});

// Chart options
const barChartOption = computed(() => ({
  tooltip: {
    trigger: 'axis',
    formatter: (params: any) => {
      return `${params[0].name}: ฿${Number(params[0].value).toLocaleString('th-TH', { 
        minimumFractionDigits: 2, 
        maximumFractionDigits: 2 
      })}`;
    }
  },
  xAxis: {
    type: 'category',
    data: dailyData.value.map(item => item.name),
    axisLabel: {
      color: darkMode.value ? '#e5e7eb' : '#1f2937'
    }
  },
  yAxis: {
    type: 'value',
    axisLabel: {
      formatter: (value: number) => {
        return `฿${value.toLocaleString('th-TH', { 
          minimumFractionDigits: 0, 
          maximumFractionDigits: 0 
        })}`;
      },
      color: darkMode.value ? '#e5e7eb' : '#1f2937'
    }
  },
  series: [
    {
      data: dailyData.value.map(item => item.amount),
      type: 'bar',
      itemStyle: {
        color: function(params: any) {
          // Generate gradient color based on value
          const value = params.value;
          const maxValue = Math.max(...dailyData.value.map(item => item.amount));
          const colorStop = Math.max(0.3, value / maxValue);
          return {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [{
              offset: 0, 
              color: darkMode.value ? '#1d4ed8' : '#3b82f6'
            }, {
              offset: 1, 
              color: darkMode.value ? '#1e40af' : '#93c5fd'
            }],
          };
        }
      },
      emphasis: {
        itemStyle: {
          color: darkMode.value ? '#2563eb' : '#60a5fa'
        }
      },
      barWidth: '60%',
      showBackground: true,
      backgroundStyle: {
        color: darkMode.value ? '#374151' : '#f3f4f6'
      },
      label: {
        show: true,
        position: 'top',
        formatter: (params: any) => {
          return `฿${Number(params.value).toLocaleString('th-TH', { 
            minimumFractionDigits: 0, 
            maximumFractionDigits: 0 
          })}`;
        },
        color: darkMode.value ? '#e5e7eb' : '#1f2937'
      },
      animation: true
    }
  ]
}));

const lineChartOption = computed(() => ({
  tooltip: {
    trigger: 'axis',
    formatter: (params: any) => {
      return `${params[0].name}: ฿${Number(params[0].value).toLocaleString('th-TH', { 
        minimumFractionDigits: 2, 
        maximumFractionDigits: 2 
      })}`;
    }
  },
  xAxis: {
    type: 'category',
    data: dailyData.value.map(item => item.name),
    boundaryGap: false,
    axisLabel: {
      color: darkMode.value ? '#e5e7eb' : '#1f2937'
    }
  },
  yAxis: {
    type: 'value',
    axisLabel: {
      formatter: (value: number) => {
        return `฿${value.toLocaleString('th-TH', { 
          minimumFractionDigits: 0, 
          maximumFractionDigits: 0 
        })}`;
      },
      color: darkMode.value ? '#e5e7eb' : '#1f2937'
    }
  },
  series: [
    {
      data: dailyData.value.map(item => item.amount),
      type: 'line',
      smooth: true,
      symbol: 'circle',
      symbolSize: 8,
      itemStyle: {
        color: darkMode.value ? '#3b82f6' : '#2563eb'
      },
      lineStyle: {
        width: 4,
        shadowColor: darkMode.value ? 'rgba(59, 130, 246, 0.5)' : 'rgba(37, 99, 235, 0.5)',
        shadowBlur: 10,
        shadowOffsetY: 5
      },
      areaStyle: {
        color: {
          type: 'linear',
          x: 0,
          y: 0,
          x2: 0,
          y2: 1,
          colorStops: [{
            offset: 0,
            color: darkMode.value ? 'rgba(59, 130, 246, 0.8)' : 'rgba(59, 130, 246, 0.6)'
          }, {
            offset: 1,
            color: darkMode.value ? 'rgba(59, 130, 246, 0.1)' : 'rgba(59, 130, 246, 0.1)'
          }]
        }
      },
      animation: true
    }
  ]
}));

const pieChartOption = computed(() => ({
  tooltip: {
    trigger: 'item',
    formatter: (params: any) => {
      return `${params.name}: ฿${Number(params.value).toLocaleString('th-TH', { 
        minimumFractionDigits: 2, 
        maximumFractionDigits: 2 
      })} (${params.percent}%)`;
    }
  },
  legend: {
    top: '5%',
    left: 'center',
    textStyle: {
      color: darkMode.value ? '#e5e7eb' : '#1f2937'
    }
  },
  series: [
    {
      name: 'ยอดรวม',
      type: 'pie',
      radius: ['40%', '70%'],
      avoidLabelOverlap: false,
      itemStyle: {
        borderRadius: 10,
        borderColor: darkMode.value ? '#1f2937' : '#ffffff',
        borderWidth: 2
      },
      label: {
        show: false,
        position: 'center'
      },
      emphasis: {
        label: {
          show: true,
          fontSize: '18',
          fontWeight: 'bold',
          formatter: (params: any) => {
            return `${params.name}\n฿${Number(params.value).toLocaleString('th-TH', { 
              minimumFractionDigits: 2, 
              maximumFractionDigits: 2 
            })}`;
          },
          color: darkMode.value ? '#e5e7eb' : '#1f2937'
        },
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      },
      labelLine: {
        show: false
      },
      data: bankSummaries.value.map(bank => ({
        value: bank.totalAmount,
        name: bank.bank,
        itemStyle: {
          color: bank.color
        }
      }))
    }
  ]
}));

// Methods
const fetchData = async () => {
  loading.value = true;
  try {
    // Mock data
    const mockTransactions: Transaction[] = [
      {
        bank: "SCB",
        transaction_id: "20ac5631-1fea-42a4-8f39-8b594740d31d",
        amount: 100.50,
        timestamp: "2025-04-27T15:30:00Z"
      },
      {
        bank: "SCB",
        transaction_id: "e49cd8ae-5148-496d-b630-01b1c10312ee",
        amount: 98.11,
        timestamp: "2025-04-27T14:20:00Z"
      },
      {
        bank: "TTB",
        transaction_id: "2639e588-9c4b-46ce-b158-e9549d334490",
        amount: 3500.00,
        timestamp: "2025-04-24T01:30:00Z"
      },
      {
        bank: "KBANK",
        transaction_id: "5b4ac8bd-11bd-4c85-8499-05a612a2cd3a",
        amount: 5000.00,
        timestamp: "2025-04-23T22:22:00Z"
      },
      {
        bank: "SCB",
        transaction_id: "f1a5d8e9-c1f2-4a3b-b4c5-d6e7f8a9b0c1",
        amount: 750.25,
        timestamp: "2025-04-27T10:15:00Z"
      },
      {
        bank: "KBANK",
        transaction_id: "b2c3d4e5-f6g7-8h9i-j0k1-l2m3n4o5p6q7",
        amount: 1250.75,
        timestamp: "2025-04-26T09:45:00Z"
      },
      {
        bank: "TTB",
        transaction_id: "r8s9t0u1-v2w3-x4y5-z6a7-b8c9d0e1f2g3",
        amount: 2100.30,
        timestamp: "2025-04-25T16:20:00Z"
      }
    ];

    transactions.value = mockTransactions;

    // Calculate summary by bank
    const bankData: Record<string, { count: number; totalAmount: number; }> = {};
    mockTransactions.forEach(txn => {
      if (!bankData[txn.bank]) {
        bankData[txn.bank] = { count: 0, totalAmount: 0 };
      }
      bankData[txn.bank].count += 1;
      bankData[txn.bank].totalAmount += txn.amount;
    });

    // Map to summaries
    const summaries = Object.keys(bankData).map((bank, index) => ({
      bank,
      count: bankData[bank].count,
      totalAmount: bankData[bank].totalAmount,
      color: (BANK_COLORS as any)[bank] || ['#0088FE', '#00C49F', '#FFBB28', '#FF8042'][index % 4]
    }));
    
    bankSummaries.value = summaries;

    // Calculate daily data
    const days = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];
    const dailyTxns = Array(7).fill(0).map((_, i) => ({
      name: days[i],
      amount: Math.random() * 10000 + 1000
    }));
    dailyData.value = dailyTxns;

    // Calculate overall stats
    const totalTxns = mockTransactions.length;
    const totalAmt = mockTransactions.reduce((sum, txn) => sum + txn.amount, 0);
    stats.value = {
      totalTransactions: totalTxns,
      totalAmount: totalAmt,
      avgAmount: totalAmt / totalTxns,
      growth: 12.5
    };

  } catch (error) {
    console.error('Error fetching transactions:', error);
  } finally {
    loading.value = false;
  }
};

const formatDate = (dateString: string) => {
  const date = new Date(dateString);
  return new Intl.DateTimeFormat('th-TH', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date);
};

const formatCurrency = (amount: number) => {
  return amount.toLocaleString('th-TH', { 
    minimumFractionDigits: 2, 
    maximumFractionDigits: 2 
  });
};

const toggleDarkMode = () => {
  darkMode.value = !darkMode.value;
};

const openTransactionDetails = (txn: Transaction) => {
  selectedTransaction.value = txn;
  showTransactionDetails.value = true;
};

const closeTransactionDetails = () => {
  showTransactionDetails.value = false;
  setTimeout(() => {
    selectedTransaction.value = null;
  }, 300);
};

const refreshData = () => {
  fetchData();
};

// Lifecycle hooks
onMounted(() => {
  fetchData();
});
</script>

<template>
  <section :class="[
    'min-h-screen transition-colors duration-300',
    
  ]">
    <div class="justify-between items-center p-4">

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <!-- Total Transactions Card -->
        <div :class="[
          'rounded-lg shadow p-6 transition-all duration-300 transform hover:scale-105',
          darkMode ? 'bg-gray-800' : 'bg-white'
        ]">
          <div class="flex items-center">
            <div :class="[
              'p-3 rounded-full',
              darkMode ? 'bg-blue-900 text-blue-300' : 'bg-blue-100 text-blue-600'
            ]">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
              </svg>
            </div>
            <div class="ml-4">
              <p :class="[
                'text-sm font-medium',
                'text-gray-500'
              ]">รายการทั้งหมด</p>
              <p class="text-2xl font-semibold">{{ stats.totalTransactions.toLocaleString() }}</p>
            </div>
          </div>
          <div class="flex items-center mt-4 text-sm">
            <span class="text-green-500 flex items-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M12 7a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0V8.414l-4.293 4.293a1 1 0 01-1.414 0L8 10.414l-4.293 4.293a1 1 0 01-1.414-1.414l5-5a1 1 0 011.414 0L11 10.586 14.586 7H12z" clip-rule="evenodd" />
              </svg>
              8.2%
            </span>
            <span :class="darkMode ? 'text-gray-400' : 'text-gray-500'" class="ml-2">จากสัปดาห์ที่แล้ว</span>
          </div>
        </div>

        <!-- Total Amount Card -->
        <div :class="[
          'rounded-lg shadow p-6 transition-all duration-300 transform hover:scale-105',
          darkMode ? 'bg-gray-800' : 'bg-white'
        ]">
          <div class="flex items-center">
            <div :class="[
              'p-3 rounded-full',
              darkMode ? 'bg-green-900 text-green-300' : 'bg-green-100 text-green-600'
            ]">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <div class="ml-4">
              <p :class="[
                'text-sm font-medium',
                darkMode ? 'text-gray-400' : 'text-gray-500'
              ]">ยอดรวมทั้งหมด</p>
              <p class="text-2xl font-semibold">฿{{ formatCurrency(stats.totalAmount) }}</p>
            </div>
          </div>
          <div class="flex items-center mt-4 text-sm">
            <span class="text-green-500 flex items-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M12 7a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0V8.414l-4.293 4.293a1 1 0 01-1.414 0L8 10.414l-4.293 4.293a1 1 0 01-1.414-1.414l5-5a1 1 0 011.414 0L11 10.586 14.586 7H12z" clip-rule="evenodd" />
              </svg>
              12.5%
            </span>
            <span :class="darkMode ? 'text-gray-400' : 'text-gray-500'" class="ml-2">จากสัปดาห์ที่แล้ว</span>
          </div>
        </div>

        <!-- Average Amount Card -->
        <div :class="[
          'rounded-lg shadow p-6 transition-all duration-300 transform hover:scale-105',
          darkMode ? 'bg-gray-800' : 'bg-white'
        ]">
          <div class="flex items-center">
            <div :class="[
              'p-3 rounded-full',
              darkMode ? 'bg-yellow-900 text-yellow-300' : 'bg-yellow-100 text-yellow-600'
            ]">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 8v8m-4-5v5m-4-2v2m-2 4h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
            </div>
            <div class="ml-4">
              <p :class="[
                'text-sm font-medium',
                darkMode ? 'text-gray-400' : 'text-gray-500'
              ]">ยอดเฉลี่ยต่อรายการ</p>
              <p class="text-2xl font-semibold">฿{{ formatCurrency(stats.avgAmount) }}</p>
            </div>
          </div>
          <div class="flex items-center mt-4 text-sm">
            <span class="text-red-500 flex items-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M12 13a1 1 0 100 2h5a1 1 0 001-1V9a1 1 0 10-2 0v2.586l-4.293-4.293a1 1 0 00-1.414 0L8 9.586 3.707 5.293a1 1 0 00-1.414 1.414l5 5a1 1 0 001.414 0L11 9.414 14.586 13H12z" clip-rule="evenodd" />
              </svg>
              3.2%
            </span>
            <span :class="darkMode ? 'text-gray-400' : 'text-gray-500'" class="ml-2">จากสัปดาห์ที่แล้ว</span>
          </div>
        </div>

        <!-- Growth Card -->
        <div :class="[
          'rounded-lg shadow p-6 transition-all duration-300 transform hover:scale-105',
          darkMode ? 'bg-gray-800' : 'bg-white'
        ]">
          <div class="flex items-center">
            <div :class="[
              'p-3 rounded-full',
              darkMode ? 'bg-purple-900 text-purple-300' : 'bg-purple-100 text-purple-600'
            ]">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
              </svg>
            </div>
            <div class="ml-4">
              <p :class="[
                'text-sm font-medium',
                darkMode ? 'text-gray-400' : 'text-gray-500'
              ]">การเติบโต</p>
              <p class="text-2xl font-semibold">{{ stats.growth.toFixed(1) }}%</p>
            </div>
          </div>
          <div class="flex items-center mt-4 text-sm">
            <span class="text-green-500 flex items-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M12 7a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0V8.414l-4.293 4.293a1 1 0 01-1.414 0L8 10.414l-4.293 4.293a1 1 0 01-1.414-1.414l5-5a1 1 0 011.414 0L11 10.586 14.586 7H12z" clip-rule="evenodd" />
              </svg>
              2.4%
            </span>
            <span :class="darkMode ? 'text-gray-400' : 'text-gray-500'" class="ml-2">จากเดือนที่แล้ว</span>
          </div>
        </div>
      </div>

      <!-- Charts Section -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-8">
        <!-- Daily Transactions Chart -->
        <div :class="[
          'rounded-lg shadow p-6 lg:col-span-2 transition-all duration-300',
          darkMode ? 'bg-gray-800' : 'bg-white'
        ]">
          <div class="flex flex-col sm:flex-row sm:items-center justify-between mb-6">
            <h2 class="text-lg font-semibold mb-2 sm:mb-0">รายการรายวัน</h2>
            <div class="flex flex-wrap gap-2">
              <div class="flex space-x-2">
                <select 
                  v-model="dateRange"
                  :class="[
                    'text-sm rounded-md shadow-sm focus:ring-2 focus:outline-none',
                    darkMode ? 'bg-gray-700 border-gray-600 text-white focus:border-blue-500 focus:ring-blue-500' : 'border-gray-300 focus:border-blue-300 focus:ring-blue-200'
                  ]"
                >
                  <option value="today">วันนี้</option>
                  <option value="week">สัปดาห์นี้</option>
                  <option value="month">เดือนนี้</option>
                </select>
                <div class="flex">
                  <button
                    @click="chartViewMode = 'bar'"
                    :class="[
                      'px-3 py-1 text-sm rounded-l-md focus:outline-none transition-colors',
                      chartViewMode === 'bar' ? 
                        (darkMode ? 'bg-blue-600 text-white' : 'bg-blue-500 text-white') : 
                        (darkMode ? 'bg-gray-700 text-gray-300 hover:bg-gray-600' : 'bg-gray-200 text-gray-700 hover:bg-gray-300')
                    ]"
                  >
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 13v-1m4 1v-3m4 3V8M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z" />
                    </svg>
                  </button>
                  <button
                    @click="chartViewMode = 'line'"
                    :class="[
                      'px-3 py-1 text-sm rounded-r-md focus:outline-none transition-colors',
                      chartViewMode === 'line' ? 
                        (darkMode ? 'bg-blue-600 text-white' : 'bg-blue-500 text-white') : 
                        (darkMode ? 'bg-gray-700 text-gray-300 hover:bg-gray-600' : 'bg-gray-200 text-gray-700 hover:bg-gray-300')
                    ]"
                  >
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12l3-3 3 3 4-4M8 21l4-4 4 4M3 4h18M4 4h16v12a1 1 0 01-1 1H5a1 1 0 01-1-1V4z" />
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>
          <div>
            <div class="h-80">
              <v-chart v-if="chartViewMode === 'bar'" :option="barChartOption" autoresize />
              <v-chart v-else :option="lineChartOption" autoresize />
            </div>
          </div>
        </div>

        <!-- Bank Distribution Chart -->
        <div :class="[
          'rounded-lg shadow p-6 transition-all duration-300',
          darkMode ? 'bg-gray-800' : 'bg-white'
        ]">
          <h2 class="text-lg font-semibold mb-6">สัดส่วนตามธนาคาร</h2>
          <div class="h-80">
            <v-chart :option="pieChartOption" autoresize />
          </div>
        </div>
      </div>

      <!-- Transaction Table -->
      <div :class="[
        'rounded-lg shadow p-6 transition-all duration-300',
        darkMode ? 'bg-gray-800' : 'bg-white'
      ]">
        <div class="flex flex-col sm:flex-row sm:items-center justify-between mb-6">
          <h2 class="text-lg font-semibold mb-2 sm:mb-0">รายการธุรกรรมล่าสุด</h2>
          <div class="flex flex-wrap gap-2">
            <div class="relative">
              <input
                v-model="searchTerm"
                type="text"
                :class="[
                  'pl-10 pr-4 py-2 rounded-md shadow-sm text-sm focus:ring-2 focus:outline-none w-full',
                  darkMode ? 'bg-gray-700 border-gray-600 text-white focus:border-blue-500 focus:ring-blue-500' : 'border-gray-300 focus:border-blue-300 focus:ring-blue-200'
                ]"
                placeholder="ค้นหารายการ..."
              >
              <div class="absolute left-3 top-2.5 text-gray-400">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
              </div>
            </div>
            <select
              v-model="selectedBank"
              :class="[
                'text-sm rounded-md shadow-sm focus:ring-2 focus:outline-none',
                darkMode ? 'bg-gray-700 border-gray-600 text-white focus:border-blue-500 focus:ring-blue-500' : 'border-gray-300 focus:border-blue-300 focus:ring-blue-200'
              ]"
            >
              <option value="all">ทั้งหมด</option>
              <option value="SCB">SCB</option>
              <option value="KBANK">KBANK</option>
              <option value="TTB">TTB</option>
            </select>
          </div>
        </div>

        <div class="overflow-x-auto">
          <table :class="[
              'min-w-full divide-y',
              darkMode ? 'divide-gray-700' : 'divide-gray-200'
          ]">
            <thead :class="darkMode ? 'bg-gray-700' : 'bg-gray-100'">
              <tr>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider" :class="darkMode ? 'text-gray-300' : 'text-gray-500'">ธนาคาร</th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider" :class="darkMode ? 'text-gray-300' : 'text-gray-500'">เลขที่รายการ</th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider" :class="darkMode ? 'text-gray-300' : 'text-gray-500'">จำนวนเงิน</th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider" :class="darkMode ? 'text-gray-300' : 'text-gray-500'">วันที่และเวลา</th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider" :class="darkMode ? 'text-gray-300' : 'text-gray-500'">สถานะ</th>
              </tr>
            </thead>
            <tbody :class="'divide-y divide-gray-200'">
              <tr 
                v-for="txn in filteredTransactions" 
                :key="txn.transaction_id"
                :class="[
                  'transition-colors cursor-pointer hover:bg-opacity-10',
                  darkMode ? 'hover:bg-blue-900' : 'hover:bg-blue-100'
                ]"
                @click="openTransactionDetails(txn)"
              >
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div :class="[
                      'flex-shrink-0 h-8 w-8 rounded-full flex items-center justify-center',
                      {
                        'bg-purple-100 text-purple-800': txn.bank === 'SCB',
                        'bg-green-100 text-green-800': txn.bank === 'KBANK',
                        'bg-blue-100 text-blue-800': txn.bank === 'TTB'
                      }
                    ]">
                      <span class="font-medium">{{ txn.bank.substring(0, 1) }}</span>
                    </div>
                    <div class="ml-4">
                      <div :class="[
                        'font-medium',
                        darkMode ? 'text-white' : 'text-gray-900'
                      ]">{{ txn.bank }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div :class="darkMode ? 'text-gray-200' : 'text-gray-700'">{{ txn.transaction_id.substring(0, 8) }}...</div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div :class="[
                    'font-medium',
                    txn.amount > 1000 ? 'text-green-600' : 'text-blue-600'
                  ]">฿{{ formatCurrency(txn.amount) }}</div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div :class="darkMode ? 'text-gray-200' : 'text-gray-700'">{{ formatDate(txn.timestamp) }}</div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap  text-sm font-medium">
                  <button 
                    @click.stop="openTransactionDetails(txn)"
                    :class="[
                      'inline-flex items-center px-3 py-1 border border-transparent rounded-md shadow-sm text-xs font-medium focus:outline-none focus:ring-2 focus:ring-offset-2',
                      darkMode ? 
                        'bg-green-600 hover:bg-green-700 text-white focus:ring-green-500' : 
                        'bg-green-500 hover:bg-green-600 text-white focus:ring-green-400'
                    ]"
                  >
                    สำเร็จ <i class="fas fa-check ml-1"></i>
                  </button>
                </td>
              </tr>
              <tr v-if="filteredTransactions.length === 0">
                <td :class="darkMode ? 'text-gray-400' : 'text-gray-500'" class="px-6 py-10 text-center" colspan="5">
                  <div class="flex flex-col items-center">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-10 w-10 mb-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    <p class="text-lg font-medium">ไม่พบรายการที่ค้นหา</p>
                    <p class="mt-1">ลองเปลี่ยนเงื่อนไขการค้นหาหรือลองอีกครั้งในภายหลัง</p>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="mt-4 flex justify-between items-center">
          <div :class="darkMode ? 'text-gray-400' : 'text-gray-500'" class="text-sm">
            แสดง {{ filteredTransactions.length }} จาก {{ transactions.length }} รายการ
          </div>
         
        </div>
      </div>
    </div>
  </section>
</template>