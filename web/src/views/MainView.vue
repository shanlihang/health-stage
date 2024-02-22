<script setup lang="ts">
import { RouterView } from 'vue-router';
import {ref,reactive} from 'vue'
import MenuAside from '../components/layout/MenuAside.vue'
import HeadAside from '../components/layout/HeadAside.vue'
import MenuItem from '../components/base/MenuItem.vue'

interface menuList{
    id:number
    content:string,
    src:string,
}

const flag = ref<number>(1)

const list:Array<menuList> = reactive([
    {id:1,content:'首页',src:'/src/assets/pic/menus/index.svg'},
    {id:2,content:'医疗管理',src:'/src/assets/pic/menus/list.svg'},
    {id:3,content:'库存管理',src:'/src/assets/pic/menus/site.svg'},
    {id:4,content:'系统设置',src:'/src/assets/pic/menus/settings.svg'},
    {id:5,content:'数据大屏',src:'/src/assets/pic/menus/data.svg'}
])

const selectMenu = (item:menuList) => {
    flag.value = item.id
}

const selectAccountCenter = () => {
    flag.value = 0
}

</script>

<template>
  <div class="container">
    <div class="side">
        <MenuAside @chooseAccountCenter="selectAccountCenter">
            <MenuItem :class="item.id == flag ? 'select' : ''" @click="selectMenu(item)" :content="item.content" :src="item.src" v-for="item in list" :key="item.id" />
        </MenuAside>
    </div>
    <div class="head">
        <HeadAside title="测试组件系统"/>
    </div>
    <div class="main">
        <div class="router">
            <img src="/src/assets/pic/header/router.svg" width="20" height="20" style="margin-right:10px">
            <a-breadcrumb>
                <a-breadcrumb-item>Home</a-breadcrumb-item>
                <a-breadcrumb-item><a href="">Application Center</a></a-breadcrumb-item>
                <a-breadcrumb-item><a href="">Application List</a></a-breadcrumb-item>
                <a-breadcrumb-item>An Application</a-breadcrumb-item>
            </a-breadcrumb>
        </div>
        <div class="show">
            <RouterView></RouterView>
        </div>
        
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
    background-color: rgba(214, 214, 214, 0.3);
}
.router{
    width: 98%;
    height: 5%;
    margin: 0 auto;
    display: flex;
    align-items: center;
}
.show{
    width: 98%;
    height: 93%;
    margin: 0 auto;
    background-color: #fff;
}
.select{
    animation: changeBkg 0.2s linear forwards;
}

@keyframes changeBkg {
    to{
        background-color: rgba(207, 207, 166, 0.3);
    }
}
</style>
