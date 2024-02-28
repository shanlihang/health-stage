import GDService from '../../utils/GDRequest'
import GDRequest from '../../utils/GDRequest'


//获取当前定位
export function getCurrentAddress(params:Object){
    return GDRequest({
        url:'/v3/ip',
        method:'get',
        params:params
    })
}

//输入提示
export function getInputTips(params:Object){
    return GDService({
        url:'/v3/assistant/inputtips',
        method:'get',
        params:params
    })
}