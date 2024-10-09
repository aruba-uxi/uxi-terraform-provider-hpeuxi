# WirelessNetworksItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Ssid** | **string** |  | 
**Security** | **NullableString** |  | 
**IpVersion** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Hidden** | **bool** |  | 
**BandLocking** | **string** |  | 
**DnsLookupDomain** | **NullableString** |  | 
**DisableEdns** | **bool** |  | 
**UseDns64** | **bool** |  | 
**ExternalConnectivity** | **bool** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "uxi/wireless-network"]

## Methods

### NewWirelessNetworksItem

`func NewWirelessNetworksItem(id string, name string, ssid string, security NullableString, ipVersion string, createdAt time.Time, updatedAt time.Time, hidden bool, bandLocking string, dnsLookupDomain NullableString, disableEdns bool, useDns64 bool, externalConnectivity bool, ) *WirelessNetworksItem`

NewWirelessNetworksItem instantiates a new WirelessNetworksItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessNetworksItemWithDefaults

`func NewWirelessNetworksItemWithDefaults() *WirelessNetworksItem`

NewWirelessNetworksItemWithDefaults instantiates a new WirelessNetworksItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WirelessNetworksItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WirelessNetworksItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WirelessNetworksItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WirelessNetworksItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WirelessNetworksItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WirelessNetworksItem) SetName(v string)`

SetName sets Name field to given value.


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


### GetCreatedAt

`func (o *WirelessNetworksItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WirelessNetworksItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WirelessNetworksItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WirelessNetworksItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WirelessNetworksItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WirelessNetworksItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


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


### GetType

`func (o *WirelessNetworksItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WirelessNetworksItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WirelessNetworksItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WirelessNetworksItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


