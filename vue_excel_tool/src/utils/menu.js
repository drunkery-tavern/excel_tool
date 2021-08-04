import router from "../router";
import store from "../store";
import Layout from "../layout/index.vue";
import home from "../views/home.vue";
import matchInactive from "../views/matchInactive.vue";
import matchExcel from "../views/matchExcel.vue";
import scheduleSplit from "../views/scheduleSplit.vue";

export function generaMenu() {
  let userMenuList = [
    {
      children: [
        {
          component: home,
          icon: "el-icon-s-home",
          name: "首页",
          path: '/index'
        }
      ],
      component: Layout,
      path: '/index'
    },
    {
      children: [
        {
          component: matchInactive,
          icon: "el-icon-s-finance",
          name: "匹配用户",
          path: '/match-inactive'
        }
      ],
      component: Layout,
      path: '/match-inactive'
    },

    {
      children: [
        {
          component: matchExcel,
          icon: "el-icon-s-order",
          name: "表格合并",
          path: '/match-excel'
        }
      ],
      component: Layout,
      path: '/match-excel'
    },

    {
      children: [
        {
          component: scheduleSplit,
          icon: "el-icon-s-management",
          name: "班期拆分",
          path: '/schedule-split'
        }
      ],
      component: Layout,
      path: '/schedule-split'
    },

  ];

  userMenuList.forEach(item => {
    // 添加侧边栏菜单
    store.commit("saveUserMenuList", userMenuList);
    // 添加菜单到路由
    router.addRoutes(userMenuList);
  });
}

function loadView(view) {
  // 路由懒加载
  return (resolve) => require([`@/views${view}`], resolve);
  // return () => import(`@/views${view}`)
}
