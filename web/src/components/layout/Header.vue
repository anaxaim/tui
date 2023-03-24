<template>
  <header class="header">
    <div class="header__left">
      <img src="@/assets/images/tui.png" alt="Logo" class="header__icon">
      <span class="header__title">Tui</span>
    </div>
    <div class="header__right">
      <a :href="githubInfo.project" target="_blank">
        <github-one size="32" :fill="['#333']" />
      </a>
      <Logout size="32" @click="logout" class="header__logout"/>
    </div>
  </header>
</template>

<script setup>
/*
  imports
*/
  import { GithubOne, Logout } from '@icon-park/vue-next';
  import { githubInfo } from '@/config.js';
  import { getUser, delUser } from '@/utils';
  import { useRouter } from 'vue-router'
  import { ElNotification } from "element-plus"
  import request from "@/axios";

  const user = getUser();
  const router = useRouter()

/*
  logout
*/
  function logout() {
    request.delete('/api/v1/auth/token').then(() => {
      console.log("logout success")
      ElNotification.success({
        title: 'Logout Success',
        message: 'Bye~ ' + user.name,
        showClose: true,
        duration: 1500,
      })
      delUser();
      router.push('/login');
    }).catch((error) => {
      console.log(error)
    })
  }
</script>

<style lang="scss" scoped>
  .header {
    background: #fff;
    padding: 0 20px;
    height: 65px;
    display: flex;
    align-items: center;
    position: relative;
    border-bottom: 1px solid #d1d5db;
  }

  .header__left {
    display: flex;
    align-items: center;
    margin-left: 0.1rem;
  }

  .header__icon {
    width: 3rem;
    margin-right: 0.5rem;
  }

  .header__title {
    font-weight: 500;
    font-size: 1.75rem;
    margin: 0;
  }

  .header__right {
    display: flex;
    align-items: center;
    margin-left: auto;
  }

  .header__logout {
    margin-left: 5px;
    cursor: pointer;
  }
</style>
