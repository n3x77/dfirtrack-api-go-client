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
	"time"
)

// Reportitem struct for Reportitem
type Reportitem struct {
	ReportitemId *int32 `json:"reportitem_id,omitempty"`
	Case NullableInt32 `json:"case,omitempty"`
	Headline int32 `json:"headline"`
	Notestatus *int32 `json:"notestatus,omitempty"`
	System int32 `json:"system"`
	Tag *[]int32 `json:"tag,omitempty"`
	ReportitemSubheadline NullableString `json:"reportitem_subheadline"`
	ReportitemNote string `json:"reportitem_note"`
	ReportitemCreateTime *time.Time `json:"reportitem_create_time,omitempty"`
	ReportitemCreatedByUserId int32 `json:"reportitem_created_by_user_id"`
	ReportitemModifyTime *time.Time `json:"reportitem_modify_time,omitempty"`
	ReportitemModifiedByUserId int32 `json:"reportitem_modified_by_user_id"`
}

// NewReportitem instantiates a new Reportitem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportitem(headline int32, system int32, reportitemSubheadline NullableString, reportitemNote string, reportitemCreatedByUserId int32, reportitemModifiedByUserId int32, ) *Reportitem {
	this := Reportitem{}
	this.Headline = headline
	this.System = system
	this.ReportitemSubheadline = reportitemSubheadline
	this.ReportitemNote = reportitemNote
	this.ReportitemCreatedByUserId = reportitemCreatedByUserId
	this.ReportitemModifiedByUserId = reportitemModifiedByUserId
	return &this
}

// NewReportitemWithDefaults instantiates a new Reportitem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportitemWithDefaults() *Reportitem {
	this := Reportitem{}
	return &this
}

// GetReportitemId returns the ReportitemId field value if set, zero value otherwise.
func (o *Reportitem) GetReportitemId() int32 {
	if o == nil || o.ReportitemId == nil {
		var ret int32
		return ret
	}
	return *o.ReportitemId
}

// GetReportitemIdOk returns a tuple with the ReportitemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reportitem) GetReportitemIdOk() (*int32, bool) {
	if o == nil || o.ReportitemId == nil {
		return nil, false
	}
	return o.ReportitemId, true
}

// HasReportitemId returns a boolean if a field has been set.
func (o *Reportitem) HasReportitemId() bool {
	if o != nil && o.ReportitemId != nil {
		return true
	}

	return false
}

// SetReportitemId gets a reference to the given int32 and assigns it to the ReportitemId field.
func (o *Reportitem) SetReportitemId(v int32) {
	o.ReportitemId = &v
}

// GetCase returns the Case field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Reportitem) GetCase() int32 {
	if o == nil || o.Case.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Case.Get()
}

// GetCaseOk returns a tuple with the Case field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Reportitem) GetCaseOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Case.Get(), o.Case.IsSet()
}

// HasCase returns a boolean if a field has been set.
func (o *Reportitem) HasCase() bool {
	if o != nil && o.Case.IsSet() {
		return true
	}

	return false
}

// SetCase gets a reference to the given NullableInt32 and assigns it to the Case field.
func (o *Reportitem) SetCase(v int32) {
	o.Case.Set(&v)
}
// SetCaseNil sets the value for Case to be an explicit nil
func (o *Reportitem) SetCaseNil() {
	o.Case.Set(nil)
}

// UnsetCase ensures that no value is present for Case, not even an explicit nil
func (o *Reportitem) UnsetCase() {
	o.Case.Unset()
}

// GetHeadline returns the Headline field value
func (o *Reportitem) GetHeadline() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value
// and a boolean to check if the value has been set.
func (o *Reportitem) GetHeadlineOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Headline, true
}

// SetHeadline sets field value
func (o *Reportitem) SetHeadline(v int32) {
	o.Headline = v
}

// GetNotestatus returns the Notestatus field value if set, zero value otherwise.
func (o *Reportitem) GetNotestatus() int32 {
	if o == nil || o.Notestatus == nil {
		var ret int32
		return ret
	}
	return *o.Notestatus
}

// GetNotestatusOk returns a tuple with the Notestatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reportitem) GetNotestatusOk() (*int32, bool) {
	if o == nil || o.Notestatus == nil {
		return nil, false
	}
	return o.Notestatus, true
}

// HasNotestatus returns a boolean if a field has been set.
func (o *Reportitem) HasNotestatus() bool {
	if o != nil && o.Notestatus != nil {
		return true
	}

	return false
}

// SetNotestatus gets a reference to the given int32 and assigns it to the Notestatus field.
func (o *Reportitem) SetNotestatus(v int32) {
	o.Notestatus = &v
}

// GetSystem returns the System field value
func (o *Reportitem) GetSystem() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.System
}

// GetSystemOk returns a tuple with the System field value
// and a boolean to check if the value has been set.
func (o *Reportitem) GetSystemOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.System, true
}

// SetSystem sets field value
func (o *Reportitem) SetSystem(v int32) {
	o.System = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *Reportitem) GetTag() []int32 {
	if o == nil || o.Tag == nil {
		var ret []int32
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reportitem) GetTagOk() (*[]int32, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *Reportitem) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []int32 and assigns it to the Tag field.
func (o *Reportitem) SetTag(v []int32) {
	o.Tag = &v
}

// GetReportitemSubheadline returns the ReportitemSubheadline field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Reportitem) GetReportitemSubheadline() string {
	if o == nil || o.ReportitemSubheadline.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReportitemSubheadline.Get()
}

// GetReportitemSubheadlineOk returns a tuple with the ReportitemSubheadline field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Reportitem) GetReportitemSubheadlineOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReportitemSubheadline.Get(), o.ReportitemSubheadline.IsSet()
}

// SetReportitemSubheadline sets field value
func (o *Reportitem) SetReportitemSubheadline(v string) {
	o.ReportitemSubheadline.Set(&v)
}

// GetReportitemNote returns the ReportitemNote field value
func (o *Reportitem) GetReportitemNote() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ReportitemNote
}

// GetReportitemNoteOk returns a tuple with the ReportitemNote field value
// and a boolean to check if the value has been set.
func (o *Reportitem) GetReportitemNoteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReportitemNote, true
}

// SetReportitemNote sets field value
func (o *Reportitem) SetReportitemNote(v string) {
	o.ReportitemNote = v
}

// GetReportitemCreateTime returns the ReportitemCreateTime field value if set, zero value otherwise.
func (o *Reportitem) GetReportitemCreateTime() time.Time {
	if o == nil || o.ReportitemCreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ReportitemCreateTime
}

// GetReportitemCreateTimeOk returns a tuple with the ReportitemCreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reportitem) GetReportitemCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.ReportitemCreateTime == nil {
		return nil, false
	}
	return o.ReportitemCreateTime, true
}

// HasReportitemCreateTime returns a boolean if a field has been set.
func (o *Reportitem) HasReportitemCreateTime() bool {
	if o != nil && o.ReportitemCreateTime != nil {
		return true
	}

	return false
}

// SetReportitemCreateTime gets a reference to the given time.Time and assigns it to the ReportitemCreateTime field.
func (o *Reportitem) SetReportitemCreateTime(v time.Time) {
	o.ReportitemCreateTime = &v
}

// GetReportitemCreatedByUserId returns the ReportitemCreatedByUserId field value
func (o *Reportitem) GetReportitemCreatedByUserId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ReportitemCreatedByUserId
}

// GetReportitemCreatedByUserIdOk returns a tuple with the ReportitemCreatedByUserId field value
// and a boolean to check if the value has been set.
func (o *Reportitem) GetReportitemCreatedByUserIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReportitemCreatedByUserId, true
}

// SetReportitemCreatedByUserId sets field value
func (o *Reportitem) SetReportitemCreatedByUserId(v int32) {
	o.ReportitemCreatedByUserId = v
}

// GetReportitemModifyTime returns the ReportitemModifyTime field value if set, zero value otherwise.
func (o *Reportitem) GetReportitemModifyTime() time.Time {
	if o == nil || o.ReportitemModifyTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ReportitemModifyTime
}

// GetReportitemModifyTimeOk returns a tuple with the ReportitemModifyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reportitem) GetReportitemModifyTimeOk() (*time.Time, bool) {
	if o == nil || o.ReportitemModifyTime == nil {
		return nil, false
	}
	return o.ReportitemModifyTime, true
}

// HasReportitemModifyTime returns a boolean if a field has been set.
func (o *Reportitem) HasReportitemModifyTime() bool {
	if o != nil && o.ReportitemModifyTime != nil {
		return true
	}

	return false
}

// SetReportitemModifyTime gets a reference to the given time.Time and assigns it to the ReportitemModifyTime field.
func (o *Reportitem) SetReportitemModifyTime(v time.Time) {
	o.ReportitemModifyTime = &v
}

// GetReportitemModifiedByUserId returns the ReportitemModifiedByUserId field value
func (o *Reportitem) GetReportitemModifiedByUserId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ReportitemModifiedByUserId
}

// GetReportitemModifiedByUserIdOk returns a tuple with the ReportitemModifiedByUserId field value
// and a boolean to check if the value has been set.
func (o *Reportitem) GetReportitemModifiedByUserIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReportitemModifiedByUserId, true
}

// SetReportitemModifiedByUserId sets field value
func (o *Reportitem) SetReportitemModifiedByUserId(v int32) {
	o.ReportitemModifiedByUserId = v
}

func (o Reportitem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReportitemId != nil {
		toSerialize["reportitem_id"] = o.ReportitemId
	}
	if o.Case.IsSet() {
		toSerialize["case"] = o.Case.Get()
	}
	if true {
		toSerialize["headline"] = o.Headline
	}
	if o.Notestatus != nil {
		toSerialize["notestatus"] = o.Notestatus
	}
	if true {
		toSerialize["system"] = o.System
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["reportitem_subheadline"] = o.ReportitemSubheadline.Get()
	}
	if true {
		toSerialize["reportitem_note"] = o.ReportitemNote
	}
	if o.ReportitemCreateTime != nil {
		toSerialize["reportitem_create_time"] = o.ReportitemCreateTime
	}
	if true {
		toSerialize["reportitem_created_by_user_id"] = o.ReportitemCreatedByUserId
	}
	if o.ReportitemModifyTime != nil {
		toSerialize["reportitem_modify_time"] = o.ReportitemModifyTime
	}
	if true {
		toSerialize["reportitem_modified_by_user_id"] = o.ReportitemModifiedByUserId
	}
	return json.Marshal(toSerialize)
}

type NullableReportitem struct {
	value *Reportitem
	isSet bool
}

func (v NullableReportitem) Get() *Reportitem {
	return v.value
}

func (v *NullableReportitem) Set(val *Reportitem) {
	v.value = val
	v.isSet = true
}

func (v NullableReportitem) IsSet() bool {
	return v.isSet
}

func (v *NullableReportitem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportitem(val *Reportitem) *NullableReportitem {
	return &NullableReportitem{value: val, isSet: true}
}

func (v NullableReportitem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportitem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


