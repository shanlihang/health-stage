<script setup lang="ts">
import {ref,reactive} from 'vue'


//搜索表单接口定义
interface SearchForm{
    name:string,
    remark:string,
    currentPage:number,
    pageSize:number
}

//搜索框表单
const searchForm = reactive<SearchForm>({
    name:'',
    remark:'',
    currentPage:1,
    pageSize:20
})

//物品接口定义
interface Goods{
    name:string,
    num:number,
    uint:string
    remark:string
}

//选中的商品名称
const tempName = ref<string>('')

//选中商品库存
const tempNum = ref<number>()

//入库数量
const insertNum = ref<number>()

//出库数量
const outsertNum = ref<number>()

//出/入库原因
const reason = ref<string>('')

//入库弹窗显示状态
const openInsertGoods = ref<boolean>(false)

//出库弹窗显示状态
const openOutsertGoods = ref<boolean>(false)

//弹窗是否居中显示
const centered = ref<boolean>(true)

//表格边框显示状态
const bordered = ref<boolean>(true)

//是否显示表格自带的分页
const table_page = ref<boolean>(false)

//分页每页显示数量的选项
const pageSizeOptions = ref<Array<string>>(['10','20','50','100'])

//表格数据
const data = [
  {
    id: 1,
    name: 'John Brown',
    num: 32,
    unit:'个',
    remark: 'New York No. 1 Lake Park',
  },
  {
    id: 2,
    name: 'John Brown',
    num: 32,
    unit:'个',
    remark: 'New York No. 1 Lake Park',
  },
  {
    id: 3,
    name: 'John Brown',
    num: 32,
    unit:'个',
    remark: 'New York No. 1 Lake Park',
  },
  {
    id: 4,
    name: 'John Brown',
    num: 32,
    unit:'个',
    remark: 'New York No. 1 Lake Park',
  },
  {
    id: 5,
    name: 'John Brown',
    num: 32,
    unit:'个',
    remark: 'New York No. 1 Lake Park',
  }
];

//insert入库事件
const insert = (record:Goods) => {
    tempName.value = record.name
    tempNum.value = record.num
    openInsertGoods.value = true
    
}

//outsert出库事件
const outsert = (record:Goods) => {
    tempName.value = record.name
    tempNum.value = record.num
    openOutsertGoods.value = true
}

//搜索框表单提交成功
const handleFinish = () => {}

//搜索框表单提交失败
const handleFinishFailed = () => {}

//分页数据改变时触发
const onShowSizeChange = () => {}

//入库弹窗确认
const insertGoodsOk = () => {
    openInsertGoods.value = false
}

//入库弹窗取消
const insertGoodsCancel = () => {
    openInsertGoods.value = false
}

//出库弹窗确认
const outsertGoodsOk = () => {
    openOutsertGoods.value = false
}

//出库弹窗取消
const outsertGoodsCancel = () => {
    openOutsertGoods.value = false
}
</script>

<template>
    <div class="section">
        <div class="form">
            <a-form layout="inline" :model="searchForm" @finish="handleFinish" @finishFailed="handleFinishFailed">
                <a-form-item label="物品名称">
                    <a-input v-model:value="searchForm.name" placeholder="请输入物品名称" />
                </a-form-item>
                <a-form-item label="操作人">
                    <a-input v-model:value="searchForm.remark" placeholder="请输入物品备注" />
                </a-form-item>
                <a-form-item>
                    <a-button type="primary">搜索</a-button>
                    <a-button style="margin-left: 15px;">重置</a-button>
                </a-form-item>
            </a-form>
        </div>
        <div class="table">
            <a-table :data-source="data" :bordered="bordered" :pagination="table_page" size="small">
                <a-table-column align="center" title="序号" data-index="id" />
                <a-table-column align="center" title="物品名称" data-index="name" />
                <a-table-column align="center" title="物品余量">
                    <template #default="{record}">
                        {{ record.num }} ({{ record.unit }}) 
                    </template>
                </a-table-column>
                <a-table-column align="center" title="物品备注" data-index="remark" />
                <a-table-column align="center" title="操作">
                    <template #default="{record}">
                        <a-button type="link" primary @click="insert(record)">入库</a-button>
                        <a-button type="link" danger @click="outsert(record)">出库</a-button>
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

        <a-modal v-model:open="openInsertGoods" title="入库" @ok="insertGoodsOk" @cancel="insertGoodsCancel" :centered="centered" cancelText="取消" okText="确认">
            <div>
                <p>当前商品：{{ tempName }}</p>
                <p>当前库存：{{ tempNum }}</p>
                <p style="display: flex;align-items: center;">
                    入库数量：<a-input-number v-model="insertNum" :min="0" style="width: 50%;" placeholder="请输入入库数量"></a-input-number>
                </p>
                <p>入库备注：<a-input v-model="reason" :min="0" style="width: 50%;" placeholder="请输入入库备注" /></p>
                
            </div>
        </a-modal>

        <a-modal v-model:open="openOutsertGoods" title="出库" @ok="outsertGoodsOk" @cancel="outsertGoodsCancel" :centered="centered" cancelText="取消" okText="确认">
            <div>
                <p>当前商品：{{ tempName }}</p>
                <p>当前库存：{{ tempNum }}</p>
                <p style="display: flex;align-items: center;">
                    出库数量：<a-input-number v-model="outsertNum" :min="0" style="width: 50%;" placeholder="请输入出库数量"></a-input-number>
                </p>
                <p>出库备注：<a-input-number v-model="reason" :min="0" style="width: 50%;" placeholder="请输入出库备注"></a-input-number></p>
            </div>
            
        </a-modal>
    </div>
</template>

<style scoped>
.section{
    width: 100%;
    height: 100%;
}
.form{
    width: 98%;
    height: 50px;
    display: flex;
    align-items: center;
}
.add{
    width: 98%;
    height: 50px;
    display: flex;
    align-items: center;
}
.table{
    width: 98%;
    margin-top: 15px;
}
.page{
    width: 98%;
    height: 50px;
    display: flex;
    align-items: center;
    justify-content: right;
}
</style>