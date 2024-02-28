<script setup lang="ts">
import {ref} from 'vue'
import SetWaterMark from '/src/components/system/SetWaterMark.vue'

//当前激活选项卡
const activeKey = ref<number>(1)
</script>

<template>
    
  <div class="container">
    <a-tabs
      v-model:activeKey="activeKey"
      tab-position="left"
      :style="{height:'100%'}"
    >
    <a-tab-pane tab="水印设置" :key="1"><SetWaterMark /></a-tab-pane>
    <a-tab-pane tab="位置设置" :key="2">更改我的位置</a-tab-pane>
    </a-tabs>
  </div>
</template>

<style scoped>
.container{
  width: 100%;
  height: 100%;
  background-color: #fff;
}
</style>
