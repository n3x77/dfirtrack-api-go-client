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

// Reason struct for Reason
type Reason struct {
	ReasonId *int32 `json:"reason_id,omitempty"`
	ReasonName string `json:"reason_name"`
}

// NewReason instantiates a new Reason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReason(reasonName string, ) *Reason {
	this := Reason{}
	this.ReasonName = reasonName
	return &this
}

// NewReasonWithDefaults instantiates a new Reason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReasonWithDefaults() *Reason {
	this := Reason{}
	return &this
}

// GetReasonId returns the ReasonId field value if set, zero value otherwise.
func (o *Reason) GetReasonId() int32 {
	if o == nil || o.ReasonId == nil {
		var ret int32
		return ret
	}
	return *o.ReasonId
}

// GetReasonIdOk returns a tuple with the ReasonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reason) GetReasonIdOk() (*int32, bool) {
	if o == nil || o.ReasonId == nil {
		return nil, false
	}
	return o.ReasonId, true
}

// HasReasonId returns a boolean if a field has been set.
func (o *Reason) HasReasonId() bool {
	if o != nil && o.ReasonId != nil {
		return true
	}

	return false
}

// SetReasonId gets a reference to the given int32 and assigns it to the ReasonId field.
func (o *Reason) SetReasonId(v int32) {
	o.ReasonId = &v
}

// GetReasonName returns the ReasonName field value
func (o *Reason) GetReasonName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ReasonName
}

// GetReasonNameOk returns a tuple with the ReasonName field value
// and a boolean to check if the value has been set.
func (o *Reason) GetReasonNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReasonName, true
}

// SetReasonName sets field value
func (o *Reason) SetReasonName(v string) {
	o.ReasonName = v
}

func (o Reason) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReasonId != nil {
		toSerialize["reason_id"] = o.ReasonId
	}
	if true {
		toSerialize["reason_name"] = o.ReasonName
	}
	return json.Marshal(toSerialize)
}

type NullableReason struct {
	value *Reason
	isSet bool
}

func (v NullableReason) Get() *Reason {
	return v.value
}

func (v *NullableReason) Set(val *Reason) {
	v.value = val
	v.isSet = true
}

func (v NullableReason) IsSet() bool {
	return v.isSet
}

func (v *NullableReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReason(val *Reason) *NullableReason {
	return &NullableReason{value: val, isSet: true}
}

func (v NullableReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


