/*
ORY Oathkeeper

ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.

API version: v0.40.8
Contact: hi@ory.am
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Upstream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Upstream{}

// Upstream struct for Upstream
type Upstream struct {
	// PreserveHost, if false (the default), tells ORY Oathkeeper to set the upstream request's Host header to the hostname of the API's upstream's URL. Setting this flag to true instructs ORY Oathkeeper not to do so.
	PreserveHost *bool `json:"preserve_host,omitempty"`
	// StripPath if set, replaces the provided path prefix when forwarding the requested URL to the upstream URL.
	StripPath *string `json:"strip_path,omitempty"`
	// URL is the URL the request will be proxied to.
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Upstream Upstream

// NewUpstream instantiates a new Upstream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpstream() *Upstream {
	this := Upstream{}
	return &this
}

// NewUpstreamWithDefaults instantiates a new Upstream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpstreamWithDefaults() *Upstream {
	this := Upstream{}
	return &this
}

// GetPreserveHost returns the PreserveHost field value if set, zero value otherwise.
func (o *Upstream) GetPreserveHost() bool {
	if o == nil || IsNil(o.PreserveHost) {
		var ret bool
		return ret
	}
	return *o.PreserveHost
}

// GetPreserveHostOk returns a tuple with the PreserveHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Upstream) GetPreserveHostOk() (*bool, bool) {
	if o == nil || IsNil(o.PreserveHost) {
		return nil, false
	}
	return o.PreserveHost, true
}

// HasPreserveHost returns a boolean if a field has been set.
func (o *Upstream) HasPreserveHost() bool {
	if o != nil && !IsNil(o.PreserveHost) {
		return true
	}

	return false
}

// SetPreserveHost gets a reference to the given bool and assigns it to the PreserveHost field.
func (o *Upstream) SetPreserveHost(v bool) {
	o.PreserveHost = &v
}

// GetStripPath returns the StripPath field value if set, zero value otherwise.
func (o *Upstream) GetStripPath() string {
	if o == nil || IsNil(o.StripPath) {
		var ret string
		return ret
	}
	return *o.StripPath
}

// GetStripPathOk returns a tuple with the StripPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Upstream) GetStripPathOk() (*string, bool) {
	if o == nil || IsNil(o.StripPath) {
		return nil, false
	}
	return o.StripPath, true
}

// HasStripPath returns a boolean if a field has been set.
func (o *Upstream) HasStripPath() bool {
	if o != nil && !IsNil(o.StripPath) {
		return true
	}

	return false
}

// SetStripPath gets a reference to the given string and assigns it to the StripPath field.
func (o *Upstream) SetStripPath(v string) {
	o.StripPath = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Upstream) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Upstream) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Upstream) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Upstream) SetUrl(v string) {
	o.Url = &v
}

func (o Upstream) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Upstream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreserveHost) {
		toSerialize["preserve_host"] = o.PreserveHost
	}
	if !IsNil(o.StripPath) {
		toSerialize["strip_path"] = o.StripPath
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Upstream) UnmarshalJSON(data []byte) (err error) {
	varUpstream := _Upstream{}

	err = json.Unmarshal(data, &varUpstream)

	if err != nil {
		return err
	}

	*o = Upstream(varUpstream)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "preserve_host")
		delete(additionalProperties, "strip_path")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpstream struct {
	value *Upstream
	isSet bool
}

func (v NullableUpstream) Get() *Upstream {
	return v.value
}

func (v *NullableUpstream) Set(val *Upstream) {
	v.value = val
	v.isSet = true
}

func (v NullableUpstream) IsSet() bool {
	return v.isSet
}

func (v *NullableUpstream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpstream(val *Upstream) *NullableUpstream {
	return &NullableUpstream{value: val, isSet: true}
}

func (v NullableUpstream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpstream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


