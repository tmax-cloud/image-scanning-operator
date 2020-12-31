/*


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

package v1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

type ScanningStatusType string

const (
	ScanningSuccess ScanningStatusType = "Success"
	ScanningError   ScanningStatusType = "Error"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Vulnerability struct {
	Name          string               `json:"Name,omitempty"`
	NamespaceName string               `json:"NamespaceName,omitempty"`
	Description   string               `json:"Description,omitempty"`
	Link          string               `json:"Link,omitempty"`
	Severity      string               `json:"Severity,omitempty"`
	Metadata      runtime.RawExtension `json:"Metadata,omitempty"`
	FixedBy       string               `json:"FixedBy,omitempty"`
}

type Vulnerabilities []Vulnerability

// ImageScanningSpec defines the desired state of ImageScanning
type ImageScanningSpec struct {
	ImageUrl         string        `json:"imageUrl"`
	AuthUrl          string        `json:"authUrl,omitempty"`
	Insecure         bool          `json:"insecure,omitempty"`
	ForceNonSSL      bool          `json:"forceNonSSL,omitempty"`
	Username         string        `json:"username,omitempty"`
	Password         string        `json:"password,omitempty"`
	Debug            bool          `json:"debug,omitempty"`
	SkipPing         bool          `json:"skipPing,omitempty"`
	TimeOut          time.Duration `json:"timeOut,omitempty"`
	FixableThreshold int           `json:"fixableThreshold,omitempty"`
	Webhook          bool          `json:"webhook,omitempty"`
}

// ImageScanningStatus defines the observed state of ImageScanning
type ImageScanningStatus struct {
	Message         string                     `json:"message,omitempty"`
	Reason          string                     `json:"reason,omitempty"`
	Status          ScanningStatusType         `json:"status,omitempty"`
	Summary         map[string]int             `json:"summary,omitempty"`
	Fatal           []string                   `json:"fatal,omitempty"`
	Vulnerabilities map[string]Vulnerabilities `json:"vulnerabilities,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ImageScanning is the Schema for the imagescannings API
type ImageScanning struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImageScanningSpec   `json:"spec,omitempty"`
	Status ImageScanningStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageScanningList contains a list of ImageScanning
type ImageScanningList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageScanning `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImageScanning{}, &ImageScanningList{})
}
