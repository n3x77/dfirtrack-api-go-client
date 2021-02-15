/*
 * DFIRTrack
 *
 * OpenAPI 3 - Documentation of DFIRTrack API
 *
 * API version: v0.4.7
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrackapi

import (
	"encoding/json"
)

// Systemtype struct for Systemtype
type Systemtype struct {
	SystemtypeId *int32 `json:"systemtype_id,omitempty"`
	SystemtypeName string `json:"systemtype_name"`
}

// NewSystemtype instantiates a new Systemtype object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemtype(systemtypeName string, ) *Systemtype {
	this := Systemtype{}
	this.SystemtypeName = systemtypeName
	return &this
}

// NewSystemtypeWithDefaults instantiates a new Systemtype object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemtypeWithDefaults() *Systemtype {
	this := Systemtype{}
	return &this
}

// GetSystemtypeId returns the SystemtypeId field value if set, zero value otherwise.
func (o *Systemtype) GetSystemtypeId() int32 {
	if o == nil || o.SystemtypeId == nil {
		var ret int32
		return ret
	}
	return *o.SystemtypeId
}

// GetSystemtypeIdOk returns a tuple with the SystemtypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Systemtype) GetSystemtypeIdOk() (*int32, bool) {
	if o == nil || o.SystemtypeId == nil {
		return nil, false
	}
	return o.SystemtypeId, true
}

// HasSystemtypeId returns a boolean if a field has been set.
func (o *Systemtype) HasSystemtypeId() bool {
	if o != nil && o.SystemtypeId != nil {
		return true
	}

	return false
}

// SetSystemtypeId gets a reference to the given int32 and assigns it to the SystemtypeId field.
func (o *Systemtype) SetSystemtypeId(v int32) {
	o.SystemtypeId = &v
}

// GetSystemtypeName returns the SystemtypeName field value
func (o *Systemtype) GetSystemtypeName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SystemtypeName
}

// GetSystemtypeNameOk returns a tuple with the SystemtypeName field value
// and a boolean to check if the value has been set.
func (o *Systemtype) GetSystemtypeNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SystemtypeName, true
}

// SetSystemtypeName sets field value
func (o *Systemtype) SetSystemtypeName(v string) {
	o.SystemtypeName = v
}

func (o Systemtype) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SystemtypeId != nil {
		toSerialize["systemtype_id"] = o.SystemtypeId
	}
	if true {
		toSerialize["systemtype_name"] = o.SystemtypeName
	}
	return json.Marshal(toSerialize)
}

type NullableSystemtype struct {
	value *Systemtype
	isSet bool
}

func (v NullableSystemtype) Get() *Systemtype {
	return v.value
}

func (v *NullableSystemtype) Set(val *Systemtype) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemtype) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemtype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemtype(val *Systemtype) *NullableSystemtype {
	return &NullableSystemtype{value: val, isSet: true}
}

func (v NullableSystemtype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemtype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


