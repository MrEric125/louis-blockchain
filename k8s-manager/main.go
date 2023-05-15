package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func main() {

	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = os.ExpandEnv("$HOME/.kube/config")
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, ns := range namespaces.Items {
		fmt.Printf("namepace name: %s status: %s\n", ns.Name, ns.Status.Phase)
	}
	default_ns := "nginx-pod"

	// 列出所有正在运行的 Pod
	pods, err := clientset.CoreV1().Pods(default_ns).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("--------------")

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Printf("Pod name: %s, Status: %s\n", pod.Name, pod.Status.Phase)
	}
	services, err := clientset.CoreV1().Services(default_ns).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(services.Items))
	for _, service := range services.Items {
		fmt.Printf("Pod name: %s, namespace: %s clusterId:%s type: %s ;externalIp:%s ;port: %s \n", service.Name, service.Namespace, service.Spec.ClusterIP, service.Spec.Type, service.Spec.ExternalIPs, service.Spec.Ports)
	}

}
