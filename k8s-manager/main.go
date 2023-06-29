package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	kubeConfig := os.Getenv("KUBECONFIG")
	if kubeConfig == "" {
		kubeConfig = os.ExpandEnv("$HOME/.kube/config")
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("current context config:%s ;apiPath:%s", config.Host, config.APIPath)
	fmt.Println()
	// 获取 客户端请求
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// 获取当前上下文信息
	contextConfig, err := clientcmd.LoadFromFile(kubeConfig)
	if err != nil {
		panic(err.Error())
	}
	var default_ns string

	for name, context := range contextConfig.Contexts {
		if name == contextConfig.CurrentContext {
			default_ns = context.Namespace

		}
		fmt.Printf("- %s\n", name)
		fmt.Printf("  Namespace: %s\n", context.Namespace)
		fmt.Printf("  Auth Info: %s\n", context.AuthInfo)
		fmt.Printf("  Cluster: %s\n", context.Cluster)
	}

	fmt.Printf("current context:%s", contextConfig.CurrentContext)

	namespaces, err := clientSet.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	// 获取当前上下文的Namespace

	if err != nil {
		panic(err.Error())
	}
	for _, ns := range namespaces.Items {
		fmt.Printf("namepace name: %s status: %s\n", ns.Name, ns.Status.Phase)
	}
	// 列出所有正在运行的 Pod
	pods, err := clientSet.CoreV1().Pods(default_ns).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("--------------")

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Printf("Pod name: %s, Status: %s\n", pod.Name, pod.Status.Phase)
	}
	services, err := clientSet.CoreV1().Services(default_ns).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(services.Items))
	for _, service := range services.Items {
		fmt.Printf("Pod name: %s, namespace: %s clusterId:%s type: %s ;externalIp:%s ;port: %s \n", service.Name, service.Namespace, service.Spec.ClusterIP, service.Spec.Type, service.Spec.ExternalIPs, service.Spec.Ports)
	}

	deploymentName := "nginx-deployment"

	deployment, err := clientSet.AppsV1().Deployments(default_ns).Get(context.Background(), deploymentName, metav1.GetOptions{})

	if err != nil {
		panic(err.Error())
	}

	deploymentJSON, _ := json.Marshal(deployment)
	fmt.Print("-------")
	fmt.Println(string(deploymentJSON))

}

type Pool struct {
	work chan func()
	sem  chan struct{}
}

func New(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}
func (p *Pool) Schedule(task func()) {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)
	}
}
func (p *Pool) worker(task func()) {
	defer func() { <-p.sem }()
	for {
		task()
		task = <-p.work
	}
}
