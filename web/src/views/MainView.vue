<script setup lang="ts">
import { RouterView } from 'vue-router';
import MenuAside from '../components/layout/MenuAside.vue'
import HeadAside from '../components/layout/HeadAside.vue'

</script>

<template>
  <div class="container">
    <div class="side">
        <MenuAside />
    </div>
    <div class="header">
        <HeadAside />
    </div>
    <div class="main">
        <RouterView></RouterView>
    </div>
    
    
  </div>
</template>

<style scoped>
.container{
    width:100%;
    height:100vh;
    display:flex;
}
.container .side{
    width:5%;
    height:100vh;
}
.container .header{
    width:95%;
    height:8%;
}
.container .main{
    width:95%;
    height:92%;
    position: absolute;
    left:5%;
    top:8%;
    background-color:#0dd606;
}
</style>
