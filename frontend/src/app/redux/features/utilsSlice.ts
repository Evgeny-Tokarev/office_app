"use client"

import {createSlice, PayloadAction} from '@reduxjs/toolkit'

export type UtilsState = {
    loading: boolean
    error: null | {
        code: string,
        message: string
    },
    infoState: null | {
        title: string,
        text: string
    },
}

const initialState: UtilsState = {
    error: null,
    loading: false,
    infoState: null,
}

const utilsSlice = createSlice({
    name: 'error',
    initialState,
    reducers: {
        setError: (state, action) => {
            state.error = action.payload
        },
        clearError: (state) => {
            state.error = null
        },
        setLoading: (state, action: PayloadAction<boolean>) => {
            state.loading = action.payload
        },
        setInfo: (state, action) => {
            state.infoState = action.payload
        },
        clearInfo: (state, action) => {
            state.infoState = null
        }
    },
})

export const {setError, clearError, setLoading, setInfo, clearInfo} = utilsSlice.actions

export default utilsSlice.reducer