# CreateUserDTO2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** |  | 
**Email** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Password** | **string** |  | 
**Age** | **int32** |  | 
**IsActive** | **bool** |  | 
**Roles** | **[]string** |  | 

## Methods

### NewCreateUserDTO2

`func NewCreateUserDTO2(login string, password string, age int32, isActive bool, roles []string, ) *CreateUserDTO2`

NewCreateUserDTO2 instantiates a new CreateUserDTO2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserDTO2WithDefaults

`func NewCreateUserDTO2WithDefaults() *CreateUserDTO2`

NewCreateUserDTO2WithDefaults instantiates a new CreateUserDTO2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *CreateUserDTO2) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *CreateUserDTO2) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *CreateUserDTO2) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetEmail

`func (o *CreateUserDTO2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserDTO2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserDTO2) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateUserDTO2) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CreateUserDTO2) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CreateUserDTO2) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *CreateUserDTO2) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CreateUserDTO2) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CreateUserDTO2) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CreateUserDTO2) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *CreateUserDTO2) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *CreateUserDTO2) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetPassword

`func (o *CreateUserDTO2) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUserDTO2) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUserDTO2) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAge

`func (o *CreateUserDTO2) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *CreateUserDTO2) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *CreateUserDTO2) SetAge(v int32)`

SetAge sets Age field to given value.


### GetIsActive

`func (o *CreateUserDTO2) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateUserDTO2) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateUserDTO2) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetRoles

`func (o *CreateUserDTO2) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateUserDTO2) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateUserDTO2) SetRoles(v []string)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


