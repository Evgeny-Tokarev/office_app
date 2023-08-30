import './globals.css'
import {Providers} from "@/app/redux/provider";
import React from 'react';
import type {Metadata} from 'next';
import {Lato} from "next/font/google";
import ThemeRegistry from "@/app/registry";
import Navbar from "@/components/Navbar";
import {Paper, Typography} from "@mui/material";

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
            <Navbar/>
            <Paper
                sx={{
                    flexGrow: 1
                }}>
                {children}
            </Paper>
        </ThemeRegistry>
    </Providers>
    </body>
    </html>)
}
