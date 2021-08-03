import Vue from 'vue'
import App from './App.vue'
import router from './router'
import VueClipboard from 'vue-clipboard2'
import uploader from 'vue-simple-uploader'
import store from "./store";
import NProgress from "nprogress";

import {deleteRequest, getRequest, postKeyValueRequest, postRequest, putRequest} from "./utils/api";

import {
  Aside,
  Avatar,
  Breadcrumb,
  BreadcrumbItem,
  Button,
  Card,
  Carousel,
  CarouselItem,
  Checkbox,
  Col,
  Container, Divider,
  Dropdown,
  DropdownItem,
  DropdownMenu,
  Empty,
  Form,
  FormItem,
  Header,
  Icon,
  Image,
  Input,
  Loading,
  Main,
  Menu,
  MenuItem,
  Message,
  MessageBox,
  Option,
  Progress, Radio, RadioGroup,
  Row,
  Scrollbar,
  Select,
  Step,
  Steps,
  TabPane,
  Tabs,
  Tag,
  Upload
} from 'element-ui';
import {generaMenu} from "./utils/menu";
import axios from "axios";

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
Vue.use(Checkbox);
Vue.use(Container);
Vue.use(Aside);
Vue.use(DropdownItem);
Vue.use(DropdownMenu);
Vue.use(Breadcrumb);
Vue.use(BreadcrumbItem);
Vue.use(Dropdown);
Vue.use(Menu);
Vue.use(MenuItem);
Vue.use(Header);
Vue.use(Avatar);
Vue.use(Main);
Vue.use(Scrollbar);
Vue.use(Carousel);
Vue.use(CarouselItem);
Vue.use(Image);
Vue.use(Row);
Vue.use(Col);
Vue.use(Tag);
Vue.use(Steps);
Vue.use(Step);
Vue.use(Progress);
Vue.use(Empty);
Vue.use(Radio);
Vue.use(RadioGroup);
Vue.use(Divider);


Vue.prototype.axios = axios;
Vue.prototype.postRequest = postRequest;
Vue.prototype.postKeyValueRequest = postKeyValueRequest;
Vue.prototype.putRequest = putRequest;
Vue.prototype.deleteRequest = deleteRequest;
Vue.prototype.getRequest = getRequest;

Vue.prototype.$message = Message;
Vue.prototype.$confirm = MessageBox.confirm;
VueClipboard.config.autoSetContainer = true; // add this line
Vue.use(VueClipboard);
Vue.config.productionTip = false;

Vue.use(uploader);


NProgress.configure({
  easing: "ease", // 动画方式
  speed: 500, // 递增进度条的速度
  showSpinner: true, // 是否显示加载ico
  trickleSpeed: 200, // 自动递增间隔
  minimum: 0.3 // 初始化时的最小百分比
});

//路由守卫
router.beforeEach((to, from, next) => {
  NProgress.start();
  if (to.path === "/") {
    next();
  } else if (!store.state.username) {
    next({path: "/"});
  } else {
    next();
  }
});

router.afterEach(() => {
  NProgress.done();
});

new Vue({
  router,
  store,
  render: h => h(App),
  created() {
    // 刷新页面查询用户菜单
    if (store.state.username != null) {
      generaMenu();
    }
  }
}).$mount('#app');
