<script setup lang="ts">
import {ref,onMounted} from 'vue'
import {useRouter} from 'vue-router'

//获取路由对象
const router = useRouter()

//绑定搜索框内容
const search = ref<string>('')

//点击搜索按钮触发
const onSearch = () => {

}

onMounted(() => {
})
</script>

<template>

  <div class="container">
    <div class="search">
      <a-input-search
      v-model:value="search"
      placeholder="请输入菜单名"
      enter-button="搜索"
      size="large"
      @search="onSearch"
    />
    </div>
    <div class="divide">
      <a-divider style="color: rgb(87, 216, 13);font-weight: bold;">菜单列表</a-divider>
    </div>
    
    <div class="section"></div>
    
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
