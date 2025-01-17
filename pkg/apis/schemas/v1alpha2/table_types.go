/*
Copyright 2019 Replicated, Inc.

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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TableSchema struct {
	Postgres *SQLTableSchema `json:"postgres,omitempty" yaml:"postgres,omitempty"`
	Mysql    *SQLTableSchema `json:"mysql,omitempty" yaml:"mysql,omitempty"`
}

// TableSpec defines the desired state of Table
type TableSpec struct {
	Database string       `json:"database" yaml:"database"`
	Name     string       `json:"name" yaml:"name"`
	Requires []string     `json:"requires,omitempty" yaml:"requires,omitempty"`
	Schema   *TableSchema `json:"schema" yaml:"schema"`
}

// TableStatus defines the observed state of Table
type TableStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Table is the Schema for the tables API
// +k8s:openapi-gen=true
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TableSpec   `json:"spec,omitempty"`
	Status TableStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TableList contains a list of Table
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Table `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}
