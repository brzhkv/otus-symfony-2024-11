# TweetJsonld

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**PhoneUserJsonldContext**](PhoneUserJsonldContext.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Author** | Pointer to **interface{}** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewTweetJsonld

`func NewTweetJsonld() *TweetJsonld`

NewTweetJsonld instantiates a new TweetJsonld object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTweetJsonldWithDefaults

`func NewTweetJsonldWithDefaults() *TweetJsonld`

NewTweetJsonldWithDefaults instantiates a new TweetJsonld object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *TweetJsonld) GetContext() PhoneUserJsonldContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TweetJsonld) GetContextOk() (*PhoneUserJsonldContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TweetJsonld) SetContext(v PhoneUserJsonldContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *TweetJsonld) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *TweetJsonld) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TweetJsonld) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TweetJsonld) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TweetJsonld) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TweetJsonld) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TweetJsonld) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TweetJsonld) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TweetJsonld) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *TweetJsonld) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TweetJsonld) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TweetJsonld) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TweetJsonld) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthor

`func (o *TweetJsonld) GetAuthor() interface{}`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *TweetJsonld) GetAuthorOk() (*interface{}, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *TweetJsonld) SetAuthor(v interface{})`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *TweetJsonld) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *TweetJsonld) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *TweetJsonld) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetText

`func (o *TweetJsonld) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TweetJsonld) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TweetJsonld) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TweetJsonld) HasText() bool`

HasText returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TweetJsonld) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TweetJsonld) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TweetJsonld) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TweetJsonld) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TweetJsonld) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TweetJsonld) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TweetJsonld) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TweetJsonld) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


