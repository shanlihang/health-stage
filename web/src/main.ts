import { createApp } from 'vue'
import App from './App.vue'
import router from './router/router'
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/reset.css';
import BinDatav from 'bin-datav'
import 'bin-datav/lib/styles/index.css'
import {createPinia} from 'pinia'

const pinia = createPinia()
const app = createApp(App)

app.use(router)
app.use(Antd)
app.use(BinDatav)
app.use(pinia)

app.mount('#app')