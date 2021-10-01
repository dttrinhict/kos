package k8s

import (
	"flag"
	flagger "github.com/fluxcd/flagger/pkg/client/clientset/versioned"
	istio "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	k8sdomain "kos/domain/k8s"
	"log"
	"os"
	"path/filepath"
)

var k8sClient *K8sClient

type K8sClient struct {
	clientset *kubernetes.Clientset
	istioClentSet *istio.Clientset
	flaggerClient *flagger.Clientset
}

func KubernetesConfig() (config *rest.Config, err error){
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

func NewK8s() k8sdomain.K8sDomain {
	config, err := KubernetesConfig()
	if err != nil {
		log.Panicln("Kubernetes config is error: ", err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panicln("Failed to create K8s clientset", err.Error())
	}
	istioClientSet, err := istio.NewForConfig(config)
	if err != nil {
		log.Panicln("Failed to create istio clientset", err.Error())
	}
	flaggerClientSet, err := flagger.NewForConfig(config)
	if err != nil {
		log.Panicln("Failed to create flagger clientset", err.Error())
	}
	return &K8sClient{
		clientset: clientset,
		istioClentSet: istioClientSet,
		flaggerClient: flaggerClientSet,
	}
}

