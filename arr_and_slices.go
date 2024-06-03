//go 1.10.4

package main
import "fmt"
func main() {
  
  //Arrays
  var bookings[50] string
  var firstName string
  var lastName string
  
  fmt.Scan(&firstName)
  fmt.Scan(&lastName)
  bookings[0] = firstName + " " + lastName
  fmt.Printf("Whole arr: %v", bookings)
  fmt.Printf("Element at 0th index: %v", bookings[0])


  //Slices
  //When we don't know the size of arr or lets Say 1 user hooked all 60 tickets in that case an arr of 60 limit having only 1 user for all 60 bookings is not efficient thats why we use slices.
  //Think of it as modified arr with a flexible capacity
  var booking[60] string{}
  booking = append(booking, firstName + " " + lastName)
  //You cam get the value of slices in same way as you were getting from arr
  
  
  //Alt synatx for declaring a variable
  var example string = arr
  example := arr

}