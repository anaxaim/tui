import { createApp } from 'vue';

/* element-plus */
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';

import App from './App.vue';
import router from './router';
import './assets/styles/global.scss';

const app = createApp(App);
app
  .use(ElementPlus)
  .use(router)
  .mount('#app');
