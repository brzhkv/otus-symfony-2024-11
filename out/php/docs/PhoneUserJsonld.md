# PhoneUserJsonld

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**PhoneUserJsonldContext**](PhoneUserJsonldContext.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Phone** | Pointer to **string** |  | [optional] 
**Id** | Pointer to [**PhoneUserId**](PhoneUserId.md) |  | [optional] 
**Login** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Tweets** | Pointer to [**[]TweetJsonld**](TweetJsonld.md) |  | [optional] 
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

### NewPhoneUserJsonld

`func NewPhoneUserJsonld() *PhoneUserJsonld`

NewPhoneUserJsonld instantiates a new PhoneUserJsonld object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneUserJsonldWithDefaults

`func NewPhoneUserJsonldWithDefaults() *PhoneUserJsonld`

NewPhoneUserJsonldWithDefaults instantiates a new PhoneUserJsonld object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *PhoneUserJsonld) GetContext() PhoneUserJsonldContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PhoneUserJsonld) GetContextOk() (*PhoneUserJsonldContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PhoneUserJsonld) SetContext(v PhoneUserJsonldContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *PhoneUserJsonld) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *PhoneUserJsonld) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhoneUserJsonld) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhoneUserJsonld) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PhoneUserJsonld) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PhoneUserJsonld) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PhoneUserJsonld) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PhoneUserJsonld) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PhoneUserJsonld) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPhone

`func (o *PhoneUserJsonld) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PhoneUserJsonld) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PhoneUserJsonld) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PhoneUserJsonld) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetId

`func (o *PhoneUserJsonld) GetId() PhoneUserId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhoneUserJsonld) GetIdOk() (*PhoneUserId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhoneUserJsonld) SetId(v PhoneUserId)`

SetId sets Id field to given value.

### HasId

`func (o *PhoneUserJsonld) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogin

`func (o *PhoneUserJsonld) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *PhoneUserJsonld) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *PhoneUserJsonld) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *PhoneUserJsonld) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PhoneUserJsonld) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PhoneUserJsonld) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PhoneUserJsonld) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PhoneUserJsonld) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PhoneUserJsonld) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PhoneUserJsonld) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PhoneUserJsonld) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PhoneUserJsonld) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTweets

`func (o *PhoneUserJsonld) GetTweets() []TweetJsonld`

GetTweets returns the Tweets field if non-nil, zero value otherwise.

### GetTweetsOk

`func (o *PhoneUserJsonld) GetTweetsOk() (*[]TweetJsonld, bool)`

GetTweetsOk returns a tuple with the Tweets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweets

`func (o *PhoneUserJsonld) SetTweets(v []TweetJsonld)`

SetTweets sets Tweets field to given value.

### HasTweets

`func (o *PhoneUserJsonld) HasTweets() bool`

HasTweets returns a boolean if a field has been set.

### GetAuthors

`func (o *PhoneUserJsonld) GetAuthors() []string`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *PhoneUserJsonld) GetAuthorsOk() (*[]string, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *PhoneUserJsonld) SetAuthors(v []string)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *PhoneUserJsonld) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetFollowers

`func (o *PhoneUserJsonld) GetFollowers() []string`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *PhoneUserJsonld) GetFollowersOk() (*[]string, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *PhoneUserJsonld) SetFollowers(v []string)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *PhoneUserJsonld) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetSubscriptionAuthors

`func (o *PhoneUserJsonld) GetSubscriptionAuthors() []string`

GetSubscriptionAuthors returns the SubscriptionAuthors field if non-nil, zero value otherwise.

### GetSubscriptionAuthorsOk

`func (o *PhoneUserJsonld) GetSubscriptionAuthorsOk() (*[]string, bool)`

GetSubscriptionAuthorsOk returns a tuple with the SubscriptionAuthors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAuthors

`func (o *PhoneUserJsonld) SetSubscriptionAuthors(v []string)`

SetSubscriptionAuthors sets SubscriptionAuthors field to given value.

### HasSubscriptionAuthors

`func (o *PhoneUserJsonld) HasSubscriptionAuthors() bool`

HasSubscriptionAuthors returns a boolean if a field has been set.

### GetSubscriptionFollowers

`func (o *PhoneUserJsonld) GetSubscriptionFollowers() []string`

GetSubscriptionFollowers returns the SubscriptionFollowers field if non-nil, zero value otherwise.

### GetSubscriptionFollowersOk

`func (o *PhoneUserJsonld) GetSubscriptionFollowersOk() (*[]string, bool)`

GetSubscriptionFollowersOk returns a tuple with the SubscriptionFollowers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionFollowers

`func (o *PhoneUserJsonld) SetSubscriptionFollowers(v []string)`

SetSubscriptionFollowers sets SubscriptionFollowers field to given value.

### HasSubscriptionFollowers

`func (o *PhoneUserJsonld) HasSubscriptionFollowers() bool`

HasSubscriptionFollowers returns a boolean if a field has been set.

### GetDeletedAt

`func (o *PhoneUserJsonld) GetDeletedAt() interface{}`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PhoneUserJsonld) GetDeletedAtOk() (*interface{}, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PhoneUserJsonld) SetDeletedAt(v interface{})`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *PhoneUserJsonld) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *PhoneUserJsonld) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *PhoneUserJsonld) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetAvatarLink

`func (o *PhoneUserJsonld) GetAvatarLink() interface{}`

GetAvatarLink returns the AvatarLink field if non-nil, zero value otherwise.

### GetAvatarLinkOk

`func (o *PhoneUserJsonld) GetAvatarLinkOk() (*interface{}, bool)`

GetAvatarLinkOk returns a tuple with the AvatarLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarLink

`func (o *PhoneUserJsonld) SetAvatarLink(v interface{})`

SetAvatarLink sets AvatarLink field to given value.

### HasAvatarLink

`func (o *PhoneUserJsonld) HasAvatarLink() bool`

HasAvatarLink returns a boolean if a field has been set.

### SetAvatarLinkNil

`func (o *PhoneUserJsonld) SetAvatarLinkNil(b bool)`

 SetAvatarLinkNil sets the value for AvatarLink to be an explicit nil

### UnsetAvatarLink
`func (o *PhoneUserJsonld) UnsetAvatarLink()`

UnsetAvatarLink ensures that no value is present for AvatarLink, not even an explicit nil
### GetPassword

`func (o *PhoneUserJsonld) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PhoneUserJsonld) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PhoneUserJsonld) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PhoneUserJsonld) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAge

`func (o *PhoneUserJsonld) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *PhoneUserJsonld) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *PhoneUserJsonld) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *PhoneUserJsonld) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetIsActive

`func (o *PhoneUserJsonld) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PhoneUserJsonld) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PhoneUserJsonld) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PhoneUserJsonld) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetRoles

`func (o *PhoneUserJsonld) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PhoneUserJsonld) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PhoneUserJsonld) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PhoneUserJsonld) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetToken

`func (o *PhoneUserJsonld) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PhoneUserJsonld) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PhoneUserJsonld) SetToken(v interface{})`

SetToken sets Token field to given value.

### HasToken

`func (o *PhoneUserJsonld) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *PhoneUserJsonld) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *PhoneUserJsonld) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetIsProtected

`func (o *PhoneUserJsonld) GetIsProtected() interface{}`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *PhoneUserJsonld) GetIsProtectedOk() (*interface{}, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *PhoneUserJsonld) SetIsProtected(v interface{})`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *PhoneUserJsonld) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### SetIsProtectedNil

`func (o *PhoneUserJsonld) SetIsProtectedNil(b bool)`

 SetIsProtectedNil sets the value for IsProtected to be an explicit nil

### UnsetIsProtected
`func (o *PhoneUserJsonld) UnsetIsProtected()`

UnsetIsProtected ensures that no value is present for IsProtected, not even an explicit nil
### GetDeletedAtInFuture

`func (o *PhoneUserJsonld) GetDeletedAtInFuture() string`

GetDeletedAtInFuture returns the DeletedAtInFuture field if non-nil, zero value otherwise.

### GetDeletedAtInFutureOk

`func (o *PhoneUserJsonld) GetDeletedAtInFutureOk() (*string, bool)`

GetDeletedAtInFutureOk returns a tuple with the DeletedAtInFuture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAtInFuture

`func (o *PhoneUserJsonld) SetDeletedAtInFuture(v string)`

SetDeletedAtInFuture sets DeletedAtInFuture field to given value.

### HasDeletedAtInFuture

`func (o *PhoneUserJsonld) HasDeletedAtInFuture() bool`

HasDeletedAtInFuture returns a boolean if a field has been set.

### GetActive

`func (o *PhoneUserJsonld) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PhoneUserJsonld) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PhoneUserJsonld) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PhoneUserJsonld) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUserIdentifier

`func (o *PhoneUserJsonld) GetUserIdentifier() string`

GetUserIdentifier returns the UserIdentifier field if non-nil, zero value otherwise.

### GetUserIdentifierOk

`func (o *PhoneUserJsonld) GetUserIdentifierOk() (*string, bool)`

GetUserIdentifierOk returns a tuple with the UserIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentifier

`func (o *PhoneUserJsonld) SetUserIdentifier(v string)`

SetUserIdentifier sets UserIdentifier field to given value.

### HasUserIdentifier

`func (o *PhoneUserJsonld) HasUserIdentifier() bool`

HasUserIdentifier returns a boolean if a field has been set.

### GetProtected

`func (o *PhoneUserJsonld) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *PhoneUserJsonld) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *PhoneUserJsonld) SetProtected(v bool)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *PhoneUserJsonld) HasProtected() bool`

HasProtected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


