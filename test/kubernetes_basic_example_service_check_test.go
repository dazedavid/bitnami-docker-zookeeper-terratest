package test

import (
	"testing"
	"time"

	"github.com/reiver/go-telnet"

	"github.com/gruntwork-io/terratest/modules/k8s"
)

func TestKubernetes(t *testing.T) {

	namespaceName := "zookeeper"
	kubeResourcePath := "../kubernetes.yml"

	options := k8s.NewKubectlOptions("", "", namespaceName)

	k8s.CreateNamespace(t, options, namespaceName)
	defer k8s.DeleteNamespace(t, options, namespaceName)

	defer k8s.KubectlDelete(t, options, kubeResourcePath)
	k8s.KubectlApply(t, options, kubeResourcePath)
	k8s.WaitUntilServiceAvailable(t, options, "zookeeper", 10, 5*time.Second)
	service := k8s.GetService(t, options, "zookeeper")
	endpoint := k8s.GetServiceEndpoint(t, options, service, 2181)
	var caller telnet.Caller = telnet.StandardCaller

	telnet.DialToAndCall(endpoint, caller)
}
