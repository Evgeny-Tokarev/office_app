import {
    Box,
    Button,
    TextField,
    Input,
    InputLabel
} from "@mui/material"
import * as React from "react";
import {
    createEmployee, fetchEmployees, EmployeesState, updateEmployee, saveImage
} from "@/app/redux/features/employeeSlice";
import {Employee} from "@/app/models"
import {ModalContext} from "@/components/ModalProvider";
import {useDispatch, useSelector} from "react-redux";
import {ThunkDispatch} from "@reduxjs/toolkit";
import {AnyAction} from "redux";
import {RootState} from "@/app/redux/store";
import {useTheme as useNextTheme} from "next-themes"
import {style} from "./OfficeForm"

export default function EmployeeForm({onCloseModal}: {
    onCloseModal: () => void
}) {
    const {theme} = useNextTheme()
    const titleInput = React.useRef<HTMLInputElement | null>(null)
    const [titleInputError, setTitleInputError] = React.useState(false)
    const [ageInputError, setAgeInputError] = React.useState(false)

    const {
        modalProps
    } = React.useContext(ModalContext)
    const dispatch = useDispatch<ThunkDispatch<EmployeesState, unknown, AnyAction>>()
    const {
        employees} = useSelector((state: RootState) => state.employees)
    const employee: Employee | null = React.useMemo(() => {
        if (modalProps.formProps && modalProps.formProps.id < 0 && modalProps.formProps?.office_id) {
            return {
                office_id: modalProps.formProps.office_id,
                id: -1,
                name: '',
                age: 0,
                created_at: '',
                photo: ''
            }
        }
        if (!employees.length && modalProps.formProps && modalProps.formProps.office_id !== undefined) dispatch(fetchEmployees(modalProps.formProps.office_id))
        return employees.find(employee => employee.id === modalProps.formProps?.id) || employees[0]
    }, [employees, modalProps.formProps?.id])
    const [name, setName] = React.useState('')
    const [age, setAge] = React.useState(0)
    const [image, setImage] = React.useState('')
    const [imageFile, setImageFile] = React.useState<File | null>(null)
    React.useEffect(() => {
        if (employee) {
            setName(employee.name || '')
            setAge(employee.age || 0)
            setImage(employee.photo
                ? employee.photo
                : theme === 'dark'
                    ? '/image-placeholder-dk.svg'
                    : '/image-placeholder-lt.svg')
        }
    }, [employee, theme])
    const saveEmployee = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (name.length < 3 || age <= 0 || age > 120) {
            setTitleInputError(name.length < 3)
            setAgeInputError(age <=0 && age > 120)
            return
        }
        if (modalProps.formProps && modalProps.formProps.id < 0 && modalProps.formProps?.office_id) {
            const res = await dispatch(createEmployee({
                office_id: modalProps.formProps.office_id,
                name: name,
                age: age,
                id: -1,
                created_at: ''
            }))
            if (res.meta.requestStatus === 'fulfilled') {
                if (res.payload && 'id' in res.payload && imageFile) {
                    const id = res.payload?.id
                    await saveImageFile(id)
                }
                await dispatch(fetchEmployees(modalProps.formProps.office_id))
                onCloseModal()
            }
            return
        }
        if (modalProps.formProps && modalProps.formProps.id && modalProps.formProps?.office_id) {
            const res = await dispatch(updateEmployee({
                id: modalProps.formProps.id,
                office_id: modalProps.formProps.office_id,
                name: name,
                age: age,
                created_at: employee.created_at,
                updated_at: employee.updated_at
            }))
            if (res.meta.requestStatus === 'fulfilled') {
                if (imageFile) await saveImageFile(modalProps.formProps.id)
                await dispatch(fetchEmployees(modalProps.formProps.office_id))
                onCloseModal()
            }
        }
    }
    const saveImageFile = async(id: number) => {
        if (imageFile) {
            await dispatch(saveImage({id: id, image: imageFile}))
        }
    }
    const onTitleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        setName(e.target.value)
        setTitleInputError(false)
    }
    const onAgeInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        setAge(+e.target.value)
        setAgeInputError(false)
    }
    const uploadImageFile = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files && e.target.files[0]) {
            setImage(URL.createObjectURL(e.target.files[0]))
            setImageFile(e.target.files[0])
        }

    }
    return (<Box
        component="form"
        onSubmit={saveEmployee}
        mt={4}
        sx={modalProps?.type === 'employee_form' ? style.form : {"display": "none"}}
        noValidate
        autoComplete="off"
    >
        <div
            className="flex flex-col items-center w-full">
            <TextField
                error={titleInputError}
                id="name"
                label="Name"
                value={name}
                fullWidth
                inputRef={titleInput}
                helperText={titleInputError ? 'Name should have at least 3 characters' : ''}
                onChange={(e) => onTitleInputChange(e)}
            />
            <TextField
                error={ageInputError}
                id="address"
                label="Age"
                value={age > 0 ? age : ''}
                multiline
                fullWidth
                maxRows={4}
                helperText={ageInputError ? 'Age should be positive integer less than or equal 120' : ''}
                onChange={(e) => onAgeInputChange(e)}
            />
            <InputLabel sx={{height: '5rem', width: '5rem', cursor: 'pointer'}}>
                <img
                    className="w-full h-full"
                    src={image}
                    alt="Office image"
                    loading="lazy">
                </img>
                <Input type="file" sx={style.fileInput}
                       onChange={uploadImageFile}/>
            </InputLabel>
        </div>
        <Button
            variant="contained"
            sx={{mt: 4}}
            type="submit">Save</Button>
    </Box>)
}
