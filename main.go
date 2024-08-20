package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/kardianos/service"
)

func main() {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}
	for _, ctr := range containers {
		b, _ := json.Marshal(ctr)
		fmt.Println(string(b))
	}

	// Define the service configuration
	svcConfig := &service.Config{Name: "dnsmasq"}
	// Create a service object
	s, err := service.New(nil, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	// Get the service status
	status, err := s.Status()
	if err != nil {
		log.Fatal(err)
	}
	// Print the status
	fmt.Println("Service status:", status)
	if err := s.Restart(); err != nil {
		log.Fatal(err)
	}
	// Get the service status
	status, err = s.Status()
	if err != nil {
		log.Fatal(err)
	}
	// Print the status
	fmt.Println("Service status2:", status)
	fmt.Println(s.String())
}
