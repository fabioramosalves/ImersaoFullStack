package route
import (	
	"errors"
	"bufio"
	"os"
	"strings"
	"strconv"
	"encoding/json"
)

type Route struct{
	ID string `json:"routeId"`
	ClientID string `json:"clientId"`
	Positions []Position  `json:"position"`
}

type Position struct{
	Latitude float64 `json:"latitude"`
	Longitude float64  `json:"longitude"`
}

type PartialRoutePosition struct{
	ID string `json:"routeId"`
	ClientID string `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool `json:"finished"`
}

func(r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("route id not informed")
	}

	f, err := os.Open("destinations/" + r.ID + ".txt")

	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		data := strings.Split(scanner.Text(), ",")
		latitude, err := strconv.ParseFloat(data[0],64)
		if err != nil {
			return nil
		}	
		longitude, err := strconv.ParseFloat(data[1],64)
		if err != nil {
			return nil
		}
		r.Positions = append(r.Positions, Position{
			Latitude: latitude,
			Longitude: longitude,
		})
	}
	return nil
}

func (r *Route) ExportJsonPositions() ([]string, error){
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)

	for k, v := range r.Positions{
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Latitude, v.Longitude}
		route.Finished = false
		if total -1 == k {
			route.Finished = true
		}
		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))
	}
	return result, nil	
}