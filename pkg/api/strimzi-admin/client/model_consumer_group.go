/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package strimziadminclient

import (
	"encoding/json"
)

// ConsumerGroup A group of Kafka consumers
type ConsumerGroup struct {
	// Unique identifier for the consumer group
	Id string `json:"id"`
	// The list of consumers associated with this consumer group
	Consumers []Consumer `json:"consumers"`
}

// NewConsumerGroup instantiates a new ConsumerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerGroup(id string, consumers []Consumer, ) *ConsumerGroup {
	this := ConsumerGroup{}
	this.Id = id
	this.Consumers = consumers
	return &this
}

// NewConsumerGroupWithDefaults instantiates a new ConsumerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerGroupWithDefaults() *ConsumerGroup {
	this := ConsumerGroup{}
	return &this
}

// GetId returns the Id field value
func (o *ConsumerGroup) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroup) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConsumerGroup) SetId(v string) {
	o.Id = v
}

// GetConsumers returns the Consumers field value
func (o *ConsumerGroup) GetConsumers() []Consumer {
	if o == nil  {
		var ret []Consumer
		return ret
	}

	return o.Consumers
}

// GetConsumersOk returns a tuple with the Consumers field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroup) GetConsumersOk() (*[]Consumer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Consumers, true
}

// SetConsumers sets field value
func (o *ConsumerGroup) SetConsumers(v []Consumer) {
	o.Consumers = v
}

func (o ConsumerGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["consumers"] = o.Consumers
	}
	return json.Marshal(toSerialize)
}

type NullableConsumerGroup struct {
	value *ConsumerGroup
	isSet bool
}

func (v NullableConsumerGroup) Get() *ConsumerGroup {
	return v.value
}

func (v *NullableConsumerGroup) Set(val *ConsumerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerGroup(val *ConsumerGroup) *NullableConsumerGroup {
	return &NullableConsumerGroup{value: val, isSet: true}
}

func (v NullableConsumerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


