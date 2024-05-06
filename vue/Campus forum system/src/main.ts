

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router'
import 'element-plus/es/components/message/style/css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import '@/assets/styles/main.scss'

import mitt from 'mitt'

const app = createApp(App)


app.config.globalProperties.emitter = mitt()
app.use(createPinia())
app.use(router)
app.use(ElementPlus)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.mount('#app')
