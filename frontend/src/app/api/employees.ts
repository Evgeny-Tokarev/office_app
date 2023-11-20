import {
    type Employee,
    type RestApiResponse
} from "@/app/models"
import axios from "@/app/api/axios"
import axiosErrorHandler
    from "@/app/helpers/axiosErrorHandler"

export default {
    getEmployees: async function (office_id: number): Promise<RestApiResponse<{
        status: number,
        data: Employee[]
    }>> {
        try {
            const resp = await axios.get(`/employees?office_id=${office_id}`)
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
    updateEmployee: async function (office: Employee): Promise<RestApiResponse<{
        status: number,
        data: Employee
    }>> {
        try {
            const resp = await axios.put(`/employees`, {
                ...office
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
    deleteEmployee: async function (id: number): Promise<RestApiResponse<{
        status: number,
        data: Employee
    }>> {
        try {
            const resp = await axios.delete(`/employees/${id}`)
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
    createEmployee: async function (employee: Employee): Promise<RestApiResponse<{
        status: number,
        data: Employee
    }>> {
        try {
            const resp = await axios.post(`/employees`, {
                name: employee.name, age: +employee.age, office_id: +employee.office_id
            })
            return {
                data: {
                    data: resp.data,
                    status: resp.status
                }, error: undefined
            };
        } catch (err: unknown) {
            return axiosErrorHandler(err)
        }
    },
    saveImage: async function (id: number, image: File): Promise<RestApiResponse<{
        status: number,
        data: boolean
    }>> {
        try {
            const formData = new FormData()
            formData.append("image", image)
            const resp = await axios.post(`/employees/${id}/image`, formData, {
                headers: {
                    "Content-Type": "multipart/form-data"
                }
            })
            return {
                data: {
                    data: resp.data,
                    status: resp.status
                }, error: undefined
            };
        } catch (err: unknown) {
            return axiosErrorHandler(err)
        }
    }

}
