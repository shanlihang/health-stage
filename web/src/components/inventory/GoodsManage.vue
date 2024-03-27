<script setup lang="ts">
import {insertGoods,selectGoodsList,deleteGoods} from '../../api/inventory'
import {ref,reactive,onMounted} from 'vue'
import service from '../../utils/request';

//搜索表单接口定义
interface SearchForm{
    name:string,
    remark:string,
    currentPage:number,
    pageSize:number
}

//新增/编辑 物品接口定义
interface Goods{
    name:string,
    num:number,
    divide:string
    remark:string
}

//搜索框表单
const searchForm = reactive<SearchForm>({
    name:'',
    remark:'',
    currentPage:1,
    pageSize:20
})

//弹窗是否居中显示
const centered = ref<boolean>(true)

//新增物品表单
const addGoods = reactive<Goods>({
    name:'',
    num:0,
    divide:'',
    remark:''
})

//编辑物品表单
const updateGoods = reactive<Goods>({
    name:'',
    num:0,
    divide:'',
    remark:''
})

//添加物品弹窗的状态
const openAddGoods = ref<boolean>(false)

//编辑物品弹窗的状态
const openUpdateGoods = ref<boolean>(false)

//表格边框显示状态
const bordered = ref<boolean>(true)

//是否显示表格自带的分页
const table_page = ref<boolean>(false)

//分页每页显示数量的选项
const pageSizeOptions = ref<Array<string>>(['10','20','50','100'])

//表格数据
const data = reactive([])

//搜索框表单提交成功
const handleFinish = () => {}

//搜索框表单提交失败
const handleFinishFailed = () => {}

//分页数据改变时触发
const onShowSizeChange = () => {}

//编辑按钮
const editGoods = (record:Goods) => {
    openUpdateGoods.value = true;
    updateGoods.name = record.name;
    updateGoods.num = record.num;
    updateGoods.divide = record.divide;
    updateGoods.remark = record.remark
}

//确认删除
const confirmDelete = (record) => {
    deleteGoods({ID:record.ID}).then(
        res => {
            console.log(res);
            
        }
    )
}

//点击新增物品按钮事件
const addNewGoods = () => {
    openAddGoods.value = true
}

//添加弹窗中的确认
const addGoodsOk = () => {
    insertGoods(addGoods).then(res => {
        console.log(res);
    })
    if(addGoods.num != undefined){
        console.log('新增+入库');
    }else{
        console.log('新增');
    }
    resetAddForm();
    openAddGoods.value = false
}

//添加弹窗中的取消
const addGoodsCancel = () => {
    resetAddForm();
    openAddGoods.value = false;
    
}

//编辑弹窗中的确认
const updateGoodsOk = () => {
    console.log('编辑完成');
    openUpdateGoods.value = false;
}

//编辑弹窗中的取消
const updateGoodsCancel = () => {
    console.log('编辑取消');
}

const resetAddForm = () => {
    addGoods.num = 0;
    addGoods.name = '';
    addGoods.divide = '';
    addGoods.remark = '';
}

const initData = () => {
  selectGoodsList(searchForm).then(res => {
    res.data.list.forEach(element => {
      data.push(element)
    });
  })
}

onMounted(() => {
  initData()
})
</script>

<template>
    <div class="section">
        <div class="form">
            <a-form layout="inline" :model="searchForm" @finish="handleFinish" @finishFailed="handleFinishFailed">
                <a-form-item label="物品名称">
                    <a-input v-model:value="searchForm.name" placeholder="请输入物品名称" />
                </a-form-item>
                <a-form-item label="物品备注">
                    <a-input v-model:value="searchForm.remark" placeholder="请输入物品备注" />
                </a-form-item>
                <a-form-item>
                    <a-button type="primary">搜索</a-button>
                    <a-button style="margin-left: 15px;">重置</a-button>
                </a-form-item>
            </a-form>
        </div>
        <div class="add">
            <a-button type="primary" @click="addNewGoods">新增物品</a-button>
        </div>
        <div class="table">
            <a-table :data-source="data" :bordered="bordered" :pagination="table_page" size="small">
                <a-table-column align="center" title="序号" data-index="ID" />
                <a-table-column align="center" title="物品名称" data-index="Name" />
                <a-table-column align="center" title="物品余量">
                    <template #default="{record}">
                        {{ record.Nums }} ({{ record.Divide }}) 
                    </template>
                </a-table-column>
                <a-table-column align="center" title="物品备注" data-index="Remark" />
                <a-table-column align="center" title="操作">
                    <template #default="{record}">
                        <a-button type="link" primary @click="editGoods(record)">编辑</a-button>
                        <a-popconfirm title="确认删除该记录吗？" ok-text="确认" cancel-text="取消" @confirm="confirmDelete(record)">
                            <a-button type="link" danger>删除</a-button>
                        </a-popconfirm>
                        
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

        <a-modal v-model:open="openAddGoods" title="新增物品" @ok="addGoodsOk" @cancel="addGoodsCancel" :centered="centered" cancelText="取消" okText="确认">
            <a-form :model="addGoods">
                <a-form-item label="物品名称">
                    <a-input v-model="addGoods.name"></a-input>
                </a-form-item>
                <a-form-item label="物品数量">
                    <a-input-number v-model="addGoods.num" :min="0"></a-input-number>
                </a-form-item>
                <a-form-item label="物品单位">
                    <a-input v-model="addGoods.divide"></a-input>
                </a-form-item>
                <a-form-item label="物品备注">
                    <a-input v-model="addGoods.remark"></a-input>
                </a-form-item>
            </a-form>
        </a-modal>

        <a-modal v-model:open="openUpdateGoods" title="编辑物品信息" @ok="updateGoodsOk" @cancel="updateGoodsCancel" :centered="centered" cancelText="取消" okText="确认">
            <a-form :model="updateGoods">
                <a-form-item label="物品名称">
                    <a-input v-model="updateGoods.name"></a-input>
                </a-form-item>
                <a-form-item label="物品数量">
                    <a-input-number v-model="updateGoods.num" :min="0"></a-input-number>
                </a-form-item>
                <a-form-item label="物品单位">
                    <a-input v-model="updateGoods.divide"></a-input>
                </a-form-item>
                <a-form-item label="物品备注">
                    <a-input v-model="updateGoods.remark"></a-input>
                </a-form-item>
            </a-form>
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