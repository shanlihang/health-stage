<script setup lang="ts">
import { RouterView } from 'vue-router';
import {ref,reactive,onMounted} from 'vue'
import MenuAside from '../components/layout/MenuAside.vue'
import HeadAside from '../components/layout/HeadAside.vue'
import MenuItem from '../components/base/MenuItem.vue'
import {useSystemStore} from '../store/system.ts'
import {getMenuList} from '../api/system/index.ts'

interface menuList{
    ID:number
    name:string,
    icon:string,
    router:string,
    createAt:string,
    updateAt:string,
    deletedAt:string
}

//控制flag显示哪个容器
const showFlag = ref<boolean>(true)

const flag = ref<number>(1)

const list:Array<menuList> = reactive([])

//定义pinia
const store = useSystemStore()

const selectMenu = (item:menuList) => {
    flag.value = item.id
}

const selectAccountCenter = () => {
    flag.value = 0
}

const initData = () => {
    getMenuList().then(res => {
        res.data.list.forEach(element => {
            list.push(element)
        });
    })
}

onMounted(() => {
    if(location.pathname == '/login'){
        showFlag.value = !showFlag.value
    }
    initData()
})

</script>

<template>
  <div class="container" v-if="showFlag">
    <div class="side">
        <MenuAside @chooseAccountCenter="selectAccountCenter">
            <MenuItem :class="item.ID == flag ? 'select' : ''" @click="selectMenu(item)" :content="item.name" :src="item.icon" :router="item.router" :id="item.ID" v-for="item in list" :key="item.ID" />
        </MenuAside>
    </div>
    <div class="head">
        <HeadAside title="测试组件系统"/>
    </div>
    <div class="main">
        <div class="router">
            <img src="/src/assets/pic/header/router.svg" width="20" height="20" style="margin-right:10px">
            <a-breadcrumb>
                <a-breadcrumb-item><a href="">Application Center</a></a-breadcrumb-item>
                <a-breadcrumb-item><a href="">Application List</a></a-breadcrumb-item>
                <a-breadcrumb-item>An Application</a-breadcrumb-item>
            </a-breadcrumb>
        </div>
        <a-watermark :content="store.waterMark" style="width: 98%;;height: 94%;margin: 0 auto;">
            <div class="show">
                <RouterView></RouterView>
            </div>
        </a-watermark>
    </div>
  </div>
  <div class="container" v-if="!showFlag">
    <router-view></router-view>
  </div>
</template>

<style scoped>
.container{
    width:100%;
    height:100vh;
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
    background-color: #e6e6e6;
}
.router{
    width: 98%;
    height: 4%;
    margin: 0 auto;
    display: flex;
    align-items: center;
}
.show{
    width: 100%;
    height: 100%;
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
