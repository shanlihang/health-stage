import service from '../../utils/request'

//插入一个新物品
export function insertGoods(data){
    return service({
        url:'/store/addGoods',
        method:'post',
        data
    })
}

//获取物品列表
export function selectGoodsList(){
    return service({
        url:'/store/goods',
        method:'get'
    })
}

//删除一个物品
export function deleteGoods(data){
    return service({
        url:'/store/deleteGoods',
        method:'delete',
        data
    })
}