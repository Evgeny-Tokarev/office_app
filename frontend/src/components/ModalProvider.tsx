"use client"
import React, {createContext, useState} from "react"
import {ModalContextType, ModalProps} from "@/app/models"

export const initialStyleProps = {
    position: 'absolute',
    top: '50%',
    left: '50%',
    bottom: 'auto',
    right: 'auto',
    transform: 'translate(-50%, -50%)',
    minWidth: '50%',
    width: 'auto',
    height: 'auto',
    borderRadius: '0.5rem',
    boxShadow: 24,
    overflow: 'hidden',
    p: 4,
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
} as { [name: string]: string | number }

export const initialProps = {
    type: 'info',
    text: '',
    title: '',
    isPermanent: false,
    withActions: false,
    delay: 4000,
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
