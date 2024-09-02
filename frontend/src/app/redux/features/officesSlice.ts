"use client"

import {createAsyncThunk, createSlice, type SerializedError} from "@reduxjs/toolkit"
import {type Office} from '@/app/models'
import api from "@/app/api"
import {setError, setInfo, setLoading} from "@/app/redux/features/utilsSlice"
import type {AppDispatch, RootState} from "@/app/redux/store"

export type OfficesState = {
    offices: Office[],
    currentOffice: Office | null,
};

const initialState = {
    offices: [], currentOffice: null, loading: false, error: null, infoState: null
} as OfficesState


export const fetchOffices = createAsyncThunk<Office[], void, {
    rejectValue: SerializedError;
    state: RootState;
    dispatch: AppDispatch
}>('offices/fetchOffices', async (_, {dispatch, rejectWithValue}) => {
    try {
        dispatch(setLoading(true))
        const response = await api.officesApi.getOffices()
        dispatch(setLoading(false))
        if (response.data?.status === 200) return response.data.data
        dispatch(setError({
            code: response.error?.code,
            message: response.error?.message
        }))
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
    state: RootState;
    dispatch: AppDispatch
}>('offices/fetchOffice', async (id, {dispatch, rejectWithValue}) => {
    try {
        dispatch(setLoading(true))
        const response = await api.officesApi.getOffice(id)
        dispatch(setLoading(false))
        if (response.data?.status === 200) return response.data.data
        dispatch(setError({
            code: response.error?.code,
            message: response.error?.message
        }))
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
    state: RootState;
    dispatch: AppDispatch
}>('offices/saveImage', async ({id, image}, {dispatch, rejectWithValue}) => {
    try {
        dispatch(setLoading(true))
        const response = await api.officesApi.saveImage(id, image)
        dispatch(setLoading(false))
        if (response.data?.status === 200) {
            dispatch(setInfo({
                title: "Image was saved",
                text: ""
            }))
            return response.data.data
        }
        dispatch(setError({
            code: response.error?.code,
            message: response.error?.message
        }))
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
    state: RootState;
    dispatch: AppDispatch
}>('offices/updateOffice', async (office, {dispatch, rejectWithValue}) => {
    try {
        dispatch(setLoading(true))
        office.id = +office.id
        const response = await api.officesApi.updateOffice(office)
        dispatch(setLoading(false))
        if (response.data?.status === 200) {
            dispatch(setInfo({
                title: "Office was updated",
                text: `Office name: ${response.data.data.name}`
            }))
            return response.data.data
        }
        dispatch(setError({
            code: response.error?.code,
            message: response.error?.message
        }))
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
    state: RootState;
    dispatch: AppDispatch
}>('offices/deleteOffice', async (id, {dispatch, rejectWithValue}) => {
    try {
        dispatch(setLoading(true))
        const response = await api.officesApi.deleteOffice(id)
        dispatch(setLoading(false))
        if (response.data?.status === 200) {
            dispatch(setInfo({
                title: "Office was deleted",
                text: ""
            }))
            return
        }
        dispatch(setError({
            code: response.error?.code,
            message: response.error?.message
        }))
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
    state: RootState;
    dispatch: AppDispatch
}>('offices/createOffice', async (office, {dispatch, rejectWithValue}) => {
    try {
        dispatch(setLoading(true))
        const response = await api.officesApi.createOffice(office)
        dispatch(setLoading(false))
        if (response.data?.status && Math.floor(response.data.status / 100) === 2) {
            dispatch(setInfo({
                title: "New office created",
                text: `Office name: ${response.data.data.name}`
            }))
            return response.data.data
        }
        dispatch(setError({
            code: response.error?.code,
            message: response.error?.message
        }))
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
            .addCase(fetchOffices.fulfilled, (state, action) => {
                state.offices = action.payload ?? []
            })
            .addCase(fetchOffices.rejected, (state, action) => {
                state.offices = []
            })
            .addCase(fetchOffice.fulfilled, (state, action) => {
                state.currentOffice = action.payload
            })
            .addCase(fetchOffice.rejected, (state, action) => {
                state.currentOffice = null
            })
    },
})

export default offices.reducer
