package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Point struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func readCSVFile(filename string) ([]Point, error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the CSV file and convert each row to a Point struct
	points := []Point{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Convert the latitude and longitude strings to float64
		lat, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			return nil, err
		}
		lng, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}

		// Create a new Point struct and append it to the slice
		point := Point{
			Latitude:  lat,
			Longitude: lng,
		}
		points = append(points, point)
	}

	return points, nil
}

func toGeoJSON(points []Point) ([]byte, error) {
	// Create a GeoJSON FeatureCollection
	featureCollection := struct {
		Type     string        `json:"type"`
		Features []interface{} `json:"features"`
	}{
		Type:     "FeatureCollection",
		Features: make([]interface{}, len(points)),
	}

	// Convert each Point struct to a GeoJSON Feature
	for i, point := range points {
		feature := struct {
			Type       string      `json:"type"`
			Properties interface{} `json:"properties"`
			Geometry   struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"geometry"`
		}{
			Type: "Feature",
			Properties: struct {
				ID int `json:"id"`
			}{
				ID: i + 1,
			},
			Geometry: struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			}{
				Type:        "Point",
				Coordinates: []float64{point.Longitude, point.Latitude},
			},
		}
		featureCollection.Features[i] = feature
	}

	// Convert the FeatureCollection to JSON
	jsonBytes, err := json.Marshal(featureCollection)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func main() {
	// Read the CSV file
	points, err := readCSVFile("points.csv")
	if err != nil {
		panic(err)
	}

	// Convert the Points to GeoJSON
	geoJSON, err := toGeoJSON(points)
	if err != nil {
		panic(err)
	}

	// Write the GeoJSON to a file
	file, err := os.Create("points.geojson")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(geoJSON)
	if err != nil {
		panic(err)
	}

	fmt.Println("Conversion complete!")
}