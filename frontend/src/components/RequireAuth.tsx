"use client"

import React, {useEffect, useRef, useState} from 'react'
import {useDispatch, useSelector} from "react-redux"
import {AppDispatch, RootState} from "@/app/redux/store"
import {usePathname, useRouter} from "next/navigation"
import {getCurrentUser, setNullCurrentUser} from "@/app/redux/features/usersSlice"

const RequireAuth: React.FC<{ children: React.ReactNode }> = ({children}) => {
    const currentUser = useSelector((state: RootState) => state.users.currentUser)
    const [authChecked, setAuthChecked] = useState(false)
    const loading = useSelector((state: RootState) => state.utils.loading)
    const router = useRouter()
    const pathname = usePathname()
    const dispatch = useDispatch<AppDispatch>()
    const showContent = useRef(false)

    useEffect(() => {
        const token = localStorage.getItem('officeAppToken')
        if (token) {
            dispatch(getCurrentUser(token)).finally(() => setAuthChecked(true))
        } else {
            dispatch(setNullCurrentUser())
            setAuthChecked(true)
        }
    }, [])

    useEffect(() => {
        if (!authChecked) return

        if (currentUser) {
            if (pathname === '/login') {
                const redirectPath = sessionStorage.getItem("redirectAfterLogin") || '/'
                sessionStorage.removeItem("redirectAfterLogin")
                router.push(redirectPath)
            }
        } else {
            const currentPath = window.location.pathname
            if (currentPath !== '/login') {
                sessionStorage.setItem("redirectAfterLogin", currentPath)
            }
            router.push('/login')
        }
    }, [currentUser, router, authChecked])

    if (!authChecked) return
    return <>{children}</>
}

export default RequireAuth