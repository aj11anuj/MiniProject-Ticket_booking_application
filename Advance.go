//go 1.10.4

package main
import "fmt"
func main() {
  // Maps
  // A datatype like a list with same key-value data type
  // For making a slice of maps use []map
  var userData = make(map[string]string)
  userData["firstName"] = firstName
  userData["lastName"] = lastName
  userData["email"] = email
  
  // Struct
  // A datatype like list but can hold diff data type
  type userData struct{
    firstName string
    lastName string
    dob int
  }
  var userData = userData{
    firstName: firstName
    lastName: lastName
    dob: dateOfBirth
  }
  
  // Go routines - concurrency
  time.Sleep(10 * time.Second)
  // Code return below it will execute after 10sec delay
  // You can put a "go" keyword while Using this logic for performing multiple threading and making our application concurrent
}