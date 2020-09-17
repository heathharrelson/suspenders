import './util/public-path'
import './assets/css/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import buildRouter from './router'

createApp(App)
  .use(buildRouter())
  .mount('#app')
