import React from 'react'
import {APIProvider} from '@vis.gl/react-google-maps'
import MapComponent from '@/components/maps/MapComponent'

export default function Dashboard() {
    const apiKey = process.env.NEXT_PUBLIC_GOOGLE_MAPS_API_KEY
    console.log(apiKey)
    return (<div className="flex grow w-full h-full">
        <APIProvider apiKey={process.env.NEXT_PUBLIC_GOOGLE_MAPS_API_KEY || ""}>
            <MapComponent/>
        </APIProvider>
    </div>)
}
