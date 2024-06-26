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

// checks if the LoadBalancerCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerCollection{}

// LoadBalancerCollection struct for LoadBalancerCollection
type LoadBalancerCollection struct {
	Loadbalancers        []LoadBalancer `json:"loadbalancers"`
	AdditionalProperties map[string]interface{}
}

type _LoadBalancerCollection LoadBalancerCollection

// NewLoadBalancerCollection instantiates a new LoadBalancerCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerCollection(loadbalancers []LoadBalancer) *LoadBalancerCollection {
	this := LoadBalancerCollection{}
	this.Loadbalancers = loadbalancers
	return &this
}

// NewLoadBalancerCollectionWithDefaults instantiates a new LoadBalancerCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerCollectionWithDefaults() *LoadBalancerCollection {
	this := LoadBalancerCollection{}
	return &this
}

// GetLoadbalancers returns the Loadbalancers field value
func (o *LoadBalancerCollection) GetLoadbalancers() []LoadBalancer {
	if o == nil {
		var ret []LoadBalancer
		return ret
	}

	return o.Loadbalancers
}

// GetLoadbalancersOk returns a tuple with the Loadbalancers field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerCollection) GetLoadbalancersOk() ([]LoadBalancer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Loadbalancers, true
}

// SetLoadbalancers sets field value
func (o *LoadBalancerCollection) SetLoadbalancers(v []LoadBalancer) {
	o.Loadbalancers = v
}

func (o LoadBalancerCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["loadbalancers"] = o.Loadbalancers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoadBalancerCollection) UnmarshalJSON(bytes []byte) (err error) {
	varLoadBalancerCollection := _LoadBalancerCollection{}

	err = json.Unmarshal(bytes, &varLoadBalancerCollection)

	if err != nil {
		return err
	}

	*o = LoadBalancerCollection(varLoadBalancerCollection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "loadbalancers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoadBalancerCollection struct {
	value *LoadBalancerCollection
	isSet bool
}

func (v NullableLoadBalancerCollection) Get() *LoadBalancerCollection {
	return v.value
}

func (v *NullableLoadBalancerCollection) Set(val *LoadBalancerCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerCollection(val *LoadBalancerCollection) *NullableLoadBalancerCollection {
	return &NullableLoadBalancerCollection{value: val, isSet: true}
}

func (v NullableLoadBalancerCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
