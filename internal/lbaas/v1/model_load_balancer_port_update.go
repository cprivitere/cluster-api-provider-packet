/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Load Balancer Management API

Load Balancer Management API is an API for managing load balancers.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the LoadBalancerPortUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerPortUpdate{}

// LoadBalancerPortUpdate struct for LoadBalancerPortUpdate
type LoadBalancerPortUpdate struct {
	// A name for the port
	Name *string `json:"name,omitempty"`
	// Listing port
	Number *int32 `json:"number,omitempty"`
	// Add the provided pool ids to the port
	AddPoolIds []string `json:"add_pool_ids,omitempty"`
	// Remove the provided pool ids to the port
	RemovePoolIds        []string `json:"remove_pool_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoadBalancerPortUpdate LoadBalancerPortUpdate

// NewLoadBalancerPortUpdate instantiates a new LoadBalancerPortUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerPortUpdate() *LoadBalancerPortUpdate {
	this := LoadBalancerPortUpdate{}
	return &this
}

// NewLoadBalancerPortUpdateWithDefaults instantiates a new LoadBalancerPortUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerPortUpdateWithDefaults() *LoadBalancerPortUpdate {
	this := LoadBalancerPortUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoadBalancerPortUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPortUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoadBalancerPortUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoadBalancerPortUpdate) SetName(v string) {
	o.Name = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *LoadBalancerPortUpdate) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPortUpdate) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *LoadBalancerPortUpdate) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *LoadBalancerPortUpdate) SetNumber(v int32) {
	o.Number = &v
}

// GetAddPoolIds returns the AddPoolIds field value if set, zero value otherwise.
func (o *LoadBalancerPortUpdate) GetAddPoolIds() []string {
	if o == nil || IsNil(o.AddPoolIds) {
		var ret []string
		return ret
	}
	return o.AddPoolIds
}

// GetAddPoolIdsOk returns a tuple with the AddPoolIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPortUpdate) GetAddPoolIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AddPoolIds) {
		return nil, false
	}
	return o.AddPoolIds, true
}

// HasAddPoolIds returns a boolean if a field has been set.
func (o *LoadBalancerPortUpdate) HasAddPoolIds() bool {
	if o != nil && !IsNil(o.AddPoolIds) {
		return true
	}

	return false
}

// SetAddPoolIds gets a reference to the given []string and assigns it to the AddPoolIds field.
func (o *LoadBalancerPortUpdate) SetAddPoolIds(v []string) {
	o.AddPoolIds = v
}

// GetRemovePoolIds returns the RemovePoolIds field value if set, zero value otherwise.
func (o *LoadBalancerPortUpdate) GetRemovePoolIds() []string {
	if o == nil || IsNil(o.RemovePoolIds) {
		var ret []string
		return ret
	}
	return o.RemovePoolIds
}

// GetRemovePoolIdsOk returns a tuple with the RemovePoolIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPortUpdate) GetRemovePoolIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RemovePoolIds) {
		return nil, false
	}
	return o.RemovePoolIds, true
}

// HasRemovePoolIds returns a boolean if a field has been set.
func (o *LoadBalancerPortUpdate) HasRemovePoolIds() bool {
	if o != nil && !IsNil(o.RemovePoolIds) {
		return true
	}

	return false
}

// SetRemovePoolIds gets a reference to the given []string and assigns it to the RemovePoolIds field.
func (o *LoadBalancerPortUpdate) SetRemovePoolIds(v []string) {
	o.RemovePoolIds = v
}

func (o LoadBalancerPortUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerPortUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.AddPoolIds) {
		toSerialize["add_pool_ids"] = o.AddPoolIds
	}
	if !IsNil(o.RemovePoolIds) {
		toSerialize["remove_pool_ids"] = o.RemovePoolIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadBalancerPortUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varLoadBalancerPortUpdate := _LoadBalancerPortUpdate{}

	err = json.Unmarshal(bytes, &varLoadBalancerPortUpdate)

	if err != nil {
		return err
	}

	*o = LoadBalancerPortUpdate(varLoadBalancerPortUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "number")
		delete(additionalProperties, "add_pool_ids")
		delete(additionalProperties, "remove_pool_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadBalancerPortUpdate struct {
	value *LoadBalancerPortUpdate
	isSet bool
}

func (v NullableLoadBalancerPortUpdate) Get() *LoadBalancerPortUpdate {
	return v.value
}

func (v *NullableLoadBalancerPortUpdate) Set(val *LoadBalancerPortUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerPortUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerPortUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerPortUpdate(val *LoadBalancerPortUpdate) *NullableLoadBalancerPortUpdate {
	return &NullableLoadBalancerPortUpdate{value: val, isSet: true}
}

func (v NullableLoadBalancerPortUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerPortUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}