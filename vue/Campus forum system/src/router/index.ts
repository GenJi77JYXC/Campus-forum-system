import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    },
    {
      path:'/layout',
      name:'layout',
      component: ()=> import('../layout/index.vue')
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue')
    },
    {
      path: '/user/signin',
      name:'user-signin',
      component: () => import('../views/user/SignIn.vue')
    },
    {
      path: '/user/signup',
      name:'user-signup',
      component: () => import('../views/user/SignUp.vue')
    },
    {
      path: '/layouts',
      name: 'layouts',
      component: () => import('../layouts/admin.vue')
    }
  ]
})

export default router
