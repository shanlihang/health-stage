<script setup lang="ts">
import {ref,reactive} from 'vue'
import type { Dayjs } from 'dayjs';

//定义时间范围类型
type RangeValue = [Dayjs,Dayjs]

//搜索表单接口
interface SearchForm{
    name:string,
    accountId:string,
    time:string,
    remark:string,
    pageSize:number,
    currentPage:number
}

//表格数据接口
interface TableData{
    id:number,
    name:string,
    personId:number,
    type:number,//0为入库，1为出库
    num:number,
    time:string,
    remark:string,
}

//搜索表单
const searchForm = reactive<SearchForm>({
    name:'',
    accountId:'',
    time:'',
    remark:'',
    pageSize:20,
    currentPage:1
})

//分页规格
const pageSizeOptions = ref<Array<string>>(['10','20','50','100'])

//表格是否显示边框
const bordered = ref<boolean>(true)

//是否使用表格自带的分页
const table_page = ref<boolean>(false)

//操作人列表
const options = reactive([
    {value:1,label:'小明'},
    {value:2,label:'小红'},
    {value:3,label:'小黑'},
    {value:4,label:'小绿'},
    {value:5,label:'小蓝'},
])

//表格数据
const tableData = reactive<Array<TableData>>([
    {id:1,name:'sfasf',type:0,personId:2,num:50,time:'2024-01-25',remark:'提货给xxx'},
    {id:2,name:'sfasf',type:0,personId:1,num:50,time:'2024-01-25',remark:'提货给xxx'},
    {id:3,name:'sfasf',type:1,personId:45,num:30,time:'2024-01-25',remark:'进货'},
    {id:4,name:'sfasf',type:0,personId:2,num:50,time:'2024-01-25',remark:'提货给xxx'},
    {id:6,name:'sfasf',type:1,personId:5,num:50,time:'2024-01-25',remark:'提货给xxx'}
])

//表单提交成功的回调
const handleFinish = () => {}

//表单提交失败的回调
const handleFinishFailed = () => {}

//删除操作 只有超级管理员有权限
const deleteRecord = (id:number) => {
    console.log(id);
}

//分页数据改变触发
const onShowSizeChange = () => {}

</script>

<template>
    <div class="container">
        <div class="form">
            <a-form layout="inline" :model="searchForm" @finish="handleFinish" @finishFailed="handleFinishFailed">
                <a-form-item label="物品名称" style="margin-top: 20px;">
                    <a-input v-model:value="searchForm.name" placeholder="请输入物品名称" />
                </a-form-item>
                <a-form-item label="操作人" style="margin-top: 20px;">
                    <a-select
                        v-model:value="searchForm.accountId"
                        :options="options"
                        style="width: 150px;"
                        
                    ></a-select>
                </a-form-item>
                <a-form-item label="操作时间" style="margin-top: 20px;">
                    <a-range-picker v-model:value="searchForm.time" show-time />
                </a-form-item>
                <a-form-item label="物品备注" style="margin-top: 20px;">
                    <a-input v-model:value="searchForm.remark" placeholder="请输入物品备注" />
                </a-form-item>
                <a-form-item style="margin-top: 20px;">
                    <a-button type="primary">搜索</a-button>
                    <a-button style="margin-left: 15px;">重置</a-button>
                </a-form-item>
            </a-form>
        </div>
        <div class="table">
            <a-table :data-source="tableData" :bordered="bordered" :pagination="table_page" size="small">
                <a-table-column align="center" title="序号" data-index="id" />
                <a-table-column align="center" title="物品名称" data-index="name" />
                <a-table-column align="center" title="操作人" data-index="personId" />
                <a-table-column align="center" title="操作类型">
                    <template #default="{record}">
                        <a-tag :color="record.type == 0 ? 'blue' : 'green'">{{ record.type == 0 ? '入库' : '出库' }}</a-tag>
                    </template>
                </a-table-column>
                <a-table-column align="center" title="操作数量" data-index="num" />
                <a-table-column align="center" title="操作时间" data-index="time" />
                <a-table-column align="center" title="操作备注" data-index="remark" />
                <a-table-column align="center" title="操作">
                    <template #default="{record}">
                        <a-button type="link" danger @click="deleteRecord(record.id)">删除</a-button>
                    </template>
                </a-table-column>
            </a-table>
        </div>
        <div class="page">
            <a-pagination
                v-model:current="searchForm.currentPage"
                v-model:pageSize="searchForm.pageSize"
                :pageSizeOptions="pageSizeOptions"
                show-size-changer
                :total="500"
                @showSizeChange="onShowSizeChange"
            />
        </div>
    </div>
</template>

<style scoped>
.container{
    width: 100%;
    height: 100%;
}
.form{
    width: 98%;
}
.table{
    width: 98%;
    margin-top: 20px;
}
.page{
    width: 98%;
    height: 50px;
    display: flex;
    align-items: center;
    justify-content: right;
}
</style>