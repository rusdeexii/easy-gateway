// main.ts
import { createApp } from 'vue';
import { createPinia } from 'pinia';
import App from './App.vue';
import router from './router';
import { useAuthStore } from './stores/auth.store';
import './assets/css/main.css';
import VueSweetalert2 from 'vue-sweetalert2';
import 'sweetalert2/dist/sweetalert2.min.css';

const app = createApp(App);
const pinia = createPinia();

app.use(pinia);
app.use(VueSweetalert2);
// ✅ รอให้ authStore.init() เสร็จก่อน แล้วค่อย mount app
const authStore = useAuthStore();
authStore.init().finally(() => {
  app.use(router);
  app.mount('#app');
});

