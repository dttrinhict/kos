package k8s

import (
	k8sdomain "kos/domain/k8s"
)

type K8sUseCaseImpl struct {
	K8sDomain k8sdomain.K8sDomain
}

type K8sUseCase interface {
	K8sDeployment(d *k8sdomain.Deployment) (*k8sdomain.DeploymentResponse, error)
	K8sConfigMap(cm *k8sdomain.ConfigMap) (*k8sdomain.ConfigMapResponse, error)
	K8sSecret(sec *k8sdomain.Secret) (*k8sdomain.SecretResponse, error)
	K8sCronjob(cj *k8sdomain.CronJob) (*k8sdomain.CronjobRespone, error)
	K8sHorizontalPodAutoscaler(hpa *k8sdomain.HorizontalPodAutoscaler) (*k8sdomain.HorizontalPodAutoscalerResponse, error)
	K8sService(s *k8sdomain.Service) (*k8sdomain.ServiceResponse, error)
	K8sVirtualService(vs *k8sdomain.VirtualService) (*k8sdomain.VirtualServiceResponse, error)
	K8sJob(job *k8sdomain.Job) (*k8sdomain.JobResponse, error)
	K8sPodsList(ns string) (*[]k8sdomain.PodResponse, error)
	K8sClusterRole(cr *k8sdomain.ClusterRole) (*k8sdomain.ClusterRoleResponse, error)
	K8sClusterRoleBinding(crb *k8sdomain.ClusterRoleBinding) (*k8sdomain.ClusterRoleBindingRespone, error)
	K8sRole(role *k8sdomain.Role) (*k8sdomain.RoleResponse, error)
	K8sRoleBinding(rb *k8sdomain.RoleBinding) (*k8sdomain.RoleBindingResponse, error)
	K8sServiceAccount(sa *k8sdomain.ServiceAccount) (*k8sdomain.ServiceAccountResponse, error)
	K8sCanaryFlagger(cf *k8sdomain.CanaryFlagger) (*k8sdomain.CanaryFlaggerResponse, error)
}

func NewK8sUseCase(d *k8sdomain.K8sDomain) K8sUseCase {
	return K8sUseCaseImpl{
		K8sDomain: *d,
	}
}

func (k K8sUseCaseImpl) K8sDeployment(d *k8sdomain.Deployment) (*k8sdomain.DeploymentResponse, error) {
	return k.K8sDomain.K8sDeployment(d)
}

func (k K8sUseCaseImpl) K8sConfigMap(cm *k8sdomain.ConfigMap) (*k8sdomain.ConfigMapResponse, error) {
	return k.K8sDomain.K8sConfigMap(cm)
}

func (k K8sUseCaseImpl) K8sSecret(sec *k8sdomain.Secret) (*k8sdomain.SecretResponse, error) {
	return k.K8sDomain.K8sSecret(sec)
}

func (k K8sUseCaseImpl) K8sCronjob(cj *k8sdomain.CronJob) (*k8sdomain.CronjobRespone, error) {
	return k.K8sDomain.K8sCronjob(cj)
}

func (k K8sUseCaseImpl) K8sHorizontalPodAutoscaler(hpa *k8sdomain.HorizontalPodAutoscaler) (*k8sdomain.HorizontalPodAutoscalerResponse, error) {
	return k.K8sDomain.K8sHorizontalPodAutoscaler(hpa)
}

func (k K8sUseCaseImpl) K8sService(s *k8sdomain.Service) (*k8sdomain.ServiceResponse, error) {
	return k.K8sDomain.K8sService(s)
}

func (k K8sUseCaseImpl) K8sVirtualService(vs *k8sdomain.VirtualService) (*k8sdomain.VirtualServiceResponse, error) {
	return k.K8sDomain.K8sVirtualService(vs)
}

func (k K8sUseCaseImpl) K8sJob(job *k8sdomain.Job) (*k8sdomain.JobResponse, error) {
	return k.K8sDomain.K8sJob(job)
}

func (k K8sUseCaseImpl) K8sPodsList(ns string) (*[]k8sdomain.PodResponse, error) {
	return k.K8sDomain.K8sPodsList(ns)
}

func (k K8sUseCaseImpl) K8sClusterRole(cr *k8sdomain.ClusterRole) (*k8sdomain.ClusterRoleResponse, error) {
	return k.K8sDomain.K8sClusterRole(cr)
}

func (k K8sUseCaseImpl) K8sClusterRoleBinding(crb *k8sdomain.ClusterRoleBinding) (*k8sdomain.ClusterRoleBindingRespone, error) {
	return k.K8sDomain.K8sClusterRoleBinding(crb)
}

func (k K8sUseCaseImpl) K8sRole(role *k8sdomain.Role) (*k8sdomain.RoleResponse, error) {
	return k.K8sDomain.K8sRole(role)
}

func (k K8sUseCaseImpl) K8sRoleBinding(rb *k8sdomain.RoleBinding) (*k8sdomain.RoleBindingResponse, error) {
	return k.K8sDomain.K8sRoleBinding(rb)
}

func (k K8sUseCaseImpl) K8sServiceAccount(sa *k8sdomain.ServiceAccount) (*k8sdomain.ServiceAccountResponse, error) {
	return k.K8sDomain.K8sServiceAccount(sa)
}

func (k K8sUseCaseImpl) K8sCanaryFlagger(cf *k8sdomain.CanaryFlagger) (*k8sdomain.CanaryFlaggerResponse, error) {
	return k.K8sDomain.K8sCanaryFlagger(cf)
}