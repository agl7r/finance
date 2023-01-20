import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import { Quasar } from 'quasar'
import quasarLang from 'quasar/lang/ru'

import '@quasar/extras/material-icons/material-icons.css'
import 'quasar/src/css/index.sass'

import App from './App.vue'
import routes from './routes.js'

const router = createRouter({
  history: createWebHistory(),
  routes,
})

const app = createApp(App)

app.use(router)

app.use(Quasar, {
  plugins: {},
  lang: quasarLang,
})

app.mount('#app')
