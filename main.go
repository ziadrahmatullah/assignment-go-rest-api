// Golang program to illustrate the usage of 
// time.Date() function 

// Including main package 
package main 

// Importing fmt and time 
import "fmt"
import "time"

// Calling main 
func main() { 

	// Calling Date() method 
	// with all its parameters 
	tm := time.Date(2020, time.April, 
		11, 21, 34, 01, 0, time.UTC) 

	// Using Local() for location and printing 
	// the stated time and date in UTC 
	fmt.Printf("%s", tm.Local()) 
} 
