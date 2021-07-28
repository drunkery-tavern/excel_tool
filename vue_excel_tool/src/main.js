import Vue from 'vue'
import App from './App.vue'
import router from './router'
import VueClipboard from 'vue-clipboard2'
import uploader from 'vue-simple-uploader'
import store from "./store";
import NProgress from "nprogress";

import {UTable, UTableColumn,} from 'umy-ui';

import {
  Aside,
  Avatar,
  Breadcrumb,
  BreadcrumbItem,
  Button,
  Card,
  Checkbox,
  Container,
  Dropdown,
  DropdownItem,
  DropdownMenu,
  Form,
  FormItem,
  Header,
  Icon,
  Input,
  Loading,
  Main,
  Menu,
  MenuItem,
  Message,
  Option,
  Scrollbar,
  Select,
  TabPane,
  Tabs,
  Upload,
  Carousel, CarouselItem, Image, Row, Col, Tag, Steps, Step, Progress, Empty
} from 'element-ui';
import {generaMenu} from "./utils/menu";

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


Vue.prototype.$message = Message;
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
