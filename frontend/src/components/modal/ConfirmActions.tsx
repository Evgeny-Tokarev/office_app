import {Box, Button, type Theme} from '@mui/material';
import * as React from "react";

const style = {
    actions: {
        width: '100%',
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        mt: 'auto'
    },
}
export default function ConfirmActions({onAction}: {onAction: (condition: boolean) => void}) {
    return (
        <Box
            sx={style.actions}
            aria-label="modal action buttons">
            <Button
                onClick={() => onAction(true)}
                variant="contained"
                color='error'>YES</Button>
            <Button
                onClick={() => onAction(false)}
                variant="contained">NO</Button>
        </Box>
    )
}
