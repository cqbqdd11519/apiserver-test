module github.com/cqbqdd11519/apiserver-test

go 1.14

require (
	github.com/gorilla/mux v1.7.4
	github.com/operator-framework/operator-sdk v0.17.1
	github.com/tmax-cloud/approval-watcher v0.0.0-20200806001512-2279c28ea568
	k8s.io/api v0.17.6
	k8s.io/apimachinery v0.17.6
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/kube-aggregator v0.17.3
	knative.dev/pkg v0.0.0-20200623024526-fb0320d9287e
	sigs.k8s.io/controller-runtime v0.5.2
)

replace k8s.io/client-go => k8s.io/client-go v0.17.4
