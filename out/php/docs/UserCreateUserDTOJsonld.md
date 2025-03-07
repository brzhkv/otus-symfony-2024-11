# UserCreateUserDTOJsonld

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** |  | 
**Email** | Pointer to **interface{}** |  | [optional] 
**Phone** | Pointer to **interface{}** |  | [optional] 
**Password** | **string** |  | 
**Age** | **int32** |  | 
**IsActive** | **bool** |  | 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUserCreateUserDTOJsonld

`func NewUserCreateUserDTOJsonld(login string, password string, age int32, isActive bool, ) *UserCreateUserDTOJsonld`

NewUserCreateUserDTOJsonld instantiates a new UserCreateUserDTOJsonld object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateUserDTOJsonldWithDefaults

`func NewUserCreateUserDTOJsonldWithDefaults() *UserCreateUserDTOJsonld`

NewUserCreateUserDTOJsonldWithDefaults instantiates a new UserCreateUserDTOJsonld object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *UserCreateUserDTOJsonld) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *UserCreateUserDTOJsonld) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *UserCreateUserDTOJsonld) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetEmail

`func (o *UserCreateUserDTOJsonld) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreateUserDTOJsonld) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreateUserDTOJsonld) SetEmail(v interface{})`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserCreateUserDTOJsonld) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserCreateUserDTOJsonld) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserCreateUserDTOJsonld) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *UserCreateUserDTOJsonld) GetPhone() interface{}`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserCreateUserDTOJsonld) GetPhoneOk() (*interface{}, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserCreateUserDTOJsonld) SetPhone(v interface{})`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserCreateUserDTOJsonld) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *UserCreateUserDTOJsonld) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *UserCreateUserDTOJsonld) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetPassword

`func (o *UserCreateUserDTOJsonld) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCreateUserDTOJsonld) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCreateUserDTOJsonld) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAge

`func (o *UserCreateUserDTOJsonld) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *UserCreateUserDTOJsonld) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *UserCreateUserDTOJsonld) SetAge(v int32)`

SetAge sets Age field to given value.


### GetIsActive

`func (o *UserCreateUserDTOJsonld) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UserCreateUserDTOJsonld) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UserCreateUserDTOJsonld) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetRoles

`func (o *UserCreateUserDTOJsonld) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserCreateUserDTOJsonld) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserCreateUserDTOJsonld) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserCreateUserDTOJsonld) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


