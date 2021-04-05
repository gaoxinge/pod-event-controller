package main

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/gaoxinge/pod-event-controller/pkg"
)

func main() {
	config := rest.Config{
		Host: "http://127.0.0.1:8001",
	}

	clientSet, err := kubernetes.NewForConfig(&config)
	if err != nil {
		log.Printf("new client set with error %v\n", err)
		return
	}

	controller := pkg.NewPodEventController(clientSet)

	stop := make(chan struct{})
	defer close(stop)

	err = controller.Run(stop)
	if err != nil {
		log.Printf("run with error %v\n", err)
	}

	select {}
}
