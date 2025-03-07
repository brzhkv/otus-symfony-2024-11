/*
My App

This is an awesome app!

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiPhoneUsersGetCollection200ResponseSearchMappingInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPhoneUsersGetCollection200ResponseSearchMappingInner{}

// ApiPhoneUsersGetCollection200ResponseSearchMappingInner struct for ApiPhoneUsersGetCollection200ResponseSearchMappingInner
type ApiPhoneUsersGetCollection200ResponseSearchMappingInner struct {
	Type *string `json:"@type,omitempty"`
	Variable *string `json:"variable,omitempty"`
	Property interface{} `json:"property,omitempty"`
	Required *bool `json:"required,omitempty"`
}

// NewApiPhoneUsersGetCollection200ResponseSearchMappingInner instantiates a new ApiPhoneUsersGetCollection200ResponseSearchMappingInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPhoneUsersGetCollection200ResponseSearchMappingInner() *ApiPhoneUsersGetCollection200ResponseSearchMappingInner {
	this := ApiPhoneUsersGetCollection200ResponseSearchMappingInner{}
	return &this
}

// NewApiPhoneUsersGetCollection200ResponseSearchMappingInnerWithDefaults instantiates a new ApiPhoneUsersGetCollection200ResponseSearchMappingInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPhoneUsersGetCollection200ResponseSearchMappingInnerWithDefaults() *ApiPhoneUsersGetCollection200ResponseSearchMappingInner {
	this := ApiPhoneUsersGetCollection200ResponseSearchMappingInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) SetType(v string) {
	o.Type = &v
}

// GetVariable returns the Variable field value if set, zero value otherwise.
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) GetVariable() string {
	if o == nil || IsNil(o.Variable) {
		var ret string
		return ret
	}
	return *o.Variable
}

// GetVariableOk returns a tuple with the Variable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) GetVariableOk() (*string, bool) {
	if o == nil || IsNil(o.Variable) {
		return nil, false
	}
	return o.Variable, true
}

// HasVariable returns a boolean if a field has been set.
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) HasVariable() bool {
	if o != nil && !IsNil(o.Variable) {
		return true
	}

	return false
}

// SetVariable gets a reference to the given string and assigns it to the Variable field.
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) SetVariable(v string) {
	o.Variable = &v
}

// GetProperty returns the Property field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) GetProperty() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) GetPropertyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return &o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given interface{} and assigns it to the Property field.
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) SetProperty(v interface{}) {
	o.Property = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) SetRequired(v bool) {
	o.Required = &v
}

func (o ApiPhoneUsersGetCollection200ResponseSearchMappingInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPhoneUsersGetCollection200ResponseSearchMappingInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["@type"] = o.Type
	}
	if !IsNil(o.Variable) {
		toSerialize["variable"] = o.Variable
	}
	if o.Property != nil {
		toSerialize["property"] = o.Property
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	return toSerialize, nil
}

type NullableApiPhoneUsersGetCollection200ResponseSearchMappingInner struct {
	value *ApiPhoneUsersGetCollection200ResponseSearchMappingInner
	isSet bool
}

func (v NullableApiPhoneUsersGetCollection200ResponseSearchMappingInner) Get() *ApiPhoneUsersGetCollection200ResponseSearchMappingInner {
	return v.value
}

func (v *NullableApiPhoneUsersGetCollection200ResponseSearchMappingInner) Set(val *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPhoneUsersGetCollection200ResponseSearchMappingInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPhoneUsersGetCollection200ResponseSearchMappingInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPhoneUsersGetCollection200ResponseSearchMappingInner(val *ApiPhoneUsersGetCollection200ResponseSearchMappingInner) *NullableApiPhoneUsersGetCollection200ResponseSearchMappingInner {
	return &NullableApiPhoneUsersGetCollection200ResponseSearchMappingInner{value: val, isSet: true}
}

func (v NullableApiPhoneUsersGetCollection200ResponseSearchMappingInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPhoneUsersGetCollection200ResponseSearchMappingInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


