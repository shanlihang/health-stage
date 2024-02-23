import{defineStore} from 'pinia'

export const useSystemStore = defineStore('sys',{
    state:() => {
        return{
            waterMark:'Ant Design Vue'
        }
    }
}) 