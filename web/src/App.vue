<template>
  <div v-if="route.path !== '/login'">
    <Header/>
    <div class="menu"><Menu/></div>
  </div>
  <div class="menu__items"><RouterView/></div>
</template>

<script setup>
/*
  imports
*/
  import Header from '@/components/layout/Header.vue';
  import Menu from '@/components/layout/Menu.vue';
  import { RouterView, useRoute } from 'vue-router'

  const route = useRoute();
</script>

<style lang="scss">
  #app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  .menu {
    display: flex;
    flex-direction: row;
    flex-grow: 0.8;
    overflow: hidden;
  }

  .menu__items {
    overflow: scroll;
    display: flex;
    margin-left: 230px;
  }
</style>
