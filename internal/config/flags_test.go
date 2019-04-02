/**
* Copyright 2018 Comcast Cable Communications Management, LLC
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
* http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package config

import (
	"testing"
)

func TestLoadFlags(t *testing.T) {
	c := NewConfig()
	a := []string{
		"-origin",
		"http://prometheus.example.com:9090",
		"-proxy-port",
		"9091",
		"-metrics-port",
		"9092",
		"-origin-type",
		"prometheus",
	}

	// it should read command line flags
	c.parseFlags("trickster-test", a)
	c.loadFlags()

	if defaultOriginURL != a[1] {
		t.Errorf("wanted \"%s\". got \"%s\".", a[1], defaultOriginURL)
	}
	if defaultOriginType != a[7] {
		t.Errorf("wanted \"%s\". got \"%s\".", a[1], defaultOriginType)
	}
	if c.ProxyServer.ListenPort != 9091 {
		t.Errorf("wanted \"%d\". got \"%d\".", 9091, c.ProxyServer.ListenPort)
	}
	if c.Metrics.ListenPort != 9092 {
		t.Errorf("wanted \"%d\". got \"%d\".", 9092, c.Metrics.ListenPort)
	}
}