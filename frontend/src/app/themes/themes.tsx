import {createTheme} from "@mui/material/styles";

export const lightTheme = createTheme({
    components: {
        MuiTypography: {
            styleOverrides: {
                root: {
                    color: 'black',
                }
            }
        }
    }, palette: {
        mode: 'light',
    },
});
export const darkTheme = createTheme({
    components: {
        MuiTypography: {
            styleOverrides: {
                root: {
                    color: 'white',
                }
            }
        }
    }, palette: {
        mode: 'dark',

    },

});
