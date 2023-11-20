import {
    type Office,
    type RestApiResponse
} from "@/app/models"
import axios from "@/app/api/axios"
import axiosErrorHandler
    from "@/app/helpers/axiosErrorHandler"

export default {
    getOffice: async function (id: number): Promise<RestApiResponse<{
        status: number,
        data: Office
    }>> {
        try {
            const resp = await axios.get(`/offices/${id}`)
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
    getOffices: async function (): Promise<RestApiResponse<{
        status: number,
        data: Office[]
    }>> {
        try {
            const resp = await axios.get(`/offices`)
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
    updateOffice: async function (office: Office): Promise<RestApiResponse<{
        status: number,
        data: Office
    }>> {
        try {
            const resp = await axios.put(`/offices`, {
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
    deleteOffice: async function (id: number): Promise<RestApiResponse<{
        status: number,
        data: Office
    }>> {
        try {
            const resp = await axios.delete(`/offices/${id}`)
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
    createOffice: async function (office: Office): Promise<RestApiResponse<{
        status: number,
        data: Office
    }>> {
        try {
            const resp = await axios.post(`/offices`, {
                name: office.name, address: office.address
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
            const resp = await axios.post(`/offices/${id}/image`, formData, {
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
