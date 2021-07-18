import Vue from 'vue'
import App from './App.vue'
import router from './router'

import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import AFTableColumn from 'af-table-column'
import UmyUi from 'umy-ui'
import 'umy-ui/lib/theme-chalk/index.css';// 引入样式

Vue.use(UmyUi);
Vue.use(AFTableColumn)
Vue.use(ElementUI);

Vue.config.productionTip = false


new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
