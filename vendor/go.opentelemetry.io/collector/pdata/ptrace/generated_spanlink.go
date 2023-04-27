// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package ptrace

import (
	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/internal/data"
	otlptrace "go.opentelemetry.io/collector/pdata/internal/data/protogen/trace/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// SpanLink is a pointer from the current span to another span in the same trace or in a
// different trace.
// See Link definition in OTLP: https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/trace/v1/trace.proto
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanLink function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanLink struct {
	orig *otlptrace.Span_Link
}

func newSpanLink(orig *otlptrace.Span_Link) SpanLink {
	return SpanLink{orig}
}

// NewSpanLink creates a new empty SpanLink.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewSpanLink() SpanLink {
	return newSpanLink(&otlptrace.Span_Link{})
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms SpanLink) MoveTo(dest SpanLink) {
	*dest.orig = *ms.orig
	*ms.orig = otlptrace.Span_Link{}
}

// TraceID returns the traceid associated with this SpanLink.
func (ms SpanLink) TraceID() pcommon.TraceID {
	return pcommon.TraceID(ms.orig.TraceId)
}

// SetTraceID replaces the traceid associated with this SpanLink.
func (ms SpanLink) SetTraceID(v pcommon.TraceID) {
	ms.orig.TraceId = data.TraceID(v)
}

// SpanID returns the spanid associated with this SpanLink.
func (ms SpanLink) SpanID() pcommon.SpanID {
	return pcommon.SpanID(ms.orig.SpanId)
}

// SetSpanID replaces the spanid associated with this SpanLink.
func (ms SpanLink) SetSpanID(v pcommon.SpanID) {
	ms.orig.SpanId = data.SpanID(v)
}

// TraceState returns the tracestate associated with this SpanLink.
func (ms SpanLink) TraceState() pcommon.TraceState {
	return pcommon.TraceState(internal.NewTraceState(&ms.orig.TraceState))
}

// Attributes returns the Attributes associated with this SpanLink.
func (ms SpanLink) Attributes() pcommon.Map {
	return pcommon.Map(internal.NewMap(&ms.orig.Attributes))
}

// DroppedAttributesCount returns the droppedattributescount associated with this SpanLink.
func (ms SpanLink) DroppedAttributesCount() uint32 {
	return ms.orig.DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this SpanLink.
func (ms SpanLink) SetDroppedAttributesCount(v uint32) {
	ms.orig.DroppedAttributesCount = v
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms SpanLink) CopyTo(dest SpanLink) {
	dest.SetTraceID(ms.TraceID())
	dest.SetSpanID(ms.SpanID())
	ms.TraceState().CopyTo(dest.TraceState())
	ms.Attributes().CopyTo(dest.Attributes())
	dest.SetDroppedAttributesCount(ms.DroppedAttributesCount())
}