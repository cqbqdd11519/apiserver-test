module github.com/cqbqdd11519/apiserver-test

go 1.14

require (
	github.com/gorilla/mux v1.7.4
	github.com/operator-framework/operator-sdk v0.17.1
	github.com/tmax-cloud/approval-watcher v0.0.0-20200806001512-2279c28ea568
	k8s.io/apimachinery v0.17.6
	sigs.k8s.io/controller-runtime v0.5.2
)

replace k8s.io/client-go => k8s.io/client-go v0.17.4
