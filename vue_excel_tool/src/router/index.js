import Vue from 'vue'
import VueRouter from 'vue-router'
import login from "../views/login";

Vue.use(VueRouter);

const routes = [
    {
        path: "/",
        name: "登录",
        hidden: true,
        component: login
    },
];

const createRouter = () =>
    new VueRouter({
        mode: 'history',
        routes: routes
    });

const router = createRouter();

export function resetRouter() {
  const newRouter = createRouter();
  router.matcher = newRouter.matcher;
}

const originalPush = router.push;
router.push = function push(location) {
    return originalPush.call(this, location).catch(err => err)
};

export default router
