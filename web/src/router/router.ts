import {createRouter,createWebHistory} from 'vue-router'
import HomePage from '../views/HomePage.vue'
import LoginPage from '../views/LoginPage.vue'
const router = createRouter({
    history:createWebHistory(),
    routes:[
        {
            path:'/',
            component:HomePage
        },
        {
            path:'/login',
            component:LoginPage
        },
        {
            path:'/account',
            component:() => import('../views/AccountCenter.vue')
        },
        {
            path:'/data',
            component:() => import('../views/DataShow.vue')
        },
        {
            path:'/dataOther',
            component:() => import('../views/DataShowOther.vue')
        },
        {
            path:'/medical',
            component:() => import('../views/MedicalManage.vue')
        },
        {
            path:'/setting',
            component:() => import('../views/SystemSetting.vue')
        },
        {
            path:'/inventory',
            component:() => import('../views/InventoryManage.vue')
        }
    ]
})

export default router