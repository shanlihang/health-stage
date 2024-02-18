<script setup lang="ts"></script>

<template>
  <div class="container">
    <div class="logo">
      <img src="../../assets/vue.svg" width="50px" height="50px" />
    </div>
    <div class="menuList">
      <slot></slot>
    </div>
    <div class="user">
      <a-tooltip placement="top" color="#0884d6">
        <template #title>
          <span>用户中心</span>
        </template>
        <img
          src="../../assets/pic/menus/user-group.svg"
          width="80px"
          height="80px"
        />
      </a-tooltip>
    </div>
  </div>
</template>

<style scoped>
.container {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
  background-color: rgb(0, 0, 0);
}
.container .logo {
  width: 70px;
  height: 70px;
  border-radius: 50%;
  overflow: hidden;
  display: flex;
  justify-content: center;
  align-items: center;
}
.container .menuList {
  width: 90%;
  height: 70%;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
  padding: 80% 0;
}
.container .user {
  width: 60px;
  height: 60px;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
  border-radius: 50%;
}
</style>
