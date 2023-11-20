"use client"

import React from "react";
import {ModalContext} from "@/components/ModalProvider";
import {useDispatch, useSelector} from "react-redux";
import {ThunkDispatch} from "@reduxjs/toolkit";
import {
    fetchEmployees, EmployeesState, createEmployee
} from "@/app/redux/features/employeeSlice"
import {fetchOffice} from "@/app/redux/features/officesSlice"
import {AnyAction} from "redux";
import {RootState} from "@/app/redux/store"
import {Employee} from "@/app/models";
import {IconButton, List, Typography} from "@mui/material";
import {AddCircleOutlined} from "@mui/icons-material"
import dynamic from "next/dynamic";

const NoSSREmployeeItem = dynamic(() => import("@/components/EmployeeItem"), {ssr: false})

export default function Employees({params}: {
    params: { id: number }
}) {
    const {
        setOpenModal, setModalProps
    } = React.useContext(ModalContext)
    const dispatch = useDispatch<ThunkDispatch<EmployeesState, unknown, AnyAction>>()
    const {
        employees
    } = useSelector((state: RootState) => state.employees)
    const {
        currentOffice
    } = useSelector((state: RootState) => state.offices)


    React.useEffect(() => {
        dispatch(fetchOffice(params.id))
        dispatch(fetchEmployees(params.id))
    }, [])

    const createNewEmployee = () => {
        const cb = async (props: unknown) => {
            if (props && typeof props === 'object') {
                const partialEmployee: Partial<Employee> = {
                    id: -1,
                    name: "",
                    age: 0,
                    office_id: params.id,
                    created_at: "", ...props as Partial<Employee>
                }
                const res = await dispatch(createEmployee(partialEmployee as Employee))
                if (res.meta.requestStatus === 'fulfilled') {
                    dispatch(fetchEmployees(params.id))
                    setOpenModal(false)
                }
            }
        }
        setOpenModal(true)
        setModalProps({
            type: 'employee_form',
            title: `Create new employee record`,
            isPermanent: true,
            actionCallback: cb,
            formProps: {
                id: -1,
                office_id: params.id
            }
        })
    }
    return (<List sx={{
        width: '100%',
        bgcolor: 'background.paper',
        minHeight: '100vh'
    }}>
        <Typography
            align="center"
            variant="h2"
            mt={2}
            sx={{fontWeight: '400'}}>
            {currentOffice?.name ?? 'tr'}
        </Typography>
        <div className="flex items-center">
            <IconButton
                size="large"
                edge="start"
                aria-label="create new office"
                sx={{ml: 'auto', mr: 4}}
                onClick={createNewEmployee}
            >
                <AddCircleOutlined sx={{fontSize: '4rem'}}/>
            </IconButton>
        </div>
        {employees.map(employee => <NoSSREmployeeItem
            employee={employee}
            key={employee.id}/>)}
    </List>)
}
