package k8s

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	domain "kos/domain/k8s"
	"log"
	"os"
	"path/filepath"
)

var client *Client

type Client struct {
	clientset *kubernetes.Clientset
}

func Config() (config *rest.Config, err error){
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	if _, err := os.Stat(*kubeconfig); os.IsNotExist(err) {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}

func Kubernetes() domain.K8s {
	config, err := Config()
	if err != nil {
		log.Panicln("Kubernetes config is error: ", err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panicln("Failed to create K8s clientset", err.Error())
	}
	return &Client{
		clientset: clientset,
	}
}

