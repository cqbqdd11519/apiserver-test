package utils

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func Client(options client.Options) (client.Client, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	c, err := client.New(cfg, options)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func AddSchemes(opt *client.Options, gv schema.GroupVersion, types... runtime.Object) {
	if opt.Scheme == nil {
		opt.Scheme = runtime.NewScheme()
	}
	opt.Scheme.AddKnownTypes(gv, types...)
}
