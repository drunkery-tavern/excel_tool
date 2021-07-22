import router from "../router";
import store from "../store";

export function generaMenu() {
  let userMenuList = [
    {
      children: [
        {
          component: () => import("../views/home.vue"),
          icon: "el-icon-s-home",
          name: "首页",
          path: "/index"
        }
      ],
      component: () => import("../layout/index.vue"),
      path: "/index"
    },
    {
      children: [
        {
          component: () => import("../views/matchInactive.vue"),
          icon: "el-icon-s-finance",
          name: "@文本生成",
          path: "/match-inactive"
        }
      ],
      component: () => import("../layout/index.vue"),
      path: "/match-inactive"
    },

    {
      children: [
        {
          component: () => import("../views/matchExcel.vue"),
          icon: "el-icon-s-order",
          name: "excel匹配合并",
          path: "/match-excel"
        }
      ],
      component: () => import("../layout/index.vue"),
      path: "/match-excel"
    },

  ];

  userMenuList.forEach(item => {
    // 添加侧边栏菜单
    store.commit("saveUserMenuList", userMenuList);
    // 添加菜单到路由
    router.addRoutes(userMenuList);
  });
}

// function loadView(view) {
//   // 路由懒加载
//   return (resolve) => require([`@/views${view}`], resolve);
//   // return () => import(`@/views${view}`)
// }
