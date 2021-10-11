package k8s

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	domain "kos/domain/k8s"
)

func (k Client) ListPods(namespace string) (listPods []domain.PodResponse, err error) {
	podClient := k.clientset.CoreV1().Pods(namespace)
	var pod domain.PodResponse
	list, err := podClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}else{
		for _, p:= range list.Items {
			pod.Pod = p
			listPods = append(listPods, pod)
		}
	}
	return listPods, err
}

func (k Client) Deployment(d *domain.Deployment) (*domain.DeploymentResponse, error) {
	panic("implement me")
}

func (k Client) ConfigMap(cm *domain.ConfigMap) (*domain.ConfigMapResponse, error) {
	panic("implement me")
}

func (k Client) Secret(sec *domain.Secret) (*domain.SecretResponse, error) {
	panic("implement me")
}

func (k Client) Cronjob(c *domain.CronJob) (*domain.CronjobRespone, error) {
	panic("implement me")
}

func (k Client) HorizontalPodAutoscaler(h *domain.HorizontalPodAutoscaler) (*domain.HorizontalPodAutoscalerResponse, error) {
	panic("implement me")
}

func (k Client) Service(h *domain.Service) (*domain.ServiceResponse, error) {
	panic("implement me")
}

func (k Client) Job(j *domain.Job) (*domain.JobResponse, error) {
	panic("implement me")
}

func (k Client) ClusterRole(cr *domain.ClusterRole) (*domain.ClusterRoleResponse, error) {
	panic("implement me")
}

func (k Client) ClusterRoleBinding(crb *domain.ClusterRoleBinding) (*domain.ClusterRoleBindingRespone, error) {
	panic("implement me")
}

func (k Client) Role(r *domain.Role) (*domain.RoleResponse, error) {
	panic("implement me")
}

func (k Client) RoleBinding(rb *domain.RoleBinding) (*domain.RoleBindingResponse, error) {
	panic("implement me")
}

func (k Client) ServiceAccount(sa *domain.ServiceAccount) (*domain.ServiceAccountResponse, error) {
	panic("implement me")
}