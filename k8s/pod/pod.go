package pod

import (
	"context"
	"fmt"

	"github.com/junchil/k8sclient/k8s"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Pod struct {
	K8s *k8s.Client
}

type PodService interface {
	List(ctx context.Context, namespace string) error
}

var _ PodService = (*Pod)(nil)

func New() (*Pod, error) {
	client, err := k8s.New()
	if err != nil {
		return nil, err
	}

	return &Pod{K8s: client}, nil
}

func (p *Pod) List(ctx context.Context, namespace string) error {
	pods, err := p.K8s.Clientset.CoreV1().Pods(namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		fmt.Printf("error getting pods: %v\n", err)
		return err
	}

	for _, pod := range pods.Items {
		fmt.Printf("Pod name: %s\n", pod.Name)
	}
	return nil
}
