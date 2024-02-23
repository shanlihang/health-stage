<script setup lang="ts">
import {ref,reactive} from 'vue'

interface Setwatermark{
    content:string,
    flag:boolean
}

//输入框内容
const form = reactive<Setwatermark>({
    content:'',
    flag:true
})

//弹窗开关控制
const open = ref<boolean>(false)

//更改水印设置
const toUpdate = () => {}

//预览水印效果
const toShow = () => {
    open.value = true
}

//弹窗确认按钮
const handleOk = () => {
    open.value = false
}
</script>

<template>
  <div class="container">
    <a-card style="width: 90%;margin-top: 20px;" title="水印设置">
        <div class="form">
        <a-form :model="form">
            <a-form-item label="水印文本">
                <a-input v-model:value="form.content" style="width: 50%;" placeholder="请输入想要设置的水印文本"/>
            </a-form-item>
            <a-form-item label="水印开关">
                <a-switch v-model:checked="form.flag" />
            </a-form-item>
        </a-form>
    </div>
    <div style="width: 100%;display: flex;justify-content: space-around;">
        <div>
            <a-button type="primary" @click="toUpdate">确认</a-button>
            <a-button style="margin-left: 15px;" @click="toShow">预览</a-button>
        </div>
    </div>
    </a-card>
    <a-modal v-model:open="open" title="预览水印效果" @ok="handleOk" :centered="true" okText="确认更改" cancelText="取消">
        <a-watermark :content="form.content">
            <img style="border: 1px solid #000;" src="/src/assets/watermark.png" width="470" height="300">
  </a-watermark>
      
    </a-modal>
  </div>
</template>

<style scoped>
.container{
  width: 100%;
  height: 100%;
  display: flex;flex-direction: column;
  justify-content: center;
  align-items: center;
}
.form{
    width: 50%;
    height: 100px;
    margin-top: 20px;
}
</style>
