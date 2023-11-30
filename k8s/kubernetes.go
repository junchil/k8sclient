package k8s

import (
	"fmt"

	"github.com/junchil/k8sclient/k8s/kubeconfig"
	"k8s.io/client-go/kubernetes"
)

type Client struct {
	Clientset *kubernetes.Clientset
}

func New() (*Client, error) {
	kubeConfig, err := kubeconfig.ReadKubeConfig()
	if err != nil {
		fmt.Printf("error getting Kubernetes config: %v\n", err)
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		fmt.Printf("error getting Kubernetes clientset: %v\n", err)
		return nil, err
	}

	return &Client{Clientset: clientset}, nil
}
