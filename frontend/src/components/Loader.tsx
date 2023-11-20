"use client"

import {Box, CircularProgress} from "@mui/material"
import React from 'react'
import {LoaderContext} from "@/components/LoaderProvider"
import {useSelector} from "react-redux"
import {RootState} from "@/app/redux/store"

export default function Loader() {
    const {
        showLoader, setShowLoader} = React.useContext(LoaderContext)
    const {loading, error} = useSelector((state: RootState) => state.offices)
    React.useEffect(() => {
        setShowLoader(loading)
    }, [loading])
    if (!showLoader) return null
    return (<Box
            sx={{
                position: 'fixed',
                top: '50%',
                left: '50%',
                transform: 'translate(-50%, -50%)'
            }}
        >
        <CircularProgress color="inherit"/>
    </Box>)
}
