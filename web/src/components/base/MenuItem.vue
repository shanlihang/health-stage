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
    <a-tooltip placement="right">
        <template #title>
          <span>{{ props.content }}</span>
        </template>
        <img :src="props.src" alt="" style="width:30px;height:30px;">
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
  background-color: rgb(49, 50, 51);
}
</style>
