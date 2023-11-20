import { configureStore, combineReducers } from "@reduxjs/toolkit"
import officesSlice from "@/app/redux/features/officesSlice"
import employeeSlice
    from "@/app/redux/features/employeeSlice";

const rootReducer = combineReducers({
    offices: officesSlice,
    employees: employeeSlice,
})
export const store = configureStore({
    reducer: rootReducer,
    devTools: process.env.NODE_ENV !== "production",
})

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
