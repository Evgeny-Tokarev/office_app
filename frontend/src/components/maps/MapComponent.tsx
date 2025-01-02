/// <reference types="@types/google.maps" />
"use client"

import React, {useRef, useState, useEffect} from 'react'
import {Map, useMap, type MapMouseEvent} from '@vis.gl/react-google-maps'
import PoiMarkers, {type Poi} from '@/components/maps/PoiMarker'
import useLocation from "@/app/hooks/useLocation";

export default function MapComponent() {
    const [pois, setPois] = useState<Poi[]>([])
    const [homePois, setHomePois] = useState<Poi[]>([])
    const map = useMap()
    const dataLayer = useRef<google.maps.Data | null>(null);
    const [currentPosition, setCurrentPosition] = useState<google.maps.LatLngLiteral>({lat: 0, lng: 0})
    const [geoJsonData, setGeoJsonData] = useState<GeoJSON.FeatureCollection | null>(null);
    const [location, accuracy, error] = useLocation(true, 100, 10);

    useEffect(() => {
        console.log(location, error, accuracy)
        if (location?.lat && location.lng) {
            setCurrentPosition({
                lat: location.lat,
                lng: location.lng
            })

            setHomePois([{key: "home", location: {lat: location.lat, lng: location.lng}}])
        }

    }, [location, error, accuracy])

    React.useEffect(() => {
        const fetchGeoJson = async () => {
            try {
                const response = await fetch("/geoData/manitoba.geojson");
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const data = await response.json();
                setGeoJsonData(data);
                console.log("GeoJSON Data:", data);
            } catch (error) {
                console.error("Error fetching GeoJSON:", error);
            }
        }

        fetchGeoJson();
    }, [])

    React.useEffect(() => {
        if (!map) return
        console.log(map, geoJsonData)
    }, [map])
    const addMarker = (e: MapMouseEvent) => {
        const latLng = e.detail.latLng;
        if (!latLng) {
            console.warn("No latLng found in the click event.", latLng);
            return;
        }

        setPois((prev) => [
            ...prev,
            {
                key: `poi-${prev.length + 1}`,
                location: { lat: latLng.lat, lng: latLng.lng },
            },
        ]);
    };

    //function to highlights predefined area using geoJson
    const highlightArea = () => {
        if (!map) return;
        dataLayer.current = new google.maps.Data();
        dataLayer.current.addGeoJson(geoJsonData as object);

        dataLayer.current.setStyle({
            fillColor: 'blue',
            fillOpacity: 0.4,
            strokeColor: 'blue',
            strokeWeight: 2,
        });

        dataLayer.current.setMap(map);
    };
    if (currentPosition.lng === 0 && currentPosition.lat === 0) return null
    return (
        <Map
            style={{width: '100%', height: 'auto', flex: '1 1 100%'}}
            mapId='process.env.NEXT_PUBLIC_MAP_ID'
            defaultCenter={{lat: currentPosition.lat, lng: currentPosition.lng}}
            defaultZoom={12}
            gestureHandling={'greedy'}
            disableDefaultUI={true}
            onClick={addMarker}
        >
            <PoiMarkers pois={pois}/>
            <PoiMarkers pois={homePois} pinStyleName={'home'}/>
        </Map>
    )
}