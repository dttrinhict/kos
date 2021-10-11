package flagger

import (
	"flag"
	flagger "github.com/fluxcd/flagger/pkg/client/clientset/versioned"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	domain "kos/domain/flagger"
	"log"
	"os"
	"path/filepath"
)

var client *Client

type Client struct {
	clientset *flagger.Clientset
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

func Flagger() domain.Flagger {
	config, err := Config()
	if err != nil {
		log.Panicln("Kubernetes config is error: ", err.Error())
	}
	flaggerClientSet, err := flagger.NewForConfig(config)
	if err != nil {
		log.Panicln("Failed to create flagger clientset", err.Error())
	}
	return &Client{
		clientset: flaggerClientSet,
	}
}

func (c Client) CanaryFlagger(cf *domain.CanaryFlagger) (*domain.CanaryFlaggerResponse, error) {
	panic("implement me")
}