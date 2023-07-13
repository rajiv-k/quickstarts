# kind - Kubernetes in Docker

<img src="../assets/kind.png" alt="kind logo" width="200"/>

kind is a tool for running local Kubernetes clusters using Docker container “nodes”.
kind was primarily designed for testing Kubernetes itself, but may be used for local development or CI.


## tl;dr

```sh
$ kind create cluster
```


## Install

```sh
$ go install sigs.k8s.io/kind@v0.20.0
```

This will put `kind` in `$(go env GOPATH)/bin` directory. You may need to add that directory to your `$PATH` if you encounter the error kind: command not found after installation.

## Creating cluster


To configure the kind cluster creation, we can also give it a config file.

```sh
## Create a config file for kind
$ cat <<! >kind-config.yaml
apiVersion: kind.x-k8s.io/v1alpha4
kind: Cluster
name: foo
nodes:
  - role: control-plane
    # add a mount from /path/to/my/files on the host to /files on the node
    extraMounts:
      - hostPath: /path/to/my/files
        containerPath: /files
    extraPortMappings:
    - containerPort: 80
      hostPort: 80
      protocol: TCP
    - containerPort: 443
      hostPort: 443
      protocol: TCP
!

## Use this config file while creating the cluster
$ kind create cluster --config ./kind-config.yaml

Creating cluster "kind" ...
 ✓ Ensuring node image (kindest/node:v1.24.0) 🖼
 ✓ Preparing nodes 📦
 ✓ Writing configuration 📜
 ✓ Starting control-plane 🕹️
 ✓ Installing CNI 🔌
 ✓ Installing StorageClass 💾
Set kubectl context to "kind-kind"
You can now use your cluster with:

kubectl cluster-info --context kind-kind

Not sure what to do next? 😅  Check out https://kind.sigs.k8s.io/docs/user/quick-start/

```
For more info refer: https://kind.sigs.k8s.io/docs/user/configuration/
