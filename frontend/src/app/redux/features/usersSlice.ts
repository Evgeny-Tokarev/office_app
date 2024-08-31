"use client"

import {createAsyncThunk, createSlice, type SerializedError, ThunkDispatch} from "@reduxjs/toolkit"
import {type User} from '@/app/models'
import {setError, type ErrorSatate, clearError} from "@/app/redux/features/errorSlice"
import api from "@/app/api"
import type { AppDispatch, RootState } from "@/app/redux/store"

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
    rejectValue: SerializedError; dispatch: AppDispatch
}>('users/login', async ({userName, email, password}, {dispatch, rejectWithValue}) => {
    try {
        const response = await api.usersApi.login(userName, email, password)
        if (response.data?.status === 200) {
            dispatch(clearError())
            return response.data.data
        }
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
            localStorage.removeItem('officeAppToken')
        },
    }, extraReducers: builder => {
        builder
            .addCase(login.pending, state => {
                state.loading = true
            })
            .addCase(login.fulfilled, (state, action) => {
                state.loading = false
                state.currentUser = action.payload.user
                localStorage.setItem('officeAppToken', action.payload.token)
            })
            .addCase(login.rejected, (state, action) => {
                console.log("login rejected")
                state.loading = false
                state.currentUser = null
            })
            .addCase(getCurrentUser.pending, state => {
                state.loading = true
            })
            .addCase(getCurrentUser.fulfilled, (state, action) => {
                state.loading = false
                state.currentUser = action.payload
            })
            .addCase(getCurrentUser.rejected, (state, action) => {
                state.loading = false
                state.currentUser = null
            })

    },
})

export default users.reducer
export const {logOut} = users.actions
