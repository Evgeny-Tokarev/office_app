'use client'

import React, {useEffect} from 'react'
import {useServerInsertedHTML} from 'next/navigation'
import createCache, {type Options} from '@emotion/cache'
import {CacheProvider} from '@emotion/react'
import {
    ThemeProvider as MuiThemeProvider, THEME_ID
} from '@mui/material/styles'
import {lightTheme, darkTheme} from "@/app/themes/themes"
import {
    ThemeProvider as NextThemeProvider,
    useTheme as useNextTheme
} from "next-themes"
import {CssBaseline} from "@mui/material";

export default function ThemeRegistry(props: {
    children: React.ReactNode, options: Options
}) {
    const {options, children} = props

    const [{cache, flush}] = React.useState(() => {
        const cache = createCache(options)
        cache.compat = true
        const prevInsert = cache.insert
        let inserted: string[] = []
        cache.insert = (...args) => {
            const serialized = args[1]
            if (cache.inserted[serialized.name] === undefined) {
                inserted.push(serialized.name)
            }
            return prevInsert(...args)
        }
        const flush = () => {
            const prevInserted = inserted
            inserted = []
            return prevInserted
        }
        return {cache, flush}
    })

    useServerInsertedHTML(() => {
        const names = flush()
        if (names.length === 0) {
            return null
        }
        let styles = ''
        for (const name of names) {
            styles += cache.inserted[name]
        }
        return (<style
            key={cache.key}
            data-emotion={`${cache.key} ${names.join(' ')}`}
            dangerouslySetInnerHTML={{
                __html: options.prepend ? `@layer emotion {${styles}}` : styles,
            }}
        />)
    })


    return (<CacheProvider value={cache}>
        <NextThemeProvider
            attribute="class">
            <MTP>{children}</MTP>
        </NextThemeProvider>
    </CacheProvider>)
}

function MTP({
                 children,
             }: {
    children: React.ReactNode
}) {
    const {theme} = useNextTheme()
    const [mounted, setMounted] = React.useState(false)
    const [mTheme, setMTheme] = React.useState(theme === 'light' || theme === '' ? lightTheme : darkTheme)
    useEffect(() => {
        setMTheme(theme === 'light' || theme === '' ? lightTheme : darkTheme)
        setMounted(true)
    }, [])

    useEffect(() => {
        setMTheme(theme === 'light' || theme === '' ? lightTheme : darkTheme)
    }, [theme])
    if (!mounted) return null
    return (<MuiThemeProvider
        theme={{ [THEME_ID]: mTheme }}>
        <CssBaseline/>
        {children}
    </MuiThemeProvider>)

}
