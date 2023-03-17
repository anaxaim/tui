<template>
  <div class="flex flex-col h-full">
    <el-menu class="w-[14rem] py-[1rem] flex-grow overflow-x-hidden overflow-y-auto"
      text-color="#4b5563" active-text-color="#10b981" unique-opened :collapse="isCollapse" router>

      <el-menu-item :index="menu.name" v-for="menu in menuList" :key="menu.title">
        <el-icon size="14">
          <component :is="menu.icon" />
        </el-icon>
        <span class="font-bold">{{ menu.title }}</span>
      </el-menu-item>

    </el-menu>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { ApplicationTwo, CategoryManagement } from '@icon-park/vue-next';

const isCollapse = ref(false);

const menuList = [
  {
    name: '/modules',
    icon: ApplicationTwo,
    title: 'Modules',
  },
  {
    name: '/stacks',
    icon: CategoryManagement,
    title: 'Stacks'
  }
]

</script>

<style scoped>
.el-menu {
  border: none;
}
</style>
