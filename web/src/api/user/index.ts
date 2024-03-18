import service from '../../utils/request'

//获取当前用户信息
export function getCurrentUser(){
    return service({
        url:'/user/current',
        method:'get'

    })
} 


//获取所有用户信息
export function getAllUser(){
    return service({
        url:'/user/all',
        method:'get'
    })
} 