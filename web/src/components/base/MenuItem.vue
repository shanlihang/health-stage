<script setup lang="ts">
import {useRouter} from 'vue-router'

interface MenuItem{
  id:number
  content:string,
  src:string,
  router:string
}

const props = defineProps<MenuItem>()

</script>

<template>
  <div class="section">
    <a-tooltip placement="right" color="#0884d6">
        <template #title>
          <span>{{ props.content }}</span>
        </template>
        <router-link :to="props.router">
          <img :src="props.src" width="40px" height="40px">
        </router-link>
        
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