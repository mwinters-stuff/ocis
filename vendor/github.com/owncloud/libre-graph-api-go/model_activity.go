/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Activity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Activity{}

// Activity Represents activity.
type Activity struct {
	// Activity ID.
	Id       string           `json:"id"`
	Times    ActivityTimes    `json:"times"`
	Template ActivityTemplate `json:"template"`
}

type _Activity Activity

// NewActivity instantiates a new Activity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivity(id string, times ActivityTimes, template ActivityTemplate) *Activity {
	this := Activity{}
	this.Id = id
	this.Times = times
	this.Template = template
	return &this
}

// NewActivityWithDefaults instantiates a new Activity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityWithDefaults() *Activity {
	this := Activity{}
	return &this
}

// GetId returns the Id field value
func (o *Activity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Activity) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Activity) SetId(v string) {
	o.Id = v
}

// GetTimes returns the Times field value
func (o *Activity) GetTimes() ActivityTimes {
	if o == nil {
		var ret ActivityTimes
		return ret
	}

	return o.Times
}

// GetTimesOk returns a tuple with the Times field value
// and a boolean to check if the value has been set.
func (o *Activity) GetTimesOk() (*ActivityTimes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Times, true
}

// SetTimes sets field value
func (o *Activity) SetTimes(v ActivityTimes) {
	o.Times = v
}

// GetTemplate returns the Template field value
func (o *Activity) GetTemplate() ActivityTemplate {
	if o == nil {
		var ret ActivityTemplate
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *Activity) GetTemplateOk() (*ActivityTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *Activity) SetTemplate(v ActivityTemplate) {
	o.Template = v
}

func (o Activity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Activity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["times"] = o.Times
	toSerialize["template"] = o.Template
	return toSerialize, nil
}

func (o *Activity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"times",
		"template",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varActivity := _Activity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivity)

	if err != nil {
		return err
	}

	*o = Activity(varActivity)

	return err
}

type NullableActivity struct {
	value *Activity
	isSet bool
}

func (v NullableActivity) Get() *Activity {
	return v.value
}

func (v *NullableActivity) Set(val *Activity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivity(val *Activity) *NullableActivity {
	return &NullableActivity{value: val, isSet: true}
}

func (v NullableActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}