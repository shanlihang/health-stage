import { createApp } from 'vue'
import App from './App.vue'
import router from './router/router'
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/reset.css';
import BinDatav from 'bin-datav'
import 'bin-datav/lib/styles/index.css'

createApp(App).use(router).use(Antd).use(BinDatav).mount('#app')