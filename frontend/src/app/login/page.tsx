"use client"
import React from 'react'
import UserForm from "@/components/modal/UserForm"
import {Box, Typography} from "@mui/material"
import {createPortal} from "react-dom"

const style = {
    layout: {
        position: 'absolute',
        top: '0',
        left: '0',
        right: '0',
        bottom: '0',
        width: '100%',
        height: '100%',
        transform: 'none',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: 'white',
        zIndex: 1000
    },
    wrapper: {
        minWidth: 320,
        width: '50%',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center'
    }
}

export default function Login() {
    return createPortal(<Box sx={(theme) => ({
        bgcolor: theme.palette.mode === 'light' ? 'white' : 'black',
        ...style.layout
    })}>
        <Box sx={{...style.wrapper}}>
            <Typography
                id="modal-modal-title"
                variant="h6"
                component="h3">
                Enter your credentials
            </Typography>
            <Typography
                id="modal-modal-description"
                sx={{mt: 2}}>
                Please enter your username, email and password
            </Typography>
            <UserForm onCloseModal={() => {
            }}/>
        </Box>
    </Box>, document.body)
}
