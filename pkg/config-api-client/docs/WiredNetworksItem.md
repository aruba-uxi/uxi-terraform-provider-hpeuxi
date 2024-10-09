# WiredNetworksItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**IpVersion** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Security** | **NullableString** |  | 
**DnsLookupDomain** | **NullableString** |  | 
**DisableEdns** | **bool** |  | 
**UseDns64** | **bool** |  | 
**ExternalConnectivity** | **bool** |  | 
**VLanId** | **NullableInt32** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "uxi/wired-network"]

## Methods

### NewWiredNetworksItem

`func NewWiredNetworksItem(id string, name string, ipVersion string, createdAt time.Time, updatedAt time.Time, security NullableString, dnsLookupDomain NullableString, disableEdns bool, useDns64 bool, externalConnectivity bool, vLanId NullableInt32, ) *WiredNetworksItem`

NewWiredNetworksItem instantiates a new WiredNetworksItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWiredNetworksItemWithDefaults

`func NewWiredNetworksItemWithDefaults() *WiredNetworksItem`

NewWiredNetworksItemWithDefaults instantiates a new WiredNetworksItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WiredNetworksItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WiredNetworksItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WiredNetworksItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WiredNetworksItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WiredNetworksItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WiredNetworksItem) SetName(v string)`

SetName sets Name field to given value.


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


### GetCreatedAt

`func (o *WiredNetworksItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WiredNetworksItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WiredNetworksItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WiredNetworksItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WiredNetworksItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WiredNetworksItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


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


### GetVLanId

`func (o *WiredNetworksItem) GetVLanId() int32`

GetVLanId returns the VLanId field if non-nil, zero value otherwise.

### GetVLanIdOk

`func (o *WiredNetworksItem) GetVLanIdOk() (*int32, bool)`

GetVLanIdOk returns a tuple with the VLanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVLanId

`func (o *WiredNetworksItem) SetVLanId(v int32)`

SetVLanId sets VLanId field to given value.


### SetVLanIdNil

`func (o *WiredNetworksItem) SetVLanIdNil(b bool)`

 SetVLanIdNil sets the value for VLanId to be an explicit nil

### UnsetVLanId
`func (o *WiredNetworksItem) UnsetVLanId()`

UnsetVLanId ensures that no value is present for VLanId, not even an explicit nil
### GetType

`func (o *WiredNetworksItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WiredNetworksItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WiredNetworksItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WiredNetworksItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


