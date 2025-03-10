/*
My App

This is an awesome app!

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the TweetJsonld type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TweetJsonld{}

// TweetJsonld 
type TweetJsonld struct {
	Context *PhoneUserJsonldContext `json:"@context,omitempty"`
	Id *string `json:"@id,omitempty"`
	Type *string `json:"@type,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Author interface{} `json:"author,omitempty"`
	Text *string `json:"text,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewTweetJsonld instantiates a new TweetJsonld object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTweetJsonld() *TweetJsonld {
	this := TweetJsonld{}
	return &this
}

// NewTweetJsonldWithDefaults instantiates a new TweetJsonld object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTweetJsonldWithDefaults() *TweetJsonld {
	this := TweetJsonld{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TweetJsonld) GetContext() PhoneUserJsonldContext {
	if o == nil || IsNil(o.Context) {
		var ret PhoneUserJsonldContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetJsonld) GetContextOk() (*PhoneUserJsonldContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TweetJsonld) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given PhoneUserJsonldContext and assigns it to the Context field.
func (o *TweetJsonld) SetContext(v PhoneUserJsonldContext) {
	o.Context = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TweetJsonld) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetJsonld) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TweetJsonld) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TweetJsonld) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TweetJsonld) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetJsonld) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TweetJsonld) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TweetJsonld) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TweetJsonld) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetJsonld) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TweetJsonld) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TweetJsonld) SetId(v int32) {
	o.Id = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TweetJsonld) GetAuthor() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TweetJsonld) GetAuthorOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return &o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *TweetJsonld) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given interface{} and assigns it to the Author field.
func (o *TweetJsonld) SetAuthor(v interface{}) {
	o.Author = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *TweetJsonld) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetJsonld) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *TweetJsonld) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *TweetJsonld) SetText(v string) {
	o.Text = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TweetJsonld) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetJsonld) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TweetJsonld) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TweetJsonld) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TweetJsonld) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TweetJsonld) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TweetJsonld) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TweetJsonld) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o TweetJsonld) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TweetJsonld) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["@context"] = o.Context
	}
	if !IsNil(o.Id) {
		toSerialize["@id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["@type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableTweetJsonld struct {
	value *TweetJsonld
	isSet bool
}

func (v NullableTweetJsonld) Get() *TweetJsonld {
	return v.value
}

func (v *NullableTweetJsonld) Set(val *TweetJsonld) {
	v.value = val
	v.isSet = true
}

func (v NullableTweetJsonld) IsSet() bool {
	return v.isSet
}

func (v *NullableTweetJsonld) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTweetJsonld(val *TweetJsonld) *NullableTweetJsonld {
	return &NullableTweetJsonld{value: val, isSet: true}
}

func (v NullableTweetJsonld) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTweetJsonld) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


