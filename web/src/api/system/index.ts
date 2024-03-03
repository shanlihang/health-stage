import service from '../../utils/request'

//测试跨域
export function getCurrentAddress(){
    return service({
        url:'/system/getMenu',
        method:'get',
    })
}