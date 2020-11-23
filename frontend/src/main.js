import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import 'uikit/dist/css/uikit.min.css'
import 'uikit/dist/js/uikit.min.js'
import 'uikit/dist/js/uikit-icons.min.js'
// There has to be another way
import UIkit from 'uikit';
import Icons from 'uikit/dist/js/uikit-icons';
UIkit.use(Icons);

createApp(App).use(store).use(router).mount('#app');
