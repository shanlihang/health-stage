<script setup lang="ts">
const props = defineProps({
    name:{
        default:'菜单项1'
    },
    src:{
        default:"/src/assets/vue.svg"
    }
})
</script>

<template>

  <div class="cardItem">
    <div class="pic">
        <img :src="props.src" width="40px" height="40px" style="background-color:#fff">
    </div>
    <div class="name">
        {{ props.name }}
    </div>
  </div>
</template>

<style scoped>
.cardItem{
    width: 18%;
    height: 50px;
    background-color: #fdfbfb;
    margin-top: 20px;
    display: flex;
    align-items: center;
    box-shadow: rgba(50, 50, 105, 0.15) 0px 2px 5px 0px, rgba(0, 0, 0, 0.05) 0px 1px 1px 0px;
}
.pic{
    margin: 0 15px 0 10px;
}
.cardItem:hover{
    background-color: rgb(233, 233, 233);
}
</style>
