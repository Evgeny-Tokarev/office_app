import Axios from 'axios';
import Logger from "@/app/helpers/logger";
const axios = Axios.create({
    baseURL: 'http://localhost:13005/'
})
// axios.interceptors.request.use(config => {
// 	if (Logger.isDebug()) {
// 		console.groupCollapsed(`API request: ${config.url}`)
// 		console.log(`Method: ${config.method}`)
// 		console.log(`URL: ${config.url}`)
// 		console.log(`Query: ${JSON.stringify(config.params)}`)
// 		console.log(`Body: ${JSON.stringify(config.data)}`)
// 		console.groupEnd()
// 	}
//
// 	return config
// }, error => error)
//
// axios.interceptors.response.use(
// 	response => {
// 		if (Logger.isDebug()) {
// 			console.groupCollapsed(`API response: ${response.config.url}`)
// 			console.log(`Method: ${response.config.method}`)
// 			console.log(`Url: ${response.config.url}`)
// 			console.log(`Status: ${response.status}`)
// 			console.log(response)
// 			console.groupEnd()
// 		}
//
// 		return response
// 	},
// 	error => {
// 		console.log(error.response)
// 		return Promise.reject(error)
// 	}
// )
export default axios
