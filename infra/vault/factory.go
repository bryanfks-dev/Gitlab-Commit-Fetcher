package vault

type Vault struct {
	SvcName string
}

func NewVault(svcName string) *Vault {
	return &Vault{
		SvcName: svcName,
	}
}
