/*
My App

This is an awesome app!

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateUserDTO2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserDTO2{}

// CreateUserDTO2 struct for CreateUserDTO2
type CreateUserDTO2 struct {
	Login string `json:"login"`
	Email NullableString `json:"email,omitempty"`
	Phone NullableString `json:"phone,omitempty"`
	Password string `json:"password"`
	Age int32 `json:"age"`
	IsActive bool `json:"isActive"`
	Roles []string `json:"roles"`
}

type _CreateUserDTO2 CreateUserDTO2

// NewCreateUserDTO2 instantiates a new CreateUserDTO2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserDTO2(login string, password string, age int32, isActive bool, roles []string) *CreateUserDTO2 {
	this := CreateUserDTO2{}
	this.Login = login
	this.Password = password
	this.Age = age
	this.IsActive = isActive
	this.Roles = roles
	return &this
}

// NewCreateUserDTO2WithDefaults instantiates a new CreateUserDTO2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserDTO2WithDefaults() *CreateUserDTO2 {
	this := CreateUserDTO2{}
	return &this
}

// GetLogin returns the Login field value
func (o *CreateUserDTO2) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *CreateUserDTO2) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *CreateUserDTO2) SetLogin(v string) {
	o.Login = v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserDTO2) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserDTO2) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *CreateUserDTO2) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *CreateUserDTO2) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *CreateUserDTO2) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *CreateUserDTO2) UnsetEmail() {
	o.Email.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserDTO2) GetPhone() string {
	if o == nil || IsNil(o.Phone.Get()) {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserDTO2) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *CreateUserDTO2) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *CreateUserDTO2) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *CreateUserDTO2) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *CreateUserDTO2) UnsetPhone() {
	o.Phone.Unset()
}

// GetPassword returns the Password field value
func (o *CreateUserDTO2) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreateUserDTO2) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreateUserDTO2) SetPassword(v string) {
	o.Password = v
}

// GetAge returns the Age field value
func (o *CreateUserDTO2) GetAge() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Age
}

// GetAgeOk returns a tuple with the Age field value
// and a boolean to check if the value has been set.
func (o *CreateUserDTO2) GetAgeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Age, true
}

// SetAge sets field value
func (o *CreateUserDTO2) SetAge(v int32) {
	o.Age = v
}

// GetIsActive returns the IsActive field value
func (o *CreateUserDTO2) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *CreateUserDTO2) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *CreateUserDTO2) SetIsActive(v bool) {
	o.IsActive = v
}

// GetRoles returns the Roles field value
func (o *CreateUserDTO2) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *CreateUserDTO2) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *CreateUserDTO2) SetRoles(v []string) {
	o.Roles = v
}

func (o CreateUserDTO2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserDTO2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["login"] = o.Login
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	toSerialize["password"] = o.Password
	toSerialize["age"] = o.Age
	toSerialize["isActive"] = o.IsActive
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

func (o *CreateUserDTO2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"login",
		"password",
		"age",
		"isActive",
		"roles",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateUserDTO2 := _CreateUserDTO2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateUserDTO2)

	if err != nil {
		return err
	}

	*o = CreateUserDTO2(varCreateUserDTO2)

	return err
}

type NullableCreateUserDTO2 struct {
	value *CreateUserDTO2
	isSet bool
}

func (v NullableCreateUserDTO2) Get() *CreateUserDTO2 {
	return v.value
}

func (v *NullableCreateUserDTO2) Set(val *CreateUserDTO2) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserDTO2) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserDTO2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserDTO2(val *CreateUserDTO2) *NullableCreateUserDTO2 {
	return &NullableCreateUserDTO2{value: val, isSet: true}
}

func (v NullableCreateUserDTO2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserDTO2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


