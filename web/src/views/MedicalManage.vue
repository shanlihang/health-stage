<script setup lang="ts">
import {ref,onMounted} from 'vue'


onMounted(() => {
})
</script>

<template>

  <div class="container">
    医疗管理
  </div>
</template>

<style scoped>
.container{
  width: 100%;
  height: 100%;
  background-color: #fff;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  border-radius: 9px;
}
.search{
  width: 90%;
  height: 8%;
  display: flex;
  justify-content: center;
  align-items: center;
}
.divide{
  width: 90%;
  height: 7%;
}
.section{
  width: 90%;
  height: 85%;
  background-color: blue;
}
</style>
