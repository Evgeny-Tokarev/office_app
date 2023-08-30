"use client";

import * as React from 'react';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import Avatar from '@mui/material/Avatar';
import Typography from '@mui/material/Typography';
import {Office} from "@/app/models"
import {IconButton} from "@mui/material";
import {DeleteForever, Edit} from "@mui/icons-material";
import {useRouter} from "next/navigation";

export default function OfficeItem({office}: {
    office: Office
}) {
    const router = useRouter()
    const deleteItem = (office: Office) => {
        console.log("Delete-", office.id)
    }
    const editItem = (office: Office) => {
        console.log("Edit-", office.id)
        router.push(`/offices/${office.id}`)
    }
    return (<ListItem alignItems="flex-start" sx={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        gap: '1rem'
    }}>
        <ListItemAvatar>
            <Avatar alt="Office photo"
                    src={office.photo}/>
        </ListItemAvatar>
        <ListItemText
            primary={<Typography
                component="h3"
                variant="h5"
                color="text.primary"
            >{office.name}</Typography>}
            secondary={<React.Fragment>
                <Typography
                    component="span"
                    variant="body1"
                    color="text.primary"
                >
                    {office.address}
                </Typography>
                <br/>
                <Typography
                    component="span"
                    variant="body2"
                    color="text.primary">
                    Created {office.created_at}</Typography>
            </React.Fragment>}
        />
        <div className="flex justify-between items-center">
            <IconButton
                size="large"
                edge="start"
                color="inherit"
                aria-label="edit"
                sx={{mr: 2}}
                onClick={() => editItem(office)}
            >
                <div
                    className="flex justify-between items-center">
                    <Edit sx={{mr: 1}}/>
                    <Typography
                        variant="h6"
                        component="div"
                        sx={{flexGrow: 1}}
                    >
                        Edit
                    </Typography>
                </div>
            </IconButton>
            <IconButton
                size="large"
                edge="start"
                color="inherit"
                aria-label="delete"
                sx={{mr: 2}}
                onClick={() => deleteItem(office)}
            >
                <div
                    className="flex justify-between items-center">
                    <DeleteForever sx={{mr: 1}}/>
                    <Typography
                        variant="h6"
                        component="div"
                        sx={{flexGrow: 1}}
                    >
                        Delete
                    </Typography>
                </div>
            </IconButton>
        </div>
    </ListItem>)
}
