<script setup lang="ts">
import {ref,reactive} from 'vue'

interface Goods{
    name:string
}
const formData = reactive<Goods>({
    name:''
})

//当前激活选项卡
const activeKey = ref<number>(1)
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





    <!-- <div class="search">
        <a-form layout="inline" :model="formData">
            <a-form-item label="物品名称">
                <a-input v-model:value="formData.name" placeholder="请输入物品名称" />
            </a-form-item>
            <a-form-item>
                <a-button type="primary">搜索</a-button>
            </a-form-item>
        </a-form>
    </div>
    <div class="action">
        <a-button type="primary">Primary</a-button>
        <a-button type="dashed">Dashed</a-button>
    </div>
    <div class="table">

    </div> -->
  </div>
</template>

<style scoped>
.container{
  width: 100%;
  height: 100%;
  border-radius: 9px;
  background-color: #fff;
}
.search{
    height: 5%;
    width: 98%;
    margin: 0 auto;
    background-color: aqua;
}
.action{
    
    margin: 20px 20px 0;
    background-color: aqua;
}
</style>
