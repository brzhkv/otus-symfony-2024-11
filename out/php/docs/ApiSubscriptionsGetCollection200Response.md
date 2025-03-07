# ApiSubscriptionsGetCollection200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | [**[]SubscriptionJsonld**](SubscriptionJsonld.md) |  | 
**TotalItems** | Pointer to **int32** |  | [optional] 
**View** | Pointer to [**ApiPhoneUsersGetCollection200ResponseView**](ApiPhoneUsersGetCollection200ResponseView.md) |  | [optional] 
**Search** | Pointer to [**ApiPhoneUsersGetCollection200ResponseSearch**](ApiPhoneUsersGetCollection200ResponseSearch.md) |  | [optional] 

## Methods

### NewApiSubscriptionsGetCollection200Response

`func NewApiSubscriptionsGetCollection200Response(member []SubscriptionJsonld, ) *ApiSubscriptionsGetCollection200Response`

NewApiSubscriptionsGetCollection200Response instantiates a new ApiSubscriptionsGetCollection200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSubscriptionsGetCollection200ResponseWithDefaults

`func NewApiSubscriptionsGetCollection200ResponseWithDefaults() *ApiSubscriptionsGetCollection200Response`

NewApiSubscriptionsGetCollection200ResponseWithDefaults instantiates a new ApiSubscriptionsGetCollection200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *ApiSubscriptionsGetCollection200Response) GetMember() []SubscriptionJsonld`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ApiSubscriptionsGetCollection200Response) GetMemberOk() (*[]SubscriptionJsonld, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ApiSubscriptionsGetCollection200Response) SetMember(v []SubscriptionJsonld)`

SetMember sets Member field to given value.


### GetTotalItems

`func (o *ApiSubscriptionsGetCollection200Response) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *ApiSubscriptionsGetCollection200Response) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *ApiSubscriptionsGetCollection200Response) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *ApiSubscriptionsGetCollection200Response) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetView

`func (o *ApiSubscriptionsGetCollection200Response) GetView() ApiPhoneUsersGetCollection200ResponseView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *ApiSubscriptionsGetCollection200Response) GetViewOk() (*ApiPhoneUsersGetCollection200ResponseView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *ApiSubscriptionsGetCollection200Response) SetView(v ApiPhoneUsersGetCollection200ResponseView)`

SetView sets View field to given value.

### HasView

`func (o *ApiSubscriptionsGetCollection200Response) HasView() bool`

HasView returns a boolean if a field has been set.

### GetSearch

`func (o *ApiSubscriptionsGetCollection200Response) GetSearch() ApiPhoneUsersGetCollection200ResponseSearch`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ApiSubscriptionsGetCollection200Response) GetSearchOk() (*ApiPhoneUsersGetCollection200ResponseSearch, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ApiSubscriptionsGetCollection200Response) SetSearch(v ApiPhoneUsersGetCollection200ResponseSearch)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ApiSubscriptionsGetCollection200Response) HasSearch() bool`

HasSearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


