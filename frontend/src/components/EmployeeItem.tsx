"use client"

import * as React from 'react'
import {
    Typography,
    Avatar,
    ListItemAvatar,
    ListItemText,
    ListItem
} from '@mui/material/'
import {Employee} from "@/app/models"
import {IconButton} from "@mui/material"
import {DeleteForever, Edit} from "@mui/icons-material"
import convertDate from "@/app/helpers/dateConverter"
import {
    ModalContext
} from "@/components/ModalProvider"
import {useDispatch} from "react-redux"
import {ThunkDispatch} from "@reduxjs/toolkit"
import {
    fetchEmployees,
    EmployeesState,
    deleteEmployee,
    updateEmployee
} from "@/app/redux/features/employeeSlice"
import {type AnyAction} from "redux"
import {useRouter} from "next/navigation"
import {useTheme as useNextTheme} from "next-themes";


export default function EmployeeItem({employee}: {
    employee: Employee
}) {
    const {theme} = useNextTheme()
    const dispatch = useDispatch<ThunkDispatch<EmployeesState, unknown, AnyAction>>()
    const [photo, setPhoto] = React.useState('')
    const {
        setOpenModal, setModalProps
    } = React.useContext(ModalContext) ?? {}
    const deleteItem = (e: React.MouseEvent, employee: Employee) => {
        e.stopPropagation()
        const cb = async (result: boolean) => {
            if (result) {
                const res = await dispatch(deleteEmployee(employee.id))
                if (res.meta.requestStatus === 'fulfilled') dispatch(fetchEmployees(employee.office_id))
            }
            setOpenModal(false)
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
    const editItem = (e: React.MouseEvent, employee: Employee) => {
        e.stopPropagation()
        const cb = async (result: boolean) => {
            if (result) {
                const res = await dispatch(updateEmployee(employee))
                if (res.meta.requestStatus === 'fulfilled') dispatch(fetchEmployees(employee.office_id))
            }
            setOpenModal(false)
        }
        setOpenModal(true)
        setModalProps({
            type: 'employee_form',
            title: `Edit employee ${employee.name}`,
            isPermanent: true,
            actionCallback: cb,
            formProps: {
                id: employee.id,
                office_id: employee.office_id,
            }
        })
    }
    React.useEffect(() => {
        if (employee) {
            setPhoto(employee.photo
                ? employee.photo
                : theme === 'dark'
                    ? '/image-placeholder-dk.svg'
                    : '/image-placeholder-lt.svg')
        }
    }, [employee, theme])
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
        }}>
        <ListItemAvatar>
            <Avatar
                alt="Employee photo"
                src={photo}/>
        </ListItemAvatar>
        <ListItemText>
            <Typography
                component="h3"
                variant="h5"
                color="text.primary"
            >{employee.name} {employee.age}</Typography>
        </ListItemText>
        <div className="flex justify-between items-center">
            <IconButton
                size="large"
                edge="start"
                aria-label="edit"
                sx={{mr: 2}}
                onClick={(e) => editItem(e, employee)}
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
                onClick={(e) => deleteItem(e, employee)}
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
