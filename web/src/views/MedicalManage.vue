<script setup lang="ts">
import {ref,onMounted} from 'vue'

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
      <a-tab-pane v-for="i in 30" :key="i" :tab="`Tab-${i}`">Content of tab {{ i }}</a-tab-pane>
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
