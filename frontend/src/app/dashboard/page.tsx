import AddIcon from '@mui/icons-material/Add';
import {Fab, Typography} from "@mui/material";
import {Paper} from "@mui/material";

export default function Dashboard() {
    return (<Paper elevation={3} sx={{flexGrow: 1, display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'space-between'}}>
        <Typography
            variant="h1">
            Dashboard
        </Typography>
        <Fab aria-label="add">
            <AddIcon/>
        </Fab>
    </Paper>)
}
