/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// BalancedNics struct for BalancedNics
type BalancedNics struct {
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path).
	Href *string `json:"href,omitempty"`
	// Array of items in the collection.
	Items *[]Nic `json:"items,omitempty"`
	// The offset (if specified in the request).
	Offset *float32 `json:"offset,omitempty"`
	// The limit (if specified in the request).
	Limit *float32         `json:"limit,omitempty"`
	Links *PaginationLinks `json:"_links,omitempty"`
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BalancedNics) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalancedNics) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *BalancedNics) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *BalancedNics) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for Type will be returned
func (o *BalancedNics) GetType() *Type {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalancedNics) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *BalancedNics) SetType(v Type) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *BalancedNics) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BalancedNics) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalancedNics) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *BalancedNics) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *BalancedNics) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []Nic will be returned
func (o *BalancedNics) GetItems() *[]Nic {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalancedNics) GetItemsOk() (*[]Nic, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *BalancedNics) SetItems(v []Nic) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *BalancedNics) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// GetOffset returns the Offset field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *BalancedNics) GetOffset() *float32 {
	if o == nil {
		return nil
	}

	return o.Offset

}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalancedNics) GetOffsetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Offset, true
}

// SetOffset sets field value
func (o *BalancedNics) SetOffset(v float32) {

	o.Offset = &v

}

// HasOffset returns a boolean if a field has been set.
func (o *BalancedNics) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// GetLimit returns the Limit field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *BalancedNics) GetLimit() *float32 {
	if o == nil {
		return nil
	}

	return o.Limit

}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalancedNics) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Limit, true
}

// SetLimit sets field value
func (o *BalancedNics) SetLimit(v float32) {

	o.Limit = &v

}

// HasLimit returns a boolean if a field has been set.
func (o *BalancedNics) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for PaginationLinks will be returned
func (o *BalancedNics) GetLinks() *PaginationLinks {
	if o == nil {
		return nil
	}

	return o.Links

}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalancedNics) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil {
		return nil, false
	}

	return o.Links, true
}

// SetLinks sets field value
func (o *BalancedNics) SetLinks(v PaginationLinks) {

	o.Links = &v

}

// HasLinks returns a boolean if a field has been set.
func (o *BalancedNics) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

func (o BalancedNics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableBalancedNics struct {
	value *BalancedNics
	isSet bool
}

func (v NullableBalancedNics) Get() *BalancedNics {
	return v.value
}

func (v *NullableBalancedNics) Set(val *BalancedNics) {
	v.value = val
	v.isSet = true
}

func (v NullableBalancedNics) IsSet() bool {
	return v.isSet
}

func (v *NullableBalancedNics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalancedNics(val *BalancedNics) *NullableBalancedNics {
	return &NullableBalancedNics{value: val, isSet: true}
}

func (v NullableBalancedNics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalancedNics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
