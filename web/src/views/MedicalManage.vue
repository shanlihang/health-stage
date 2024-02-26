<script setup lang="ts">
import {ref,onMounted} from 'vue'
import CommunityManage from '/src/components/medical/CommunityManage.vue'

//当前激活的选项卡
const activeKey = ref<number>(1)

onMounted(() => {
})
</script>

<template>

  <div class="container">
    <a-tabs
      v-model:activeKey="activeKey"
      tab-position="left"
      :style="{ height: '100%' }"
    >
      <a-tab-pane :key="1" tab="社区管理"><CommunityManage /></a-tab-pane>
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
