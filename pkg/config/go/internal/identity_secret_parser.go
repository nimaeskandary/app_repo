package internal

import config_types "github.com/nimaeskandary/app_repo/pkg/config/go/types"

type identitySecretParser struct{}

func NewIdentitySecretParser() config_types.SecretParser {
	return &identitySecretParser{}
}

func (p *identitySecretParser) Parse(raw string) (string, error) {
	return raw, nil
}
