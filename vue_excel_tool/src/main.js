import Vue from 'vue'
import App from './App.vue'
import router from './router'

// import 'element-ui/lib/theme-chalk/index.css';
import VueClipboard from 'vue-clipboard2'

import {
  UTableColumn,
  UTable,
} from 'umy-ui';

import {
  Button,
  Input,
  Card,
  Icon,
  Select,
  Form,
  Loading,
  Message,
  Tabs,
  TabPane,
  FormItem,
  Upload,
} from 'element-ui';

Vue.use(UTableColumn);
Vue.use(UTable);

Vue.use(Button);
Vue.use(Input);
Vue.use(Card);
Vue.use(Icon);
Vue.use(Select);
Vue.use(Form);
Vue.use(Loading);
Vue.use(Tabs);
Vue.use(TabPane);
Vue.use(FormItem);
Vue.use(Upload);


Vue.prototype.$message = Message;
VueClipboard.config.autoSetContainer = true; // add this line
Vue.use(VueClipboard);
Vue.config.productionTip = false;

new Vue({
  router,
  render: h => h(App)
}).$mount('#app');
