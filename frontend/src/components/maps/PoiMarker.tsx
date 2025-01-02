/// <reference types="@types/google.maps" />
import React, {useEffect, useState, useRef} from 'react';
import {
    AdvancedMarker,
    useMap,
    Pin
} from '@vis.gl/react-google-maps';
import {MarkerClusterer} from '@googlemaps/markerclusterer';
import type {Marker} from '@googlemaps/markerclusterer';

export type Poi = { key: string, location: google.maps.LatLngLiteral }

interface PinStyles {
    background: string;
    glyphColor: string;
    borderColor: string;
}

export type PinStyleName = 'default' | 'home' | 'warning'

const pinStyles: Record<PinStyleName, PinStyles> = {
    default: {
        background: '#FBBC04',
        glyphColor: '#000',
        borderColor: '#000',
    },
    home: {
        background: '#34A853',
        glyphColor: '#ffffff',
        borderColor: '#000',
    },
    warning: {
        background: '#EA4335',
        glyphColor: '#ffffff',
        borderColor: '#ff0000',
    },
};
export default function PoiMarkers({pois, pinStyleName = 'default'}: {
    pois: Poi[];
    pinStyleName?: PinStyleName;
}) {
    const map = useMap();
    const [markers, setMarkers] = useState<{ [key: string]: Marker }>({});
    const clusterer = useRef<MarkerClusterer | null>(null);

    // Initialize MarkerClusterer, if the map has changed
    useEffect(() => {
        if (!map) return;
        if (!clusterer.current) {
            clusterer.current = new MarkerClusterer({map});
        }
    }, [map]);

    // Update markers, if the markers array has changed
    useEffect(() => {
        clusterer.current?.clearMarkers();
        clusterer.current?.addMarkers(Object.values(markers));
    }, [markers]);

    const setMarkerRef = (marker: Marker | null, key: string) => {
        if (marker && markers[key]) return;
        if (!marker && !markers[key]) return;

        setMarkers(prev => {
            if (marker) {
                return {...prev, [key]: marker};
            } else {
                const newMarkers = {...prev};
                delete newMarkers[key];
                return newMarkers;
            }
        });
    };

    return (
        <>
            {pois.map((poi: Poi) => (
                <AdvancedMarker
                    key={poi.key}
                    position={poi.location}
                    ref={marker => setMarkerRef(marker, poi.key)}
                >
                    <Pin {...pinStyles[pinStyleName]}/>
                </AdvancedMarker>
            ))}
        </>
    );
};