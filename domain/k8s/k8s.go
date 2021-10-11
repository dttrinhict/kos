package k8s

import (
	appsv1 "k8s.io/api/apps/v1"
	hpav1 "k8s.io/api/autoscaling/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Deployment struct {
	TypeMeta   map[string]string       `json:",inline"`
	ObjectMeta metav1.ObjectMeta       `json:"metadata"`
	Spec       appsv1.DeploymentSpec   `json:"spec,omitempty"`
	Status     appsv1.DeploymentStatus `json:"status,omitempty"`
}

type CronJob struct {
	TypeMeta   map[string]string          `json:",inline"`
	ObjectMeta metav1.ObjectMeta          `json:"metadata"`
	Spec       batchv1beta1.CronJobSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status     batchv1beta1.CronJobStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type Service struct {
	TypeMeta   map[string]string    `json:",inline"`
	ObjectMeta metav1.ObjectMeta    `json:"metadata"`
	Spec       corev1.ServiceSpec   `json:"spec,omitempty"`
	Status     corev1.ServiceStatus `json:"status,omitempty"`
}

type ServiceResponse struct {
	Service *corev1.Service `json:"service"`
}

type DeploymentResponse struct {
	Deployment *appsv1.Deployment `json:"deployment"`
}

type CronjobRespone struct {
	Cronjob *batchv1beta1.CronJob `json:"cronjob"`
}

type HorizontalPodAutoscaler struct {
	TypeMeta   map[string]string                   `json:",inline"`
	ObjectMeta metav1.ObjectMeta                   `json:"metadata,omitempty"`
	Spec       hpav1.HorizontalPodAutoscalerSpec   `json:"spec,omitempty"`
	Status     hpav1.HorizontalPodAutoscalerStatus `json:"status,omitempty"`
}

type HorizontalPodAutoscalerResponse struct {
	HorizontalPodAutoscaler *hpav1.HorizontalPodAutoscaler `json:"horizontalPodAutoscaler,omitempty"`
}

type ConfigMap struct {
	TypeMeta   map[string]string `json:",inline"`
	ObjectMeta metav1.ObjectMeta `json:"metadata"`
	Immutable  *bool             `json:"immutable,omitempty"`
	Data       map[string]string `json:"data,omitempty"`
	BinaryData map[string][]byte `json:"binaryData,omitempty"`
}

type ConfigMapResponse struct {
	ConfigMap *corev1.ConfigMap
}

type Secret struct {
	TypeMeta   map[string]string `json:",inline"`
	ObjectMeta metav1.ObjectMeta `json:"metadata"`
	Immutable  *bool             `json:"immutable,omitempty"`
	Data       map[string][]byte `json:"data,omitempty"`
	StringData map[string]string `json:"stringData,omitempty"`
	Type       corev1.SecretType `json:"type,omitempty"`
}

type SecretResponse struct {
	Secret *corev1.Secret
}

type Job struct {
	TypeMeta   map[string]string `json:",inline"`
	ObjectMeta metav1.ObjectMeta `json:"metadata"`
	Spec       batchv1.JobSpec   `json:"spec,omitempty"`
	Status     batchv1.JobStatus `json:"status,omitempty"`
}

type JobResponse struct {
	Job *batchv1.Job `json:"job"`
}

type Pod struct {
	TypeMeta   map[string]string `json:",inline"`
	ObjectMeta metav1.ObjectMeta `json:"metadata"`
	Spec       corev1.PodSpec    `json:"spec,omitempty"`
	Status     corev1.PodStatus  `json:"status,omitempty"`
}

type PodResponse struct {
	Pod corev1.Pod `json:"pod,omitempty"`
}

type ClusterRole struct {
	TypeMeta        map[string]string       `json:",inline"`
	ObjectMeta      metav1.ObjectMeta       `json:"metadata"`
	Rules           []rbacv1.PolicyRule     `json:"rules"`
	AggregationRule *rbacv1.AggregationRule `json:"aggregationRule,omitempty"`
}

type ClusterRoleResponse struct {
	ClusterRole *rbacv1.ClusterRole `json:"clusterRole"`
}

type Role struct {
	TypeMeta   map[string]string   `json:",inline"`
	ObjectMeta metav1.ObjectMeta   `json:"metadata"`
	Rules      []rbacv1.PolicyRule `json:"rules"`
}

type RoleResponse struct {
	Role *rbacv1.Role `json:"role"`
}

type RoleBinding struct {
	TypeMeta   map[string]string `json:",inline"`
	ObjectMeta metav1.ObjectMeta `json:"metadata"`
	Subjects   []rbacv1.Subject  `json:"subjects,omitempty"`
	RoleRef    rbacv1.RoleRef    `json:"roleRef"`
}

type RoleBindingResponse struct {
	RoleBinding *rbacv1.RoleBinding `json:"role"`
}

type ClusterRoleBinding struct {
	TypeMeta   map[string]string `json:",inline"`
	ObjectMeta metav1.ObjectMeta `json:"metadata"`
	Subjects   []rbacv1.Subject  `json:"subjects,omitempty"`
	RoleRef    rbacv1.RoleRef    `json:"roleRef"`
}

type ClusterRoleBindingRespone struct {
	ClusterRoleBinding *rbacv1.ClusterRoleBinding `json:"clusterRoleBinding"`
}

type ServiceAccount struct {
	TypeMeta                     map[string]string             `json:",inline"`
	ObjectMeta                   metav1.ObjectMeta             `json:"metadata"`
	Secrets                      []corev1.ObjectReference      `json:"secrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=secrets"`
	ImagePullSecrets             []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty" protobuf:"bytes,3,rep,name=imagePullSecrets"`
	AutomountServiceAccountToken *bool                         `json:"automountServiceAccountToken,omitempty" protobuf:"varint,4,opt,name=automountServiceAccountToken"`
}

type ServiceAccountResponse struct {
	ServiceAccount *corev1.ServiceAccount `json:"serviceAccount"`
}

type K8s interface {
	Deployment(d *Deployment) (*DeploymentResponse, error)
	ConfigMap(cm *ConfigMap) (*ConfigMapResponse, error)
	Secret(sec *Secret) (*SecretResponse, error)
	Cronjob(c *CronJob) (*CronjobRespone, error)
	HorizontalPodAutoscaler(h *HorizontalPodAutoscaler) (*HorizontalPodAutoscalerResponse, error)
	Service(h *Service) (*ServiceResponse, error)
	Job(j *Job) (*JobResponse, error)
	ListPods(namespace string) ([]PodResponse, error)
	ClusterRole(cr *ClusterRole) (*ClusterRoleResponse, error)
	ClusterRoleBinding(crb *ClusterRoleBinding) (*ClusterRoleBindingRespone, error)
	Role(r *Role) (*RoleResponse, error)
	RoleBinding(rb *RoleBinding) (*RoleBindingResponse, error)
	ServiceAccount(sa *ServiceAccount) (*ServiceAccountResponse, error)
}