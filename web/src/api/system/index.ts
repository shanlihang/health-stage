import service from '../../utils/request'

//获取系统菜单
export function getMenuList(){
    return service({
        url:'/system/menu',
        method:'get'
    })
}