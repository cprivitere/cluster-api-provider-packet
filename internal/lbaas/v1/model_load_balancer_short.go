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

// checks if the LoadBalancerShort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerShort{}

// LoadBalancerShort struct for LoadBalancerShort
type LoadBalancerShort struct {
	// ID of a load balancer
	Id string `json:"id"`
	// A name for the load balancer
	Name string `json:"name"`
	// Date and time of creation
	CreatedAt time.Time `json:"created_at"`
	// Date and time of last update
	UpdatedAt            time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _LoadBalancerShort LoadBalancerShort

// NewLoadBalancerShort instantiates a new LoadBalancerShort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerShort(id string, name string, createdAt time.Time, updatedAt time.Time) *LoadBalancerShort {
	this := LoadBalancerShort{}
	this.Id = id
	this.Name = name
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewLoadBalancerShortWithDefaults instantiates a new LoadBalancerShort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerShortWithDefaults() *LoadBalancerShort {
	this := LoadBalancerShort{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *LoadBalancerShort) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadBalancerShort) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoadBalancerShort) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *LoadBalancerShort) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerShort) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LoadBalancerShort) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *LoadBalancerShort) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerShort) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LoadBalancerShort) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *LoadBalancerShort) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerShort) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *LoadBalancerShort) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o LoadBalancerShort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerShort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadBalancerShort) UnmarshalJSON(bytes []byte) (err error) {
	varLoadBalancerShort := _LoadBalancerShort{}

	err = json.Unmarshal(bytes, &varLoadBalancerShort)

	if err != nil {
		return err
	}

	*o = LoadBalancerShort(varLoadBalancerShort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadBalancerShort struct {
	value *LoadBalancerShort
	isSet bool
}

func (v NullableLoadBalancerShort) Get() *LoadBalancerShort {
	return v.value
}

func (v *NullableLoadBalancerShort) Set(val *LoadBalancerShort) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerShort) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerShort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerShort(val *LoadBalancerShort) *NullableLoadBalancerShort {
	return &NullableLoadBalancerShort{value: val, isSet: true}
}

func (v NullableLoadBalancerShort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerShort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
