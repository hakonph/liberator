package nais_io_v1_test

import (
	v1 "github.com/nais/liberator/pkg/apis/nais.io/v1"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestMaskinportenClient_CalculateHash(t *testing.T) {
	actual, err := minimalMaskinportenClient().Hash()
	assert.NoError(t, err)
	assert.Equal(t, "98105fd6e1607430", actual)
}

func TestIDPortenClient_Hash(t *testing.T) {
	actual, err := minimalIDPortenClient().Hash()
	assert.NoError(t, err)
	assert.Equal(t, "8b5ebee90b513411", actual)
}

func minimalMaskinportenClient() *v1.MaskinportenClient {
	return &v1.MaskinportenClient{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "test-app",
			Namespace:   "test-namespace",
			ClusterName: "test-cluster",
		},
		Spec: v1.MaskinportenClientSpec{
			Scopes: nil,
		},
	}
}

func minimalIDPortenClient() *v1.IDPortenClient {
	return &v1.IDPortenClient{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "test-app",
			Namespace:   "test-namespace",
			ClusterName: "test-cluster",
		},
		Spec: v1.IDPortenClientSpec{
			ClientURI:   "",
			RedirectURI: "https://test.com",
			SecretName:  "test",
		},
	}
}
