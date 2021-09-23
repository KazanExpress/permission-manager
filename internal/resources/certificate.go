package resources

import (
	"context"
	"encoding/base64"
	"fmt"
	"sighupio/permission-manager/internal/config"

	cv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "sigs.k8s.io/controller-runtime"
)

// getCaBase64 returns the base64 encoding of the Kubernetes cluster api-server CA
func (r *Manager) getCaBase64(cluster config.ClusterConfig, username string) (string, error) {

	switch cluster.CASource {
	case "kubeconfig":

		kConfig, err := runtime.GetConfig()

		if err != nil {
			return "", fmt.Errorf("unable to get kubeconfig: %w", err)
		}

		return base64.StdEncoding.EncodeToString(kConfig.CAData), nil
	case "serviceaccount":
		secrets, err := r.kubeclient.CoreV1().Secrets(cluster.Namespace).List(context.Background(), v1.ListOptions{})
		if err != nil {
			return "", fmt.Errorf("secrets list: %w", err)
		}

		for _, secret := range secrets.Items {
			if secret.Type != cv1.SecretTypeServiceAccountToken {
				continue
			}

			saName, ok := secret.Annotations[cv1.ServiceAccountNameKey]
			if !ok || saName != username {
				continue
			}

			caCrt, ok := secret.Data[cv1.ServiceAccountRootCAKey]
			if !ok {
				return "", fmt.Errorf("service account secret does not contain ca.crt")
			}

			return base64.StdEncoding.EncodeToString(caCrt), nil
		}

		return "", fmt.Errorf("no secrets found for service account")

	default:
		return "", fmt.Errorf("unknown ca source: %s", cluster.CASource)

	}

}
