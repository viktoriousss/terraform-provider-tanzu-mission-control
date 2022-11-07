/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/
// Code generated by go-swagger; DO NOT EDIT.

package credentialmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1AccountCredentialData Credential data.
//
// swagger:model vmware.tanzu.manage.v1alpha1.account.credential.Data
type VmwareTanzuManageV1alpha1AccountCredentialData struct {

	// AWS credential.
	AwsCredential *VmwareTanzuManageV1alpha1AccountCredentialTypeAwsSpec `json:"awsCredential,omitempty"`

	// Generic credential.
	GenericCredential string `json:"genericCredential,omitempty"`

	// Key Value credential.
	KeyValue *VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec `json:"keyValue,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1AccountCredentialData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1AccountCredentialData) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1AccountCredentialData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec KeyValue Type credential stored in Account Manager.
//
// swagger:model vmware.tanzu.manage.v1alpha1.account.credential.type.keyvalue.Spec
type VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpec struct {

	// Data contains the secret data. Each key must consist of alphanumeric
	// characters, '-', '_' or '.'.
	Data map[string]strfmt.Base64 `json:"data,omitempty"`

	// Type of Secret.
	// The default value is SECRET_TYPE_OPAQUE.
	Type *VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType `json:"type,omitempty"`
}

// VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType Type of Secret.
//
//   - SECRET_TYPE_UNSPECIFIED: SECRET_TYPE_UNSPECIFIED is default.
//   - OPAQUE_SECRET_TYPE: SECRET_TYPE_OPAQUE maps to the k8s secret type OPAQUE.
//
// It is arbitrary user-defined data.
//   - DOCKERCONFIGJSON_SECRET_TYPE: DOCKERCONFIGJSON_SECRET_TYPE maps to  Kubernetes secrets type kubernetes.io/dockerconfigjson.
//
// swagger:model vmware.tanzu.manage.v1alpha1.account.credential.type.keyvalue.Spec.SecretType
type VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType string

func NewVmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType(value VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType) *VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType.
func (m VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType) Pointer() *VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretTypeSECRETTYPEUNSPECIFIED captures enum value "SECRET_TYPE_UNSPECIFIED"
	VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretTypeSECRETTYPEUNSPECIFIED VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType = "SECRET_TYPE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretTypeOPAQUESECRETTYPE captures enum value "OPAQUE_SECRET_TYPE"
	VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretTypeOPAQUESECRETTYPE VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType = "OPAQUE_SECRET_TYPE"

	// VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretTypeDOCKERCONFIGJSONSECRETTYPE captures enum value "DOCKERCONFIGJSON_SECRET_TYPE"
	VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretTypeDOCKERCONFIGJSONSECRETTYPE VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType = "DOCKERCONFIGJSON_SECRET_TYPE"
)

// for schema
var vmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretTypeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretType
	if err := json.Unmarshal([]byte(`["SECRET_TYPE_UNSPECIFIED","OPAQUE_SECRET_TYPE","DOCKERCONFIGJSON_SECRET_TYPE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretTypeEnum = append(vmwareTanzuManageV1alpha1AccountCredentialTypeKeyvalueSpecSecretTypeEnum, v)
	}
}
