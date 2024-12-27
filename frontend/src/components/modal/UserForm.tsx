"use client"

import * as React from "react"
import {MouseEvent} from "react"
import {useDispatch} from "react-redux"
import {useTheme as useNextTheme} from "next-themes"
import {AppDispatch} from "@/app/redux/store"
import {login, register} from "@/app/redux/features/usersSlice"
import {Autocomplete, Box, Button, TextField} from "@mui/material"
import {useRouter} from "next/navigation"


const emailReg = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

export const style = {
    form: {
        '& .MuiTextField-root': {
            m: 1,
        },
        width: '100%',
        maxWidth: '30rem',
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center'
    }, fileInput: {
        clip: 'rect(0 0 0 0)',
        clipPath: 'inset(50%)',
        height: 1,
        overflow: 'hidden',
        position: 'absolute',
        bottom: 0,
        left: 0,
        whiteSpace: 'nowrap',
        width: 1
    },
    buttonWrapper: {
        width: '100%',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between'
    }
}

const roles = [
    'user', 'admin'
]
export default function UserForm({onCloseModal}: {
    onCloseModal: () => void
}) {
    // const {
    //     openModal, setOpenModal, modalProps, setModalProps
    // } = React.useContext(ModalContext)
    const {theme} = useNextTheme()
    const [roleValue, setRoleValue] = React.useState<string>(roles[0])
    const [showRoles, setShowRoles] = React.useState<boolean>(false)
    const titleInput = React.useRef<HTMLInputElement | null>(null)
    const [nameInputError, setNameInputError] = React.useState(false)
    const [emailInputError, setEmailInputError] = React.useState(false)
    const [passwordInputError, setPasswordInputError] = React.useState(false)
    const dispatch = useDispatch<AppDispatch>()
    const router = useRouter()

    // const user: User | null = React.useMemo(() => {
    //     if (modalProps.formProps && modalProps.formProps.id < 0) return {
    //         id: -1,
    //         name: '',
    //         email: '',
    //         password: ''
    //     }
    //     return null
    // }, [modalProps.formProps?.id])

    const [name, setName] = React.useState('')
    const [password, setPassword] = React.useState('')
    const [email, setEmail] = React.useState('')


    const handleFocus = () => {
        setNameInputError(false)
        setEmailInputError(false)
        setPasswordInputError(false)
    }
    const checkForErrors = () => {
        if (name.length < 3 || !emailReg.test(email) || password.length < 3) {
            setNameInputError(name.length < 3)
            setEmailInputError(!emailReg.test(email))
            setPasswordInputError(password.length < 3)
            return true
        }
        return false
    }

    const onLoginUserClick = async (e: MouseEvent) => {
        e.preventDefault()
        if (!showRoles) {
            if (checkForErrors()) return
            await dispatch(login({userName: name, email: email, password: password}))
        } else setShowRoles(false)
    }

    const onRegisterUserClick = async (e: MouseEvent) => {
        e.preventDefault()
        if (showRoles) {
            if (checkForErrors()) return
            await dispatch(register({userName: name, email: email, password: password, role: roleValue}))
        } else setShowRoles(true)
    }
    const onNameInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        setName(e.target.value)
        setNameInputError(false)
    }
    const onEmailInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        setEmail(e.target.value)
        setEmailInputError(false)
    }
    const onPasswordInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        setPassword(e.target.value)
        setPasswordInputError(false)
    }

    const onRoleChange = (e: React.SyntheticEvent, value: string) => {
        console.log(value)
    }

    return (<Box
        component="form"
        mt={4}
        sx={style.form}
        noValidate
        autoComplete="off"
    >
        <div
            className="flex flex-col items-center w-full">
            <TextField
                error={nameInputError}
                id="name"
                label="Name"
                value={name}
                fullWidth
                inputRef={titleInput}
                helperText={nameInputError ? 'Name should have at least 3 characters' : ' '}
                onChange={(e) => onNameInputChange(e)}
                onFocus={handleFocus}
            />
            <TextField
                error={emailInputError}
                id="email"
                label="Email"
                value={email}
                multiline
                fullWidth
                maxRows={4}
                helperText={emailInputError ? 'Address should have at least 3 characters' : ' '}
                onChange={(e) => onEmailInputChange(e)}
                onFocus={handleFocus}
            />
            <TextField
                error={passwordInputError}
                id="password"
                label="Password"
                value={password}
                multiline
                fullWidth
                maxRows={4}
                helperText={emailInputError ? 'Password should have at least 3 characters' : ' '}
                onChange={(e) => onPasswordInputChange(e)}
                onFocus={handleFocus}
            />
            {showRoles && <Autocomplete
                disableClearable={true}
                onChange={(event: any, newValue: string) => {
                    setRoleValue(newValue)
                }}
                renderOption={(props, option) => {
                    return (
                        <li
                            {...props}
                            key={option}>
                            {option}
                        </li>
                    )
                }}
                renderInput={(params) =>
                    <TextField
                        {...params}
                        label="Role"/>}
                options={roles}
                value={roleValue}
                id="role"
                fullWidth
            />}

        </div>
        <Box
            component="div"
            mt={4}
            sx={style.buttonWrapper}
        >
            <Button
                component="button"
                variant={!showRoles ? "contained" : "outlined"}
                onClick={onLoginUserClick}>Login</Button>
            <Button
                component="button"
                variant={showRoles ? "contained" : "outlined"}
                onClick={onRegisterUserClick}>Register</Button>
        </Box>
    </Box>)
}
