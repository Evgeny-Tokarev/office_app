"use client"
import React, {createContext, useState} from "react"
import {ModalContextType, ModalProps} from "@/app/models"
import type {Theme} from "@mui/material";

export const initialStyleProps = {
    bgcolor: (theme: Theme)  => theme.palette.background.paper,
    position: 'absolute',
    top: '50%',
    left: '50%',
    bottom: 'auto',
    right: 'auto',
    transform: 'translate(-50%, -50%)',
    minWidth: '50%',
    width: 'auto',
    height: 'auto',
    maxWidth: '100%',
    maxHeight: '100%',
    borderRadius: '0.5rem',
    boxShadow: 24,
    overflow: 'hidden',
    p: 4,
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
} as { [name: string]: string | number | Function}

export const initialProps = {
    type: 'info',
    text: 'Initial text',
    title: 'Initial title',
    isPermanent: false,
    withActions: false,
    delay: 4000,
    closable: true,
    style: {...initialStyleProps}
} as ModalProps
export const ModalContext = createContext<ModalContextType>({
    openModal: false,
    setOpenModal: () => {
    },
    modalProps: {...initialProps},
    setModalProps: () => {
    }
})

export function ModalContextProvider({children}: {
    children: React.ReactNode
}) {
    const [openModal, setOpenModal] = useState<boolean>(false)
    const [modalProps, setModalProps] = useState<ModalProps>({
        type: 'info',
        text: '',
        title: '',
        isPermanent: false,
        withActions: false,
        delay: 4000
    })

    return (
        <ModalContext.Provider
            value={{
                openModal,
                setOpenModal,
                modalProps,
                setModalProps
            }}
        >
            {children}
        </ModalContext.Provider>
    )
}
