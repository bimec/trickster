/*
 * Copyright 2018 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package testutil provides utilities for running Mockser as an HTTP server during unit testing
package testutil

import (
	"net/http/httptest"

	"github.com/tricksterproxy/mockster/pkg/routes"
)

// NewTestServer launches a Test Prometheus Server (for unit testing)
func NewTestServer() *httptest.Server {
	return httptest.NewServer(routes.GetRouter())
}
