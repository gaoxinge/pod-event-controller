package pkg

import (
	"errors"
	"log"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type PodEventController struct {
	clientSet kubernetes.Interface
}

func NewPodEventController(clientSet kubernetes.Interface) *PodEventController {
	podEventController := PodEventController{
		clientSet: clientSet,
	}
	return &podEventController
}

func (podEventController *PodEventController) Run(stopCh chan struct{}) error {
	informerFactory := informers.NewSharedInformerFactory(podEventController.clientSet, 30*time.Second)
	podInformer := informerFactory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    podEventController.add,
		UpdateFunc: podEventController.update,
		DeleteFunc: podEventController.delete,
	})
	informerFactory.Start(stopCh)
	if !cache.WaitForCacheSync(stopCh, podInformer.Informer().HasSynced) {
		return errors.New("run controller with fail to sync")
	}
	return nil
}

func (podEventController *PodEventController) add(obj interface{}) {
	pod := obj.(*v1.Pod)
	log.Printf("add %s\n", pod.Name)
}

func (podEventController *PodEventController) update(oldObj, newObj interface{}) {
	pod1 := oldObj.(*v1.Pod)
	pod2 := oldObj.(*v1.Pod)
	log.Printf("update old %s to new %s\n", pod1.Name, pod2.Name)
}

func (podEventController *PodEventController) delete(obj interface{}) {
	pod := obj.(*v1.Pod)
	log.Printf("delete %s\n", pod.Name)
}
