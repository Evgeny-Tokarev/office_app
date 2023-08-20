"use client"

import Link from "next/link";
import Icon from '@mui/material/Icon';
import Switch from '@mui/material/Switch';
import React from "react";
import { useTheme } from 'next-themes'

const label = {inputProps: {'aria-label': 'Switch demo'}};

export default function Navbar() {
    const { theme, setTheme } = useTheme()
    const [mounted, setMounted] = React.useState(false)

    React.useEffect(() => {
        setMounted(true)
    }, [])

    if (!mounted) {
        return null
    }

    return (<nav
            className="flex justify-between items-center bg-slate-300 dark:bg-slate-800 px-8 py-3">
        <Link
            className="text-black dark:text-white font-bold"
            href={"/"}>
            Dashboard
        </Link>
        <div
                className="flex justify-between items-center">
            <Icon sx={{ color: theme === 'light' ? 'black' : 'white' }}>light_mode</Icon>
            <Switch
                {...label}
                defaultChecked
                color="default"
                onChange={() => setTheme(theme === 'dark' ? 'light' : 'dark')}/>
            <Icon sx={{ color: theme === 'light' ? 'black' : 'white' }}>dark_mode_icon</Icon>
        </div>

    </nav>);
}
