import service from '../../utils/request'

//插入一个社区
export function insertCommunity(data){
    return service({
        url:'/medical/community',
        method:'post',
        data
    })
}

//获取社区列表
export function selectCommunity(){
    return service({
        url:'/medical/community',
        method:'get'
    })
}

//根据id删除社区
export function deleteCommunityById(params){
    return service({
        url:'/medical/community',
        method:'delete',
        params
    })
}