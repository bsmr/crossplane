/*
Copyright 2018 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gke

import (
	"testing"

	"github.com/go-test/deep"

	"golang.org/x/oauth2/google"
)

func TestNewClusterClient(t *testing.T) {
	type want struct {
		err error
		res *ClusterClient
	}
	tests := []struct {
		name string
		args *google.Credentials
		want want
	}{
		{name: "Test", args: &google.Credentials{}, want: want{res: &ClusterClient{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClusterClient(tt.args)
			if diff := deep.Equal(err, tt.want.err); diff != nil {
				t.Errorf("NewClusterClient() error = %v, want.err %v\n%s", err, tt.want.err, diff)
				return
			}
			if diff := deep.Equal(got, tt.want.res); diff != nil {
				t.Errorf("NewClusterClient() = %v, want %v\n%s", got, tt.want.res, diff)
			}
		})
	}
}
