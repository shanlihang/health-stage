import {createRouter,createWebHistory} from 'vue-router'
import HelloWorld from '../components/HelloWorld.vue'
import MainView from '../views/MainView.vue'
const router = createRouter({
    history:createWebHistory(),
    routes:[
        {
            path:'/',
            component:HelloWorld
        }
    ]
})

export default router