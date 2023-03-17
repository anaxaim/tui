<template>
  <el-header>
    <el-row class="flex h-full w-full" justify="center" align="middle">
      <el-col :span="4" class="flex text-center items-center content-center">
            <img class="w-[2.5rem] mx-[0.5rem]" src="@/assets/tui.png"  alt=""/>
            <span class="font-bold font-mono text-2xl pl-[0.5rem]">Tui</span>
      </el-col>
    </el-row>
  </el-header>
</template>

<script setup>
import { getUser, delUser } from '@/utils';
import request from '@/axios';
import { ElNotification } from "element-plus";
import { useRouter } from 'vue-router';

const user = getUser();
const router = useRouter();

function logout() {
  let lg = function () {
    console.log("logout success")
    ElNotification.success({
      title: 'Logout Success',
      message: 'Bye~ ' + user.name,
      showClose: true,
      duration: 1500,
    })
    delUser();
    router.push('/login');
  }

  request.delete("/api/v1/auth/token").then(() => {
    lg();
  }).catch((error) => {
    console.log(error)
  })
}

</script>
