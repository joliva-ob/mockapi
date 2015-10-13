package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "github.com/gorilla/mux"
    "io/ioutil"
)

/**
* Main method
*/
func main() {

    // Checking parameters
    if( !validateParams() ){
        fmt.Printf("Invalid parameters\n -f Path to the apiblueprint file\n -p Server port\n")
        fmt.Printf("Check them and try it again.\n\nExample: go run mockapi.go -p ./myapi.apib -p 8000\n")

        os.Exit(3)
    }

    // Loading the api blueprint file
    file := loadBlueprintFile( os.Args[2] )

    // Starting the server listener to defined endpoints
    logStartServer( os.Args[4] )

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/todos/{todoId}", TodoShow)
    
    var port string = ":"
    port += os.Args[4]    
    log.Fatal(http.ListenAndServe(port, router))
}


func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}

func logStartServer(port string){
    fmt.Printf("Server started listening port %s ...\n",port)       
}

/**
* Method to check the input parameters
*/
func validateParams() bool {

    var isValid bool = true
    size := len(os.Args)
    
    if ( size != 5 ){
        isValid = false
    } else if ( os.Args[1] != "-f" || os.Args[3] != "-p" ){
        isValid = false
    }    
    
    return isValid
}

/**
* Read the api blueprint file
*/
func loadBlueprintFile(file string) []byte{
    // It read the whole file !
    file, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    return file
}
