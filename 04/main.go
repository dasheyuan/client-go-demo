package main

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	//config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return
	}
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"
	//client
	client, err := rest.RESTClientFor(config)
	if err != nil {
		return
	}

	//get data
	var pods v1.PodList
	err = client.Get().Namespace("default").Resource("pods").Do(context.TODO()).Into(&pods)
	if err != nil {
		panic(err)
	}

}
