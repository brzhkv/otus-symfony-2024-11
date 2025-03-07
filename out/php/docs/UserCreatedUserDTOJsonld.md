# UserCreatedUserDTOJsonld

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**PhoneUserJsonldContext**](PhoneUserJsonldContext.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **int32** |  | [optional] 
**Login** | Pointer to **string** |  | [optional] 
**AvatarLink** | Pointer to **interface{}** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **interface{}** |  | [optional] 
**Email** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewUserCreatedUserDTOJsonld

`func NewUserCreatedUserDTOJsonld() *UserCreatedUserDTOJsonld`

NewUserCreatedUserDTOJsonld instantiates a new UserCreatedUserDTOJsonld object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreatedUserDTOJsonldWithDefaults

`func NewUserCreatedUserDTOJsonldWithDefaults() *UserCreatedUserDTOJsonld`

NewUserCreatedUserDTOJsonldWithDefaults instantiates a new UserCreatedUserDTOJsonld object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *UserCreatedUserDTOJsonld) GetContext() PhoneUserJsonldContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UserCreatedUserDTOJsonld) GetContextOk() (*PhoneUserJsonldContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UserCreatedUserDTOJsonld) SetContext(v PhoneUserJsonldContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *UserCreatedUserDTOJsonld) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *UserCreatedUserDTOJsonld) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserCreatedUserDTOJsonld) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserCreatedUserDTOJsonld) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserCreatedUserDTOJsonld) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserCreatedUserDTOJsonld) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserCreatedUserDTOJsonld) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserCreatedUserDTOJsonld) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserCreatedUserDTOJsonld) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *UserCreatedUserDTOJsonld) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserCreatedUserDTOJsonld) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserCreatedUserDTOJsonld) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserCreatedUserDTOJsonld) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogin

`func (o *UserCreatedUserDTOJsonld) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *UserCreatedUserDTOJsonld) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *UserCreatedUserDTOJsonld) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *UserCreatedUserDTOJsonld) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetAvatarLink

`func (o *UserCreatedUserDTOJsonld) GetAvatarLink() interface{}`

GetAvatarLink returns the AvatarLink field if non-nil, zero value otherwise.

### GetAvatarLinkOk

`func (o *UserCreatedUserDTOJsonld) GetAvatarLinkOk() (*interface{}, bool)`

GetAvatarLinkOk returns a tuple with the AvatarLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarLink

`func (o *UserCreatedUserDTOJsonld) SetAvatarLink(v interface{})`

SetAvatarLink sets AvatarLink field to given value.

### HasAvatarLink

`func (o *UserCreatedUserDTOJsonld) HasAvatarLink() bool`

HasAvatarLink returns a boolean if a field has been set.

### SetAvatarLinkNil

`func (o *UserCreatedUserDTOJsonld) SetAvatarLinkNil(b bool)`

 SetAvatarLinkNil sets the value for AvatarLink to be an explicit nil

### UnsetAvatarLink
`func (o *UserCreatedUserDTOJsonld) UnsetAvatarLink()`

UnsetAvatarLink ensures that no value is present for AvatarLink, not even an explicit nil
### GetRoles

`func (o *UserCreatedUserDTOJsonld) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserCreatedUserDTOJsonld) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserCreatedUserDTOJsonld) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserCreatedUserDTOJsonld) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserCreatedUserDTOJsonld) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserCreatedUserDTOJsonld) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserCreatedUserDTOJsonld) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserCreatedUserDTOJsonld) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserCreatedUserDTOJsonld) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserCreatedUserDTOJsonld) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserCreatedUserDTOJsonld) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserCreatedUserDTOJsonld) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPhone

`func (o *UserCreatedUserDTOJsonld) GetPhone() interface{}`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserCreatedUserDTOJsonld) GetPhoneOk() (*interface{}, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserCreatedUserDTOJsonld) SetPhone(v interface{})`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserCreatedUserDTOJsonld) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *UserCreatedUserDTOJsonld) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *UserCreatedUserDTOJsonld) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *UserCreatedUserDTOJsonld) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreatedUserDTOJsonld) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreatedUserDTOJsonld) SetEmail(v interface{})`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserCreatedUserDTOJsonld) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserCreatedUserDTOJsonld) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserCreatedUserDTOJsonld) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


