import {Icon, Switch} from "@mui/material";
import React from "react";
import {useTheme} from "next-themes";

const label = {inputProps: {'aria-label': 'Switch demo'}};

export default function ThemeSwitcher() {
    const {theme, setTheme} = useTheme()
    const [mounted, setMounted] = React.useState(false)

    React.useEffect(() => {
        setMounted(true)
    }, [])

    if (!mounted) {
        return null
    }
    return (<div
        className="flex justify-between items-center">
        <Icon>light_mode</Icon>
        <Switch
            {...label}
            defaultChecked
            color="default"
            onChange={() => setTheme(theme === 'dark' ? 'light' : 'dark')}/>
        <Icon>dark_mode_icon</Icon>
    </div>)
}
