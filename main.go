package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	//	"time"

	//	"k8s.io/apimachinery/pkg/api/errors"
//	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
//	api "k8s.io/client-go/kubernetes/typed/batch/v1"
	batchv1client "k8s.io/client-go/kubernetes/typed/batch/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error)
	}
	fmt.Printf("there are %d pods\n", len(pods.Items))

	batchClient := batchv1client.NewForConfigOrDie(config)
	jobsClient := batchClient.Jobs("default")
	jobList, err := jobsClient.List(metav1.ListOptions{})
	fmt.Printf("there are %d jobs\n", len(jobList.Items))
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
