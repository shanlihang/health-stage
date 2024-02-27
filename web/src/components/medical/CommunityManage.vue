<script setup lang="ts">
import {ref,onMounted,onUnmounted} from 'vue'
import AMapLoader from "@amap/amap-jsapi-loader";

const address = ref<string>('')

//新增社区弹窗
const addCommunityModal = ref<boolean>(false)

//定义地图map
let map = null

//新增社区按钮点击事件
const addCommunity = () => {
    AMapLoader.load({
    key: "cf9e9bc4d8c1e744c1298de4af957ad5", // 申请好的Web端开发者Key，首次调用 load 时必填
    version: "2.0", // 指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
  })
    .then((AMap:any) => {
      map = new AMap.Map("map", {
        // 设置地图容器id
        zoom:14, // 初始化地图级别
        center: [116.917177,36.672156], // 初始化地图中心点位置，不加会定位到当前城市
      });
      //此处添加插件
    })
    .catch((e:any) => {
      console.log(e);
    });
    addCommunityModal.value = true
}

//新增社区弹窗确认事件
const addCommunityOk = () => {
    addCommunityModal.value = false
    map?.destroy();
}

//新增社区弹窗取消事件
const addCommunityCancel = () => {
    addCommunityModal.value = false
    map?.destroy();
}

//新增社区弹窗搜索事件
const onSearch = () => {

onMounted(() => {
  
})

onUnmounted(() => {
  
});
}
</script>

<template>
    <div class="container">
        <a-button type="primary" @click="addCommunity">新增社区</a-button>


        <a-modal v-model:open="addCommunityModal" style="width: 60%;" title="新增合作社区" @ok="addCommunityOk" @cancel="addCommunityCancel" okText="确认" cancelText="取消">
            <a-input-search
                v-model:value="address"
                placeholder="请输入地址"
                enter-button
                @search="onSearch"
            />
            <div class="map" id="map"></div>
        </a-modal>
    </div>
</template>

<style scoped>
.map{
    width: 100%;
    height: 500px;
    margin-top: 20px;
    border: 1px solid #000000;
}
</style>