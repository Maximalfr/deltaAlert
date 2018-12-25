package jsonparser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Struct for parsing (extract) directly the json from the file
type GpsAlert struct {
	Type    int `json:"type"`
	EventID struct {
		Value string `json:"value"`
	} `json:"event_id"`
	CreateTime struct {
		Seconds int `json:"seconds"`
		Nanos   int `json:"nanos"`
	} `json:"create_time"`
	UpdateTime struct {
		Seconds int `json:"seconds"`
		Nanos   int `json:"nanos"`
	} `json:"update_time"`
	TTL        int `json:"ttl"`
	GeoDisplay struct {
		Geometry struct {
			Type        int       `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Properties struct {
			Length struct {
				Kind struct {
					NumberValue int `json:"NumberValue"`
				} `json:"Kind"`
			} `json:"length"`
		} `json:"properties"`
	} `json:"geo_display"`
	Confirmations int `json:"confirmations"`
	Status        int `json:"status"`
	Properties    struct {
		EventCoords struct {
			Kind struct {
				ListValue struct {
					Values []struct {
						Kind struct {
							NumberValue float64 `json:"NumberValue"`
						} `json:"Kind"`
					} `json:"values"`
				} `json:"ListValue"`
			} `json:"Kind"`
		} `json:"event_coords"`
		WayIds struct {
			Kind struct {
				ListValue struct {
					Values []struct {
						Kind struct {
							NumberValue int `json:"NumberValue"`
						} `json:"Kind"`
					} `json:"values"`
				} `json:"ListValue"`
			} `json:"Kind"`
		} `json:"way_ids"`
		WayLength struct {
			Kind struct {
				NumberValue int `json:"NumberValue"`
			} `json:"Kind"`
		} `json:"way_length"`
	} `json:"properties"`
	Revision      int `json:"revision"`
	Score         int `json:"score"`
	MergeCategory int `json:"merge_category"`
	GeoMerge      struct {
		Geometry struct {
			Type        int       `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Properties struct {
			Length struct {
				Kind struct {
					NumberValue int `json:"NumberValue"`
				} `json:"Kind"`
			} `json:"length"`
		} `json:"properties"`
	} `json:"geo_merge"`
	Country         string `json:"country"`
	MatchedPosition struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
		Heading   int     `json:"heading"`
		Ref       string  `json:"ref"`
		NodeLeft  int     `json:"node_left"`
		NodeRight int     `json:"node_right"`
		Way       int     `json:"way"`
	} `json:"matched_position"`
	WayIds []int `json:"way_ids"`
	UeType int   `json:"ue_type"`
}

// Read and transform the json file to the GpsAlert Struct
func JsonRead(filename string) (GpsAlert, error) {
	var j GpsAlert

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return j, err
	}
	json.Unmarshal(data, &j)

	return j, nil
}
