"use client";

import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {type Office} from '@/app/models';
import api from "@/app/api";
import {createAsyncThunk} from "@reduxjs/toolkit";

type OfficesState = {
    offices: Office[],
    loading: boolean,
    error: null | string,
};

const initialState = {
    offices: [], loading: false, error: null,
} as OfficesState;

const dateOptions = {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: 'numeric',
    minute: 'numeric',
    hour12: true
} as Intl.DateTimeFormatOptions

export const fetchOffices = createAsyncThunk('offices/fetchOffices', async () => {
    const response = await api.officesApi.getOffices();
    response.data.map(item => {
        item.created_at = new Date(String(item.created_at).replace("UTC", "GMT")).toLocaleString("en-CA", dateOptions)
        item.updated_at = new Date(String(item.updated_at).replace("UTC", "GMT")).toLocaleString("en-CA", dateOptions)
        return item
    })
    return response.data;
});

export const offices = createSlice({
    name: "offices",
    initialState,
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase(fetchOffices.pending, state => {
                state.loading = true;
                state.error = null;
            })
            .addCase(fetchOffices.fulfilled, (state, action) => {
                state.loading = false;
                state.offices = action.payload;
                state.error = null;
            })
            .addCase(fetchOffices.rejected, (state, action) => {
                state.loading = false;
                state.offices = [];
                state.error = action.error.message || 'An error occurred';
            });
    },
})

export default offices.reducer;
