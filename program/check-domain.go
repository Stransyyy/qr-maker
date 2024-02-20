package main

import (
	"context"
	"log"
	"os"

	godaddy "github.com/oze4/godaddygo"
)

// DomainCheck checks if a domain is available. It uses the godaddyugo api using the Available method. Meaning it checks if the domain is available. but we use that to our advantage to check if the domain is not available.
func DomainCheck(domain string) (string, error) {

	// Create a new Godaddy Client
	api, err := godaddy.NewDevelopment(os.Getenv("GODADDY_API_KEY"), os.Getenv("SECRET"))
	if err != nil {
		panic(err)
	}

	// Check if a domain is available
	available, err := api.V1().CheckAvailability(context.Background(), domain, true)
	if err != nil {
		log.Printf("Cannot check availability: %s \n", err)
	}

	url := available.Domain

	if available.Available == false {
		return url, nil
	}

	return url, nil
}
