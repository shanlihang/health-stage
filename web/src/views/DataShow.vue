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

//我的定位
const myAddress = ref({
  row: 116.917177,
  col: 36.672156,
});

//地图原点坐标
const mapCircle = reactive([
  { row: 116.897468, col: 36.718323 },
  { row: 116.947781, col: 36.684456 },
]);

//存储圆点标记实例
const arr = reactive([]);

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

//创建地图标记圆点的方法
const createCircle = (AMap,row,col,color) => {
  let flag = new AMap.CircleMarker({
    center: new AMap.LngLat(row, col),
    radius: 10,
    strokeColor: "white",
    strokeWeight: 2,
    strokeOpacity: 0.5,
    fillColor:color,
    fillOpacity: 0.5,
    zIndex: 10,
    cursor: "pointer",
    content:'123456'//自定义窗台展示信息
  });
  map.add(flag);
  arr.push(flag);
};

onMounted(() => {
  AMapLoader.load({
    key: "cf9e9bc4d8c1e744c1298de4af957ad5", // 申请好的Web端开发者Key，首次调用 load 时必填
    version: "2.0", // 指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
  })
    .then((AMap: any) => {
      map = new AMap.Map("map", {
        mapStyle: "amap://styles/darkblue",
        // center: [116.917177, 36.672156], // 初始化地图中心点位置，不加会定位到当前城市
      });
      //此处添加插件

      // 创建我的标记
      createCircle(AMap,myAddress.value.row,myAddress.value.col,'rgba(0,255,0,1)');

      //创建合作社区标记
      for (let index in mapCircle) {
        createCircle(AMap,mapCircle[index].row,mapCircle[index].col,'rgba(0,0,255,1)');
      }
      map.setFitView(arr);
      arr.forEach(element => {
        element.on("mouseover",function(e){
          
          new AMap.InfoWindow({
            content: e.target._opts.content,
            anchor: 'bottom-center',
          }).open(map,[e.target._opts.center.KL,e.target._opts.center.kT]);
          
        })
      })
    })
    .catch((e: any) => {
      console.log(e);
    });
});

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
        <div class="toggle">
          <a-tooltip
            placement="bottom"
            :getPopupContainer="(e) => e.parentNode"
          >
            <template #title>
              <span>切换面板</span>
            </template>
            <router-link to="/dataOther">
              <img
                src="/src/assets/pic/header/toggle.svg"
                width="30"
                height="30"
              />
            </router-link>
          </a-tooltip>
        </div>
        <div class="refresh">
          <a-tooltip
            placement="bottom"
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
          <div style="width: 100%; height: 100%">
            <bv-border-box name="border10">
              <div
                style="
                  width: 100%;
                  height: 100%;
                  display: flex;
                  justify-content: center;
                  align-items: center;
                "
              >
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
  width: 200px;
  height: 100%;
  position: absolute;
  right: 0;
}
.btns .refresh,
.full,
.toggle {
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
  font-size: 12px;
  color: #000;
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
