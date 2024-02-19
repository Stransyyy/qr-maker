package main

import (
	"context"
	"os"

	godaddy "github.com/oze4/godaddygo"
)

func domain_check(domain string) ([]string, error) {
	var domains []string

	// Create a new Godaddy Client
	api, err := godaddy.NewDevelopment(os.Getenv("GODADDY_API_KEY"), os.Getenv("SECRET"))
	if err != nil {
		panic(err)
	}

	// Check if a domain is available
	available, err := api.V1().CheckAvailability(context.Background(), domain, false)
	if err != nil {
		panic(err)
	}

	if available.Available == true {
		domains = append(domains, available.Domain)
	}

	return domains, err
}
