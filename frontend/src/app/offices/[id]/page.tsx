"use client";

import React from 'react'
import {Box, TextField, Button} from "@mui/material";
import {useDispatch, useSelector} from "react-redux";
import {
    fetchOffices
} from "@/app/redux/features/officesSlice";
import {RootState} from '@/app/redux/store';
import {ThunkDispatch} from "@reduxjs/toolkit";
import {Office} from "@/app/models";
import {usePathname} from 'next/navigation';

export default function EditOffice({params}: {
    params: { id: string | number }
}) {
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();
    const {
        offices, loading, error
    } = useSelector((state: RootState) => state.offices);
    const office: Office | undefined = React.useMemo(() => {
        if (!offices.length) dispatch(fetchOffices());
        return offices.find(of => of.id === params.id) || offices[0];
    }, [offices, params.id]);
    const [name, setName] = React.useState('');
    const [address, setAddress] = React.useState('');
    React.useEffect(() => {
        if (office) {
            setName(office.name || '');
            setAddress(office.address || '');
        }
    }, [office]);
    const pathname = usePathname()
    const saveOffice = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        console.log(params.id, name, address)
    }
    return (<Box
            component="form"
            onSubmit={saveOffice}
            mt={4}
            sx={{
                '& .MuiTextField-root': {
                    m: 1, width: '25ch'
                },
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center'
            }}
            noValidate
            autoComplete="off"
        >
            <div className="flex flex-col items-center">
                <TextField
                    error={name.length < 3}
                    id="name"
                    label="Name"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                />
                <TextField
                    id="address"
                    label="Address"
                    value={address}
                    multiline
                    maxRows={4}
                    onChange={(e) => setAddress(e.target.value)}
                />
            </div>
            <Button variant="contained"
                    type="submit">Save</Button>
        </Box>

    )
}
