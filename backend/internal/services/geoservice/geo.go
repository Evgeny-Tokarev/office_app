package geoservice

import (
	"encoding/json"
	"fmt"
	"github.com/evgeny-tokarev/office_app/backend/internal/client"
)

type GeoService struct {
	GeocodingClient *client.GeocodingClient
}

func NewGeoService(geoClient *client.GeocodingClient) *GeoService {
	return &GeoService{
		GeocodingClient: geoClient,
	}
}

func (gs *GeoService) GetCoordinates(address string) (json.RawMessage, error) {
	location, err := gs.GeocodingClient.GetCoordinates(address)
	if err != nil {
		return nil, fmt.Errorf("error getting coordinates for address %s: %v", address, err)
	}

	if location.Lat == 0 && location.Lng == 0 {
		return nil, fmt.Errorf("invalid coordinates for address %s", address)
	}

	jsonLocation, err := json.Marshal(location)
	if err != nil {
		return nil, err
	}

	return jsonLocation, nil
}
