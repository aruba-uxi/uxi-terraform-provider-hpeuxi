# WiredNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** |  | 
**Alias** | **string** |  | 
**IpVersion** | **string** |  | 
**SensorCount** | **int32** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**Disabled** | **bool** |  | 
**Security** | **NullableString** |  | 
**DnsLookupDomain** | **NullableString** |  | 
**DisableEdns** | **bool** |  | 
**UseDns64** | **bool** |  | 
**ExternalConnectivity** | **bool** |  | 
**VlanId** | **NullableInt32** |  | 

## Methods

### NewWiredNetwork

`func NewWiredNetwork(uid string, alias string, ipVersion string, sensorCount int32, updatedAt time.Time, createdAt time.Time, disabled bool, security NullableString, dnsLookupDomain NullableString, disableEdns bool, useDns64 bool, externalConnectivity bool, vlanId NullableInt32, ) *WiredNetwork`

NewWiredNetwork instantiates a new WiredNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWiredNetworkWithDefaults

`func NewWiredNetworkWithDefaults() *WiredNetwork`

NewWiredNetworkWithDefaults instantiates a new WiredNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *WiredNetwork) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *WiredNetwork) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *WiredNetwork) SetUid(v string)`

SetUid sets Uid field to given value.


### GetAlias

`func (o *WiredNetwork) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *WiredNetwork) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *WiredNetwork) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetIpVersion

`func (o *WiredNetwork) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *WiredNetwork) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *WiredNetwork) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.


### GetSensorCount

`func (o *WiredNetwork) GetSensorCount() int32`

GetSensorCount returns the SensorCount field if non-nil, zero value otherwise.

### GetSensorCountOk

`func (o *WiredNetwork) GetSensorCountOk() (*int32, bool)`

GetSensorCountOk returns a tuple with the SensorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorCount

`func (o *WiredNetwork) SetSensorCount(v int32)`

SetSensorCount sets SensorCount field to given value.


### GetUpdatedAt

`func (o *WiredNetwork) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WiredNetwork) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WiredNetwork) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *WiredNetwork) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WiredNetwork) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WiredNetwork) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDisabled

`func (o *WiredNetwork) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *WiredNetwork) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *WiredNetwork) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetSecurity

`func (o *WiredNetwork) GetSecurity() string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *WiredNetwork) GetSecurityOk() (*string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *WiredNetwork) SetSecurity(v string)`

SetSecurity sets Security field to given value.


### SetSecurityNil

`func (o *WiredNetwork) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *WiredNetwork) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetDnsLookupDomain

`func (o *WiredNetwork) GetDnsLookupDomain() string`

GetDnsLookupDomain returns the DnsLookupDomain field if non-nil, zero value otherwise.

### GetDnsLookupDomainOk

`func (o *WiredNetwork) GetDnsLookupDomainOk() (*string, bool)`

GetDnsLookupDomainOk returns a tuple with the DnsLookupDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsLookupDomain

`func (o *WiredNetwork) SetDnsLookupDomain(v string)`

SetDnsLookupDomain sets DnsLookupDomain field to given value.


### SetDnsLookupDomainNil

`func (o *WiredNetwork) SetDnsLookupDomainNil(b bool)`

 SetDnsLookupDomainNil sets the value for DnsLookupDomain to be an explicit nil

### UnsetDnsLookupDomain
`func (o *WiredNetwork) UnsetDnsLookupDomain()`

UnsetDnsLookupDomain ensures that no value is present for DnsLookupDomain, not even an explicit nil
### GetDisableEdns

`func (o *WiredNetwork) GetDisableEdns() bool`

GetDisableEdns returns the DisableEdns field if non-nil, zero value otherwise.

### GetDisableEdnsOk

`func (o *WiredNetwork) GetDisableEdnsOk() (*bool, bool)`

GetDisableEdnsOk returns a tuple with the DisableEdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEdns

`func (o *WiredNetwork) SetDisableEdns(v bool)`

SetDisableEdns sets DisableEdns field to given value.


### GetUseDns64

`func (o *WiredNetwork) GetUseDns64() bool`

GetUseDns64 returns the UseDns64 field if non-nil, zero value otherwise.

### GetUseDns64Ok

`func (o *WiredNetwork) GetUseDns64Ok() (*bool, bool)`

GetUseDns64Ok returns a tuple with the UseDns64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDns64

`func (o *WiredNetwork) SetUseDns64(v bool)`

SetUseDns64 sets UseDns64 field to given value.


### GetExternalConnectivity

`func (o *WiredNetwork) GetExternalConnectivity() bool`

GetExternalConnectivity returns the ExternalConnectivity field if non-nil, zero value otherwise.

### GetExternalConnectivityOk

`func (o *WiredNetwork) GetExternalConnectivityOk() (*bool, bool)`

GetExternalConnectivityOk returns a tuple with the ExternalConnectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnectivity

`func (o *WiredNetwork) SetExternalConnectivity(v bool)`

SetExternalConnectivity sets ExternalConnectivity field to given value.


### GetVlanId

`func (o *WiredNetwork) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *WiredNetwork) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *WiredNetwork) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.


### SetVlanIdNil

`func (o *WiredNetwork) SetVlanIdNil(b bool)`

 SetVlanIdNil sets the value for VlanId to be an explicit nil

### UnsetVlanId
`func (o *WiredNetwork) UnsetVlanId()`

UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


