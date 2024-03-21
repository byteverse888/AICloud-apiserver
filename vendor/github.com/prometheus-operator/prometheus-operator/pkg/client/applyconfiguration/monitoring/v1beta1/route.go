// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// RouteApplyConfiguration represents an declarative configuration of the Route type for use
// with apply.
type RouteApplyConfiguration struct {
	Receiver            *string                     `json:"receiver,omitempty"`
	GroupBy             []string                    `json:"groupBy,omitempty"`
	GroupWait           *string                     `json:"groupWait,omitempty"`
	GroupInterval       *string                     `json:"groupInterval,omitempty"`
	RepeatInterval      *string                     `json:"repeatInterval,omitempty"`
	Matchers            []MatcherApplyConfiguration `json:"matchers,omitempty"`
	Continue            *bool                       `json:"continue,omitempty"`
	Routes              []v1.JSON                   `json:"routes,omitempty"`
	MuteTimeIntervals   []string                    `json:"muteTimeIntervals,omitempty"`
	ActiveTimeIntervals []string                    `json:"activeTimeIntervals,omitempty"`
}

// RouteApplyConfiguration constructs an declarative configuration of the Route type for use with
// apply.
func Route() *RouteApplyConfiguration {
	return &RouteApplyConfiguration{}
}

// WithReceiver sets the Receiver field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Receiver field is set to the value of the last call.
func (b *RouteApplyConfiguration) WithReceiver(value string) *RouteApplyConfiguration {
	b.Receiver = &value
	return b
}

// WithGroupBy adds the given value to the GroupBy field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the GroupBy field.
func (b *RouteApplyConfiguration) WithGroupBy(values ...string) *RouteApplyConfiguration {
	for i := range values {
		b.GroupBy = append(b.GroupBy, values[i])
	}
	return b
}

// WithGroupWait sets the GroupWait field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GroupWait field is set to the value of the last call.
func (b *RouteApplyConfiguration) WithGroupWait(value string) *RouteApplyConfiguration {
	b.GroupWait = &value
	return b
}

// WithGroupInterval sets the GroupInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GroupInterval field is set to the value of the last call.
func (b *RouteApplyConfiguration) WithGroupInterval(value string) *RouteApplyConfiguration {
	b.GroupInterval = &value
	return b
}

// WithRepeatInterval sets the RepeatInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RepeatInterval field is set to the value of the last call.
func (b *RouteApplyConfiguration) WithRepeatInterval(value string) *RouteApplyConfiguration {
	b.RepeatInterval = &value
	return b
}

// WithMatchers adds the given value to the Matchers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Matchers field.
func (b *RouteApplyConfiguration) WithMatchers(values ...*MatcherApplyConfiguration) *RouteApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMatchers")
		}
		b.Matchers = append(b.Matchers, *values[i])
	}
	return b
}

// WithContinue sets the Continue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Continue field is set to the value of the last call.
func (b *RouteApplyConfiguration) WithContinue(value bool) *RouteApplyConfiguration {
	b.Continue = &value
	return b
}

// WithRoutes adds the given value to the Routes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Routes field.
func (b *RouteApplyConfiguration) WithRoutes(values ...v1.JSON) *RouteApplyConfiguration {
	for i := range values {
		b.Routes = append(b.Routes, values[i])
	}
	return b
}

// WithMuteTimeIntervals adds the given value to the MuteTimeIntervals field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MuteTimeIntervals field.
func (b *RouteApplyConfiguration) WithMuteTimeIntervals(values ...string) *RouteApplyConfiguration {
	for i := range values {
		b.MuteTimeIntervals = append(b.MuteTimeIntervals, values[i])
	}
	return b
}

// WithActiveTimeIntervals adds the given value to the ActiveTimeIntervals field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ActiveTimeIntervals field.
func (b *RouteApplyConfiguration) WithActiveTimeIntervals(values ...string) *RouteApplyConfiguration {
	for i := range values {
		b.ActiveTimeIntervals = append(b.ActiveTimeIntervals, values[i])
	}
	return b
}
