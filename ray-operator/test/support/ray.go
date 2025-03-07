package support

import (
	"github.com/onsi/gomega"
	rayv1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func RayJob(t Test, namespace, name string) func(g gomega.Gomega) *rayv1.RayJob {
	return func(g gomega.Gomega) *rayv1.RayJob {
		job, err := t.Client().Ray().RayV1().RayJobs(namespace).Get(t.Ctx(), name, metav1.GetOptions{})
		g.Expect(err).NotTo(gomega.HaveOccurred())
		return job
	}
}

func GetRayJob(t Test, namespace, name string) *rayv1.RayJob {
	t.T().Helper()
	return RayJob(t, namespace, name)(t)
}

func RayJobStatus(job *rayv1.RayJob) rayv1.JobStatus {
	return job.Status.JobStatus
}

func RayJobDeploymentStatus(job *rayv1.RayJob) rayv1.JobDeploymentStatus {
	return job.Status.JobDeploymentStatus
}

func GetRayJobId(t Test, namespace, name string) string {
	t.T().Helper()
	job := RayJob(t, namespace, name)(t)
	return job.Status.JobId
}

func RayCluster(t Test, namespace, name string) func(g gomega.Gomega) *rayv1.RayCluster {
	return func(g gomega.Gomega) *rayv1.RayCluster {
		cluster, err := t.Client().Ray().RayV1().RayClusters(namespace).Get(t.Ctx(), name, metav1.GetOptions{})
		g.Expect(err).NotTo(gomega.HaveOccurred())
		return cluster
	}
}

func RayClusterOrError(t Test, namespace, name string) func(g gomega.Gomega) (*rayv1.RayCluster, error) {
	return func(g gomega.Gomega) (*rayv1.RayCluster, error) {
		return t.Client().Ray().RayV1().RayClusters(namespace).Get(t.Ctx(), name, metav1.GetOptions{})
	}
}

func GetRayCluster(t Test, namespace, name string) *rayv1.RayCluster {
	t.T().Helper()
	return RayCluster(t, namespace, name)(t)
}

func RayClusterState(cluster *rayv1.RayCluster) rayv1.ClusterState {
	return cluster.Status.State
}
