<script setup lang="ts">
const props = defineProps({
    name:{
        default:'菜单分类'
    }
})
</script>

<template>
  <div class="card">
    <div class="flag"></div>
    <div>{{ props.name }}</div>
  </div>
  <div class="items">
    <slot></slot>
  </div>
</template>

<style scoped>
.card{
    /* width: 100%; */
    margin: 0 20px;
    height: 40px;
    font-size: 16px;
    font-weight: bold;
    display: flex;
    align-items: center;
    border-bottom: 2px solid #494949;
    margin-top: 30px;
    
}
.flag{
    width: 5px;
    height: 30px;
    margin: 0 15px;
    border-radius: 3px;
    background-color: darkmagenta;
}
.items{
    margin: 0 20px;
    display: flex;
    justify-content: space-around;
    flex-wrap: wrap;
    align-items: center;
}
</style>
