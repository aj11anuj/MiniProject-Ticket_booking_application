//go 1.10.4

package main
import {
  "fmt"
  "strings"
}
// Operators
func main() {
  var isValidname bool = len(firstName) >= 2 && len(lastName) >= 2
  isValidemail := strings.Contains(email, "@")
  isValidticket := userTickets > 0 && userTickets <= remainingTickets 
  isValidcity := city == "Singapore" || city == "London"
  
  if isValidname && isValidemail && isValidticket && isValidcity {
    
  }

}