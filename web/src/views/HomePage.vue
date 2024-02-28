<script setup lang="ts">
import LastMenu from "/src/components/base/LastMenu.vue";
import { ref, reactive } from "vue";
import { ScanOutlined } from "@ant-design/icons-vue";

//快捷访问菜单接口
interface QuickMenu {
  id: number;
  name: string;
  src?: string;
}

//数据卡片信息
const dataCard = reactive([
  {title:'合作社区数量',value:10},
  {title:'系统参检人数',value:1000},
  {title:'今日新增人数',value:50},
  {title:'今日检测人数',value:10},
])

//地图原点坐标
const mapCircle = reactive([
  {row:116.897468,col:36.718323},
  {row:116.947781,col:36.684456},
  {row:116.811754,col:40.208198}

])

//二维码弹窗
const scanOpen = ref<boolean>(false);

//二维码网址
const text = ref<string>("https://www.antdv.com/");

//点击悬浮按钮打开二维码弹窗
const openScanModal = () => {
  scanOpen.value = true;
};

//二维码弹窗处理函数
const scanModalhandle = () => {
  scanOpen.value = false;
};
</script>

<template>
  <div class="container">
    <div class="quick">
      <a-card
        size="small"
        style="width: 100%"
        :hoverable="true"
      >
        <a-row>
          <a-col :span="6" v-for="(item,index) in dataCard" :key="index">
            <a-statistic
              :title="item.title"
              :value="item.value"
              style="text-align: center;"
            />
          </a-col>
        </a-row>
      </a-card>
    </div>

    <div class="quick">
      <a-card
        size="small"
        title="快捷访问"
        style="width: 100%; height: 400px"
        :hoverable="true"
      >
      </a-card>
    </div>
    <div class="quick">
      <a-card
        size="small"
        title="快捷访问"
        style="width: 100%; height: 400px"
        :hoverable="true"
      >
      </a-card>
    </div>
    <div class="quick">
      <a-card
        size="small"
        title="快捷访问"
        style="width: 100%; height: 400px"
        :hoverable="true"
      >
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

    <a-modal
      v-model:open="scanOpen"
      :centered="true"
      okText="完成"
      cancelText="取消"
      title="扫码查看检测数据"
      @ok="scanModalhandle"
    >
      <a-alert
        message="注：用户扫码查看个人数据（需使用身份证号进行身份验证）"
        type="info"
      />
      <div class="imgs">
        <a-qrcode :value="text" />
        <img src="/src/assets/pic/other/scan.svg" width="150" height="150" />
      </div>
    </a-modal>
  </div>
</template>

<style scoped>
.container {
  width: 100%;
  height: 100%;
  background-color: #fff;
  display: flex;
  flex-direction: column;
  align-items: center;
  overflow-y: scroll;
}
.quick {
  width: 96%;
  margin-top: 30px;
}
.quick .menus {
  width: 100%;
  height: 70px;
  /* background-color: aqua; */
  display: flex;
  align-items: center;
  justify-content: space-around;
}
.imgs {
  width: 100%;
  height: 150px;
  margin-top: 20px;
  display: flex;
  align-items: center;
  justify-content: space-around;
}
</style>
