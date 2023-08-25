generate:
	GOFLAGS=-mod=mod go run sigs.k8s.io/controller-tools/cmd/controller-gen@v0.8.0 object:headerFile="hack/boilerplate.go.txt" paths="./api/v1alpha1/"
	GOFLAGS=-mod=mod go run sigs.k8s.io/controller-tools/cmd/controller-gen@v0.8.0 crd paths="./api/v1alpha1" output:crd:artifacts:config=.
