package vaultrepository

import (
	"main/infra/logger"
	"main/infra/vault"
)

type vaultRepository struct {
	logger logger.Logger
	vault  *vault.Vault
}

type VaultRepository interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	Delete(key string) error
}

func NewVaultRepository(
	logger logger.Logger,
	vault *vault.Vault,
) VaultRepository {
	return &vaultRepository{
		logger: logger,
		vault:  vault,
	}
}
