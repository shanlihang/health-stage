import service from '../../utils/request'

//插入一个新物品
export function insertGoods(data){
    return service({
        url:'/store/addGoods',
        method:'POST',
        data:data
    })
}

//获取物品列表
export function selectGoodsList(){
    return service({
        url:'/store/goods',
        method:'get'
    })
}