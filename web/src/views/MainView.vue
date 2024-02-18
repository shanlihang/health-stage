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
    isSelect:boolean
}

const flag = ref<number>(1)

const list:Array<menuList> = reactive([
    {id:1,content:'首页',src:'/src/assets/pic/menus/index.svg',isSelect:true},
    {id:2,content:'医疗管理',src:'/src/assets/pic/menus/list.svg',isSelect:false},
    {id:3,content:'库存管理',src:'/src/assets/pic/menus/site.svg',isSelect:false},
    {id:4,content:'系统设置',src:'/src/assets/pic/menus/settings.svg',isSelect:false},
    {id:5,content:'数据大屏',src:'/src/assets/pic/menus/data.svg',isSelect:false}
])

const selectMenu = (item:menuList) => {
    flag.value = item.id
}

</script>

<template>
  <div class="container">
    <div class="side">
        <MenuAside>
            <MenuItem :class="item.id == flag ? 'select' : ''" @click="selectMenu(item)" :content="item.content" :src="item.src" v-for="item in list" :key="item.id" />
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
.select{
    animation: changeBkg 0.2s linear forwards;
}

@keyframes changeBkg {
    to{
        border: 1px rgb(13, 238, 152) solid;
        /* background-color: dimgray; */
    }
}
</style>
