<script setup lang="ts">
import {useRouter} from 'vue-router'

interface MenuItem{
  id:number
  content:string,
  src:string,
  router:string
}

const props = defineProps<MenuItem>()

//获取路由对象
const router = useRouter()

//路由跳转事件
const urlToggle = (id:number) => {
  if(props.router == '/'){
    router.push('/')
  }else if(props.router == '/setting'){
    router.push('/setting')
  } else {
    router.push({
      path:props.router,
      query:{
        id:props.id
      }
    })
  }
}
</script>

<template>
  <div class="section">
    <a-tooltip placement="right" color="#0884d6">
        <template #title>
          <span>{{ props.content }}</span>
        </template>
        <img :src="props.src" width="40px" height="40px" @click="urlToggle(props.id)">
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
