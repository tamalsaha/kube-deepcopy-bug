# kube-deepcopy-bug

```console
$ ./hack/codegen.sh 
++ go env GOPATH
+ GOPATH=/home/tamal/go
+ PACKAGE_NAME=github.com/tamalsaha/kube-deepcopy-bug
+ REPO_ROOT=/home/tamal/go/src/github.com/tamalsaha/kube-deepcopy-bug
+ DOCKER_REPO_ROOT=/go/src/github.com/tamalsaha/kube-deepcopy-bug
+ DOCKER_CODEGEN_PKG=/go/src/k8s.io/code-generator
+ pushd /home/tamal/go/src/github.com/tamalsaha/kube-deepcopy-bug
~/go/src/github.com/tamalsaha/kube-deepcopy-bug ~/go/src/github.com/tamalsaha/kube-deepcopy-bug
++ id -u
++ id -g
+ docker run --rm -ti -u 1000:1000 -v /home/tamal/go/src/github.com/tamalsaha/kube-deepcopy-bug:/go/src/github.com/tamalsaha/kube-deepcopy-bug -w /go/src/github.com/tamalsaha/kube-deepcopy-bug appscode/gengo:release-1.9 /go/src/k8s.io/code-generator/generate-groups.sh deepcopy github.com/tamalsaha/kube-deepcopy-bug/client github.com/tamalsaha/kube-deepcopy-bug/apis demo:v1alpha1 --go-header-file /go/src/github.com/tamalsaha/kube-deepcopy-bug/hack/gengo/boilerplate.go.txt
Generating deepcopy funcs

+ go build ./...
# github.com/tamalsaha/kube-deepcopy-bug/apis/demo/v1alpha1
apis/demo/v1alpha1/zz_generated.deepcopy.go:37:9: (*in).DeepCopyInto undefined (type *AggregatePlan has no field or method DeepCopyInto)
+ popd
~/go/src/github.com/tamalsaha/kube-deepcopy-bug

```
