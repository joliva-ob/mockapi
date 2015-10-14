
package blueprint

import (
	  "fmt"
    "os"
    "bufio"
    "strings"
)

const(
  BeginEndpointStr = "## "
)


/**
 * Load and process the api blueprint file to get the endpoints, 
 * requests and responses
 */
func ProcessApiBlueprint(filepath string) map[string]string {

    var endpoint string
    var sresponse string
    isrecordingresponse := false
    bpm := make( map[string]string )

    apibfile, err := os.Open( filepath )
    if err != nil {
        fmt.Printf("error opening file: %v\n",err)
        os.Exit(1)
    }

    reader := bufio.NewReader(apibfile)
    line, e := Readln(reader)
    for e == nil {
        
        if isrecordingresponse {
            
            if strings.HasPrefix(line, "+") || strings.HasPrefix(line, "#"){
                
                isrecordingresponse = false
                bpm[endpoint] = sresponse // store the response to its key endpoint
            } else {

                sresponse += line
            }
        } else {
          
          if strings.HasPrefix(line, BeginEndpointStr) {

              startindex := strings.Index(line, "[")
              endindex := strings.Index(line, "{")
              endpoint = line[startindex+1:endindex] // store the endpoint key temporally
          } else if strings.HasPrefix(line, "+ Response 200") {

              isrecordingresponse = true
          }
        }  
        
        line,e = Readln(reader)
    }

    return bpm
}


/**
 * Load and process the api blueprint file to get the endpoints, 
 * requests and responses
 */
func ProcessApiBlueprintTest(filepath string) map[string]string {

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