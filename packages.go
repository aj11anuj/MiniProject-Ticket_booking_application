//go 1.10.4

package main
import "fmt"
func main() {
  fmt.Printf("Hello, Dcoder!")
  //You can make multiple files with same package name for using them as a single application
  //Make sure that each file consist of different function name
  //For running you have to run all the files with same package at once, otherwise you will get error that xyz var is undefined
  //go run main.go master.go etc..
}
