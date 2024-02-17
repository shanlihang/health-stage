<script setup lang="ts">

</script>

<template>
  <div class="container">
      <div class="logo">
      </div>
      <div class="menuList">
        <slot></slot>
      </div>
      <div class="user"></div>
  </div>
</template>

<style scoped>
.container{
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;
    background-color: rgb(0, 0, 0);
}
.container .logo{
    width: 80px;
    height: 9%;
    border-radius: 50%;
    background-color: antiquewhite;
}
.container .menuList{
    width: 90%;
    height: 70%;
    background-color: aqua;
}
.container .user{
    width: 80px;
    height: 9%;
    background-color: antiquewhite;
    border-radius: 50%;
}
</style>