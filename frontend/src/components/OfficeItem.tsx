"use client"

import * as React from 'react'
import {
    Typography,
    Avatar,
    ListItemAvatar,
    ListItemText,
    ListItem
} from '@mui/material/'
import {Office} from "@/app/models"
import {IconButton} from "@mui/material"
import {DeleteForever, Edit} from "@mui/icons-material"
import convertDate from "@/app/helpers/dateConverter"
import {
    ModalContext
} from "@/components/ModalProvider"
import {useDispatch} from "react-redux"
import {ThunkDispatch} from "@reduxjs/toolkit"
import {
    deleteOffice, fetchOffices, OfficesState, updateOffice
} from "@/app/redux/features/officesSlice"
import {type AnyAction} from "redux"
import {useRouter} from "next/navigation"


export default function OfficeItem({office}: {
    office: Office
}) {
    const dispatch = useDispatch<ThunkDispatch<OfficesState, unknown, AnyAction>>()
    const router = useRouter()
    const {
        setOpenModal, setModalProps
    } = React.useContext(ModalContext) ?? {}
    const deleteItem = (e: React.MouseEvent, office: Office) => {
        e.stopPropagation()
        const cb = async (result: boolean) => {
            setOpenModal(false)
            if (result) {
                const res = await dispatch(deleteOffice(office.id))
                if (res.meta.requestStatus === 'fulfilled') {
                    dispatch(fetchOffices())
                }
            }
        }
        setOpenModal(true)
        setModalProps({
            type: 'warn',
            title: 'Warning',
            text: 'Do you want to delete item?',
            isPermanent: true,
            withActions: true,
            actionCallback: cb
        })
    }
    const editItem = (e: React.MouseEvent, office: Office) => {
        e.stopPropagation()
        setOpenModal(false)
        const cb = async (result: boolean) => {
            if (result) {
                const res = await dispatch(updateOffice(office))
                if (res.meta.requestStatus === 'fulfilled') {
                    dispatch(fetchOffices())
                }
            }
        }
        setOpenModal(true)
        setModalProps({
            type: 'office_form',
            title: `Edit office ${office.name}`,
            isPermanent: true,
            actionCallback: cb,
            formProps: {
                id: office.id
            }
        })
    }
    const openOffice = () => {
        router.push(`offices/${office.id}`)
    }
    return (<ListItem
        alignItems="flex-start"
        sx={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'space-between',
            gap: '1rem',
            borderBottom: '1px solid',
            borderColor: 'text.primary',
            cursor: 'pointer',
            '&:hover': {
                bgcolor: (theme) => theme.palette.action.hover
            }
        }}
        onClick={openOffice}>
        <ListItemAvatar>
            <Avatar
                alt="Office photo"
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
                    Created {convertDate(office.created_at)}</Typography>
            </React.Fragment>}
        />
        <div className="flex justify-between items-center">
            <IconButton
                size="large"
                edge="start"
                aria-label="edit"
                sx={{mr: 2}}
                onClick={(e) => editItem(e, office)}
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
                aria-label="delete"
                sx={{mr: 2}}
                onClick={(e) => deleteItem(e, office)}
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
