import axios from 'axios'

const service = axios.create({
    baseURL:'http://localhost:9090',
    timeout:50000
})

export default service