import {Icon, Switch} from "@mui/material"
import React from "react"
import {useTheme} from "next-themes"

const label = {inputProps: {'aria-label': 'Theme switcher'}}

export default function ThemeSwitcher() {
    const [checked, setChecked] = React.useState(true)
    const {theme, setTheme} = useTheme()
    const [mounted, setMounted] = React.useState(false)

    React.useEffect(() => {
        console.log("Theme: ", theme)
        setMounted(true)
        setChecked(theme === 'dark')
    }, [])
    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setChecked(event.target.checked)
        setTheme(event.target.checked ? 'dark' : 'light')
    }

    if (!mounted) {
        return null
    }
    return (<div
        className="flex justify-between items-center ml-auto">
        <Icon>light_mode</Icon>
        <Switch
            {...label}
            checked={checked}
            color="default"
            onChange={handleChange}/>
        <Icon>dark_mode_icon</Icon>
    </div>)
}
