<template>
  <el-menu text-color="#000" active-text-color="#10b981" router>
    <el-menu-item v-for="menu in menuList" :key="menu.title" :index="menu.name">
      <el-icon size="20">
        <component :is="menu.icon" />
      </el-icon>
      <span class="menu__title">{{ menu.title }}</span>
    </el-menu-item>
  </el-menu>
</template>

<script setup>
/*
  imports
*/
import { CloudStorage, AllApplication } from '@icon-park/vue-next';

/*
  menu
*/
const menuList = [
  {
    name: '/modules',
    icon: CloudStorage,
    title: 'Modules',
  },
  {
    name: '/stacks',
    icon: AllApplication,
    title: 'Stacks',
  },
];
</script>

<style lang="scss" scoped>
  .el-menu {
    height: 100%;
    position: fixed;
  }

  .el-menu {
    width: 200px;
    border: none;
  }

  .el-menu-item {
    margin-left: 0.4rem;
    margin-top: 0.75rem;
  }

  .menu__title {
    font-weight: 530;
    font-size: 1.2rem;
    margin-left: 0.4rem;
  }
</style>
