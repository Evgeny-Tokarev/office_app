import {combineReducers, configureStore} from "@reduxjs/toolkit"
import officesSlice from "@/app/redux/features/officesSlice"
import employeeSlice from "@/app/redux/features/employeeSlice"
import usersSlice from "@/app/redux/features/usersSlice"

const rootReducer = combineReducers({
    offices: officesSlice,
    employees: employeeSlice,
    users: usersSlice
})
export const store = configureStore({
    reducer: rootReducer,
    devTools: process.env.NODE_ENV !== "production",
})

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
