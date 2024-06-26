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
	"time"
)

// checks if the LoadBalancerPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerPort{}

// LoadBalancerPort struct for LoadBalancerPort
type LoadBalancerPort struct {
	// ID of a port
	Id *string `json:"id,omitempty"`
	// Date and time of creation
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time of last update
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// A name for the port
	Name *string `json:"name,omitempty"`
	// Port Number
	Number *int32 `json:"number,omitempty"`
	// ID of the load balancer this port is assigned to
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
	// A list of pool IDs assigned to the port
	PoolIds              []string `json:"pool_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoadBalancerPort LoadBalancerPort

// NewLoadBalancerPort instantiates a new LoadBalancerPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerPort() *LoadBalancerPort {
	this := LoadBalancerPort{}
	return &this
}

// NewLoadBalancerPortWithDefaults instantiates a new LoadBalancerPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerPortWithDefaults() *LoadBalancerPort {
	this := LoadBalancerPort{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LoadBalancerPort) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPort) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LoadBalancerPort) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LoadBalancerPort) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LoadBalancerPort) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPort) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LoadBalancerPort) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LoadBalancerPort) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LoadBalancerPort) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPort) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LoadBalancerPort) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *LoadBalancerPort) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoadBalancerPort) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPort) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoadBalancerPort) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoadBalancerPort) SetName(v string) {
	o.Name = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *LoadBalancerPort) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPort) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *LoadBalancerPort) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *LoadBalancerPort) SetNumber(v int32) {
	o.Number = &v
}

// GetLoadbalancerId returns the LoadbalancerId field value if set, zero value otherwise.
func (o *LoadBalancerPort) GetLoadbalancerId() string {
	if o == nil || IsNil(o.LoadbalancerId) {
		var ret string
		return ret
	}
	return *o.LoadbalancerId
}

// GetLoadbalancerIdOk returns a tuple with the LoadbalancerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPort) GetLoadbalancerIdOk() (*string, bool) {
	if o == nil || IsNil(o.LoadbalancerId) {
		return nil, false
	}
	return o.LoadbalancerId, true
}

// HasLoadbalancerId returns a boolean if a field has been set.
func (o *LoadBalancerPort) HasLoadbalancerId() bool {
	if o != nil && !IsNil(o.LoadbalancerId) {
		return true
	}

	return false
}

// SetLoadbalancerId gets a reference to the given string and assigns it to the LoadbalancerId field.
func (o *LoadBalancerPort) SetLoadbalancerId(v string) {
	o.LoadbalancerId = &v
}

// GetPoolIds returns the PoolIds field value if set, zero value otherwise.
func (o *LoadBalancerPort) GetPoolIds() []string {
	if o == nil || IsNil(o.PoolIds) {
		var ret []string
		return ret
	}
	return o.PoolIds
}

// GetPoolIdsOk returns a tuple with the PoolIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPort) GetPoolIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PoolIds) {
		return nil, false
	}
	return o.PoolIds, true
}

// HasPoolIds returns a boolean if a field has been set.
func (o *LoadBalancerPort) HasPoolIds() bool {
	if o != nil && !IsNil(o.PoolIds) {
		return true
	}

	return false
}

// SetPoolIds gets a reference to the given []string and assigns it to the PoolIds field.
func (o *LoadBalancerPort) SetPoolIds(v []string) {
	o.PoolIds = v
}

func (o LoadBalancerPort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.LoadbalancerId) {
		toSerialize["loadbalancer_id"] = o.LoadbalancerId
	}
	if !IsNil(o.PoolIds) {
		toSerialize["pool_ids"] = o.PoolIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadBalancerPort) UnmarshalJSON(bytes []byte) (err error) {
	varLoadBalancerPort := _LoadBalancerPort{}

	err = json.Unmarshal(bytes, &varLoadBalancerPort)

	if err != nil {
		return err
	}

	*o = LoadBalancerPort(varLoadBalancerPort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "name")
		delete(additionalProperties, "number")
		delete(additionalProperties, "loadbalancer_id")
		delete(additionalProperties, "pool_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadBalancerPort struct {
	value *LoadBalancerPort
	isSet bool
}

func (v NullableLoadBalancerPort) Get() *LoadBalancerPort {
	return v.value
}

func (v *NullableLoadBalancerPort) Set(val *LoadBalancerPort) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerPort) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerPort(val *LoadBalancerPort) *NullableLoadBalancerPort {
	return &NullableLoadBalancerPort{value: val, isSet: true}
}

func (v NullableLoadBalancerPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
