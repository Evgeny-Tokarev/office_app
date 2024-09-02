"use client"

import {createAsyncThunk, createSlice, type SerializedError} from "@reduxjs/toolkit"
import {type User} from '@/app/models'
import {clearError, setError, setLoading} from "@/app/redux/features/utilsSlice"
import api from "@/app/api"
import type {AppDispatch, RootState} from "@/app/redux/store"

export type UserState = {
    currentUser: User | null | undefined,
};

const initialState = {
    currentUser: undefined
} as UserState


export const login = createAsyncThunk<{ user: User, token: string }, {
    userName: string,
    email: string,
    password: string
}, { rejectValue: SerializedError; state: RootState; dispatch: AppDispatch }>('users/login', async ({
                                                                                                        userName,
                                                                                                        email,
                                                                                                        password
                                                                                                    }, {
                                                                                                        dispatch,
                                                                                                        rejectWithValue
                                                                                                    }) => {
    try {
        dispatch(setLoading(true))
        const response = await api.usersApi.login(userName, email, password)
        dispatch(setLoading(false))
        console.log(response)
        if (response.data?.status === 200) {
            dispatch(clearError())
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

export const getCurrentUser = createAsyncThunk<User, string, {
    rejectValue: SerializedError;
    state: RootState;
    dispatch: AppDispatch
}>('users/getCurrentUser', async (token, {dispatch, rejectWithValue}) => {
    try {
        dispatch(setLoading(true))
        const response = await api.usersApi.getCurrentUser(token)
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

export const users = createSlice({
    name: "users", initialState, reducers: {
        logOut: (state) => {
            state.currentUser = null
            localStorage.removeItem('officeAppToken')
        },

        setNullCurrentUser: (state) => {
            state.currentUser = null
        },
    }, extraReducers: builder => {
        builder
            .addCase(login.fulfilled, (state, action) => {
                state.currentUser = action.payload.user
                localStorage.setItem('officeAppToken', action.payload.token)
            })
            .addCase(login.rejected, (state, action) => {
                state.currentUser = null
            })
            .addCase(getCurrentUser.fulfilled, (state, action) => {
                state.currentUser = action.payload
            })
            .addCase(getCurrentUser.rejected, (state, action) => {
                state.currentUser = null
            })

    },
})

export default users.reducer
export const {logOut, setNullCurrentUser} = users.actions
