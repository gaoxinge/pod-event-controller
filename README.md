# pod event controller

[kubernetes/sample-controller](https://github.com/kubernetes/sample-controller) is not simple.

## docker

You can build and push image by yourself

```shell script
docker build -f docker/Dockerfile -t gaoxinge/pod-event-controller .
docker push gaoxinge/pod-event-controller:latest
```

or directly use

- [docker hub](https://hub.docker.com/r/gaoxinge/pod-event-controller)

## yaml

```shell script
kubectl create -f yaml/pod-event-controller.yaml
kubectl create serviceaccount pod-event-controller
kubectl create clusterrolebinding pod-event-controller --clusterrole=cluster-admin --serviceaccount=default:pod-event-controller
kubectl delete clusterrolebinding pod-event-controller
kubectl delete serviceaccount pod-event-controller
kubectl delete deployment pod-event-controller
```

## test

```shell script
kubectl get pod # pod-event-controller-5b8568f979-x7shr
kubectl logs -f pod-event-controller-5b8568f979-x7shr -c main
```

## TODO

- [ ] add pod event controller test