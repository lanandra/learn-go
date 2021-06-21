package main

import "fmt"

func main () {

        // Ask user input
        fmt.Println("Please input company name:")

        var companyName string

        fmt.Scanln(&companyName)

        // Check whether user input is match with company name or not
        if companyName == "company" {
                fmt.Println("Our company name")
        } else {
                fmt.Println("Not our company name")
        }
}
