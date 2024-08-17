import {Box, Button, Input, InputLabel, TextField} from "@mui/material"
import * as React from "react"
import {createOffice, fetchOffices, OfficesState, saveImage, updateOffice} from "@/app/redux/features/officesSlice"
import {Office} from "@/app/models"
import {ModalContext} from "@/components/ModalProvider"
import {useDispatch, useSelector} from "react-redux"
import {ThunkDispatch} from "@reduxjs/toolkit"
import {AnyAction} from "redux"
import {RootState} from "@/app/redux/store"
import {useTheme as useNextTheme} from "next-themes"


export const style = {
    form: {
        '& .MuiTextField-root': {
            m: 1
        },
        width: '100%',
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
    }
}
export default function OfficeForm({onCloseModal}: {
    onCloseModal: () => void
}) {
    const {
        openModal, setOpenModal, modalProps, setModalProps
    } = React.useContext(ModalContext)
    const {theme} = useNextTheme()
    const titleInput = React.useRef<HTMLInputElement | null>(null)
    const [titleInputError, setTitleInputError] = React.useState(false)
    const [addressInputError, setAddressInputError] = React.useState(false)
    const dispatch = useDispatch<ThunkDispatch<OfficesState, unknown, AnyAction>>()
    const {
        offices, loading, error
    } = useSelector((state: RootState) => state.offices)
    const office: Office | null = React.useMemo(() => {
        if (modalProps.formProps && modalProps.formProps.id < 0) return {
            id: -1,
            name: '',
            address: '',
            created_at: '',
            photo: ''
        }
        if (!offices.length) dispatch(fetchOffices())
        return offices.find(of => of.id === modalProps.formProps?.id) || offices[0]
    }, [offices, modalProps.formProps?.id])
    const [name, setName] = React.useState('')
    const [address, setAddress] = React.useState('')
    const [image, setImage] = React.useState('')
    const [imageFile, setImageFile] = React.useState<File | null>(null)
    React.useEffect(() => {
        if (office) {
            setName(office.name || '')
            setAddress(office.address || '')
            setImage(office.photo ? office.photo : theme === 'dark' ? '/image-placeholder-dk.svg' : '/image-placeholder-lt.svg')
        }
    }, [office, theme])
    const saveOffice = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (name.length < 3 || address.length < 3) {
            setTitleInputError(name.length < 3)
            setAddressInputError(address.length < 3)
            return
        }
        onCloseModal()
        if (modalProps.formProps && modalProps.formProps.id < 0) {
            const res = await dispatch(createOffice({
                name: name,
                address: address,
                id: -1,
                created_at: ''
            }))
            console.log(res)
            if (res.meta.requestStatus === 'fulfilled') {
                if (res.payload && 'id' in res.payload && imageFile) {
                    const id = res.payload?.id
                    await saveImageFile(id)
                }
                await dispatch(fetchOffices())
            }
            return
        }
        if (modalProps.formProps && modalProps.formProps.id) {
            const res = await dispatch(updateOffice({
                id: modalProps.formProps.id,
                name: name,
                address: address,
                created_at: office.created_at,
                updated_at: office.updated_at
            }))
            if (res.meta.requestStatus === 'fulfilled') {
                if (imageFile) await saveImageFile(modalProps.formProps.id)
                await dispatch(fetchOffices())
            }
        }
    }
    const saveImageFile = async (id: number) => {
        if (imageFile) {
            await dispatch(saveImage({
                id: id, image: imageFile
            }))
        }
    }
    const onTitleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        setName(e.target.value)
        setTitleInputError(false)
    }
    const onAddressInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        setAddress(e.target.value)
        setAddressInputError(false)
    }
    const uploadImageFile = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files && e.target.files[0]) {
            setImage(URL.createObjectURL(e.target.files[0]))
            setImageFile(e.target.files[0])
        }

    }
    return (<Box
        component="form"
        onSubmit={saveOffice}
        mt={4}
        sx={modalProps?.type === 'office_form' ? style.form : {"display": "none"}}
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
                error={addressInputError}
                id="address"
                label="Address"
                value={address}
                multiline
                fullWidth
                maxRows={4}
                helperText={addressInputError ? 'Address should have at least 3 characters' : ''}
                onChange={(e) => onAddressInputChange(e)}
            />
            <InputLabel sx={{
                height: '5rem',
                width: '5rem',
                cursor: 'pointer'
            }}>
                <img
                    className="w-full h-full"
                    src={image}
                    alt="Office image"
                    loading="lazy">
                </img>
                <Input
type="file"
sx={style.fileInput}
                       onChange={uploadImageFile}/>
            </InputLabel>
        </div>
        <Button
            variant="contained"
            sx={{mt: 4}}
            type="submit">Save</Button>
    </Box>)
}
