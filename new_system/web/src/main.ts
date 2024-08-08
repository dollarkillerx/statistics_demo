import './index.css'
import naive from 'naive-ui'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import VueApexCharts from "vue3-apexcharts";

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(naive)
app.use(VueApexCharts)

app.mount('#app')
