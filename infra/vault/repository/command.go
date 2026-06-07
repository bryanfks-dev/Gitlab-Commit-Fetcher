package vaultrepository

import (
	vaulterrors "main/infra/vault/errors"

	"github.com/zalando/go-keyring"
)

func (r *vaultRepository) Set(key string, value []byte) error {
	valStr := string(value)
	if err := keyring.Set(r.vault.SvcName, key, valStr); err != nil {
		r.logger.Errorf("Error setting key in vault: %v", err)
		return vaulterrors.ErrSetItemFailed
	}
	return nil
}

func (r *vaultRepository) Delete(key string) error {
	if err := keyring.Delete(r.vault.SvcName, key); err != nil {
		r.logger.Errorf("Error deleting key from vault: %v", err)
		return vaulterrors.ErrDeleteItemFailed
	}
	return nil
}
