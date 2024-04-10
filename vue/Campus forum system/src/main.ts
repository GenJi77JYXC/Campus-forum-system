

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router'
import 'element-plus/es/components/message/style/css'
import { ElMessage } from 'element-plus'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(ElementPlus)

app.mount('#app')
