package clientset

import (
	"fmt"
	"os"
	"path/filepath"

	"ezone.xxxxx.com/xxxxx/xxxxx/communal/lib"
	clientversioned "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned"
	"golang.org/x/net/context"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var promOperator *clientversioned.Clientset

// 构建k8s集群的rest.Config对象
func buildClusterConfig() (*rest.Config, error) {
	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" && os.Getenv("KUBERNETES_SERVICE_PORT") != "" {
		config, err := rest.InClusterConfig()
		if err != nil {
			lib.Log.TagError(lib.GetTraceContext(context.Background()), lib.DLTagMySqlFailed, map[string]interface{}{
				"message": fmt.Errorf("create cluster config failed, %s", err).Error(),
			})
			return nil, err
		}
		return config, nil
	}

	home := homedir.HomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(context.Background()), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("build config from flags failed, %s", err).Error(),
		})
		return nil, err
	}

	return config, nil
}

// 通过rest.Config对象构建ClientSet 对象
func newPromOperatorClient(cfg *rest.Config) (*clientversioned.Clientset, error) {
	var err error
	promOperator, err = clientversioned.NewForConfig(cfg)
	if err != nil {
		lib.Log.TagError(lib.GetTraceContext(context.Background()), lib.DLTagMySqlFailed, map[string]interface{}{
			"message": fmt.Errorf("new ClientSet failed by cluster config, %s", err).Error(),
		})
		return nil, err
	}

	return promOperator, nil
}

// PromOperatorClientSet 获取Prometheus Operator的ClientSet
func PromOperatorClientSet() (*clientversioned.Clientset, error) {
	var err error
	var cfg *rest.Config
	if promOperator == nil {
		cfg, err = buildClusterConfig()
		if err != nil {
			return nil, err
		}
		promOperator, err = newPromOperatorClient(cfg)
		if err != nil {
			return nil, err
		}
	}

	return promOperator, nil
}

func ClientSet() *clientversioned.Clientset {
	return promOperator
}
