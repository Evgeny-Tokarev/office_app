import Axios from 'axios'
import Logger from "@/app/helpers/logger"

const axios = Axios.create({
    baseURL: 'http://office.local:8000'
})
axios.interceptors.request.use(config => {
    if (Logger.isDebugMode()) {
        console.groupCollapsed(`API request: ${config.url}`)
        console.log(`Method: ${config.method}`)
        console.log(`URL: ${config.url}`)
        console.log(`Query: ${JSON.stringify(config.params)}`)
        console.log(`Body: ${JSON.stringify(config.data)}`)
        console.groupEnd()
    }

    if (!config.headers?.noToken) {
        const token = localStorage.getItem('officeAppToken')

        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`
        }
    }

    delete config.headers?.noToken

    return config
}, error => error)

axios.interceptors.response.use(
    response => {
        if (Logger.isDebugMode()) {
            console.groupCollapsed(`API response: ${response.config.url}`)
            console.log(`Method: ${response.config.method}`)
            console.log(`Url: ${response.config.url}`)
            console.log(`Status: ${response.status}`)
            console.log(response)
            console.groupEnd()
        }

        return response
    },
    error => {
        console.log(error)
        return Promise.reject(error)
    }
)
export default axios
