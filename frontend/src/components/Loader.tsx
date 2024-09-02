"use client"

import {Box, CircularProgress} from "@mui/material"
import React from 'react'
import {useSelector} from "react-redux"
import {RootState} from "@/app/redux/store"
import {createPortal} from "react-dom"

export default function Loader() {
    const [
        showLoader, setShowLoader
    ] = React.useState(false)
    const {loading} = useSelector((state: RootState) => state.utils)
    React.useEffect(() => {
        setShowLoader(loading)
    }, [loading])
    if (!showLoader) return null
    return showLoader ? createPortal(<Box
        sx={{
            position: 'fixed',
            top: '50%',
            left: '50%',
            transform: 'translate(-50%, -50%)',
            zIndex: '1500'
        }}
    >
        <CircularProgress color="inherit"/>
    </Box>, document.body) : <p>No loader</p>
}
