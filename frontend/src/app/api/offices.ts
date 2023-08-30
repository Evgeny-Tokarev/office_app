import {type Office} from "@/app/models";
import axios from "@/app/api/axios";
import {createAsyncThunk} from '@reduxjs/toolkit'

export default {
    getOffices: async function (): Promise<{
        status: number; data: Office[]
    }> {
        try {
            const resp = await axios.get(`/offices`)
            return {
                data: resp.data, status: resp.status
            }
        } catch (err: any) {
            return err.response.status
        }
    }
}
