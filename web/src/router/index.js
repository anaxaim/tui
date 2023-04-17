import { createRouter, createWebHistory } from 'vue-router';
import Login from 'views/Login.vue';
import Modules from 'views/Modules.vue';
import Credentials from 'views/Credentials.vue';
import { getUser } from '@/utils';

const routes = [
  {
    path: '/',
    redirect: '/login',
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
  {
    path: '/modules',
    name: 'Modules',
    component: Modules,
  },
  {
    path: '/credentials',
    name: 'Credentials',
    component: Credentials,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  let isAuthenticated = false;
  const user = getUser();
  if (user && user.name) {
    isAuthenticated = true;
  }

  if (!isAuthenticated && to.name !== 'Login') {
    next({ name: 'Login' });
  } else if (isAuthenticated && to.name === 'Login') {
    next({ name: 'Modules' });
  } else {
    next();
  }
});

export default router;
