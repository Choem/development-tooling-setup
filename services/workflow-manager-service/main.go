package main

import (
	"context"
	"fmt"
	"log"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, configErr := rest.InClusterConfig()
	if configErr != nil {
		log.Fatalln(configErr)
	}

	clientset, clientsetErr := kubernetes.NewForConfig(config)
	if clientsetErr != nil {
		log.Fatalln(clientsetErr)
	}

	for {
		pods, podsErr := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
		if podsErr != nil {
			panic(podsErr.Error())
		}

		fmt.Printf("There are %d pods in this cluster\n", len(pods.Items))

		time.Sleep(10 * time.Second)
	}
}
