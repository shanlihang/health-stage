<script setup lang="ts">

const props = defineProps({
    title:{
        default:'一体机设备信息管理系统',
    }
})
</script>

<template>
  <div class="container">
    <div class="title">
        {{ props.title }}
    </div>
  </div>
</template>

<style scoped>
.container{
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: rgb(0, 0, 0);
}
.container .title{
    font-size: 24px;
    color: #fff;
    font-weight: 800;
    letter-spacing: .2rem;
}
</style>