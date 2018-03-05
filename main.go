package main

// go build -ldflags "-X main.version=0.0.1"

import(

  "fmt"
  // "log"
   "os"

  u "github.com/ardeshir/version"
)


var ( 
 debug bool = false
 version string = "0.0.0"
 )
 

func main() {

  fmt.Println("We're OK!")


  if debugTrue() {
    u.V(version)
  }

}

// Function to check env variable DEFAULT_DEBUG bool
func debugTrue() bool {
    
     if os.Getenv("DEFAULT_DEBUG") != "" {
        return true
     }  
     return false 
}
