package v1

import (
	k8smeta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Consumer is a specification for a Consumer resource
type Consumer struct {
	k8smeta.TypeMeta   `json:",inline"`
	k8smeta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConsumerSpec `json:"spec"`
	Status Status       `json:"status"`
}

// ConsumerSpec is the spec for a Consumer resource
type ConsumerSpec struct {
	Servers           []string          `json:"servers"`
	CredentialsSecret CredentialsSecret `json:"credentialsSecret"`

	StreamName     string `json:"streamName"`
	DurableName    string `json:"durableName"`
	DeliverSubject string `json:"deliverSubject"`
	AckPolicy      string `json:"ackPolicy"`
	FilterSubject  string `json:"filterSubject"`
	ReplayPolicy   string `json:"replayPolicy"`
	SampleFreq     string `json:"sampleFreq"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConsumerList is a list of Consumer resources
type ConsumerList struct {
	k8smeta.TypeMeta `json:",inline"`
	k8smeta.ListMeta `json:"metadata"`

	Items []Consumer `json:"items"`
}
