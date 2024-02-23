<script setup lang="ts">
import {ref} from 'vue'
import SetWaterMark from '/src/components/system/SetWaterMark.vue'

//选项卡激活页面
const activeKey = ref<number>(1)
</script>

<template>
    
  <div class="container">
    <a-tabs
      v-model:activeKey="activeKey"
      tab-position="left"
      size="large"
      :style="{ width:'100%',height:'100%'}"
    >
    <a-tab-pane tab="水印设置"><SetWaterMark /></a-tab-pane>
    </a-tabs>
  </div>
</template>

<style scoped>
.container{
  width: 100%;
  height: 100%;
  border-radius: 9px;
  background-color: #fff;
}
</style>
