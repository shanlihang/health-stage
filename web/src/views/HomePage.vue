<script setup lang="ts">
import LastMenu from '/src/components/base/LastMenu.vue'
import { reactive } from 'vue';

//快捷访问菜单接口
interface QuickMenu{
  id:number,
  name:string,
  src?:string
}

const quickMenu = reactive<Array<QuickMenu>>([
  {id:1,name:'数据大屏'},
  {id:2,name:'检测结果'},
  {id:3,name:'数据大屏'},
  {id:4,name:'检测结果'},
])
</script>

<template>

  <div class="container">
    <div class="quick">
      <a-card size="small" title="快捷访问" style="width: 100%" hoverable="true">
        <a-alert
          message="快捷菜单选项可在系统设置中配置"
          banner
          closable
        />
        <div class="menus">
          <LastMenu v-for="item in quickMenu" :key="item.id" :name="item.name"></LastMenu>
        </div>


          
        
      </a-card>
    </div>
    
  </div>
</template>

<style scoped>
.container{
  width: 100%;
  height: 100%;
  background-color: #fff;
  display: flex;
  flex-direction: column;
  align-items: center;
}
.quick{
  width: 96%;
  margin-top: 30px;
}
.quick .menus{
  width: 100%;
  height: 70px;
  /* background-color: aqua; */
  display: flex;
  align-items: center;
  justify-content: space-around;
}
</style>
