// Code generated by go generate; DO NOT EDIT.
package system

import (
	"net/url"

	"github.com/containers/podman/v3/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *EventsOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *EventsOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithFilters set field Filters to given value
func (o *EventsOptions) WithFilters(value map[string][]string) *EventsOptions {
	o.Filters = value
	return o
}

// GetFilters returns value of field Filters
func (o *EventsOptions) GetFilters() map[string][]string {
	if o.Filters == nil {
		var z map[string][]string
		return z
	}
	return o.Filters
}

// WithSince set field Since to given value
func (o *EventsOptions) WithSince(value string) *EventsOptions {
	o.Since = &value
	return o
}

// GetSince returns value of field Since
func (o *EventsOptions) GetSince() string {
	if o.Since == nil {
		var z string
		return z
	}
	return *o.Since
}

// WithStream set field Stream to given value
func (o *EventsOptions) WithStream(value bool) *EventsOptions {
	o.Stream = &value
	return o
}

// GetStream returns value of field Stream
func (o *EventsOptions) GetStream() bool {
	if o.Stream == nil {
		var z bool
		return z
	}
	return *o.Stream
}

// WithUntil set field Until to given value
func (o *EventsOptions) WithUntil(value string) *EventsOptions {
	o.Until = &value
	return o
}

// GetUntil returns value of field Until
func (o *EventsOptions) GetUntil() string {
	if o.Until == nil {
		var z string
		return z
	}
	return *o.Until
}
