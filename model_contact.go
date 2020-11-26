/*
 * DFIRTrack
 *
 * OpenAPI 3 - Documentation of DFIRTrack API
 *
 * API version: 0.4.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfirtrack-api-client

import (
	"encoding/json"
)

// Contact struct for Contact
type Contact struct {
	ContactId *int32 `json:"contact_id,omitempty"`
	ContactName string `json:"contact_name"`
	ContactEmail string `json:"contact_email"`
	ContactPhone NullableString `json:"contact_phone,omitempty"`
}

// NewContact instantiates a new Contact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContact(contactName string, contactEmail string, ) *Contact {
	this := Contact{}
	this.ContactName = contactName
	this.ContactEmail = contactEmail
	return &this
}

// NewContactWithDefaults instantiates a new Contact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactWithDefaults() *Contact {
	this := Contact{}
	return &this
}

// GetContactId returns the ContactId field value if set, zero value otherwise.
func (o *Contact) GetContactId() int32 {
	if o == nil || o.ContactId == nil {
		var ret int32
		return ret
	}
	return *o.ContactId
}

// GetContactIdOk returns a tuple with the ContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetContactIdOk() (*int32, bool) {
	if o == nil || o.ContactId == nil {
		return nil, false
	}
	return o.ContactId, true
}

// HasContactId returns a boolean if a field has been set.
func (o *Contact) HasContactId() bool {
	if o != nil && o.ContactId != nil {
		return true
	}

	return false
}

// SetContactId gets a reference to the given int32 and assigns it to the ContactId field.
func (o *Contact) SetContactId(v int32) {
	o.ContactId = &v
}

// GetContactName returns the ContactName field value
func (o *Contact) GetContactName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ContactName
}

// GetContactNameOk returns a tuple with the ContactName field value
// and a boolean to check if the value has been set.
func (o *Contact) GetContactNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContactName, true
}

// SetContactName sets field value
func (o *Contact) SetContactName(v string) {
	o.ContactName = v
}

// GetContactEmail returns the ContactEmail field value
func (o *Contact) GetContactEmail() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value
// and a boolean to check if the value has been set.
func (o *Contact) GetContactEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContactEmail, true
}

// SetContactEmail sets field value
func (o *Contact) SetContactEmail(v string) {
	o.ContactEmail = v
}

// GetContactPhone returns the ContactPhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Contact) GetContactPhone() string {
	if o == nil || o.ContactPhone.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContactPhone.Get()
}

// GetContactPhoneOk returns a tuple with the ContactPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Contact) GetContactPhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContactPhone.Get(), o.ContactPhone.IsSet()
}

// HasContactPhone returns a boolean if a field has been set.
func (o *Contact) HasContactPhone() bool {
	if o != nil && o.ContactPhone.IsSet() {
		return true
	}

	return false
}

// SetContactPhone gets a reference to the given NullableString and assigns it to the ContactPhone field.
func (o *Contact) SetContactPhone(v string) {
	o.ContactPhone.Set(&v)
}
// SetContactPhoneNil sets the value for ContactPhone to be an explicit nil
func (o *Contact) SetContactPhoneNil() {
	o.ContactPhone.Set(nil)
}

// UnsetContactPhone ensures that no value is present for ContactPhone, not even an explicit nil
func (o *Contact) UnsetContactPhone() {
	o.ContactPhone.Unset()
}

func (o Contact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContactId != nil {
		toSerialize["contact_id"] = o.ContactId
	}
	if true {
		toSerialize["contact_name"] = o.ContactName
	}
	if true {
		toSerialize["contact_email"] = o.ContactEmail
	}
	if o.ContactPhone.IsSet() {
		toSerialize["contact_phone"] = o.ContactPhone.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableContact struct {
	value *Contact
	isSet bool
}

func (v NullableContact) Get() *Contact {
	return v.value
}

func (v *NullableContact) Set(val *Contact) {
	v.value = val
	v.isSet = true
}

func (v NullableContact) IsSet() bool {
	return v.isSet
}

func (v *NullableContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContact(val *Contact) *NullableContact {
	return &NullableContact{value: val, isSet: true}
}

func (v NullableContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


