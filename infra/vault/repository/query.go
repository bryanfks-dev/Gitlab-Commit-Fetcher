package vaultrepository

import (
	"errors"
	vaulterrors "main/infra/vault/errors"

	"github.com/zalando/go-keyring"
)

func (r *vaultRepository) Get(key string) ([]byte, error) {
	item, err := keyring.Get(r.vault.SvcName, key)
	if errors.Is(err, keyring.ErrNotFound) {
		r.logger.Infof("Key '%s' not found in vault", key)
		return nil, nil
	}
	if err != nil {
		r.logger.Errorf("Error getting key from vault: %v", err)
		return nil, vaulterrors.ErrGetItemFailed
	}
	return []byte(item), nil
}
