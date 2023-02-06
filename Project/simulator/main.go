package main
import (
	"fmt"
	route2 "github.com/fabioramosalves/ImersaoFullStack-simulator/application/route"
	"github.com/joho/godotenv"
)
func init(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("error loading .env file")
	}
}

func main(){
	route := route2.Route{
		ID: 	  "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringJson, _  := route.ExportJsonPositions()

	fmt.Println(stringJson[0])
}