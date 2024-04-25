package googlecloud

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

type GetSecretRequest struct {
	SecretName string
}
type SecretsManager interface {
	Get(ctx context.Context, req GetSecretRequest) ([]byte, error)
}

type secretManager struct {
	client *secretmanager.Client
}

var _ SecretsManager = &secretManager{}

func NewSecretManager() (SecretsManager, error) {
	client, err := secretmanager.NewClient(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve secret manager client: %w", err)
	}
	return &secretManager{client: client}, nil
}

func (s *secretManager) Get(ctx context.Context, req GetSecretRequest) ([]byte, error) {
	secret, err := s.client.AccessSecretVersion(
		ctx,
		&secretmanagerpb.AccessSecretVersionRequest{
			Name: req.SecretName,
		},
	)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to get secret: %w", err)
	}

	return secret.Payload.Data, nil
}
