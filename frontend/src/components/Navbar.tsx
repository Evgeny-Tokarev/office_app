"use client"

import Link from "next/link";
import {
    AppBar,
    Box,
    Toolbar,
    Button,
    IconButton,
    Icon,
    Switch,
    Typography,
    Menu,
    MenuItem,
    useTheme as useMTheme
} from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import React from "react";
import {useTheme} from 'next-themes'

const label = {inputProps: {'aria-label': 'Switch demo'}};
const options = ['None', 'Atria', 'Callisto', 'Dione', 'Ganymede', 'Hangouts Call', 'Luna', 'Oberon', 'Phobos', 'Pyxis', 'Sedna', 'Titania', 'Triton', 'Umbriel',];

const ITEM_HEIGHT = 48;
export default function Navbar() {
    const {theme, setTheme} = useTheme()
    const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
    const open = Boolean(anchorEl);
    const [mounted, setMounted] = React.useState(false)

    const toggleMenuButton = (event: React.MouseEvent<HTMLElement>) => {
        setAnchorEl(event.currentTarget);
    }
    const handleClose = () => {
        setAnchorEl(null);
    };

    React.useEffect(() => {
        setMounted(true)
    }, [])

    if (!mounted) {
        return null
    }

    return (<Box>
        <AppBar position="static">
            <Toolbar
                sx={{background: theme === 'dark' ? 'black' : ''}}>
                <IconButton
                    size="large"
                    edge="start"
                    color="inherit"
                    aria-label="menu"
                    sx={{mr: 2}}
                    onClick={toggleMenuButton}
                >
                    <MenuIcon/>
                </IconButton>
                <Menu
                    id="long-menu"
                    MenuListProps={{
                        'aria-labelledby': 'long-button',
                    }}
                    anchorEl={anchorEl}
                    open={open}
                    onClose={handleClose}
                    sx={{
                        maxHeight: ITEM_HEIGHT * 4.5,
                        width: '20ch',
                    }}
                >
                    {options.map((option) => (
                        <MenuItem key={option}
                                  selected={option === 'Pyxis'}
                                  onClick={handleClose}>
                            {option}
                        </MenuItem>))}
                </Menu>
                <Typography variant="h6" component="div"
                            sx={{flexGrow: 1}}>
                    Menu
                </Typography>
                <div
                    className="flex justify-between items-center">
                    <Icon>light_mode</Icon>
                    <Switch
                        {...label}
                        defaultChecked
                        color="default"
                        onChange={() => setTheme(theme === 'dark' ? 'light' : 'dark')}/>
                    <Icon>dark_mode_icon</Icon>
                </div>
            </Toolbar>
        </AppBar>
    </Box>);
}
// <nav
//     className="flex justify-between items-center bg-slate-300 dark:bg-slate-800 px-8 py-3">
//
//     <Link
//         className="text-black dark:text-white font-bold"
//         href={"/"}>
//         Dashboard
//     </Link>
//     <Typography>
//         {mTheme.palette.mode}
//     </Typography>
//     <Typography>
//         {theme}
//     </Typography>
//     <div
//         className="flex justify-between items-center">
//         <Icon
//             sx={{color: theme === 'light' ? 'black' : 'white'}}>light_mode</Icon>
//         <Switch
//             {...label}
//             defaultChecked
//             color="default"
//             onChange={() => setTheme(theme === 'dark' ? 'light' : 'dark')}/>
//         <Icon
//             sx={{color: theme === 'light' ? 'black' : 'white'}}>dark_mode_icon</Icon>
//     </div>
//
// </nav>
