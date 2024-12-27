import React from 'react'
import {Map, useMap} from '@vis.gl/react-google-maps'


export default function MapComponent() {
    const map = useMap()

    React.useEffect(() => {
        if (!map) return
        console.log(map)
    }, [map])
    return (
        <Map
            style={{width: '100%', height: 'auto', flex: '1 1 100%'}}
            defaultCenter={{lat: 49.897029828583435, lng: -97.13808431942881}}
            defaultZoom={12}
            gestureHandling={'greedy'}
            disableDefaultUI={true}
        />
    )
}