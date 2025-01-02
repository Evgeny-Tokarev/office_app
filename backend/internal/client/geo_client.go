package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type GeocodingClient struct {
	apiKey string
}

func NewGeocodingClient(apiKey string) *GeocodingClient {
	return &GeocodingClient{
		apiKey: apiKey,
	}
}

type AddressComponent struct {
	ShortName          string   `json:"short_name"`
	LongName           string   `json:"long_name"`
	Types              []string `json:"types"`
	PostcodeLocalities []string `json:"postcode_localities"`
}

type Geometry struct {
	Location     LatLng       `json:"location"`
	LocationType string       `json:"location_type"`
	Viewport     LatLngBounds `json:"viewport"`
	Bounds       LatLngBounds `json:"bounds"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type LatLngBounds struct {
	NorthEast LatLng `json:"northeast"`
	SouthWest LatLng `json:"southwest"`
}

type GoogleGeocodeResponse struct {
	Types              []string           `json:"types"`
	FormattedAddress   string             `json:"formatted_address"`
	AddressComponents  []AddressComponent `json:"address_components"`
	PartialMatch       bool               `json:"partial_match"`
	PlaceID            string             `json:"place_id"`
	PostcodeLocalities []string           `json:"postcode_localities"`
	Geometry           Geometry           `json:"geometry"`
}

type GeocodingResult struct {
	Results []GoogleGeocodeResponse `json:"results"`
}

func (c *GeocodingClient) GetCoordinates(address string) (*Location, error) {
	encodedAddress := url.QueryEscape(address)
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", encodedAddress, c.apiKey)
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	var geoResp GeocodingResult
	if err := json.NewDecoder(resp.Body).Decode(&geoResp); err != nil {
		return nil, err
	}

	if len(geoResp.Results) == 0 {
		return nil, fmt.Errorf("no results found for address: %s", address)
	}

	location := geoResp.Results[0].Geometry.Location

	return &Location{
		Lat: location.Lat,
		Lng: location.Lng,
	}, nil
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
