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
            <MenuItem content="首页" src="/src/assets/pic/menus/index.svg"/>
            <MenuItem content="医疗管理" src="/src/assets/pic/menus/list.svg"/>
            <MenuItem content="库存管理" src="/src/assets/pic/menus/site.svg"/>
            <MenuItem content="系统设置" src="/src/assets/pic/menus/settings.svg"/>
            <MenuItem content="数据大屏" src="/src/assets/pic/menus/data.svg"/>
        </MenuAside>
    </div>
    <div class="head">
        <HeadAside title="测试组件系统"/>
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
