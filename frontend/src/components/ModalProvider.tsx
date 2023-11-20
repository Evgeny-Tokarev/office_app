"use client"
import React, { useState, createContext } from "react"
import {ModalProps, ModalContextType} from "@/app/models"

export const initialProps = {type: 'info', text: '', title: '', isPermanent: false, withActions: false, delay: 4000} as ModalProps
export const ModalContext = createContext<ModalContextType>({openModal: false, setOpenModal: () => {}, modalProps: initialProps, setModalProps: () => {}})

export function ModalContextProvider({ children }: {
    children: React.ReactNode
}) {
    const [openModal, setOpenModal] = useState<boolean>(false)
    const [modalProps, setModalProps] = useState<ModalProps>({type: 'info', text: '', title: '', isPermanent: false, withActions: false, delay: 4000})

    return (
        <ModalContext.Provider
            value={{
                openModal,
                setOpenModal,
                modalProps,
                setModalProps}}
        >
            {children}
        </ModalContext.Provider>
    )
}
