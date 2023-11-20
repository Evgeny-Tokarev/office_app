"use client"
import React, { useState, createContext } from "react"

export interface LoaderContextType {
    showLoader: boolean;
    setShowLoader: React.Dispatch<React.SetStateAction<boolean>>;
}

export const LoaderContext = createContext<LoaderContextType>({showLoader: false,
setShowLoader: () => {}})

export function LoaderContextProvider({ children }: {
    children: React.ReactNode
}) {
    const [showLoader, setShowLoader] = useState<boolean>(false)

    return (
        <LoaderContext.Provider
            value={{showLoader, setShowLoader
                }}
        >
            {children}
        </LoaderContext.Provider>
    )
}
