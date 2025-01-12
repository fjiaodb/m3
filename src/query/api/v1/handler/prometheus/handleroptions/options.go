// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package handleroptions

import (
	"time"

	"github.com/m3db/m3/src/x/retry"
)

// PromWriteHandlerForwardingOptions is the forwarding
// options for prometheus write handler.
type PromWriteHandlerForwardingOptions struct {
	// MaxConcurrency is the max parallel forwarding and if zero will be unlimited.
	MaxConcurrency int                                    `yaml:"maxConcurrency"`
	Timeout        time.Duration                          `yaml:"timeout"`
	Retry          *retry.Configuration                   `yaml:"retry"`
	Targets        []PromWriteHandlerForwardTargetOptions `yaml:"targets"`
}

// PromWriteHandlerForwardTargetOptions is a prometheus write
// handler forwarder target.
type PromWriteHandlerForwardTargetOptions struct {
	// URL of the target to send to.
	URL string `yaml:"url"`
	// Method defaults to POST if not set.
	Method string `yaml:"method"`
	// Headers to send along with requests to the target.
	Headers map[string]string `yaml:"headers"`
	// Transform original remote write header to Thanos Tenant `THANOS-TENANT`
	TenantHeader string `yaml:"tenantHeader"`
}
