'use client'

import React, {useEffect, useState} from 'react'
import {useServerInsertedHTML} from 'next/navigation'
import createCache, {type Options} from '@emotion/cache'
import {CacheProvider} from '@emotion/react'
import {Theme, THEME_ID, ThemeProvider as MuiThemeProvider} from '@mui/material/styles'
import {darkTheme, lightTheme} from "@/app/themes/themes"
import {ThemeProvider as NextThemeProvider, useTheme as useNextTheme} from "next-themes"
import {CssBaseline} from "@mui/material"

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
    // this part uses autodetection for the system(browser?) theme
    const {theme} = useNextTheme()
    const [mounted, setMounted] = React.useState(false)
    const [mTheme, setMTheme] = React.useState(theme === '' ? lightTheme : darkTheme)
    const [prefersDarkMode, setPrefersDarkMode] = useState(false)
    const themeMap = {
        'light': lightTheme,
        'dark': darkTheme,
        'system': prefersDarkMode ? darkTheme : lightTheme
    } as { [key: string]: Theme }

    useEffect(() => {
        const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
        setPrefersDarkMode(mediaQuery.matches)

        const handleChange = (e: MediaQueryListEvent) => setPrefersDarkMode(e.matches);
        mediaQuery.addEventListener('change', handleChange);
        setMTheme(themeMap[theme ?? 'system'] )
        setMounted(true)

        return () => mediaQuery.removeEventListener('change', handleChange);
    }, [])

    useEffect(() => {
        console.log(theme)
        setMTheme(themeMap[theme ?? 'system'] )
    }, [theme])
    if (!mounted) return null
    return (<MuiThemeProvider
        theme={{ [THEME_ID]: mTheme }}>
        <CssBaseline/>
        {children}
    </MuiThemeProvider>)

    // const {resolvedTheme} = useNextTheme()
    // const [mounted, setMounted] = React.useState(false)
    //
    // useEffect(() => {
    //     setMounted(true)
    // }, [])
    //
    // const selectedTheme = resolvedTheme === 'dark' ? darkTheme : lightTheme
    //
    // if (!mounted) return null
    //
    // return (
    //     <MuiThemeProvider theme={{[THEME_ID]: selectedTheme}}>
    //         <CssBaseline/>
    //         {children}
    //     </MuiThemeProvider>
    // )
}
