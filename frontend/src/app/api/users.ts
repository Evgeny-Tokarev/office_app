import axios from "@/app/api/axios"
import axiosErrorHandler from "@/app/helpers/axiosErrorHandler"
import type {RestApiResponse, User} from "@/app/models"

export default {

    login: async function (userName: string, email: string, password: string): Promise<RestApiResponse<{
        status: number,
        data: { user: User, token: string }
    }>> {
        try {
            const resp = await axios.post('/user/login', {
                password: password,
                email: email,
                name: userName
            })
            return {
                data: {
                    data: resp.data,
                    status: resp.status
                }, error: undefined
            }
        } catch (err: unknown) {
            return axiosErrorHandler(err)
        }
    },
    getCurrentUser: async function (token: string): Promise<RestApiResponse<{
        status: number,
        data: User
    }>> {
        try {
            const resp = await axios.get('/user/current', {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            })
            return {
                data: {
                    data: resp.data,
                    status: resp.status
                }, error: undefined
            }
        } catch (err: unknown) {
            return axiosErrorHandler(err)
        }
    },
    register: async function (userName: string, email: string, password: string, role: string): Promise<RestApiResponse<{
        status: number,
        data: { user: User, token: string }
    }>> {
        try {
            const resp = await axios.post('/user/register', {
                password: password,
                email: email,
                name: userName,
                role: role
            })
            return {
                data: {
                    data: resp.data,
                    status: resp.status
                }, error: undefined
            }
        } catch (err: unknown) {
            return axiosErrorHandler(err)
        }
    }
}