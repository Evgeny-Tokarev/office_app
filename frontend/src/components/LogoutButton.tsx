import {IconButton} from "@mui/material"
import React from "react"
import {ModalContext} from "@/components/ModalProvider"
import {useDispatch} from "react-redux"
import {logOut, UserState} from "@/app/redux/features/usersSlice"
import LogoutIcon from '@mui/icons-material/Logout'
import {ThunkDispatch} from "@reduxjs/toolkit"
import type {AnyAction} from "redux"

export default function LogoutButton() {
    const {
        setOpenModal, setModalProps
    } = React.useContext(ModalContext) ?? {}
    const dispatch = useDispatch<ThunkDispatch<UserState, unknown, AnyAction>>()
    const logOutUser = (e: React.MouseEvent | React.TouchEvent) => {
        e.stopPropagation()
        const cb = (result: boolean) => {
            if (result) {
                dispatch(logOut())
            }
            setOpenModal(false)
        }
        setModalProps({
            type: 'warn',
            title: "Are you sure",
            text: 'You\'re about to logout.',
            isPermanent: true,
            withActions: true,
            actionCallback: cb
        })
        setOpenModal(true)
    }
    return (<div
        className="flex justify-between items-center">
        <IconButton
            aria-label="logout"
            color="default"
            onClick={logOutUser}>
            <LogoutIcon/>
        </IconButton>
    </div>)
}