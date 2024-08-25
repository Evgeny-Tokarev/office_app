import React from 'react'
import {Typography} from "@mui/material"

export default function Dashboard() {
    return (<div className="flex-1 flex-col">
        <Typography
            sx={{"wordWrap": "break-word", "textAlign": "center"}}
            variant="h1">
            Dashboard
        </Typography>
    </div>)
}
