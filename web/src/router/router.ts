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
            path:'/menuList',
            component:() => import('../views/MenuList.vue')
        },
        {
            path:'/setting',
            component:() => import('../views/SystemSetting.vue')
        }
    ]
})

export default router