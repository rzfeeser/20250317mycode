/* RZFeeser | Alta3 Research
   nil in the error position means "no error" */
   
package main

import (
    "fmt"
    "errors"
)

//name of funct   //input                      //output
func rollchar(firstName string, lastName int) (string, error) {
    if lastName == 55 {
        return "", errors.New("Needs to be a string not 55")
    } else {
    return firstName + " the ", nil
  }
}

func main() {
    fmt.Println("Welcome to the Character Generator")

    playerChar, err := rollchar("Gandalf", 55)

    if err != nil {
        fmt.Println("Error while spawning your requested character.")
        fmt.Println(err) // belongs in debug log
    } else {
        fmt.Println(playerChar, "has been generated.")
    }
}

