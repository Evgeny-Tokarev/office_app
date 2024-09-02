"use client"

import React from 'react'
import {IconButton, List} from "@mui/material"
import {useDispatch, useSelector} from "react-redux"
import {fetchOffices} from "@/app/redux/features/officesSlice"
import {AddCircleOutlined} from "@mui/icons-material"
import {ModalContext} from "@/components/ModalProvider"
import {AppDispatch, RootState} from "@/app/redux/store"
import dynamic from 'next/dynamic'

const NoSSROfficeItem = dynamic(() => import("@/components/OfficeItem"), {ssr: false})

export default function Page() {
    const {
        setOpenModal, setModalProps
    } = React.useContext(ModalContext)
    const dispatch = useDispatch<AppDispatch>()
    const {
        offices
    } = useSelector((state: RootState) => state.offices)
    React.useEffect(() => {
        dispatch(fetchOffices())
    }, [])

    const createNewOffice = () => {
        setOpenModal(true)
        setModalProps({
            type: 'office_form',
            title: `Create new office`,
            isPermanent: true,
            formProps: {
                id: -1
            }
        })
    }

    return (<List sx={{
        width: '100%',
        bgcolor: 'background.paper',
        minHeight: '100vh'
    }}>
        <div className="flex items-center">
            <IconButton
                size="large"
                edge="start"
                aria-label="create new office"
                sx={{ml: 'auto', mr: 4}}
                onClick={createNewOffice}
            >
                <AddCircleOutlined sx={{fontSize: '4rem'}}/>
            </IconButton>
        </div>
        {offices.map(office => <NoSSROfficeItem
            office={office}
            key={office.id}/>)}
    </List>)
}
