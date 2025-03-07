# AddFollowersDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FollowerLoginPrefix** | **string** |  | 
**Count** | **int32** |  | 
**Async** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAddFollowersDTO

`func NewAddFollowersDTO(followerLoginPrefix string, count int32, ) *AddFollowersDTO`

NewAddFollowersDTO instantiates a new AddFollowersDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFollowersDTOWithDefaults

`func NewAddFollowersDTOWithDefaults() *AddFollowersDTO`

NewAddFollowersDTOWithDefaults instantiates a new AddFollowersDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFollowerLoginPrefix

`func (o *AddFollowersDTO) GetFollowerLoginPrefix() string`

GetFollowerLoginPrefix returns the FollowerLoginPrefix field if non-nil, zero value otherwise.

### GetFollowerLoginPrefixOk

`func (o *AddFollowersDTO) GetFollowerLoginPrefixOk() (*string, bool)`

GetFollowerLoginPrefixOk returns a tuple with the FollowerLoginPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerLoginPrefix

`func (o *AddFollowersDTO) SetFollowerLoginPrefix(v string)`

SetFollowerLoginPrefix sets FollowerLoginPrefix field to given value.


### GetCount

`func (o *AddFollowersDTO) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AddFollowersDTO) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AddFollowersDTO) SetCount(v int32)`

SetCount sets Count field to given value.


### GetAsync

`func (o *AddFollowersDTO) GetAsync() bool`

GetAsync returns the Async field if non-nil, zero value otherwise.

### GetAsyncOk

`func (o *AddFollowersDTO) GetAsyncOk() (*bool, bool)`

GetAsyncOk returns a tuple with the Async field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsync

`func (o *AddFollowersDTO) SetAsync(v bool)`

SetAsync sets Async field to given value.

### HasAsync

`func (o *AddFollowersDTO) HasAsync() bool`

HasAsync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


