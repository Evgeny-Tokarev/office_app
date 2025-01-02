"use client"

import * as React from 'react'
import {Box, IconButton, Modal, Typography, type Theme} from '@mui/material'
import {createPortal} from 'react-dom'
import {Close} from "@mui/icons-material"
import {initialProps, initialStyleProps, ModalContext} from "@/components/ModalProvider"
import {useDispatch, useSelector} from "react-redux"
import {AppDispatch, RootState} from "@/app/redux/store"
import ConfirmActions from "@/components/modal/ConfirmActions"
import OfficeForm from "@/components/modal/OfficeForm"
import UserForm from "@/components/modal/UserForm"
import EmployeeForm from "@/components/modal/EmployeeForm"
import {type ModalProps, type StyleObj} from "@/app/models"


const style: StyleObj = {
    modal: {...initialStyleProps}
}

const loginModalProps: ModalProps = {
    isPermanent: true,
    type: 'user_form',
    title: 'Please Log In',
    text: 'You need to log in to access the content.',
    closable: false,
    style: {
        top: '0',
        left: '0',
        right: '0',
        bottom: '0',
        width: '100%',
        height: '100%',
        transform: 'none',
    }
}


export default function BasicModal() {
    const {
        openModal, setOpenModal, modalProps, setModalProps
    } = React.useContext(ModalContext)
    const {
        error, infoState
    } = useSelector((state: RootState) => state.utils)
    const [initialModalProps, setInitialModalProps] = React.useState<ModalProps>({...initialProps})
    const [initialModalOpenState, setInitialModalOpenState] = React.useState(false)
    const dispatch = useDispatch<AppDispatch>()
    const [mounted, setMounted] = React.useState(false)


    React.useEffect(() => {
        if (infoState) {
            setInitialModalOpenState(openModal)
            setInitialModalProps({...modalProps})
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
        console.log("error: ", error)
        if (error) {
            setInitialModalOpenState(openModal)
            setInitialModalProps({...modalProps})
            setModalProps({
                type: 'info', title: error.code, text: error.message, isPermanent: false, closable: true,
                style: {
                    ...initialStyleProps,
                    bgcolor: 'pink'
                }
            })
            setOpenModal(true)
        }
    }, [error])


    React.useEffect(() => {
        setMounted(true)
    }, [])

    React.useEffect(() => {
        if (modalProps.style && Object.keys(modalProps.style).length > 0) {
            style.modal = {...style.modal, ...modalProps.style}
        }
        if (!modalProps.isPermanent && openModal) {
            setTimeout(() => {
                closeModal()
            }, modalProps?.delay ?? 4000)
        }
    }, [modalProps])

    const onAction = (result?: unknown) => {
        if (modalProps.actionCallback) modalProps.actionCallback(result)
    }
    const checkReasonAndClose = (reason?: string) => {
        if (reason !== 'backdropClick') closeModal()
    }
    const closeModal = () => {
        setOpenModal(initialModalOpenState)
        setModalProps({...initialModalProps})
        setInitialModalOpenState(false)
        setInitialModalProps({...initialProps})
    }
    return mounted ? createPortal(<Modal
            open={openModal ?? false}
            onClose={(_, reason) => checkReasonAndClose(reason)}
            aria-labelledby="modal-modal-title"
            aria-describedby="modal-modal-description"
        >
            <Box sx={(theme) => ({
                ...style.modal
            })}>
                <IconButton
                    size="large"
                    edge="start"
                    color="inherit"
                    aria-label="close modal"
                    sx={{ml: 'auto', display: modalProps.closable ? '' : 'none'}}
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
                {modalProps.type === 'user_form' &&
                    <UserForm onCloseModal={closeModal}/>}
            </Box>
        </Modal>
        , document.body) : null
}
