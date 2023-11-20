import './globals.css'
import {Providers} from "@/app/redux/provider"
import React from 'react'
import type {Metadata} from 'next'
import {Lato} from "next/font/google"
import ThemeRegistry from "@/app/registry"
import Navbar from "@/components/Navbar"
import Loader from "@/components/Loader"
import {Paper} from "@mui/material"
import {
    ModalContextProvider
} from "@/components/ModalProvider"
import Modal from "@/components/modal/Modal"
import {
    LoaderContextProvider
} from "@/components/LoaderProvider"

const lato = Lato({
    weight: ['400', '700'],
    subsets: ['latin'],
    display: "swap"
})


export const metadata: Metadata = {
    title: 'Dashboard', description: 'Test dashboard',
}


export default function RootLayout({
                                       children,
                                   }: {
    children: React.ReactNode
}) {
    return (<html
        lang="en"
        suppressHydrationWarning={true}>
        <body
        className={lato.className}
        suppressHydrationWarning={true}>
            <Providers>
                <ThemeRegistry
            options={{key: 'mui'}}>
                    <ModalContextProvider>
                        <LoaderContextProvider>
                            <Modal/>
                            <Navbar/>
                            <Loader />
                            <div className="flex flex-col justify-between items-stretch flex-1">

                                <Paper
                                    sx={{
                                        flexGrow: 1,
                                        display: 'flex',
                                        flexDirection: 'column'
                                    }}>
                            {children}
                                    </Paper>
                            </div>
                        </LoaderContextProvider>
                    </ModalContextProvider>
                </ThemeRegistry>
            </Providers>
        </body>
    </html>)
}
