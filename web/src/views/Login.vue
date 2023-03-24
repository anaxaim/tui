<template>
  <div class="content">
    <div class="main__form">
      <div class="icon__form">
        <img src="@/assets/images/tui.png" class="icon" alt=""/>
        <h1 class="icon__title">Tui</h1>
      </div>

      <div v-if="showLogin">
        <el-form ref="loginFormRef" :model="loginUser" size="large" :rules="rules" show-message>
          <el-form-item prop="name">
            <el-input v-model="loginUser.name" placeholder="tui">
              <template #prefix>
                <User />
              </template>
            </el-input>
          </el-form-item>

          <el-form-item prop="password">
            <el-input v-model="loginUser.password" type="password" placeholder="tui312" show-password>
              <template #prefix>
                <Lock />
              </template>
            </el-input>
          </el-form-item>
        </el-form>

        <el-button class="submit_form_btn" type="success" size="large" @click="login(loginFormRef)">SIGN IN</el-button>

        <div class="change_form_btn">
          <el-button link @click="showLogin=false">SIGN UP</el-button>
        </div>
      </div>

      <div v-if="showLogin === false">
        <el-form ref="registrationFormRef" :model="registrationUser" label-position="top" :rules="rules"
                 label-width="auto" size="large">
          <el-form-item label="Username" prop="name">
            <el-input placeholder="username" v-model="registrationUser.name" size="large"></el-input>
          </el-form-item>
          <el-form-item label="Email" prop="email">
            <el-input placeholder="email" v-model="registrationUser.email"></el-input>
          </el-form-item>
          <el-form-item label="Password" prop="password">
            <el-input placeholder="password" minlength="6" v-model="registrationUser.password"></el-input>
          </el-form-item>
        </el-form>

        <el-button class="submit_form_btn" type="success" size="large" @click="register(registrationFormRef)">SIGN UP</el-button>
        <div class="change_form_btn">
          <el-button link @click="showLogin=true">SIGN IN</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
/*
  imports
*/
  import { ElMessage, ElNotification } from "element-plus"
  import { User, Lock } from '@icon-park/vue-next'
  import { ref, reactive } from 'vue'
  import request from '@/axios'
  import { useRouter } from 'vue-router'

/*
  router
*/
  const router = useRouter();

/*
  rules
*/
  const rules = reactive({
    name: [
      { required: true, message: 'Please input username', trigger: 'blur' }
    ],
    password: [
      { required: true, message: 'Please input password', trigger: 'blur' },
      { min: 6, message: 'Length should be great than 6', trigger: 'blur' }
    ],
    email: [
      { required: true, message: 'Please input email', trigger: 'blur' },
      { type: 'email', message: 'Please input correct email address', trigger: ['blur', 'change'] },
    ]
  });

/*
  login
*/
  const loginFormRef = ref();
  const showLogin = ref(true);

  const loginUser = reactive({
    name: '',
    password: '',
  });

  const login = async (form) => {
    if (!form) {
      return
    }

    let name = loginUser.name;

    await form.validate((valid, fields) => {
      if (valid) {
        request.post('/api/v1/auth/token', {
          name: loginUser.name,
          password: loginUser.password,
          setCookie: true,
        }).then(() => {
          ElNotification.success({
            title: 'Login Success',
            message: 'Hi~ ' + name,
            showClose: true,
            duration: 1500,
          })
          router.push('/modules');
        })
      } else {
        console.log('Invalid input =>', fields)
        ElMessage({
          message: 'Invalid input',
          type: 'error',
        });
      }
    });
  };

/*
  registration
*/
  const registrationFormRef = ref();

  const registrationUser = reactive({
    name: '',
    email: '',
    password: '',
  });

  const register = async (form) => {
    if (!form) {
      return
    }

    await form.validate((valid, fields) => {
      if (valid) {
        request.post('/api/v1/auth/user', {
          name: registrationUser.name,
          password: registrationUser.password,
          email: registrationUser.email,
        }).then(() => {
          ElMessage({
            message: 'Register success',
            type: 'success',
          })
          loginUser.name = registrationUser.name;
          loginUser.password = registrationUser.password;
        })
      } else {
        console.log('Invalid input =>', fields)
        ElMessage({
          message: 'Invalid input',
          type: 'error',
        });
      }
    });
  };
</script>

<style lang="scss">
  .content {
    display: flex;
    height: 80vh;
    justify-content: center;
    align-items: center;
    flex-grow: 0.8;
  }

  .main__form {
    height: max-content;
    min-width: 16rem;
    max-width: 24rem;
    width: 25%;
    text-align: center;
    align-items: center;
  }

  .icon__form {
    display: inline-flex;
    align-items: center;
    margin-top: 1rem;
    margin-bottom: 2rem;
  }

  .icon {
    height: 3rem;
    margin-right: 0.5rem;
  }

  .icon__title {
    font-weight: bold;
    font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace;
    font-size: 1.25rem;
    line-height: 1.75rem;
  }

  .submit_form_btn {
    width: 100%;
  }

  .change_form_btn {
    text-align: right;
    margin-top: 0.25rem;
  }
</style>
