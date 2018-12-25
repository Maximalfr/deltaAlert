package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"./haversine"
	"./jsonparser"
)

func main() {
	argsWithoutProg := os.Args[1:]

	switch len(argsWithoutProg) {
	case 0:
		log.Fatal("Please, specify the filename")
		break
	case 1:
		data := getData(argsWithoutProg[0])
		getDelta(&data)
		break
	case 2: // Display the delta from Matchedpositions if two jsons are put in command line arguments.
		data1 := getData(argsWithoutProg[0])
		data2 := getData(argsWithoutProg[1])
		compareMatchedPositions(&data1, &data2)

	}

}

// Call the JsonRead from the parser and handle errors
func getData(filename string) jsonparser.GpsAlert {
	data, err := jsonparser.JsonRead(filename)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// Get the delta between the Matchedposition and the first Geodisplay.
func getDelta(data *jsonparser.GpsAlert) {
	lat := data.MatchedPosition.Latitude
	long := data.MatchedPosition.Longitude
	heading := data.MatchedPosition.Heading

	coordGeo := data.GeoDisplay.Geometry.Coordinates
	latGeo := coordGeo[1]
	longGeo := coordGeo[0]

	p1 := haversine.Coord{lat, long}
	p2 := haversine.Coord{latGeo, longGeo}

	_, km := haversine.Distance(p1, p2)

	displayMatchedPosition(lat, long, heading)
	displayGeo(latGeo, longGeo)
	displayDelta(km * 1000)
}

// Return the delta (in meters) between two positions.
// Take only Matchedposition points to calculate the delta.
func compareMatchedPositions(data1 *jsonparser.GpsAlert, data2 *jsonparser.GpsAlert) {
	// Sorry it's ugly
	lat1 := data1.MatchedPosition.Latitude
	long1 := data1.MatchedPosition.Longitude
	heading1 := data1.MatchedPosition.Heading
	lat2 := data2.MatchedPosition.Latitude
	long2 := data2.MatchedPosition.Longitude
	heading2 := data2.MatchedPosition.Heading

	p1 := haversine.Coord{lat1, long1}
	p2 := haversine.Coord{lat2, long2}
	_, km := haversine.Distance(p1, p2)

	displayMatchedPosition(lat1, long1, heading1)
	displayMatchedPosition(lat2, long2, heading2)
	displayDelta(km * 1000)
}

// Display the latitude, longitude and heading related to the MatchedPosition.
func displayMatchedPosition(latitude float64, longitude float64, heading int) {
	fmt.Println("Matched Position")
	fmt.Printf("Latitude: %f°\n", latitude)
	fmt.Printf("Longitude: %f°\n", longitude)
	fmt.Printf("Heading: %d\n", heading)
	fmt.Println("==========")
}

// Display the latitude and longitude  related to the GeoDisplay.
func displayGeo(latitude float64, longitude float64) {
	fmt.Println("Geo Display")
	fmt.Printf("Latitude: %f°\n", latitude)
	fmt.Printf("Longitude: %f°\n", longitude)
	fmt.Println("==========")
}

// Display the delta in meters and the rounded delta.
// The argument need to be in meter.
func displayDelta(delta float64) {
	fmt.Printf("delta: %fm ≈ %gm\n", delta, math.Round(delta))
}
