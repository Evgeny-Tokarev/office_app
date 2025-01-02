import { useEffect, useState } from 'react';

interface ILocation {
    lat: number;
    lng: number;
}

type UseLocationReturn = [ILocation | undefined, number | undefined, string | undefined];

const useLocation = (
    enabled: boolean,
    accuracyThreshold?: number,
    accuracyThresholdWaitTime?: number,
    options?: PositionOptions
): [ILocation | undefined, number | undefined, string | undefined] => {
    const [accuracy, setAccuracy] = useState<number>();
    const [location, setLocation] = useState<ILocation>();
    const [error, setError] = useState<string>();

    useEffect(() => {
        if (!enabled) {
            setAccuracy(undefined);
            setError(undefined);
            setLocation(undefined);
            return;
        }

        if (navigator.geolocation) {
            let timeout: NodeJS.Timeout | undefined;
            let initialCoords: ILocation | undefined;
            const geoId = navigator.geolocation.watchPosition(
                (position) => {
                    const lat = position.coords.latitude;
                    const lng = position.coords.longitude;
                    const currentAccuracy = position.coords.accuracy;

                    setAccuracy(currentAccuracy);

                    if (!initialCoords) {
                        initialCoords = { lat, lng };
                        setLocation(initialCoords);
                    }

                    if (!accuracyThreshold || currentAccuracy < accuracyThreshold) {
                        setLocation({ lat, lng });
                    }
                },
                (e) => {
                    setError(e.message);
                },
                options ?? { enableHighAccuracy: true, maximumAge: 1000, timeout: 10000 }
            );

            if (accuracyThreshold && accuracyThresholdWaitTime) {
                timeout = setTimeout(() => {
                    if (initialCoords && (!accuracy || accuracy >= accuracyThreshold)) {
                        setLocation(initialCoords);
                        setError('Failed to reach desired accuracy within time limit');
                    }
                }, accuracyThresholdWaitTime * 1000);
            }

            return () => {
                window.navigator.geolocation.clearWatch(geoId);
                if (timeout) {
                    clearTimeout(timeout);
                }
            };
        } else {
            setError('Geolocation API not available');
        }
    }, [enabled, accuracyThreshold, accuracyThresholdWaitTime, options]);

    if (!enabled) {
        return [undefined, undefined, undefined];
    }

    return [location, accuracy, error];
};



export default useLocation;
