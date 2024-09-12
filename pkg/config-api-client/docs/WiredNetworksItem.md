# WiredNetworksItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | 
**Alias** | **string** |  | 
**IpVersion** | **string** |  | 
**DatetimeUpdated** | **time.Time** |  | 
**DatetimeCreated** | **time.Time** |  | 
**Security** | **NullableString** |  | 
**DnsLookupDomain** | **NullableString** |  | 
**DisableEdns** | **bool** |  | 
**UseDns64** | **bool** |  | 
**ExternalConnectivity** | **bool** |  | 
**VlanId** | **NullableInt32** |  | 

## Methods

### NewWiredNetworksItem

`func NewWiredNetworksItem(uid string, alias string, ipVersion string, datetimeUpdated time.Time, datetimeCreated time.Time, security NullableString, dnsLookupDomain NullableString, disableEdns bool, useDns64 bool, externalConnectivity bool, vlanId NullableInt32, ) *WiredNetworksItem`

NewWiredNetworksItem instantiates a new WiredNetworksItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWiredNetworksItemWithDefaults

`func NewWiredNetworksItemWithDefaults() *WiredNetworksItem`

NewWiredNetworksItemWithDefaults instantiates a new WiredNetworksItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *WiredNetworksItem) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *WiredNetworksItem) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *WiredNetworksItem) SetUid(v string)`

SetUid sets Uid field to given value.


### GetAlias

`func (o *WiredNetworksItem) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *WiredNetworksItem) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *WiredNetworksItem) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetIpVersion

`func (o *WiredNetworksItem) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *WiredNetworksItem) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *WiredNetworksItem) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.


### GetDatetimeUpdated

`func (o *WiredNetworksItem) GetDatetimeUpdated() time.Time`

GetDatetimeUpdated returns the DatetimeUpdated field if non-nil, zero value otherwise.

### GetDatetimeUpdatedOk

`func (o *WiredNetworksItem) GetDatetimeUpdatedOk() (*time.Time, bool)`

GetDatetimeUpdatedOk returns a tuple with the DatetimeUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeUpdated

`func (o *WiredNetworksItem) SetDatetimeUpdated(v time.Time)`

SetDatetimeUpdated sets DatetimeUpdated field to given value.


### GetDatetimeCreated

`func (o *WiredNetworksItem) GetDatetimeCreated() time.Time`

GetDatetimeCreated returns the DatetimeCreated field if non-nil, zero value otherwise.

### GetDatetimeCreatedOk

`func (o *WiredNetworksItem) GetDatetimeCreatedOk() (*time.Time, bool)`

GetDatetimeCreatedOk returns a tuple with the DatetimeCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeCreated

`func (o *WiredNetworksItem) SetDatetimeCreated(v time.Time)`

SetDatetimeCreated sets DatetimeCreated field to given value.


### GetSecurity

`func (o *WiredNetworksItem) GetSecurity() string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *WiredNetworksItem) GetSecurityOk() (*string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *WiredNetworksItem) SetSecurity(v string)`

SetSecurity sets Security field to given value.


### SetSecurityNil

`func (o *WiredNetworksItem) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *WiredNetworksItem) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetDnsLookupDomain

`func (o *WiredNetworksItem) GetDnsLookupDomain() string`

GetDnsLookupDomain returns the DnsLookupDomain field if non-nil, zero value otherwise.

### GetDnsLookupDomainOk

`func (o *WiredNetworksItem) GetDnsLookupDomainOk() (*string, bool)`

GetDnsLookupDomainOk returns a tuple with the DnsLookupDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsLookupDomain

`func (o *WiredNetworksItem) SetDnsLookupDomain(v string)`

SetDnsLookupDomain sets DnsLookupDomain field to given value.


### SetDnsLookupDomainNil

`func (o *WiredNetworksItem) SetDnsLookupDomainNil(b bool)`

 SetDnsLookupDomainNil sets the value for DnsLookupDomain to be an explicit nil

### UnsetDnsLookupDomain
`func (o *WiredNetworksItem) UnsetDnsLookupDomain()`

UnsetDnsLookupDomain ensures that no value is present for DnsLookupDomain, not even an explicit nil
### GetDisableEdns

`func (o *WiredNetworksItem) GetDisableEdns() bool`

GetDisableEdns returns the DisableEdns field if non-nil, zero value otherwise.

### GetDisableEdnsOk

`func (o *WiredNetworksItem) GetDisableEdnsOk() (*bool, bool)`

GetDisableEdnsOk returns a tuple with the DisableEdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEdns

`func (o *WiredNetworksItem) SetDisableEdns(v bool)`

SetDisableEdns sets DisableEdns field to given value.


### GetUseDns64

`func (o *WiredNetworksItem) GetUseDns64() bool`

GetUseDns64 returns the UseDns64 field if non-nil, zero value otherwise.

### GetUseDns64Ok

`func (o *WiredNetworksItem) GetUseDns64Ok() (*bool, bool)`

GetUseDns64Ok returns a tuple with the UseDns64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDns64

`func (o *WiredNetworksItem) SetUseDns64(v bool)`

SetUseDns64 sets UseDns64 field to given value.


### GetExternalConnectivity

`func (o *WiredNetworksItem) GetExternalConnectivity() bool`

GetExternalConnectivity returns the ExternalConnectivity field if non-nil, zero value otherwise.

### GetExternalConnectivityOk

`func (o *WiredNetworksItem) GetExternalConnectivityOk() (*bool, bool)`

GetExternalConnectivityOk returns a tuple with the ExternalConnectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnectivity

`func (o *WiredNetworksItem) SetExternalConnectivity(v bool)`

SetExternalConnectivity sets ExternalConnectivity field to given value.


### GetVlanId

`func (o *WiredNetworksItem) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *WiredNetworksItem) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *WiredNetworksItem) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.


### SetVlanIdNil

`func (o *WiredNetworksItem) SetVlanIdNil(b bool)`

 SetVlanIdNil sets the value for VlanId to be an explicit nil

### UnsetVlanId
`func (o *WiredNetworksItem) UnsetVlanId()`

UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


