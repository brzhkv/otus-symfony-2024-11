# PostTweetDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** |  | 
**Text** | **string** |  | 
**Async** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPostTweetDTO

`func NewPostTweetDTO(userId int32, text string, ) *PostTweetDTO`

NewPostTweetDTO instantiates a new PostTweetDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTweetDTOWithDefaults

`func NewPostTweetDTOWithDefaults() *PostTweetDTO`

NewPostTweetDTOWithDefaults instantiates a new PostTweetDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PostTweetDTO) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PostTweetDTO) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PostTweetDTO) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetText

`func (o *PostTweetDTO) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostTweetDTO) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostTweetDTO) SetText(v string)`

SetText sets Text field to given value.


### GetAsync

`func (o *PostTweetDTO) GetAsync() bool`

GetAsync returns the Async field if non-nil, zero value otherwise.

### GetAsyncOk

`func (o *PostTweetDTO) GetAsyncOk() (*bool, bool)`

GetAsyncOk returns a tuple with the Async field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsync

`func (o *PostTweetDTO) SetAsync(v bool)`

SetAsync sets Async field to given value.

### HasAsync

`func (o *PostTweetDTO) HasAsync() bool`

HasAsync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


