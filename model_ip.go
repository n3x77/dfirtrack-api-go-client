/*
DFIRTrack

OpenAPI 3 - Documentation of DFIRTrack API

API version: v2.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrackapi

import (
	"encoding/json"
)

// Ip struct for Ip
type Ip struct {
	IpId *int32 `json:"ip_id,omitempty"`
	IpIp string `json:"ip_ip"`
}

// NewIp instantiates a new Ip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIp(ipIp string) *Ip {
	this := Ip{}
	this.IpIp = ipIp
	return &this
}

// NewIpWithDefaults instantiates a new Ip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpWithDefaults() *Ip {
	this := Ip{}
	return &this
}

// GetIpId returns the IpId field value if set, zero value otherwise.
func (o *Ip) GetIpId() int32 {
	if o == nil || o.IpId == nil {
		var ret int32
		return ret
	}
	return *o.IpId
}

// GetIpIdOk returns a tuple with the IpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ip) GetIpIdOk() (*int32, bool) {
	if o == nil || o.IpId == nil {
		return nil, false
	}
	return o.IpId, true
}

// HasIpId returns a boolean if a field has been set.
func (o *Ip) HasIpId() bool {
	if o != nil && o.IpId != nil {
		return true
	}

	return false
}

// SetIpId gets a reference to the given int32 and assigns it to the IpId field.
func (o *Ip) SetIpId(v int32) {
	o.IpId = &v
}

// GetIpIp returns the IpIp field value
func (o *Ip) GetIpIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpIp
}

// GetIpIpOk returns a tuple with the IpIp field value
// and a boolean to check if the value has been set.
func (o *Ip) GetIpIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IpIp, true
}

// SetIpIp sets field value
func (o *Ip) SetIpIp(v string) {
	o.IpIp = v
}

func (o Ip) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpId != nil {
		toSerialize["ip_id"] = o.IpId
	}
	if true {
		toSerialize["ip_ip"] = o.IpIp
	}
	return json.Marshal(toSerialize)
}

type NullableIp struct {
	value *Ip
	isSet bool
}

func (v NullableIp) Get() *Ip {
	return v.value
}

func (v *NullableIp) Set(val *Ip) {
	v.value = val
	v.isSet = true
}

func (v NullableIp) IsSet() bool {
	return v.isSet
}

func (v *NullableIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIp(val *Ip) *NullableIp {
	return &NullableIp{value: val, isSet: true}
}

func (v NullableIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


