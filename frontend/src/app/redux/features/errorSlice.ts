"use client"

import { createSlice } from '@reduxjs/toolkit'

export type ErrorSatate = {
    error: null |  {
        code: string,
        message: string
    },
}

const initialState = {
    error: null,
}

const errorSlice = createSlice({
    name: 'error',
    initialState,
    reducers: {
        setError: (state, action) => {
            console.log("there is an error")
            state.error = action.payload
        },
        clearError: (state) => {
            console.log("removing an error")
            state.error = null
        },
    },
})

export const { setError, clearError } = errorSlice.actions

export default errorSlice.reducer