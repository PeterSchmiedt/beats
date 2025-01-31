// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package kubernetes

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

// AssetKubernetes returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/kubernetes.
func AssetKubernetes() string {
	return "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"
}
