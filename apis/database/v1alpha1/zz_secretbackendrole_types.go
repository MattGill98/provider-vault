/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SecretBackendRoleObservation struct {

	// The path of the Database Secret Backend the role belongs to.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Database statements to execute to create and configure a user.
	CreationStatements []*string `json:"creationStatements,omitempty" tf:"creation_statements,omitempty"`

	// Specifies the configuration for the given credential_type.
	CredentialConfig map[string]*string `json:"credentialConfig,omitempty" tf:"credential_config,omitempty"`

	// Specifies the type of credential that will be generated for the role.
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// Database connection to use for this role.
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// Default TTL for leases associated with this role, in seconds.
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Maximum TTL for leases associated with this role, in seconds.
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Database statements to execute to renew a user.
	RenewStatements []*string `json:"renewStatements,omitempty" tf:"renew_statements,omitempty"`

	// Database statements to execute to revoke a user.
	RevocationStatements []*string `json:"revocationStatements,omitempty" tf:"revocation_statements,omitempty"`

	// Database statements to execute to rollback a create operation in the event of an error.
	RollbackStatements []*string `json:"rollbackStatements,omitempty" tf:"rollback_statements,omitempty"`
}

type SecretBackendRoleParameters struct {

	// The path of the Database Secret Backend the role belongs to.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Database statements to execute to create and configure a user.
	// +kubebuilder:validation:Optional
	CreationStatements []*string `json:"creationStatements,omitempty" tf:"creation_statements,omitempty"`

	// Specifies the configuration for the given credential_type.
	// +kubebuilder:validation:Optional
	CredentialConfig map[string]*string `json:"credentialConfig,omitempty" tf:"credential_config,omitempty"`

	// Specifies the type of credential that will be generated for the role.
	// +kubebuilder:validation:Optional
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// Database connection to use for this role.
	// +kubebuilder:validation:Optional
	DBName *string `json:"dbName,omitempty" tf:"db_name,omitempty"`

	// Default TTL for leases associated with this role, in seconds.
	// +kubebuilder:validation:Optional
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// Maximum TTL for leases associated with this role, in seconds.
	// +kubebuilder:validation:Optional
	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Database statements to execute to renew a user.
	// +kubebuilder:validation:Optional
	RenewStatements []*string `json:"renewStatements,omitempty" tf:"renew_statements,omitempty"`

	// Database statements to execute to revoke a user.
	// +kubebuilder:validation:Optional
	RevocationStatements []*string `json:"revocationStatements,omitempty" tf:"revocation_statements,omitempty"`

	// Database statements to execute to rollback a create operation in the event of an error.
	// +kubebuilder:validation:Optional
	RollbackStatements []*string `json:"rollbackStatements,omitempty" tf:"rollback_statements,omitempty"`
}

// SecretBackendRoleSpec defines the desired state of SecretBackendRole
type SecretBackendRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendRoleParameters `json:"forProvider"`
}

// SecretBackendRoleStatus defines the observed state of SecretBackendRole.
type SecretBackendRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendRole is the Schema for the SecretBackendRoles API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackendRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.backend)",message="backend is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.creationStatements)",message="creationStatements is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.dbName)",message="dbName is a required parameter"
	Spec   SecretBackendRoleSpec   `json:"spec"`
	Status SecretBackendRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendRoleList contains a list of SecretBackendRoles
type SecretBackendRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackendRole `json:"items"`
}

// Repository type metadata.
var (
	SecretBackendRole_Kind             = "SecretBackendRole"
	SecretBackendRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackendRole_Kind}.String()
	SecretBackendRole_KindAPIVersion   = SecretBackendRole_Kind + "." + CRDGroupVersion.String()
	SecretBackendRole_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackendRole_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackendRole{}, &SecretBackendRoleList{})
}