import {createTheme} from "@mui/material/styles"

const typography = {
    fontSize: 16,
    h1: {
        fontSize: 48,
        '@media (min-width:600px)': {
            fontSize: 56,
        },
        '@media (min-width:960px)': {
            fontSize: 64,
        },
        '@media (min-width:1280px)': {
            fontSize: 72,
        },
    },
    h2: {
        fontSize: 36,
        '@media (min-width:600px)': {
            fontSize: 42,
        },
        '@media (min-width:960px)': {
            fontSize: 48,
        },
        '@media (min-width:1280px)': {
            fontSize: 56,
        },
    },
    h3: {
        fontWeight: 'bold',
        fontSize: 18,
        '@media (min-width:600px)': {
            fontSize: 20,
        },
        '@media (min-width:960px)': {
            fontSize: 24,
        },
        '@media (min-width:1280px)': {
            fontSize: 30,
        },
    },
    body1: {
        fontSize: 16, // 1rem equivalent in pixels
        '@media (min-width:600px)': {
            fontSize: 17.6, // 1.1rem equivalent in pixels
        },
        '@media (min-width:960px)': {
            fontSize: 19.2, // 1.2rem equivalent in pixels
        },
        '@media (min-width:1280px)': {
            fontSize: 20.8, // 1.3rem equivalent in pixels
        },
    },
}

export const lightTheme = createTheme({
    palette: {
        mode: 'light',
        // primary: {
        //     main: '#1976d2',
        // },
        // secondary: {
        //     main: '#ff4081',
        // },
    },
    typography
})
export const darkTheme = createTheme({
    palette: {
        mode: 'dark',
        // primary: {
        //     main: '#90caf9',
        // },
        // secondary: {
        //     main: '#f48fb1',
        // },
    },
    typography
})
