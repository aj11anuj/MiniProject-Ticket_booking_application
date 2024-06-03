//go 1.10.4

package main
import "fmt"

//If you declare variables here then they will be accessible to all the functions
//You can do scooping for variables for their accessibility in diff block or functions
func main() {
  fmt.Printf("Hello, Dcoder!")
  var user string = "Anu"
}

func greet(customer string){
  fmt.Printf("Welcome %v", user)
}

func duplicateGreeting(grahak string) (string){
  var gr string = fmt.Printf("%v welcome", user)
  greet()
  return gr
  // func func_name(parameters)(output datatype in case of return statement)
}

