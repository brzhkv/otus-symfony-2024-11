# SubscriptionJsonld

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**PhoneUserJsonldContext**](PhoneUserJsonldContext.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Author** | Pointer to **interface{}** |  | [optional] 
**Follower** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewSubscriptionJsonld

`func NewSubscriptionJsonld() *SubscriptionJsonld`

NewSubscriptionJsonld instantiates a new SubscriptionJsonld object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionJsonldWithDefaults

`func NewSubscriptionJsonldWithDefaults() *SubscriptionJsonld`

NewSubscriptionJsonldWithDefaults instantiates a new SubscriptionJsonld object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SubscriptionJsonld) GetContext() PhoneUserJsonldContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SubscriptionJsonld) GetContextOk() (*PhoneUserJsonldContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SubscriptionJsonld) SetContext(v PhoneUserJsonldContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *SubscriptionJsonld) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *SubscriptionJsonld) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionJsonld) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionJsonld) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionJsonld) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SubscriptionJsonld) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionJsonld) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionJsonld) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubscriptionJsonld) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *SubscriptionJsonld) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionJsonld) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionJsonld) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionJsonld) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthor

`func (o *SubscriptionJsonld) GetAuthor() interface{}`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *SubscriptionJsonld) GetAuthorOk() (*interface{}, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *SubscriptionJsonld) SetAuthor(v interface{})`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *SubscriptionJsonld) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *SubscriptionJsonld) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *SubscriptionJsonld) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetFollower

`func (o *SubscriptionJsonld) GetFollower() interface{}`

GetFollower returns the Follower field if non-nil, zero value otherwise.

### GetFollowerOk

`func (o *SubscriptionJsonld) GetFollowerOk() (*interface{}, bool)`

GetFollowerOk returns a tuple with the Follower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollower

`func (o *SubscriptionJsonld) SetFollower(v interface{})`

SetFollower sets Follower field to given value.

### HasFollower

`func (o *SubscriptionJsonld) HasFollower() bool`

HasFollower returns a boolean if a field has been set.

### SetFollowerNil

`func (o *SubscriptionJsonld) SetFollowerNil(b bool)`

 SetFollowerNil sets the value for Follower to be an explicit nil

### UnsetFollower
`func (o *SubscriptionJsonld) UnsetFollower()`

UnsetFollower ensures that no value is present for Follower, not even an explicit nil
### GetCreatedAt

`func (o *SubscriptionJsonld) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionJsonld) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionJsonld) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionJsonld) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubscriptionJsonld) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionJsonld) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionJsonld) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubscriptionJsonld) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


