"use client"

import dynamic from 'next/dynamic'
import {Paper} from "@mui/material"
import React from "react"

const DashboardNoSSR = dynamic(() => import('@/app/dashboard/page'), {ssr: false})

export default function Page() {
    return (
        <DashboardNoSSR/>
    )
}
