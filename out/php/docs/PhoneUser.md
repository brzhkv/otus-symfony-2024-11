# PhoneUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** |  | [optional] 
**Id** | Pointer to [**PhoneUserId**](PhoneUserId.md) |  | [optional] 
**Login** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Tweets** | Pointer to [**[]Tweet**](Tweet.md) |  | [optional] 
**Authors** | Pointer to **[]string** |  | [optional] 
**Followers** | Pointer to **[]string** |  | [optional] 
**SubscriptionAuthors** | Pointer to **[]string** |  | [optional] 
**SubscriptionFollowers** | Pointer to **[]string** |  | [optional] 
**DeletedAt** | Pointer to **interface{}** |  | [optional] [readonly] 
**AvatarLink** | Pointer to **interface{}** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Age** | Pointer to **int32** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Token** | Pointer to **interface{}** |  | [optional] 
**IsProtected** | Pointer to **interface{}** |  | [optional] 
**DeletedAtInFuture** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] [readonly] 
**UserIdentifier** | Pointer to **string** |  | [optional] [readonly] 
**Protected** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewPhoneUser

`func NewPhoneUser() *PhoneUser`

NewPhoneUser instantiates a new PhoneUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneUserWithDefaults

`func NewPhoneUserWithDefaults() *PhoneUser`

NewPhoneUserWithDefaults instantiates a new PhoneUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *PhoneUser) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PhoneUser) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PhoneUser) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PhoneUser) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetId

`func (o *PhoneUser) GetId() PhoneUserId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhoneUser) GetIdOk() (*PhoneUserId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhoneUser) SetId(v PhoneUserId)`

SetId sets Id field to given value.

### HasId

`func (o *PhoneUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogin

`func (o *PhoneUser) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *PhoneUser) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *PhoneUser) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *PhoneUser) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PhoneUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PhoneUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PhoneUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PhoneUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PhoneUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PhoneUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PhoneUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PhoneUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTweets

`func (o *PhoneUser) GetTweets() []Tweet`

GetTweets returns the Tweets field if non-nil, zero value otherwise.

### GetTweetsOk

`func (o *PhoneUser) GetTweetsOk() (*[]Tweet, bool)`

GetTweetsOk returns a tuple with the Tweets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweets

`func (o *PhoneUser) SetTweets(v []Tweet)`

SetTweets sets Tweets field to given value.

### HasTweets

`func (o *PhoneUser) HasTweets() bool`

HasTweets returns a boolean if a field has been set.

### GetAuthors

`func (o *PhoneUser) GetAuthors() []string`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *PhoneUser) GetAuthorsOk() (*[]string, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *PhoneUser) SetAuthors(v []string)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *PhoneUser) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetFollowers

`func (o *PhoneUser) GetFollowers() []string`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *PhoneUser) GetFollowersOk() (*[]string, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *PhoneUser) SetFollowers(v []string)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *PhoneUser) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetSubscriptionAuthors

`func (o *PhoneUser) GetSubscriptionAuthors() []string`

GetSubscriptionAuthors returns the SubscriptionAuthors field if non-nil, zero value otherwise.

### GetSubscriptionAuthorsOk

`func (o *PhoneUser) GetSubscriptionAuthorsOk() (*[]string, bool)`

GetSubscriptionAuthorsOk returns a tuple with the SubscriptionAuthors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAuthors

`func (o *PhoneUser) SetSubscriptionAuthors(v []string)`

SetSubscriptionAuthors sets SubscriptionAuthors field to given value.

### HasSubscriptionAuthors

`func (o *PhoneUser) HasSubscriptionAuthors() bool`

HasSubscriptionAuthors returns a boolean if a field has been set.

### GetSubscriptionFollowers

`func (o *PhoneUser) GetSubscriptionFollowers() []string`

GetSubscriptionFollowers returns the SubscriptionFollowers field if non-nil, zero value otherwise.

### GetSubscriptionFollowersOk

`func (o *PhoneUser) GetSubscriptionFollowersOk() (*[]string, bool)`

GetSubscriptionFollowersOk returns a tuple with the SubscriptionFollowers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionFollowers

`func (o *PhoneUser) SetSubscriptionFollowers(v []string)`

SetSubscriptionFollowers sets SubscriptionFollowers field to given value.

### HasSubscriptionFollowers

`func (o *PhoneUser) HasSubscriptionFollowers() bool`

HasSubscriptionFollowers returns a boolean if a field has been set.

### GetDeletedAt

`func (o *PhoneUser) GetDeletedAt() interface{}`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PhoneUser) GetDeletedAtOk() (*interface{}, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PhoneUser) SetDeletedAt(v interface{})`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *PhoneUser) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *PhoneUser) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *PhoneUser) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetAvatarLink

`func (o *PhoneUser) GetAvatarLink() interface{}`

GetAvatarLink returns the AvatarLink field if non-nil, zero value otherwise.

### GetAvatarLinkOk

`func (o *PhoneUser) GetAvatarLinkOk() (*interface{}, bool)`

GetAvatarLinkOk returns a tuple with the AvatarLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarLink

`func (o *PhoneUser) SetAvatarLink(v interface{})`

SetAvatarLink sets AvatarLink field to given value.

### HasAvatarLink

`func (o *PhoneUser) HasAvatarLink() bool`

HasAvatarLink returns a boolean if a field has been set.

### SetAvatarLinkNil

`func (o *PhoneUser) SetAvatarLinkNil(b bool)`

 SetAvatarLinkNil sets the value for AvatarLink to be an explicit nil

### UnsetAvatarLink
`func (o *PhoneUser) UnsetAvatarLink()`

UnsetAvatarLink ensures that no value is present for AvatarLink, not even an explicit nil
### GetPassword

`func (o *PhoneUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PhoneUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PhoneUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PhoneUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAge

`func (o *PhoneUser) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *PhoneUser) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *PhoneUser) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *PhoneUser) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetIsActive

`func (o *PhoneUser) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PhoneUser) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PhoneUser) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PhoneUser) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetRoles

`func (o *PhoneUser) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PhoneUser) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PhoneUser) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PhoneUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetToken

`func (o *PhoneUser) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PhoneUser) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PhoneUser) SetToken(v interface{})`

SetToken sets Token field to given value.

### HasToken

`func (o *PhoneUser) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *PhoneUser) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *PhoneUser) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetIsProtected

`func (o *PhoneUser) GetIsProtected() interface{}`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *PhoneUser) GetIsProtectedOk() (*interface{}, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *PhoneUser) SetIsProtected(v interface{})`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *PhoneUser) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### SetIsProtectedNil

`func (o *PhoneUser) SetIsProtectedNil(b bool)`

 SetIsProtectedNil sets the value for IsProtected to be an explicit nil

### UnsetIsProtected
`func (o *PhoneUser) UnsetIsProtected()`

UnsetIsProtected ensures that no value is present for IsProtected, not even an explicit nil
### GetDeletedAtInFuture

`func (o *PhoneUser) GetDeletedAtInFuture() string`

GetDeletedAtInFuture returns the DeletedAtInFuture field if non-nil, zero value otherwise.

### GetDeletedAtInFutureOk

`func (o *PhoneUser) GetDeletedAtInFutureOk() (*string, bool)`

GetDeletedAtInFutureOk returns a tuple with the DeletedAtInFuture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAtInFuture

`func (o *PhoneUser) SetDeletedAtInFuture(v string)`

SetDeletedAtInFuture sets DeletedAtInFuture field to given value.

### HasDeletedAtInFuture

`func (o *PhoneUser) HasDeletedAtInFuture() bool`

HasDeletedAtInFuture returns a boolean if a field has been set.

### GetActive

`func (o *PhoneUser) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PhoneUser) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PhoneUser) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PhoneUser) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUserIdentifier

`func (o *PhoneUser) GetUserIdentifier() string`

GetUserIdentifier returns the UserIdentifier field if non-nil, zero value otherwise.

### GetUserIdentifierOk

`func (o *PhoneUser) GetUserIdentifierOk() (*string, bool)`

GetUserIdentifierOk returns a tuple with the UserIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentifier

`func (o *PhoneUser) SetUserIdentifier(v string)`

SetUserIdentifier sets UserIdentifier field to given value.

### HasUserIdentifier

`func (o *PhoneUser) HasUserIdentifier() bool`

HasUserIdentifier returns a boolean if a field has been set.

### GetProtected

`func (o *PhoneUser) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *PhoneUser) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *PhoneUser) SetProtected(v bool)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *PhoneUser) HasProtected() bool`

HasProtected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


