package main

import (
	"context"
	"fmt"
	"os"

	"github.com/junchil/k8sclient/modules/kubeconfig"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {
	kubeConfig, err := kubeconfig.ReadKubeConfig()
	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		fmt.Printf("error getting Kubernetes clientset: %v\n", err)
		os.Exit(1)
	}

	pods, err := clientset.CoreV1().Pods("kube-system").List(context.Background(), v1.ListOptions{})
	if err != nil {
		fmt.Printf("error getting pods: %v\n", err)
		os.Exit(1)
	}
	for _, pod := range pods.Items {
		fmt.Printf("Pod name: %s\n", pod.Name)
	}
}
