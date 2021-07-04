import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/login.vue'
import Home from '../views/home.vue'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'Login',
    component: Login
  },
    {
      path: '/home',
      name: 'Home',
      component: Home,
    }

]

const router = new VueRouter({
  routes
})

export default router
