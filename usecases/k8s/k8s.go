package k8s

import (
	domain "kos/domain/k8s"
)

type KubernetesImpl struct {
	K8sDomain domain.K8s
}

type Kubernetes interface {
	Deployment(d *domain.Deployment) (*domain.DeploymentResponse, error)
	ConfigMap(cm *domain.ConfigMap) (*domain.ConfigMapResponse, error)
	Secret(sec *domain.Secret) (*domain.SecretResponse, error)
	Cronjob(cj *domain.CronJob) (*domain.CronjobRespone, error)
	HorizontalPodAutoscaler(hpa *domain.HorizontalPodAutoscaler) (*domain.HorizontalPodAutoscalerResponse, error)
	Service(s *domain.Service) (*domain.ServiceResponse, error)
	Job(job *domain.Job) (*domain.JobResponse, error)
	ListPods(namespace string) ([]domain.PodResponse, error)
	ClusterRole(cr *domain.ClusterRole) (*domain.ClusterRoleResponse, error)
	ClusterRoleBinding(crb *domain.ClusterRoleBinding) (*domain.ClusterRoleBindingRespone, error)
	Role(role *domain.Role) (*domain.RoleResponse, error)
	RoleBinding(rb *domain.RoleBinding) (*domain.RoleBindingResponse, error)
	ServiceAccount(sa *domain.ServiceAccount) (*domain.ServiceAccountResponse, error)
}

func K8s(d *domain.K8s) Kubernetes {
	return KubernetesImpl{
		K8sDomain: *d,
	}
}

func (k KubernetesImpl) Deployment(d *domain.Deployment) (*domain.DeploymentResponse, error) {
	panic("implement me")
}

func (k KubernetesImpl) ConfigMap(cm *domain.ConfigMap) (*domain.ConfigMapResponse, error) {
	panic("implement me")
}

func (k KubernetesImpl) Secret(sec *domain.Secret) (*domain.SecretResponse, error) {
	panic("implement me")
}

func (k KubernetesImpl) Cronjob(cj *domain.CronJob) (*domain.CronjobRespone, error) {
	panic("implement me")
}

func (k KubernetesImpl) HorizontalPodAutoscaler(hpa *domain.HorizontalPodAutoscaler) (*domain.HorizontalPodAutoscalerResponse, error) {
	panic("implement me")
}

func (k KubernetesImpl) Service(s *domain.Service) (*domain.ServiceResponse, error) {
	panic("implement me")
}

func (k KubernetesImpl) Job(job *domain.Job) (*domain.JobResponse, error) {
	panic("implement me")
}

func (k KubernetesImpl) ListPods(namespace string) ([]domain.PodResponse, error) {
	return k.K8sDomain.ListPods("default")
}

func (k KubernetesImpl) ClusterRole(cr *domain.ClusterRole) (*domain.ClusterRoleResponse, error) {
	panic("implement me")
}

func (k KubernetesImpl) ClusterRoleBinding(crb *domain.ClusterRoleBinding) (*domain.ClusterRoleBindingRespone, error) {
	panic("implement me")
}

func (k KubernetesImpl) Role(role *domain.Role) (*domain.RoleResponse, error) {
	panic("implement me")
}

func (k KubernetesImpl) RoleBinding(rb *domain.RoleBinding) (*domain.RoleBindingResponse, error) {
	panic("implement me")
}

func (k KubernetesImpl) ServiceAccount(sa *domain.ServiceAccount) (*domain.ServiceAccountResponse, error) {
	panic("implement me")
}