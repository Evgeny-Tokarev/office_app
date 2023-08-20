import './globals.css'
import React from 'react'
import type {Metadata} from 'next'
import {Lato} from "next/font/google";
import ThemeRegistry from "@/app/registry";

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
            <ThemeRegistry
            options={{key: 'mui'}}>{children}</ThemeRegistry>
        </body>
    </html>)
}
