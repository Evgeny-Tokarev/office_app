"use client"

import {createAsyncThunk, createSlice, type SerializedError} from "@reduxjs/toolkit"
import {type Office} from '@/app/models'
import api from "@/app/api"

export type OfficesState = {
    offices: Office[],
    currentOffice: Office | null,
    loading: boolean,
    infoState: null | {
        title: string,
        text: string
    },
    error: null | {
        code: string,
        message: string
    },
};

const initialState = {
    offices: [], currentOffice: null, loading: false, error: null, infoState: null
} as OfficesState


export const fetchOffices = createAsyncThunk<Office[], void, {
    rejectValue: SerializedError;
}>('offices/fetchOffices', async (_, {rejectWithValue}) => {
    try {
        const response = await api.officesApi.getOffices()
        if (response.data?.status === 200) return response.data.data
        return rejectWithValue({
            code: response.error?.code,
            message: response.error?.message
        } as SerializedError)
    } catch (err) {
        return rejectWithValue(err as SerializedError)
    }
})
export const fetchOffice = createAsyncThunk<Office, number, {
    rejectValue: SerializedError;
}>('offices/fetchOffice', async (id, {rejectWithValue}) => {
    try {
        const response = await api.officesApi.getOffice(id)
        if (response.data?.status === 200) return response.data.data
        return rejectWithValue({
            code: response.error?.code,
            message: response.error?.message
        } as SerializedError)
    } catch (err) {
        return rejectWithValue(err as SerializedError)
    }
})
export const saveImage = createAsyncThunk<boolean, { id: number, image: File }, {
    rejectValue: SerializedError;
}>('offices/saveImage', async ({id, image}, {rejectWithValue}) => {
    try {
        const response = await api.officesApi.saveImage(id, image)
        if (response.data?.status === 200) return response.data.data
        return rejectWithValue({
            code: response.error?.code,
            message: response.error?.message
        } as SerializedError)
    } catch (err) {
        return rejectWithValue(err as SerializedError)
    }
})


export const updateOffice = createAsyncThunk<Office, Office, {
    rejectValue: SerializedError;
}>('offices/updateOffice', async (office, {rejectWithValue}) => {
    try {
        office.id = +office.id
        const response = await api.officesApi.updateOffice(office)
        if (response.data?.status === 200) return response.data.data
        return rejectWithValue({
            code: response.error?.code,
            message: response.error?.message
        } as SerializedError)
    } catch (err) {
        return rejectWithValue(err as SerializedError)
    }
})
export const deleteOffice = createAsyncThunk<void, number, {
    rejectValue: SerializedError;
}>('offices/deleteOffice', async (id, {rejectWithValue}) => {
    try {
        const response = await api.officesApi.deleteOffice(id)
        if (response.data?.status === 200) return
        return rejectWithValue({
            code: response.error?.code,
            message: response.error?.message
        } as SerializedError)
    } catch (err) {
        return rejectWithValue(err as SerializedError)
    }
})

export const createOffice = createAsyncThunk<Office, Office, {
    rejectValue: SerializedError;
}>('offices/createOffice', async (office, {rejectWithValue}) => {
    try {
        const response = await api.officesApi.createOffice(office)
        if (response.data?.status && Math.floor(response.data.status / 100) === 2) return response.data.data
        return rejectWithValue({
            code: response.error?.code,
            message: response.error?.message
        } as SerializedError)
    } catch (err) {
        return rejectWithValue(err as SerializedError)
    }
})
export const offices = createSlice({
    name: "offices", initialState, reducers: {}, extraReducers: builder => {
        builder
            .addCase(fetchOffices.pending, state => {
                state.loading = true
                state.error = null
            })
            .addCase(fetchOffices.fulfilled, (state, action) => {
                state.loading = false
                state.offices = action.payload ?? []
                state.error = null
            })
            .addCase(fetchOffices.rejected, (state, action) => {
                state.loading = false
                state.offices = []
                state.error = {
                    message: action.payload?.message || action.error.message || 'Unknown error occurred',
                    code: action.payload?.code || action.error.code || 'Error'
                }
            }).addCase(fetchOffice.pending, state => {
            state.loading = true
            state.error = null
        })
            .addCase(fetchOffice.fulfilled, (state, action) => {
                state.loading = false
                state.currentOffice = action.payload
                state.error = null
            })
            .addCase(fetchOffice.rejected, (state, action) => {
                state.loading = false
                state.currentOffice = null
                state.error = {
                    message: action.payload?.message || action.error.message || 'Unknown error occurred',
                    code: action.payload?.code || action.error.code || 'Error'
                }
            })
            .addCase(updateOffice.pending, state => {
                state.loading = true
                state.error = null
            })
            .addCase(updateOffice.fulfilled, (state) => {
                state.loading = false
                state.error = null
            }).addCase(updateOffice.rejected, (state, action) => {
            state.loading = false
            state.error = {
                message: action.payload?.message || action.error.message || 'Unknown error occurred',
                code: action.payload?.code || action.error.code || 'Error'
            }
        }).addCase(createOffice.pending, state => {
            state.loading = true
            state.infoState = null
            state.error = null
        })
            .addCase(createOffice.fulfilled, (state, action) => {
                state.loading = false
                state.infoState = {
                    title: "New office created",
                    text: `Office name: ${action.payload.name}`
                }
                state.error = null
            }).addCase(createOffice.rejected, (state, action) => {
            state.loading = false
            state.infoState = null
            state.error = {
                message: action.payload?.message || action.error.message || 'Unknown error occurred',
                code: action.payload?.code || action.error.code || 'Error'
            }
        }).addCase(deleteOffice.pending, state => {
            state.loading = true
            state.error = null
            state.infoState = null
        })
            .addCase(deleteOffice.fulfilled, (state) => {
                state.loading = false
                state.error = null
                state.infoState = {
                    title: "Office was deleted",
                    text: ""
                }
            }).addCase(deleteOffice.rejected, (state, action) => {
            state.loading = false
            state.infoState = null
            state.error = {
                message: action.payload?.message || action.error.message || 'Unknown error occurred',
                code: action.payload?.code || action.error.code || 'Error'
            }
        })
    },
})

export default offices.reducer
