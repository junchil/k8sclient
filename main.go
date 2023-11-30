package main

import (
	"context"

	"github.com/junchil/k8sclient/k8s/pod"
)

func main() {
	k8spod, err := pod.New()
	if err != nil {
		panic(err)
	}

	err = k8spod.List(context.TODO(), "karpenter")
	if err != nil {
		panic(err)
	}
}
