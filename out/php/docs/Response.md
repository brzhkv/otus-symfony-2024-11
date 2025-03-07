# Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tweets** | [**[]TweetDTO**](TweetDTO.md) | Твиты | 

## Methods

### NewResponse

`func NewResponse(tweets []TweetDTO, ) *Response`

NewResponse instantiates a new Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithDefaults

`func NewResponseWithDefaults() *Response`

NewResponseWithDefaults instantiates a new Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTweets

`func (o *Response) GetTweets() []TweetDTO`

GetTweets returns the Tweets field if non-nil, zero value otherwise.

### GetTweetsOk

`func (o *Response) GetTweetsOk() (*[]TweetDTO, bool)`

GetTweetsOk returns a tuple with the Tweets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweets

`func (o *Response) SetTweets(v []TweetDTO)`

SetTweets sets Tweets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


