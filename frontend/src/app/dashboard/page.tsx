"use client";

import AddIcon from '@mui/icons-material/Add';
import {Fab, Typography} from "@mui/material";
import {useTheme} from "@mui/material";
import {useTheme as useNTheme} from "next-themes";
import React, {useEffect} from "react";

export default function Dashboard() {
    const theme = useTheme();
    const nTheme = useNTheme()
    useEffect(() => {
        console.log(theme, nTheme);
    }, [theme, nTheme]);
    return (<main
        className="p-4 flex-1 flex flex-col justify-between items-center bg-white dark:bg-black">
        <Typography
            mx={'auto'}
            variant="h1"
            component="h1"
        >Header</Typography>
        <Typography
            variant="h2"
            component="h2">Dashboard</Typography>
        <Fab aria-label="add">
            <AddIcon/>
        </Fab>
    </main>)
}
