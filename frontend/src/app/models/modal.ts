import React from "react"

export type ModalType =
    'info'
    | 'office_form'
    | 'warn'
    | 'employee_form'
    | 'user_form'

export interface ModalContextType {
    openModal: boolean;
    setOpenModal: React.Dispatch<React.SetStateAction<boolean>>;
    modalProps: ModalProps,
    setModalProps: React.Dispatch<React.SetStateAction<ModalProps>>,
}

export interface ModalProps {
    type: ModalType,
    text?: string,
    title?: string,
    isPermanent?: boolean,
    withActions?: boolean,
    style?: { [name: string]: string | number },
    actionCallback?: (props?: any) => void | Promise<void>
    delay?: number,
    closable?: boolean,
    formProps?: {
        id: number, office_id?: number
    }
}
