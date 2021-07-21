import Vue from 'vue'
import App from './App.vue'
import router from './router'
import VueClipboard from 'vue-clipboard2'
import uploader from 'vue-simple-uploader'

import {UTable, UTableColumn,} from 'umy-ui';

import {
  Button,
  Card,
  Form,
  FormItem,
  Icon,
  Input,
  Loading,
  Message,
  Option,
  Select,
  TabPane,
  Tabs,
  Upload
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
Vue.use(Option);


Vue.prototype.$message = Message;
VueClipboard.config.autoSetContainer = true; // add this line
Vue.use(VueClipboard);
Vue.config.productionTip = false;

Vue.use(uploader);

new Vue({
  router,
  render: h => h(App)
}).$mount('#app');
