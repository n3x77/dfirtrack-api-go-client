/*
 * DFIRTrack
 *
 * OpenAPI 3 - Documentation of DFIRTrack API
 *
 * API version: v1.5.5
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrackapi

import (
	"encoding/json"
)

// Domain struct for Domain
type Domain struct {
	DomainId *int32 `json:"domain_id,omitempty"`
	DomainName string `json:"domain_name"`
}

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain(domainName string, ) *Domain {
	this := Domain{}
	this.DomainName = domainName
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	return &this
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *Domain) GetDomainId() int32 {
	if o == nil || o.DomainId == nil {
		var ret int32
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDomainIdOk() (*int32, bool) {
	if o == nil || o.DomainId == nil {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *Domain) HasDomainId() bool {
	if o != nil && o.DomainId != nil {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given int32 and assigns it to the DomainId field.
func (o *Domain) SetDomainId(v int32) {
	o.DomainId = &v
}

// GetDomainName returns the DomainName field value
func (o *Domain) GetDomainName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value
// and a boolean to check if the value has been set.
func (o *Domain) GetDomainNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DomainName, true
}

// SetDomainName sets field value
func (o *Domain) SetDomainName(v string) {
	o.DomainName = v
}

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DomainId != nil {
		toSerialize["domain_id"] = o.DomainId
	}
	if true {
		toSerialize["domain_name"] = o.DomainName
	}
	return json.Marshal(toSerialize)
}

type NullableDomain struct {
	value *Domain
	isSet bool
}

func (v NullableDomain) Get() *Domain {
	return v.value
}

func (v *NullableDomain) Set(val *Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomain(val *Domain) *NullableDomain {
	return &NullableDomain{value: val, isSet: true}
}

func (v NullableDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


