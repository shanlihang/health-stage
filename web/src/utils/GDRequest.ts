import axios from 'axios'

const GDService = axios.create({
    baseURL:'https://restapi.amap.com',
    timeout:50000
})

export default GDService