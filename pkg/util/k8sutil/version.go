package k8sutil

import (
	"strings"

	"github.com/coreos/go-semver/semver"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"

	"yunion.io/x/pkg/errors"
)

type ClusterVersion interface {
	GetRawVersion() string
	GetSemVerion() *semver.Version
	IsLessThan(iv string) bool
	IsGreatOrEqualThan(iv string) bool
	GetIngressGVR() schema.GroupVersionResource
}

type clusterVersion struct {
	client  kubernetes.Interface
	version *version.Info
}

func NewClusterVersion(cli kubernetes.Interface) (ClusterVersion, error) {
	ver, err := cli.Discovery().ServerVersion()
	if err != nil {
		return nil, errors.Wrap(err, "Get ServerVersion")
	}
	return &clusterVersion{
		client:  cli,
		version: ver,
	}, nil
}

func (cv clusterVersion) GetRawVersion() string {
	return cv.version.GitVersion
}

func (cv clusterVersion) GetSemVerion() *semver.Version {
	rv := strings.TrimPrefix(cv.GetRawVersion(), "v")
	return semver.New(rv)
}

func (cv clusterVersion) IsLessThan(iv string) bool {
	iv = strings.TrimPrefix(iv, "v")
	return cv.GetSemVerion().LessThan(*semver.New(iv))
}

func (cv clusterVersion) IsGreatOrEqualThan(iv string) bool {
	return !cv.IsLessThan(iv)
}

func (cv clusterVersion) GetIngressGVR() schema.GroupVersionResource {
	if cv.IsGreatOrEqualThan("1.19.0") {
		return schema.GroupVersionResource{
			Group:    "networking.k8s.io",
			Version:  "v1",
			Resource: "ingresses",
		}
	} else if cv.IsGreatOrEqualThan("1.14.0") {
		return schema.GroupVersionResource{
			Group:    "networking.k8s.io",
			Version:  "v1beta1",
			Resource: "ingresses",
		}
	}
	return schema.GroupVersionResource{
		Group:    "extensions",
		Version:  "v1beta1",
		Resource: "ingresses",
	}
}
