"use client";

import React from 'react'
import {List} from "@mui/material";
import Divider from "@mui/material/Divider";
import {useDispatch, useSelector} from "react-redux";
import {
    fetchOffices
} from "@/app/redux/features/officesSlice";
import {RootState} from '@/app/redux/store';
import {ThunkDispatch} from "@reduxjs/toolkit";
import OfficeItem from "@/components/OfficeItem";

export default function Page() {
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();
    const {
        offices,
        loading,
        error
    } = useSelector((state: RootState) => state.offices);

    React.useEffect(() => {
        dispatch(fetchOffices());
    }, []);

    return (
        <List sx={{
            width: '100%',
            bgcolor: 'background.paper'
        }}>
            {offices.map(office =>
                <OfficeItem office={office} key={office.id}/>)}
            <Divider />
        </List>
    )
}
