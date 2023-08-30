import { configureStore, combineReducers } from "@reduxjs/toolkit";
import officesSlice from "./features/officesSlice";

const rootReducer = combineReducers({
    offices: officesSlice,
});
export const store = configureStore({
    reducer: rootReducer,
    devTools: process.env.NODE_ENV !== "production",
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
