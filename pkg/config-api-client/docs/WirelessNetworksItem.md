# WirelessNetworksItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | 
**Alias** | **string** |  | 
**Ssid** | **string** |  | 
**Security** | **NullableString** |  | 
**IpVersion** | **string** |  | 
**DatetimeCreated** | **time.Time** |  | 
**DatetimeUpdated** | **time.Time** |  | 
**Hidden** | **bool** |  | 
**BandLocking** | **string** |  | 
**DnsLookupDomain** | **NullableString** |  | 
**DisableEdns** | **bool** |  | 
**UseDns64** | **bool** |  | 
**ExternalConnectivity** | **bool** |  | 

## Methods

### NewWirelessNetworksItem

`func NewWirelessNetworksItem(uid string, alias string, ssid string, security NullableString, ipVersion string, datetimeCreated time.Time, datetimeUpdated time.Time, hidden bool, bandLocking string, dnsLookupDomain NullableString, disableEdns bool, useDns64 bool, externalConnectivity bool, ) *WirelessNetworksItem`

NewWirelessNetworksItem instantiates a new WirelessNetworksItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessNetworksItemWithDefaults

`func NewWirelessNetworksItemWithDefaults() *WirelessNetworksItem`

NewWirelessNetworksItemWithDefaults instantiates a new WirelessNetworksItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *WirelessNetworksItem) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *WirelessNetworksItem) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *WirelessNetworksItem) SetUid(v string)`

SetUid sets Uid field to given value.


### GetAlias

`func (o *WirelessNetworksItem) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *WirelessNetworksItem) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *WirelessNetworksItem) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetSsid

`func (o *WirelessNetworksItem) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *WirelessNetworksItem) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *WirelessNetworksItem) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetSecurity

`func (o *WirelessNetworksItem) GetSecurity() string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *WirelessNetworksItem) GetSecurityOk() (*string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *WirelessNetworksItem) SetSecurity(v string)`

SetSecurity sets Security field to given value.


### SetSecurityNil

`func (o *WirelessNetworksItem) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *WirelessNetworksItem) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetIpVersion

`func (o *WirelessNetworksItem) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *WirelessNetworksItem) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *WirelessNetworksItem) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.


### GetDatetimeCreated

`func (o *WirelessNetworksItem) GetDatetimeCreated() time.Time`

GetDatetimeCreated returns the DatetimeCreated field if non-nil, zero value otherwise.

### GetDatetimeCreatedOk

`func (o *WirelessNetworksItem) GetDatetimeCreatedOk() (*time.Time, bool)`

GetDatetimeCreatedOk returns a tuple with the DatetimeCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeCreated

`func (o *WirelessNetworksItem) SetDatetimeCreated(v time.Time)`

SetDatetimeCreated sets DatetimeCreated field to given value.


### GetDatetimeUpdated

`func (o *WirelessNetworksItem) GetDatetimeUpdated() time.Time`

GetDatetimeUpdated returns the DatetimeUpdated field if non-nil, zero value otherwise.

### GetDatetimeUpdatedOk

`func (o *WirelessNetworksItem) GetDatetimeUpdatedOk() (*time.Time, bool)`

GetDatetimeUpdatedOk returns a tuple with the DatetimeUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeUpdated

`func (o *WirelessNetworksItem) SetDatetimeUpdated(v time.Time)`

SetDatetimeUpdated sets DatetimeUpdated field to given value.


### GetHidden

`func (o *WirelessNetworksItem) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *WirelessNetworksItem) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *WirelessNetworksItem) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetBandLocking

`func (o *WirelessNetworksItem) GetBandLocking() string`

GetBandLocking returns the BandLocking field if non-nil, zero value otherwise.

### GetBandLockingOk

`func (o *WirelessNetworksItem) GetBandLockingOk() (*string, bool)`

GetBandLockingOk returns a tuple with the BandLocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandLocking

`func (o *WirelessNetworksItem) SetBandLocking(v string)`

SetBandLocking sets BandLocking field to given value.


### GetDnsLookupDomain

`func (o *WirelessNetworksItem) GetDnsLookupDomain() string`

GetDnsLookupDomain returns the DnsLookupDomain field if non-nil, zero value otherwise.

### GetDnsLookupDomainOk

`func (o *WirelessNetworksItem) GetDnsLookupDomainOk() (*string, bool)`

GetDnsLookupDomainOk returns a tuple with the DnsLookupDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsLookupDomain

`func (o *WirelessNetworksItem) SetDnsLookupDomain(v string)`

SetDnsLookupDomain sets DnsLookupDomain field to given value.


### SetDnsLookupDomainNil

`func (o *WirelessNetworksItem) SetDnsLookupDomainNil(b bool)`

 SetDnsLookupDomainNil sets the value for DnsLookupDomain to be an explicit nil

### UnsetDnsLookupDomain
`func (o *WirelessNetworksItem) UnsetDnsLookupDomain()`

UnsetDnsLookupDomain ensures that no value is present for DnsLookupDomain, not even an explicit nil
### GetDisableEdns

`func (o *WirelessNetworksItem) GetDisableEdns() bool`

GetDisableEdns returns the DisableEdns field if non-nil, zero value otherwise.

### GetDisableEdnsOk

`func (o *WirelessNetworksItem) GetDisableEdnsOk() (*bool, bool)`

GetDisableEdnsOk returns a tuple with the DisableEdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEdns

`func (o *WirelessNetworksItem) SetDisableEdns(v bool)`

SetDisableEdns sets DisableEdns field to given value.


### GetUseDns64

`func (o *WirelessNetworksItem) GetUseDns64() bool`

GetUseDns64 returns the UseDns64 field if non-nil, zero value otherwise.

### GetUseDns64Ok

`func (o *WirelessNetworksItem) GetUseDns64Ok() (*bool, bool)`

GetUseDns64Ok returns a tuple with the UseDns64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDns64

`func (o *WirelessNetworksItem) SetUseDns64(v bool)`

SetUseDns64 sets UseDns64 field to given value.


### GetExternalConnectivity

`func (o *WirelessNetworksItem) GetExternalConnectivity() bool`

GetExternalConnectivity returns the ExternalConnectivity field if non-nil, zero value otherwise.

### GetExternalConnectivityOk

`func (o *WirelessNetworksItem) GetExternalConnectivityOk() (*bool, bool)`

GetExternalConnectivityOk returns a tuple with the ExternalConnectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnectivity

`func (o *WirelessNetworksItem) SetExternalConnectivity(v bool)`

SetExternalConnectivity sets ExternalConnectivity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


