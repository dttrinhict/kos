package isito

import (
	"flag"
	istio "istio.io/client-go/pkg/clientset/versioned"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	domain "kos/domain/istio"
	"log"
	"os"
	"path/filepath"
)

var client *Client

type Client struct {
	clientSet *istio.Clientset
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

func Istio() domain.Istio {
	config, err := Config()
	if err != nil {
		log.Panicln("Kubernetes config is error: ", err.Error())
	}
	istioClientSet, err := istio.NewForConfig(config)
	if err != nil {
		log.Panicln("Failed to create istio clientset", err.Error())
	}
	return &Client{
		clientSet: istioClientSet,
	}
}

func (c Client) VirtualService(h *domain.VirtualService) (*domain.VirtualServiceResponse, error) {
	panic("implement me")
}
