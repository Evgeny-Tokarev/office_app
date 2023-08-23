"use client";

import AddIcon from '@mui/icons-material/Add';
import {Fab, Typography} from "@mui/material";
import {useTheme} from "@mui/material";
import {useTheme as useNTheme} from "next-themes";
import React, {useEffect} from "react";

export default function Dashboard() {
    return (<main
        className="p-4 flex-1 flex flex-col justify-between items-center bg-white dark:bg-black">
        <Typography
            variant="h1">
            Dashboard
        </Typography>
        <Fab aria-label="add">
            <AddIcon/>
        </Fab>
    </main>)
}
