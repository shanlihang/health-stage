<script setup lang="ts">
import {ref,reactive} from 'vue'
import GoodsManage from '/src/components/inventory/GoodsManage.vue'
import StockManage from '/src/components/inventory/StockManage.vue'

//当前激活选项卡
const activeKey = ref<number>(1)
</script>

<template>
  <div class="container">
    <a-tabs
      v-model:activeKey="activeKey"
      tab-position="left"
      :style="{ height: '100%'}"
    >
        <a-tab-pane :key="1" tab="物品管理"><GoodsManage /></a-tab-pane>
        <a-tab-pane :key="2" tab="库存管理"><StockManage /></a-tab-pane>
        <a-tab-pane :key="3" tab="操作记录">789</a-tab-pane>
    </a-tabs>
  </div>
</template>

<style scoped>
.container{
  width: 100%;
  height: 100%;
  border-radius: 9px;
  background-color: #fff;
  overflow-y: scroll;
}
</style>
