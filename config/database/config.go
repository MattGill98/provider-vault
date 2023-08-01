package database

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/config"
)

const (
	ErrFmtNoAttribute    = "required attribute %s not found"
	ErrFmtUnexpectedType = "attribute %s had the wrong type"
)

func getNameFromFullyQualifiedID(tfstate map[string]any) (string, error) {
	id, ok := tfstate["id"]
	if !ok {
		return "", errors.Errorf(ErrFmtNoAttribute, "id")
	}
	idStr, ok := id.(string)
	if !ok {
		return "", errors.Errorf(ErrFmtUnexpectedType, "id")
	}
	words := strings.Split(idStr, "/")
	return words[len(words)-1], nil
}

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_database_secret_backend_connection", func(r *config.Resource) {
		r.ExternalName = config.TemplatedStringAsIdentifier("name", "{{ .parameters.backend }}/config/{{ .externalName }}")
		r.ExternalName.GetExternalNameFn = getNameFromFullyQualifiedID
		r.References["backend"] = config.Reference{
			TerraformName: "vault_mount",
		}
	})
	p.AddResourceConfigurator("vault_database_secret_backend_role", func(r *config.Resource) {
		r.ExternalName = config.TemplatedStringAsIdentifier("name", "{{ .parameters.backend }}/roles/{{ .externalName }}")
		r.ExternalName.GetExternalNameFn = getNameFromFullyQualifiedID
		r.References["backend"] = config.Reference{
			TerraformName: "vault_mount",
		}
	})
}
