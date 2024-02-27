<script setup lang="ts">
import {ref,onMounted} from 'vue'
import AccountHome from '/src/components/account/AccountHome.vue'
import AccountManage from '/src/components/account/AccountManage.vue'

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
      <a-tab-pane :key="1" tab="个人中心"><AccountHome /></a-tab-pane>
      <a-tab-pane :key="2" tab="用户管理"><AccountManage /></a-tab-pane>
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
