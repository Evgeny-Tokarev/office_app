import {createTheme} from "@mui/material/styles";

export const lightTheme = createTheme({
    components: {
        MuiTypography: {
            styleOverrides: {
                root: {
                    fontFamily: 'Lato',
                    color: 'blue',
                    backgroundColor: 'pink'
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
                    fontFamily: 'Lato',
                    color: 'green',
                    backgroundColor: 'yellow'
                }
            }
        }
    }, palette: {
        mode: 'dark',

    },

});
