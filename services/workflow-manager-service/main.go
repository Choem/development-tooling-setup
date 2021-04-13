package main

import (
	"fmt"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	fmt.Println("test")

	config, configErr := rest.InClusterConfig()
	if configErr != nil {
		log.Fatalln(configErr)
	}

	clientset, clientsetErr := kubernetes.NewForConfig(config)
	if clientsetErr != nil {
		log.Fatalln(clientsetErr)
	}
	fmt.Println(clientset)
	// clientset.CoreV1().Pods().List(context.TODO(), metav1.listOptions{})
}
