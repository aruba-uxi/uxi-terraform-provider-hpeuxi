# WiredNetworksGetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the wired network | 
**Name** | **string** | The name of the wired network | 
**IpVersion** | [**IpVersion**](IpVersion.md) |  | 
**CreatedAt** | **time.Time** | The creation timestamp of the wired network | 
**UpdatedAt** | **time.Time** | The last update timestamp of the wired network | 
**Security** | **NullableString** | The security settings of the wired network | 
**DnsLookupDomain** | **NullableString** | The DNS lookup domain of the wired network | 
**DisableEdns** | **bool** | Whether EDNS is disabled for the wired network | 
**UseDns64** | **bool** | Whether DNS64 is used for the wired network | 
**ExternalConnectivity** | **bool** | Whether the wired network has external connectivity | 
**VLanId** | **NullableInt32** | The VLAN ID of the wired network | 
**Type** | **string** | The type of the resource. | 

## Methods

### NewWiredNetworksGetItem

`func NewWiredNetworksGetItem(id string, name string, ipVersion IpVersion, createdAt time.Time, updatedAt time.Time, security NullableString, dnsLookupDomain NullableString, disableEdns bool, useDns64 bool, externalConnectivity bool, vLanId NullableInt32, type_ string, ) *WiredNetworksGetItem`

NewWiredNetworksGetItem instantiates a new WiredNetworksGetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWiredNetworksGetItemWithDefaults

`func NewWiredNetworksGetItemWithDefaults() *WiredNetworksGetItem`

NewWiredNetworksGetItemWithDefaults instantiates a new WiredNetworksGetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WiredNetworksGetItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WiredNetworksGetItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WiredNetworksGetItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WiredNetworksGetItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WiredNetworksGetItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WiredNetworksGetItem) SetName(v string)`

SetName sets Name field to given value.


### GetIpVersion

`func (o *WiredNetworksGetItem) GetIpVersion() IpVersion`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *WiredNetworksGetItem) GetIpVersionOk() (*IpVersion, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *WiredNetworksGetItem) SetIpVersion(v IpVersion)`

SetIpVersion sets IpVersion field to given value.


### GetCreatedAt

`func (o *WiredNetworksGetItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WiredNetworksGetItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WiredNetworksGetItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WiredNetworksGetItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WiredNetworksGetItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WiredNetworksGetItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSecurity

`func (o *WiredNetworksGetItem) GetSecurity() string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *WiredNetworksGetItem) GetSecurityOk() (*string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *WiredNetworksGetItem) SetSecurity(v string)`

SetSecurity sets Security field to given value.


### SetSecurityNil

`func (o *WiredNetworksGetItem) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *WiredNetworksGetItem) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetDnsLookupDomain

`func (o *WiredNetworksGetItem) GetDnsLookupDomain() string`

GetDnsLookupDomain returns the DnsLookupDomain field if non-nil, zero value otherwise.

### GetDnsLookupDomainOk

`func (o *WiredNetworksGetItem) GetDnsLookupDomainOk() (*string, bool)`

GetDnsLookupDomainOk returns a tuple with the DnsLookupDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsLookupDomain

`func (o *WiredNetworksGetItem) SetDnsLookupDomain(v string)`

SetDnsLookupDomain sets DnsLookupDomain field to given value.


### SetDnsLookupDomainNil

`func (o *WiredNetworksGetItem) SetDnsLookupDomainNil(b bool)`

 SetDnsLookupDomainNil sets the value for DnsLookupDomain to be an explicit nil

### UnsetDnsLookupDomain
`func (o *WiredNetworksGetItem) UnsetDnsLookupDomain()`

UnsetDnsLookupDomain ensures that no value is present for DnsLookupDomain, not even an explicit nil
### GetDisableEdns

`func (o *WiredNetworksGetItem) GetDisableEdns() bool`

GetDisableEdns returns the DisableEdns field if non-nil, zero value otherwise.

### GetDisableEdnsOk

`func (o *WiredNetworksGetItem) GetDisableEdnsOk() (*bool, bool)`

GetDisableEdnsOk returns a tuple with the DisableEdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEdns

`func (o *WiredNetworksGetItem) SetDisableEdns(v bool)`

SetDisableEdns sets DisableEdns field to given value.


### GetUseDns64

`func (o *WiredNetworksGetItem) GetUseDns64() bool`

GetUseDns64 returns the UseDns64 field if non-nil, zero value otherwise.

### GetUseDns64Ok

`func (o *WiredNetworksGetItem) GetUseDns64Ok() (*bool, bool)`

GetUseDns64Ok returns a tuple with the UseDns64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDns64

`func (o *WiredNetworksGetItem) SetUseDns64(v bool)`

SetUseDns64 sets UseDns64 field to given value.


### GetExternalConnectivity

`func (o *WiredNetworksGetItem) GetExternalConnectivity() bool`

GetExternalConnectivity returns the ExternalConnectivity field if non-nil, zero value otherwise.

### GetExternalConnectivityOk

`func (o *WiredNetworksGetItem) GetExternalConnectivityOk() (*bool, bool)`

GetExternalConnectivityOk returns a tuple with the ExternalConnectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnectivity

`func (o *WiredNetworksGetItem) SetExternalConnectivity(v bool)`

SetExternalConnectivity sets ExternalConnectivity field to given value.


### GetVLanId

`func (o *WiredNetworksGetItem) GetVLanId() int32`

GetVLanId returns the VLanId field if non-nil, zero value otherwise.

### GetVLanIdOk

`func (o *WiredNetworksGetItem) GetVLanIdOk() (*int32, bool)`

GetVLanIdOk returns a tuple with the VLanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVLanId

`func (o *WiredNetworksGetItem) SetVLanId(v int32)`

SetVLanId sets VLanId field to given value.


### SetVLanIdNil

`func (o *WiredNetworksGetItem) SetVLanIdNil(b bool)`

 SetVLanIdNil sets the value for VLanId to be an explicit nil

### UnsetVLanId
`func (o *WiredNetworksGetItem) UnsetVLanId()`

UnsetVLanId ensures that no value is present for VLanId, not even an explicit nil
### GetType

`func (o *WiredNetworksGetItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WiredNetworksGetItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WiredNetworksGetItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


