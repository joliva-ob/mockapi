
package blueprint

import (
	  "fmt"
    "os"
    "bufio"
    "strings"
)

/**
 * Load and process the api blueprint file to get the endpoints, 
 * requests and responses
 */
func ProcessApiBlueprint(filepath string) []string {

    endpointlist := make( []string, 0 )

    apibfile, err := os.Open( filepath )
    if err != nil {
        fmt.Printf("error opening file: %v\n",err)
        os.Exit(1)
    }

    reader := bufio.NewReader(apibfile)
    line, e := Readln(reader)
    for e == nil {
        
        if strings.HasPrefix(line, "## ") {

            startindex := strings.Index(line, "[")
            endindex := strings.Index(line, "{")
            endpointlist = append( endpointlist,line[startindex+1:endindex] )
        }
        line,e = Readln(reader)
    }

    return endpointlist
}


/**
 * Load and process the api blueprint file to get the endpoints, 
 * requests and responses
 * TODO load from the api blueprint file
 */
func LoadEndpointsMap(filepath string) map[string]string {

    bpm := make( map[string]string )

    apibfile, err := os.Open( "blueprint/test_buyers.json" )
    if err != nil {
        fmt.Printf("error opening file: %v\n",err)
        os.Exit(1)
    }
    var outputmock string
    reader := bufio.NewReader(apibfile)
    line, e := Readln(reader)
    for e == nil {
        outputmock += line
        line,e = Readln(reader)
    }
    bpm["/crm-api/v1.0/buyers"] = outputmock

    apibfile, err = os.Open( "blueprint/test_purchases.json" )
    if err != nil {
        fmt.Printf("error opening file: %v\n",err)
        os.Exit(1)
    }
    reader = bufio.NewReader(apibfile)
    line, e = Readln(reader)
    for e == nil {
        outputmock += line
        line,e = Readln(reader)
    }
    bpm["/crm-api/v1.0/purchases"] = outputmock

    return bpm
}


// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
  var (isPrefix bool = true
       err error = nil
       line, ln []byte
      )
  for isPrefix && err == nil {
      line, isPrefix, err = r.ReadLine()
      ln = append(ln, line...)
  }
  return string(ln),err
}	