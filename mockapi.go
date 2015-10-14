package main

import (
    "log"
    "fmt"
    "net/http"
    "os"
    "github.com/gorilla/mux"
    "github.com/joliva-ob/mockapi/blueprint"
    "time"
)

// Global variable to store mappings between requested endpoints and
// its mocked response. 
var(
    blueprintmap map[string]string
)


/**
 * Main method
 */
func main() {

    // Checking parameters
    if ( !ValidateParams() ){
        fmt.Printf("Invalid parameters\n -f Path to the apiblueprint file\n -p Server port\n")
        fmt.Printf("Check them and try it again.\n\nExample: go run mockapi.go -p ./myapi.apib -p 8000\n")

        os.Exit(3)
    }

    // Load and process the api blueprint file to get the endpoints, requests and responses
    blueprintmap = blueprint.ProcessApiBlueprint( os.Args[2] )

    // Create the router to handle mockup requests with its response properly
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index) // General welcome endpoint

    // Customized endpoint from blueprint file
    for k := range blueprintmap {
        fmt.Printf("Adding endpoint: %s\n", k)
        router.HandleFunc(k, ProcessEndpoint)    
    }
    
    // Starting server on given port number
    LogStartServer( os.Args[4] )
    port := ":"
    port += os.Args[4]    
    log.Fatal( http.ListenAndServe(port, router) ) // Start the server at listening port
}


// Welcome endpoint
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to mockapi, the mockup tool for api blueprint testing purposes.")
}


/**
 * Generic endpoint to expose resources already loaded from the given 
 * api blueprint file.
 */
func ProcessEndpoint(w http.ResponseWriter, r *http.Request) {

    // Get the mock response from the map, if exists
    status := "200"
    _, ok := blueprintmap[r.URL.Path]
    if !ok {
        status = "500" // General error temporally TODO: get the proper code, response from the api blueprint specification
    }
    outputmock := blueprintmap[r.URL.Path]
    fmt.Printf("%s [mockapi] Request received: %s with response status: %s \n", time.Now(), r.URL.Path, status)
    
    // Response mock string
    fmt.Fprintln(w, outputmock)
}


/*
//Register with: router.HandleFunc("/todos/{todoId}", TodoShow)
//
func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}
*/


func LogStartServer(port string){
    fmt.Printf("Server started successfully on port %s ...\n",port)       
}


/**
* Method to check the input parameters
*/
func ValidateParams() bool {

    var isValid bool = true
    size := len(os.Args)
    
    if ( size != 5 ){
        isValid = false
    } else if ( os.Args[1] != "-f" || os.Args[3] != "-p" ){
        isValid = false
    }    
    
    return isValid
}

