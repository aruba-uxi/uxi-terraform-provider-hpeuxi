# WirelessNetworksGetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the wireless network | 
**Name** | **string** | The name of the wireless network | 
**Ssid** | **string** | The SSID of the wireless network | 
**Security** | **NullableString** | The security type of the wireless network | 
**IpVersion** | [**IpVersion**](IpVersion.md) |  | 
**CreatedAt** | **time.Time** | The creation timestamp of the wireless network | 
**UpdatedAt** | **time.Time** | The last update timestamp of the wireless network | 
**Hidden** | **bool** | Whether the wireless network is hidden | 
**BandLocking** | **string** | The band locking setting of the wireless network | 
**DnsLookupDomain** | **NullableString** | The DNS lookup domain of the wireless network | 
**DisableEdns** | **bool** | Whether EDNS is disabled for the wireless network | 
**UseDns64** | **bool** | Whether DNS64 is used for the wireless network | 
**ExternalConnectivity** | **bool** | Whether the wireless network has external connectivity | 
**Type** | **string** | The type of the resource. | 

## Methods

### NewWirelessNetworksGetItem

`func NewWirelessNetworksGetItem(id string, name string, ssid string, security NullableString, ipVersion IpVersion, createdAt time.Time, updatedAt time.Time, hidden bool, bandLocking string, dnsLookupDomain NullableString, disableEdns bool, useDns64 bool, externalConnectivity bool, type_ string, ) *WirelessNetworksGetItem`

NewWirelessNetworksGetItem instantiates a new WirelessNetworksGetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessNetworksGetItemWithDefaults

`func NewWirelessNetworksGetItemWithDefaults() *WirelessNetworksGetItem`

NewWirelessNetworksGetItemWithDefaults instantiates a new WirelessNetworksGetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WirelessNetworksGetItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WirelessNetworksGetItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WirelessNetworksGetItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WirelessNetworksGetItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WirelessNetworksGetItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WirelessNetworksGetItem) SetName(v string)`

SetName sets Name field to given value.


### GetSsid

`func (o *WirelessNetworksGetItem) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *WirelessNetworksGetItem) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *WirelessNetworksGetItem) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetSecurity

`func (o *WirelessNetworksGetItem) GetSecurity() string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *WirelessNetworksGetItem) GetSecurityOk() (*string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *WirelessNetworksGetItem) SetSecurity(v string)`

SetSecurity sets Security field to given value.


### SetSecurityNil

`func (o *WirelessNetworksGetItem) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *WirelessNetworksGetItem) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetIpVersion

`func (o *WirelessNetworksGetItem) GetIpVersion() IpVersion`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *WirelessNetworksGetItem) GetIpVersionOk() (*IpVersion, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *WirelessNetworksGetItem) SetIpVersion(v IpVersion)`

SetIpVersion sets IpVersion field to given value.


### GetCreatedAt

`func (o *WirelessNetworksGetItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WirelessNetworksGetItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WirelessNetworksGetItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WirelessNetworksGetItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WirelessNetworksGetItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WirelessNetworksGetItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetHidden

`func (o *WirelessNetworksGetItem) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *WirelessNetworksGetItem) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *WirelessNetworksGetItem) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetBandLocking

`func (o *WirelessNetworksGetItem) GetBandLocking() string`

GetBandLocking returns the BandLocking field if non-nil, zero value otherwise.

### GetBandLockingOk

`func (o *WirelessNetworksGetItem) GetBandLockingOk() (*string, bool)`

GetBandLockingOk returns a tuple with the BandLocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandLocking

`func (o *WirelessNetworksGetItem) SetBandLocking(v string)`

SetBandLocking sets BandLocking field to given value.


### GetDnsLookupDomain

`func (o *WirelessNetworksGetItem) GetDnsLookupDomain() string`

GetDnsLookupDomain returns the DnsLookupDomain field if non-nil, zero value otherwise.

### GetDnsLookupDomainOk

`func (o *WirelessNetworksGetItem) GetDnsLookupDomainOk() (*string, bool)`

GetDnsLookupDomainOk returns a tuple with the DnsLookupDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsLookupDomain

`func (o *WirelessNetworksGetItem) SetDnsLookupDomain(v string)`

SetDnsLookupDomain sets DnsLookupDomain field to given value.


### SetDnsLookupDomainNil

`func (o *WirelessNetworksGetItem) SetDnsLookupDomainNil(b bool)`

 SetDnsLookupDomainNil sets the value for DnsLookupDomain to be an explicit nil

### UnsetDnsLookupDomain
`func (o *WirelessNetworksGetItem) UnsetDnsLookupDomain()`

UnsetDnsLookupDomain ensures that no value is present for DnsLookupDomain, not even an explicit nil
### GetDisableEdns

`func (o *WirelessNetworksGetItem) GetDisableEdns() bool`

GetDisableEdns returns the DisableEdns field if non-nil, zero value otherwise.

### GetDisableEdnsOk

`func (o *WirelessNetworksGetItem) GetDisableEdnsOk() (*bool, bool)`

GetDisableEdnsOk returns a tuple with the DisableEdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEdns

`func (o *WirelessNetworksGetItem) SetDisableEdns(v bool)`

SetDisableEdns sets DisableEdns field to given value.


### GetUseDns64

`func (o *WirelessNetworksGetItem) GetUseDns64() bool`

GetUseDns64 returns the UseDns64 field if non-nil, zero value otherwise.

### GetUseDns64Ok

`func (o *WirelessNetworksGetItem) GetUseDns64Ok() (*bool, bool)`

GetUseDns64Ok returns a tuple with the UseDns64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDns64

`func (o *WirelessNetworksGetItem) SetUseDns64(v bool)`

SetUseDns64 sets UseDns64 field to given value.


### GetExternalConnectivity

`func (o *WirelessNetworksGetItem) GetExternalConnectivity() bool`

GetExternalConnectivity returns the ExternalConnectivity field if non-nil, zero value otherwise.

### GetExternalConnectivityOk

`func (o *WirelessNetworksGetItem) GetExternalConnectivityOk() (*bool, bool)`

GetExternalConnectivityOk returns a tuple with the ExternalConnectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnectivity

`func (o *WirelessNetworksGetItem) SetExternalConnectivity(v bool)`

SetExternalConnectivity sets ExternalConnectivity field to given value.


### GetType

`func (o *WirelessNetworksGetItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WirelessNetworksGetItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WirelessNetworksGetItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


