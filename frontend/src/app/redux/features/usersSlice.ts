"use client"

import {createAsyncThunk, createSlice, type SerializedError} from "@reduxjs/toolkit"
import {type User} from '@/app/models'
import api from "@/app/api"

export type UserState = {
    currentUser: User | null,
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
    currentUser: null, loading: false, error: null, infoState: null
} as UserState


export const login = createAsyncThunk<{ user: User, token: string }, {
    userName: string,
    email: string,
    password: string
}, {
    rejectValue: SerializedError;
}>('users/login', async ({userName, email, password}, {rejectWithValue}) => {
    try {
        const response = await api.usersApi.login(userName, email, password)
        if (response.data?.status === 200) return response.data.data
        return rejectWithValue({
            code: response.error?.code,
            message: response.error?.message
        } as SerializedError)
    } catch (err) {
        return rejectWithValue(err as SerializedError)
    }
})

export const getCurrentUser = createAsyncThunk<User, string, {
    rejectValue: SerializedError;
}>('users/getCurrentUser', async (token, {rejectWithValue}) => {
    try {
        const response = await api.usersApi.getCurrentUser(token)
        console.log(response.data)
        if (response.data?.status === 200) return response.data.data
        return rejectWithValue({
            code: response.error?.code,
            message: response.error?.message
        } as SerializedError)
    } catch (err) {
        return rejectWithValue(err as SerializedError)
    }
})

export const users = createSlice({
    name: "users", initialState, reducers: {
        logOut: (state) => {
            state.currentUser = null
            state.error = null
            localStorage.removeItem('officeAppToken')
        },
    }, extraReducers: builder => {
        builder
            .addCase(login.pending, state => {
                state.loading = true
                state.error = null
            })
            .addCase(login.fulfilled, (state, action) => {
                state.loading = false
                state.currentUser = action.payload.user
                localStorage.setItem('officeAppToken', action.payload.token)
                state.error = null
            })
            .addCase(login.rejected, (state, action) => {
                state.loading = false
                state.currentUser = null
                state.error = {
                    message: action.payload?.message || action.error.message || 'Unknown error occurred',
                    code: action.payload?.code || action.error.code || 'Error'
                }
            })
            .addCase(getCurrentUser.pending, state => {
                state.loading = true
                state.error = null
            })
            .addCase(getCurrentUser.fulfilled, (state, action) => {
                state.loading = false
                state.currentUser = action.payload
                state.error = null
            })
            .addCase(getCurrentUser.rejected, (state, action) => {
                state.loading = false
                state.currentUser = null
                state.error = {
                    message: action.payload?.message || action.error.message || 'Unknown error occurred',
                    code: action.payload?.code || action.error.code || 'Error'
                }
            })

    },
})

export default users.reducer
export const {logOut} = users.actions
