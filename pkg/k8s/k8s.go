package k8s

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sdomain "kos/domain/k8s"
)

func (k K8sClient) K8sDeployment(d *k8sdomain.Deployment) (*k8sdomain.DeploymentResponse, error) {
	panic("implement me")
}

func (k K8sClient) K8sConfigMap(cm *k8sdomain.ConfigMap) (*k8sdomain.ConfigMapResponse, error) {
	panic("implement me")
}

func (k K8sClient) K8sSecret(sec *k8sdomain.Secret) (*k8sdomain.SecretResponse, error) {
	panic("implement me")
}

func (k K8sClient) K8sCronjob(c *k8sdomain.CronJob) (*k8sdomain.CronjobRespone, error) {
	panic("implement me")
}

func (k K8sClient) K8sHorizontalPodAutoscaler(h *k8sdomain.HorizontalPodAutoscaler) (*k8sdomain.HorizontalPodAutoscalerResponse, error) {
	panic("implement me")
}

func (k K8sClient) K8sService(h *k8sdomain.Service) (*k8sdomain.ServiceResponse, error) {
	panic("implement me")
}

func (k K8sClient) K8sVirtualService(h *k8sdomain.VirtualService) (*k8sdomain.VirtualServiceResponse, error) {
	panic("implement me")
}

func (k K8sClient) K8sJob(j *k8sdomain.Job) (*k8sdomain.JobResponse, error) {
	panic("implement me")
}

func (k K8sClient) K8sPodsList(ns string) (*[]k8sdomain.PodResponse, error) {
	podClient := k.clientset.CoreV1().Pods(ns)
	var responses []k8sdomain.PodResponse
	var pod k8sdomain.PodResponse
	list, err := podClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}else{
		for _, p:= range list.Items {
			pod.Pod = p
			responses = append(responses, pod)
		}
	}
	return &responses, err
}

func (k K8sClient) K8sClusterRole(cr *k8sdomain.ClusterRole) (*k8sdomain.ClusterRoleResponse, error) {
	panic("implement me")
}

func (k K8sClient) K8sClusterRoleBinding(crb *k8sdomain.ClusterRoleBinding) (*k8sdomain.ClusterRoleBindingRespone, error) {
	panic("implement me")
}

func (k K8sClient) K8sRole(r *k8sdomain.Role) (*k8sdomain.RoleResponse, error) {
	panic("implement me")
}

func (k K8sClient) K8sRoleBinding(rb *k8sdomain.RoleBinding) (*k8sdomain.RoleBindingResponse, error) {
	panic("implement me")
}

func (k K8sClient) K8sServiceAccount(sa *k8sdomain.ServiceAccount) (*k8sdomain.ServiceAccountResponse, error) {
	panic("implement me")
}

func (k K8sClient) K8sCanaryFlagger(cf *k8sdomain.CanaryFlagger) (*k8sdomain.CanaryFlaggerResponse, error) {
	panic("implement me")
}
