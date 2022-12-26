package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	kubeconfig := path.Join(home, ".kube/config")
	if envvar := os.Getenv("KUBECONFIG"); len(envvar) > 0 {
		kubeconfig = envvar
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("the kubeconfig cannot be loaded: %v\n", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	job, err := clientset.BatchV1().Jobs("default").Get(context.TODO(), "hello", metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(job.Name)
}
