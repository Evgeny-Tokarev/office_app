"use client";

import {
    AppBar,
    Box,
    Toolbar,
    IconButton,
    Typography,
    Menu,
    MenuItem,
} from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import React from "react";
import ThemeSwitcher from "@/components/ThemeSwitcher";
import {useRouter} from 'next/navigation'

const menuItems = [{
    name: 'dashboard', route: '/'
}, {
    name: 'offices', route: '/offices'
}, {name: 'users', route: '/users'}]

export default function Navbar() {
    const router = useRouter()

    const [selectedItem, setSelectedItem] = React.useState(menuItems[0].name)
    const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
    const open = Boolean(anchorEl);

    const toggleMenuButton = (e: React.MouseEvent<HTMLElement> | React.TouchEvent<HTMLElement>) => {
        setAnchorEl(e.currentTarget);
    }
    const selectItemAndClose = (item: {
        name: string,
        route: string
    }) => {
        setSelectedItem(item.name)
        router.push(`${item.route}`)
        closeMenu()
    };
    const closeMenu = () => {
        setAnchorEl(null);
    };

    return (<Box>
        <AppBar position="static">
            <Toolbar sx={{
                display: 'flex',
                justifyContent: 'space-between',
                alignItems: 'center'
            }}>
                <IconButton
                    size="large"
                    edge="start"
                    color="inherit"
                    aria-label="menu"
                    sx={{mr: 2}}
                    onClick={toggleMenuButton}
                >
                    <div
                        className="flex justify-between items-center">
                        <MenuIcon sx={{mb: 0.5, mr: 1}}/>
                        <Typography
                            variant="h6"
                            component="div"
                            sx={{flexGrow: 1}}
                        >
                            Menu
                        </Typography>
                    </div>
                </IconButton>
                <Menu
                    id="long-menu"
                    MenuListProps={{
                        'aria-labelledby': 'long-button',
                    }}
                    anchorEl={anchorEl}
                    open={open}
                    onClose={closeMenu}
                >
                    {menuItems.map((menuItem) => (<MenuItem
                        key={menuItem.name}
                        selected={menuItem.name === selectedItem}
                        onClick={() => selectItemAndClose(menuItem)}>
                        {menuItem.name}
                    </MenuItem>))}
                </Menu>
                <ThemeSwitcher/>
            </Toolbar>
        </AppBar>
    </Box>);
}
