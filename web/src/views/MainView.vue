<script setup lang="ts">
import { RouterView } from 'vue-router';
import MenuAside from '../components/layout/MenuAside.vue'
import HeadAside from '../components/layout/HeadAside.vue'
import MenuItem from '../components/base/MenuItem.vue'

</script>

<template>
  <div class="container">
    <div class="side">
        <MenuAside>
            <MenuItem content="1" src="../assets/vue.svg"/>
            <MenuItem content="2" src="../assets/vue.svg"/>
            <MenuItem content="3" src="../assets/vue.svg"/>
            <MenuItem content="4" src="../assets/vue.svg"/>
            <MenuItem content="5" src="../assets/vue.svg"/>
        </MenuAside>
    </div>
    <div class="head">
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
    height:100%;
    position: relative;
}
.container .side{
    width:5%;
    height:100%;
}
.container .head{
    width:95%;
    height: 8%;
    position: absolute;
    top: 0;
    left: 5%;
}
.container .main{
    width:95%;
    height:92%;
    position: absolute;
    left:5%;
    top:8%;
}
</style>
