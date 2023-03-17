import { createRouter, createWebHistory } from 'vue-router'
import { getUser } from '@/utils'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import("views/Home.vue"),
    redirect: '/modules',
    children: [
      {
        path: '/modules',
        name: 'Modules',
        component: () => import("views/Modules.vue")
      }
    ]
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import("views/auth/Login.vue")
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  let isAuthenticated = false;
  let user = getUser();
  if (user && user.name) {
    isAuthenticated = true;
  }

  if (!isAuthenticated && to.name !== 'Login') next({ name: 'Login' })
  else next()
})

export default router
