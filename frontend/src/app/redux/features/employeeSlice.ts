"use client"

import {
    createAsyncThunk, createSlice, type SerializedError
} from "@reduxjs/toolkit"
import {type Employee} from '@/app/models'
import api from "@/app/api"

export type EmployeesState = {
    employees: Employee[],
    loading: boolean,
    error: null | {
        code: string,
        message: string
    },
};

const initialState = {
    employees: [], loading: false, error: null,
} as EmployeesState


export const fetchEmployees = createAsyncThunk<Employee[], number, {
    rejectValue: SerializedError;
}>('offices/fetchOffices', async (office_id: number, {rejectWithValue}) => {
    try {
        const response = await api.employeesApi.getEmployees(office_id)
        if (response.data?.status === 200) return response.data.data
        return rejectWithValue({
            code: response.error?.code,
            message: response.error?.message
        } as SerializedError)
    } catch (err) {
        return rejectWithValue(err as SerializedError)
    }
})
export const saveImage = createAsyncThunk<boolean, {id: number, image: File}, {
    rejectValue: SerializedError;
}>('offices/saveImage', async ({id, image}, {rejectWithValue}) => {
    try {
        const response = await api.employeesApi.saveImage(id, image)
        if (response.data?.status === 200) return response.data.data
        return rejectWithValue({
            code: response.error?.code,
            message: response.error?.message
        } as SerializedError)
    } catch (err) {
        return rejectWithValue(err as SerializedError)
    }
})


export const updateEmployee = createAsyncThunk<Employee, Employee, {
    rejectValue: SerializedError;
}>('offices/updateOffice', async (employee, {rejectWithValue}) => {
    try {
        employee.id = +employee.id
        const response = await api.employeesApi.updateEmployee(employee)
        if (response.data?.status === 200) return response.data.data
        return rejectWithValue({
            code: response.error?.code,
            message: response.error?.message
        } as SerializedError)
    } catch (err) {
        return rejectWithValue(err as SerializedError)
    }
})
export const deleteEmployee = createAsyncThunk<void, number, {
    rejectValue: SerializedError;
}>('offices/deleteOffice', async (id, {rejectWithValue}) => {
    try {
        const response = await api.employeesApi.deleteEmployee(id)
        if (response.data?.status === 200) return
        return rejectWithValue({
            code: response.error?.code,
            message: response.error?.message
        } as SerializedError)
    } catch (err) {
        return rejectWithValue(err as SerializedError)
    }
})

export const createEmployee = createAsyncThunk<Employee, Employee, {
    rejectValue: SerializedError;
}>('offices/createOffice', async (employee, {rejectWithValue}) => {
    try {
        const response = await api.employeesApi.createEmployee(employee)
        if (response.data?.status === 200 || response.data?.status === 201) return response.data.data
        return rejectWithValue({
            code: response.error?.code,
            message: response.error?.message
        } as SerializedError)
    } catch (err) {
        return rejectWithValue(err as SerializedError)
    }
})
export const employees = createSlice({
    name: "offices", initialState, reducers: {}, extraReducers: builder => {
        builder
            .addCase(fetchEmployees.pending, state => {
                state.loading = true
                state.error = null
            })
            .addCase(fetchEmployees.fulfilled, (state, action) => {
                state.loading = false
                state.employees = action.payload ?? []
                state.error = null
            })
            .addCase(fetchEmployees.rejected, (state, action) => {
                state.loading = false
                state.employees = []
                state.error = {message: action.payload?.message || action.error.message || 'Unknown error occurred',
                    code: action.payload?.code || action.error.code || 'Error'}
            })
            .addCase(updateEmployee.pending, state => {
                state.loading = true
                state.error = null
            })
            .addCase(updateEmployee.fulfilled, (state) => {
                state.loading = false
                state.error = null
            }).addCase(updateEmployee.rejected, (state, action) => {
            state.loading = false
            state.error = {message: action.payload?.message || action.error.message || 'Unknown error occurred',
                code: action.payload?.code || action.error.code || 'Error'}
        }).addCase(createEmployee.pending, state => {
            state.loading = true
            state.error = null
        })
            .addCase(createEmployee.fulfilled, (state) => {
                state.loading = false
                state.error = null
            }).addCase(createEmployee.rejected, (state, action) => {
            state.loading = false
            state.error = {message: action.payload?.message || action.error.message || 'Unknown error occurred',
                code: action.payload?.code || action.error.code || 'Error'}
        }).addCase(deleteEmployee.pending, state => {
            state.loading = true
            state.error = null
        })
            .addCase(deleteEmployee.fulfilled, (state) => {
                state.loading = false
                state.error = null
            }).addCase(deleteEmployee.rejected, (state, action) => {
            state.loading = false
            state.error = {message: action.payload?.message || action.error.message || 'Unknown error occurred',
                code: action.payload?.code || action.error.code || 'Error'}
        })
    },
})

export default employees.reducer
