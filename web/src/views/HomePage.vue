<script setup lang="ts">
import LastMenu from '/src/components/base/LastMenu.vue'
import { ref,reactive } from 'vue';
import { ScanOutlined }from '@ant-design/icons-vue'

//快捷访问菜单接口
interface QuickMenu{
  id:number,
  name:string,
  src?:string
}

//二维码弹窗
const scanOpen = ref<boolean>(false)

//二维码网址
const text = ref<string>("https://www.antdv.com/")

//快捷菜单列表
const quickMenu = reactive<Array<QuickMenu>>([
  {id:1,name:'数据大屏'},
  {id:2,name:'检测结果'},
  {id:3,name:'数据大屏'},
  {id:4,name:'检测结果'},
])

//点击悬浮按钮打开二维码弹窗
const openScanModal = () => {
  scanOpen.value = true
}

//二维码弹窗处理函数
const scanModalhandle = () => {
  scanOpen.value = false
}
</script>

<template>

  <div class="container">
    <div class="quick">
      <a-card size="small" title="快捷访问" style="width: 100%" :hoverable="true">
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

    <div class="quick">
      <a-card size="small" title="快捷访问" style="width: 100%;height: 400px;" :hoverable="true">        
      </a-card>
    </div>
    <div class="quick">
      <a-card size="small" title="快捷访问" style="width: 100%;height: 400px;" :hoverable="true">        
      </a-card>
    </div>
    <div class="quick">
      <a-card size="small" title="快捷访问" style="width: 100%;height: 400px;" :hoverable="true">        
      </a-card>
    </div>


    <a-float-button
      @click="openScanModal"
      type="primary"
      :style="{
        right: '58px',
      }"
    >
      <template #icon>
        <ScanOutlined />
      </template>
    </a-float-button>

    <a-modal v-model:open="scanOpen" :centered="true" okText="完成" cancelText="取消" title="扫码查看检测数据" @ok="scanModalhandle">

        <a-alert message="注：用户扫码查看个人数据（需使用身份证号进行身份验证）" type="info" :closable="true"/>
        <div class="imgs">
          <a-qrcode :value="text" />
          <img src="/src/assets/pic/other/scan.svg" width="150" height="150">
        </div>
    </a-modal>
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
  border-radius: 9px;
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
.imgs{
  width: 100%;
  height: 150px;
  margin-top: 20px;
  display: flex;
  align-items: center;
  justify-content: space-around;
}
</style>
