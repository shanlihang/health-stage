<script setup lang="ts">
import { ref,reactive, onMounted, onUnmounted } from "vue";
import {getInputTips} from '../../api/gd/index.ts'

//输入提示接口
interface InputTips{
  id?:string,
  adcode?:string,
  address?:string,
  city?:Array<string>,
  district?:string,
  location?:string,
  name?:string,
  typecode?:string
}

//合作社区列表管理
interface TableData{
  id:number,
  POI_id:string,
  POI_name:string,
  POI_district:string,
  POI_adcode:string,
  POI_location:string,
  POI_address:string
}

//搜索表单接口定义
interface SearchForm{
    name:string,
    district:string,
    detail_address:string,
    currentPage:number,
    pageSize:number
}

//搜索框表单
const searchForm = reactive<SearchForm>({
    name:'',
    district:'',
    detail_address:'',
    currentPage:1,
    pageSize:20
})

//表格是否开启边框
const bordered = ref<boolean>(true)

//表格是否使用自带的分页
const table_page = ref<boolean>(false)

//表格数据
const tableData = reactive<Array<TableData>>([
  {id:1,POI_id:'6545',POI_name:'每年三个v你',POI_district:'山东省济南市历城区',POI_adcode:'safasfa',POI_address:'asfasfasfa',POI_location:'sfafa'},
  {id:2,POI_id:'6545',POI_name:'每年三个v你',POI_district:'山东省济南市历城区',POI_adcode:'safasfa',POI_address:'asfasfasfa',POI_location:'sfafa'},
  {id:3,POI_id:'6545',POI_name:'每年三个v你',POI_district:'山东省济南市历城区',POI_adcode:'safasfa',POI_address:'asfasfasfa',POI_location:'sfafa'},
  {id:4,POI_id:'6545',POI_name:'每年三个v你',POI_district:'山东省济南市历城区',POI_adcode:'safasfa',POI_address:'asfasfasfa',POI_location:'sfafa'},
  {id:5,POI_id:'6545',POI_name:'每年三个v你',POI_district:'山东省济南市历城区',POI_adcode:'safasfa',POI_address:'asfasfasfa',POI_location:'sfafa'}
])

//输入提示
const tips = reactive<Array<InputTips>>([])

//定义中转变量接收选中的地址
const tempAddress = ref<InputTips>({})

//关键字搜索框
const address = ref<string>("");

//分页每页显示数量的选项
const pageSizeOptions = ref<Array<string>>(['10','20','50','100'])

//新增社区弹窗
const addCommunityModal = ref<boolean>(false);

//社区详情抽屉
const detailDrawerFlag = ref<boolean>(false);

//打开社区详情按钮点击事件
const openDetail = () => {
  detailDrawerFlag.value = true;
}

//分页数据改变时触发
const onShowSizeChange = () => {}

//搜索提示
const searchTips = (keywords:string) => {
  tips.length = 0;
  getInputTips({
    key:'0b195d069492038cc8d9086e25346fb7',
    keywords:keywords
  }).then(res => {
    res.data.tips.forEach(element => {
      tips.push(element)
    });
  }).catch(err => {
    console.log('错误信息' + err);
  })
}

//搜索表单提交成功的回调
const handleFinish = () => {}

//搜索表单提交失败的回调
const handleFinishFailed = () => {}

//新增社区的弹窗输入事件
const change = () => {
  searchTips(address.value)
}

//新增社区按钮点击事件
const addCommunity = () => {
  addCommunityModal.value = true;
};

//新增社区弹窗确认事件
const addCommunityOk = () => {
  addCommunityModal.value = false;
  console.log(tempAddress.value.address);
  clearAddModal();
};

//新增社区弹窗取消事件
const addCommunityCancel = () => {
  addCommunityModal.value = false;
  clearAddModal();
};

//关闭抽屉
const closeDetailDrawer = () => {
  detailDrawerFlag.value = false;
}

//清空新增社区弹窗中的内容
const clearAddModal = () => {
  address.value = '';
  tempAddress.value = {};
  tips.length = 0;
}

//选中地址后的操作
const chooseAddress = (item:InputTips) => {
  tempAddress.value = item;
}

onMounted(() => {});

onUnmounted(() => {});

</script>

<template>
  <div class="container">
    <div class="form">
      <a-form layout="inline" :model="searchForm" @finish="handleFinish" @finishFailed="handleFinishFailed">
                <a-form-item label="物品名称">
                    <a-input v-model:value="searchForm.name" placeholder="请输入物品名称" />
                </a-form-item>
                <a-form-item label="所属区域">
                    <a-input v-model:value="searchForm.district" placeholder="请输入所属区域" />
                </a-form-item>
                <a-form-item label="详细地址">
                    <a-input v-model:value="searchForm.detail_address" placeholder="请输入详细地址" />
                </a-form-item>
                <a-form-item>
                    <a-button type="primary">搜索</a-button>
                    <a-button style="margin-left: 15px;">重置</a-button>
                </a-form-item>
            </a-form>
    </div>
    <div class="btn"><a-button type="primary" @click="addCommunity">新增社区</a-button></div>
    <div class="table">
      <a-table :data-source="tableData" :bordered="bordered" :pagination="table_page" size="small">
                <a-table-column align="center" title="序号" data-index="id" />
                <a-table-column align="center" title="POI编号" data-index="POI_id" />
                <a-table-column align="center" title="单位名称" data-index="POI_name" />
                
                <a-table-column align="center" title="所属区域" data-index="POI_district" />
                <a-table-column align="center" title="详细地址" data-index="POI_address" />

                <a-table-column align="center" title="操作">
                    <template #default="{record}">
                        <a-button type="link" primary @click="openDetail">详情</a-button>
                        <a-button type="link" style="color: orange;">编辑</a-button>
                        <a-popconfirm title="确认删除该记录吗？" ok-text="确认" cancel-text="取消" @confirm="deleteGoods(record)">
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

    <a-drawer title="合作社区详情" placement="bottom" :open="detailDrawerFlag" @close="closeDetailDrawer">
      描述列表
    </a-drawer>
    <a-modal
      v-model:open="addCommunityModal"
      title="新增合作社区"
      @ok="addCommunityOk"
      @cancel="addCommunityCancel"
      okText="确认"
      cancelText="取消"
    >
      <a-input
        v-model:value="address"
        placeholder="请输入名称/地址"
        @change="change"
      />
      <div class="selected">
            已选择位置：{{ tempAddress.name == undefined ? '暂未选择' : tempAddress.name}}
        </div>
      <div class="options">
          <a-list item-layout="horizontal" :data-source="tips" :bordered="true">
            <template #renderItem="{ item }">
              <a-list-item>
                <template #actions>
                    <a-button type="link" size="small" @click="chooseAddress(item)">选中</a-button>
                  </template>
                <a-list-item-meta
                  :description="item.district + ' ' + item.address"
                >
                  <template #title>
                    <span>{{ item.name }}</span>
                  </template>
                  
                </a-list-item-meta>
              </a-list-item>
            </template>
          </a-list>
      </div>
    </a-modal>
  </div>
</template>

<style scoped>
.options{
  width: 100%;
  max-height: 400px;
  overflow-y: auto;
}
.form{
  height: 50px;
  display: flex;
  align-items: center;
}
.btn{
  height: 50px;
  display: flex;
  align-items: center;
}
.selected{
  width: 100%;
  height: 50px;
  display: flex;
  align-items: center;
  font-weight: bold;
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
