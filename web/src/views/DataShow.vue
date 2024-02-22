<script setup lang="ts">
import { ref, reactive, onMounted, onUnmounted } from "vue";
import { message } from "ant-design-vue";
import screenfull from "screenfull";
import AMapLoader from "@amap/amap-jsapi-loader";

//判断是否全屏
const flag = ref<boolean>(false);

//定义ref变量
const dataSection = ref(null);

//定义地图
let map = null;

//刷新数据
const refresh = () => {
  console.log("刷新数据");
};
//全屏显示
const fullScreen = () => {
  if (!screenfull.isEnabled) {
    message.error("该浏览器不支持全屏功能");
  } else {
    screenfull.toggle(dataSection.value);
    flag.value = !flag.value;
  }
};

onMounted(() => {
  AMapLoader.load({
    key: "cf9e9bc4d8c1e744c1298de4af957ad5", // 申请好的Web端开发者Key，首次调用 load 时必填
    version: "2.0", // 指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
    plugins: [], // 需要使用的的插件列表，如比例尺'AMap.Scale'等
  })
    .then((AMap:any) => {
      map = new AMap.Map("map", {
        viewMode: "3D",
        // 设置地图容器id
        zoom:5, // 初始化地图级别
        center: [116.397428, 39.90923], // 初始化地图中心点位置
      });
    })
    .catch((e:any) => {
      console.log(e);
    });
})

onUnmounted(() => {
  map?.destroy();
});
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
          <a-tooltip
            placement="bottom"
            style="z-index: 111"
            :getPopupContainer="(e) => e.parentNode"
          >
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
          <a-tooltip
            placement="bottom"
            :getPopupContainer="(e) => e.parentNode"
          >
            <template #title>
              <span>{{ flag ? "退出全屏" : "全屏显示" }}</span>
            </template>
            <img
              @click="fullScreen"
              src="/src/assets/pic/header/full.svg"
              width="25"
              height="25"
            />
          </a-tooltip>
        </div>
      </div>
    </div>
    <div class="datas">
      <div class="fontData">
        <div class="infoCard">
          <bv-border-box name="border1">
            <div class="child">
              <span>今日新增检测数量</span>
              <span class="num">10</span>
            </div>
          </bv-border-box>
        </div>
        <div class="infoCard">
          <bv-border-box name="border1">
            <div class="child">
              <span>合作社区数量</span>
              <span class="num">10</span>
            </div>
          </bv-border-box>
        </div>
        <div class="infoCard">
          <bv-border-box name="border1">
            <div class="child">
              <span>今日高血压患者人数</span>
              <span class="num">10</span>
            </div>
          </bv-border-box>
        </div>
        <div class="infoCard">
          <bv-border-box name="border1">
            <div class="child">
              <span>高血压患者人数</span>
              <span class="num">10</span>
            </div>
          </bv-border-box>
        </div>
      </div>
      <div class="chart">
        <div class="left">
          <div class="chartItem">
            <bv-border-box name="border12"></bv-border-box>
          </div>
          <div class="chartItem">
            <bv-border-box name="border12"></bv-border-box>
          </div>
        </div>
        <div class="map">
          <div style="width: 100%; height: 100%;">
            <bv-border-box name="border10">
              <div style="width:100%;height:100%;display:flex;justify-content: center;align-items: center;">
                <div id="map"></div>
              </div>
            </bv-border-box>
          </div>
        </div>
        <div class="right">
          <div class="chartItem">
            <bv-border-box name="border12"></bv-border-box>
          </div>
          <div class="chartItem">
            <bv-border-box name="border12"></bv-border-box>
          </div>
        </div>
      </div>
    </div>
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
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  /* padding-bottom: 10px; */
}
.datas .fontData {
  height: 25%;
  display: flex;
  justify-content: space-around;
  align-items: center;
}
.infoCard {
  width: 20%;
  height: 90%;
}
.infoCard .child {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-evenly;
}
.num {
  font-size: 36px;
  font-weight: 900;
  color: crimson;
}
.chart {
  height: 70%;
  display: flex;
  justify-content: space-around;
}
.chart .left {
  width: 20%;
  height: 100%;
}
.chart .map {
  width: 50%;
  height: 100%;
}
.chart .right {
  width: 20%;
  height: 100%;
}
.chartItem {
  width: 100%;
  height: 50%;
}
#map {
  width: 98%;
  height: 98%;
  border-radius: 9px;
}
</style>
