"use client"

import * as React from 'react'
import {
    Box, Typography, Modal
} from '@mui/material'
import {createPortal} from 'react-dom'
import {IconButton} from "@mui/material"
import {Close} from "@mui/icons-material"
import {ModalContext, initialProps} from "@/components/ModalProvider"
import {useSelector} from "react-redux"
import {RootState} from "@/app/redux/store"
import {useContext, useEffect} from "react"
import {LoaderContext} from "@/components/LoaderProvider"
import ConfirmActions
    from "@/components/modal/ConfirmActions"
import OfficeForm from "@/components/modal/OfficeForm"
import EmployeeForm from "@/components/modal/EmployeeForm"
import {type StyleObj, type ModalProps} from "@/app/models"


const style: StyleObj = {
    modal: {
        position: 'absolute' as 'absolute',
        top: '50%',
        left: '50%',
        bottom: 'auto',
        right: 'auto',
        transform: 'translate(-50%, -50%)',
        minWidth: '50%',
        borderRadius: '0.5rem',
        boxShadow: 24,
        overflow: 'hidden',
        p: 4,
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    }
}


export default function BasicModal() {
    const {
        openModal, setOpenModal, modalProps, setModalProps
    } = React.useContext(ModalContext)
    const {
        error, loading, infoState
    } = useSelector((state: RootState) => state.offices)
    const {setShowLoader} = useContext(LoaderContext)
    const [initialModalProps, setInitialModalProps] = React.useState<ModalProps>(initialProps)
    const [initialModalOpenState, setInitialModalOpenState] = React.useState(false)
    React.useEffect(() => {
        if (infoState) {
            setInitialModalOpenState(openModal)
            setInitialModalProps(modalProps)
            setOpenModal(true)
            setModalProps({
                type: 'info', title: infoState.title, text: infoState.text, isPermanent: false,
                style: {
                    bgcolor: 'green'
                }
            })
        }
    }, [infoState])
    React.useEffect(() => {
        if (error) {
            setInitialModalOpenState(openModal)
            setInitialModalProps(modalProps)
            setOpenModal(true)
            setModalProps({
                type: 'info', title: error.code, text: error.message, isPermanent: false,
                style: {
                    bgcolor: 'pink'
                }
            })
        }
    }, [error])

    useEffect(() => {
        if (!loading) setShowLoader(false)
        else setShowLoader(true)
    }, [loading])

    const [mounted, setMounted] = React.useState(false)

    React.useEffect(() => {
        setMounted(true)
    }, [])
    React.useEffect(() => {
        if (modalProps.style && Object.keys(modalProps.style).length > 0) Object.entries(modalProps.style).forEach(styleEntry => style.modal[styleEntry[0]] = styleEntry[1])
        if (!modalProps.isPermanent && openModal) {
            setTimeout(() => {
                closeModal()
            }, modalProps?.delay ?? 4000)
        }
    }, [openModal, modalProps])

    const onAction = (result?: unknown) => {
        if (modalProps.actionCallback) modalProps.actionCallback(result)
    }
    const checkReasonAndClose = (reason?: string) => {
        if (reason !== 'backdropClick') closeModal()
    }
    const closeModal = () => {
        setOpenModal(initialModalOpenState)
        setModalProps(initialModalProps)
        setInitialModalOpenState(false)
        setInitialModalProps(initialProps)
    }
    return mounted ? createPortal(<Modal
        open={openModal ?? false}
        onClose={(_, reason) => checkReasonAndClose(reason)}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
    >
        <Box sx={(theme) => ({
            ...style.modal,
            bgcolor: theme.palette.mode === 'light' ? 'white': 'black'})}>
            <IconButton
                size="large"
                edge="start"
                color="inherit"
                aria-label="menu"
                sx={{ml: 'auto'}}
                onClick={closeModal}
            >
                <div
                    className="flex justify-between items-center">
                    <Close/>
                </div>
            </IconButton>
            <Typography
                id="modal-modal-title"
                variant="h6"
                component="h3">
                {modalProps.title}
            </Typography>
            <Typography
                id="modal-modal-description"
                sx={{mt: 2}}>
                {modalProps.text}
            </Typography>
            {modalProps.type === 'office_form' &&
                <OfficeForm onCloseModal={closeModal}/>}
            {modalProps.type === 'employee_form' &&
                <EmployeeForm onCloseModal={closeModal}/>}
            {modalProps.withActions &&
                <ConfirmActions onAction={onAction}/>}
        </Box>
    </Modal>
, document.body) : null
}
