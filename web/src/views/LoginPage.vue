<script setup lang="ts">

import {reactive} from 'vue'

//表单接口
interface LoginForm{
  username:string,
  password:string,
  code:string
}

//定义表单数据
const formData = reactive<LoginForm>({
  username:'',
  password:'',
  code:''
})

//表单提交成功
const onFinish = () => {}

//表单提交失败
const onFinishFailed = () => {}
</script>

<template>
  <div class="container">
    <div class="login">
      <div class="form">
        <div class="title">登录</div>
        <a-form
    :model="formData"
    name="basic"
    autocomplete="off"
    @finish="onFinish"
    @finishFailed="onFinishFailed"
    style="margin-top: 20px;"
  >
    <a-form-item
      label="账号"
      name="username"
      :rules="[{ required: true, message: '账号不能为空' }]"
      :labelCol="{ style: 'width: 70px' }"
      labelAlign="left"
    >
      <a-input v-model:value="formData.username" placeholder="请输入账号" />
    </a-form-item>

    <a-form-item
      label="密码"
      name="password"
      :rules="[{ required: true, message: '密码不能为空' }]"
      :labelCol="{ style: 'width: 70px' }"
      labelAlign="left"
    >
      <a-input-password v-model:value="formData.password" placeholder="请输入密码" />
    </a-form-item>
    <a-form-item
      label="验证码"
      name="password"
      :rules="[{ required: true, message: '验证码不能为空' }]"
      :labelCol="{ style: 'width: 70px' }"
      labelAlign="left"
      style="position: relative;"
    >
      <a-input style="width: 50%;" v-model:value="formData.code" placeholder="请输入验证码" />
      <div class="codeArea" style="position: absolute;right: 0;top: 0;"></div>
    </a-form-item>

    <a-form-item>
      <a-button type="primary" html-type="submit" style="margin-left: 50%;">登录</a-button>
    </a-form-item>
  </a-form>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container{
  width: 100%;
  height: 100%;
  background: url('/src/assets/login-bkg.jpg') no-repeat;
  background-size: cover;
  overflow: hidden;
  display: flex;
  justify-content: center;
  align-items: center;
}
.login{
  width: 700px;
  height: 350px;
  background-color: aliceblue;
  border-radius: 9px;
  background: url('/src/assets/login.png') no-repeat;
  background-size: cover;
}

.login .form{
  width: 45%;
  height: 100%;
  margin-left: 5%;
}
.login .title{
  font-size: 24px;
  font-weight: bold;
  letter-spacing: 0.1rem;
  text-align: center;
  padding-top: 20px;
  margin-top: 20px;
}
.codeArea{
  width: 40%;
  height: 100%;
  background-color: aqua;
}
</style>
