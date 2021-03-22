import App from '@/App'
import { createApp } from 'vue/dist/vue.esm-bundler'//'vue'
import VueFeather from 'vue-feather';

const app = createApp(App)
app.component(VueFeather.name, VueFeather);
app.mount('#app')
