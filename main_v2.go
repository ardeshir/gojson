package main

// go build -ldflags "-X main.version=0.0.1"
// Data link
// https://gbfs.citibikenyc.com/gbfs/en/station_status.json

import(
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"

  u "github.com/ardeshir/version"
)

// citiBikeUrl provides the station statues of CitiBike Bike sharing stations
const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

// stationData is used to unmarshal the JSON document returned from citiBikeURL
type stationData struct {
    LastUpdated int `json:"last_updated"`
    TTL         int `json:"ttl"`
    Data        struct {
                Stations []station `json:"stations"`
    } `json:"data"`
}

// station is used to unmarshall each of the station documents in stationData
type station struct {
    ID                  string `json:"station_id"`
    NumBikesAvailable      int `json:"num_bikes_available"`
    NUmBikesDisabled       int `json:"num_biles_disabled"`
    NumDocksAvailable      int `json:"num_docks_available"`
    NumDocksDisabled       int `json:"num_docks_disabled"`
    IsInstalled            int `json:"is_installed"`
    IsRenting              int `json:"is_renting"`
    IsReturning            int `json:"is_returning"`
    LastReporting          int `json:"last_reported"`
    HasAvailableKeys      bool `json:"eightd_has_available_keys"`
}

var ( 
 debug bool = false
 version string = "0.0.1"
 )
 

func main() {
  fmt.Println("We're OK!")
 //+++++++++++  main   

  // Get the JSON response from the URL
  response, err := http.Get(citiBikeURL)
  u.ErrNil(err, "Unable to http.Get(citiBikeURL)")
  
  // Defer closing of response
  defer response.Body.Close()
  
  // Read body of the response in to []byte
  body, err := ioutil.ReadAll(response.Body)
  u.ErrNil(err, "Unable to read response.body")

  // Declare a variable of type stationData
  var sd stationData 
  
  // Unmarshall the JSON data into the variable
  if err := json.Unmarshal(body, &sd); err != nil {
        log.Fatal(err)
        return
   }
  
  // Marshal the data 
  outputData, err := json.Marshal(sd)
  u.ErrNil(err, "Unable to marshal(sd)")

  // Save the marshalled data to file
  if err := ioutil.WriteFile("citibike.json", outputData, 0644); err != nil {
      log.Fatal(err)
  }
  
  // Print the first station. 
  fmt.Printf("%+v\n\n", sd.Data.Stations[0])
  
  

 //++++++++++  footer 
  if debugTrue() {
    u.V(version)
  }

}

// Function to check env variable DEFAULT_DEBUG bool is set
func debugTrue() bool {
    
     if os.Getenv("DEFAULT_DEBUG") != "" {
        return true
     }  
     return false 
}
