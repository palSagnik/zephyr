package deployment

import (
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func GetK8sClient() (*kubernetes.Clientset, error) {

    // In-Cluster Config
    // Finally the application would be running inside a pod
    var (
        restConfig  *rest.Config
        clientset   *kubernetes.Clientset
        err         error
    )

    // restConfig, err = rest.InClusterConfig()
    // if err != nil {
    //     return nil, err
    // }

    // For now we go for Out-Cluster Config
    home := homedir.HomeDir()
    kubeConfigPath := filepath.Join(home, ".kube", "config")
    restConfig, err = clientcmd.BuildConfigFromFlags("", kubeConfigPath)
    if err != nil {
        return nil, err
    }

    // creating k8s client
    clientset, err = kubernetes.NewForConfig(restConfig)
    if err != nil {
        return nil, err
    }

    return clientset, nil
}