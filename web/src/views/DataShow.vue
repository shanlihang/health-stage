<script setup lang="ts">
import { ref, reactive,onMounted } from "vue";
import { message } from 'ant-design-vue';
import screenfull from 'screenfull'

//判断是否全屏
const flag = ref<boolean>(false)

//定义ref变量
const dataSection = ref(null)



//刷新数据
const refresh = () => {
  console.log('刷新数据');
}
//全屏显示
const fullScreen = () => {
  if(!screenfull.isEnabled){
    message.error('该浏览器不支持全屏功能');
  }else{
    screenfull.toggle(dataSection.value)
    flag.value = !flag.value
  }
  
}

onMounted(() => {
})
</script>

<template>
  <div class="container" ref="dataSection">
    <div class="title">
      <div class="name">数据大屏</div>
      <bv-decorator
        name="decorator5"
        style="width: 50%; height: 100%"
      ></bv-decorator>
      <div class="btns">
        <div class="refresh">
          <a-tooltip placement="bottom" style="z-index:111" :getPopupContainer="(e) => e.parentNode">
            <template #title>
              <span>刷新数据</span>
            </template>
            <img
              src="/src/assets/pic/header/refresh.svg"
              width="30"
              height="30"
              @click="refresh"
            />
          </a-tooltip>
        </div>
        <div class="full">
          <a-tooltip placement="bottom" :getPopupContainer="(e) => e.parentNode">
            <template #title>
              <span>{{ flag ? '退出全屏' : '全屏显示' }}</span>
            </template>
            <img @click="fullScreen" src="/src/assets/pic/header/full.svg" width="25" height="25" />
          </a-tooltip>
        </div>
      </div>
    </div>
    <div class="datas"></div>
  </div>
</template>

<style scoped>
.container {
  width: 100%;
  height: 100%;
  color: #fff;
  background: url("/src/assets/bkg.jpg") no-repeat;
  background-size: cover;
}
.title {
  height: 10%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  position: relative;
}
.title .name {
  font-size: 18px;
  font-weight: bold;
  height: 20px;
  position: absolute;
  top: 15px;
}
.btns {
  display: flex;
  justify-content: space-around;
  align-items: center;
  width: 150px;
  height: 100%;
  position: absolute;
  right: 0;
}
.btns .refresh,
.full {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgb(255, 255, 255);
}
.datas {
  width: 100%;
  height: 90%;
  /* background-color: brown; */
}
</style>
