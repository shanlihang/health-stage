<script setup lang="ts">
import defineProps from 'vue'

interface MenuItem{
  content:string,
  src:string
}

const props = defineProps<MenuItem>()
</script>

<template>
  <div class="section">
    <a-tooltip placement="right" color="#0884d6">
        <template #title>
          <span>{{ props.content }}</span>
        </template>
        <img :src="props.src" width="40px" height="40px">
      </a-tooltip>
  </div>
</template>

<style scoped>
.section{
    width: 60px;
    height: 60px;
    overflow: hidden;
    border-radius: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
}
.section:hover{
  background-color: rgb(32, 33, 35);
}
</style>
