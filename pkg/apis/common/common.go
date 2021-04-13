/*
Copyright 2018 BlackRock, Inc.

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

package common

import (
	corev1 "k8s.io/api/core/v1"
)

// EventSourceType is the type of event source
type EventSourceType string

// possible event source types
var (
	MinioEvent       EventSourceType = "minio"
	CalendarEvent    EventSourceType = "calendar"
	FileEvent        EventSourceType = "file"
	ResourceEvent    EventSourceType = "resource"
	WebhookEvent     EventSourceType = "webhook"
	AMQPEvent        EventSourceType = "amqp"
	KafkaEvent       EventSourceType = "kafka"
	MQTTEvent        EventSourceType = "mqtt"
	NATSEvent        EventSourceType = "nats"
	SNSEvent         EventSourceType = "sns"
	SQSEvent         EventSourceType = "sqs"
	PubSubEvent      EventSourceType = "pubsub"
	GithubEvent      EventSourceType = "github"
	GitlabEvent      EventSourceType = "gitlab"
	HDFSEvent        EventSourceType = "hdfs"
	SlackEvent       EventSourceType = "slack"
	StorageGridEvent EventSourceType = "storagegrid"
	AzureEventsHub   EventSourceType = "azureEventsHub"
	StripeEvent      EventSourceType = "stripe"
	EmitterEvent     EventSourceType = "emitter"
	RedisEvent       EventSourceType = "redis"
	NSQEvent         EventSourceType = "nsq"
	PulsarEvent      EventSourceType = "pulsar"
	GenericEvent     EventSourceType = "generic"
)

var (
	// RecreateStrategyEventSources refers to the list of event source types
	// that need to use Recreate strategy for its Deployment
	RecreateStrategyEventSources = []EventSourceType{
		AMQPEvent,
		CalendarEvent,
		KafkaEvent,
		PubSubEvent,
		AzureEventsHub,
		NATSEvent,
		MQTTEvent,
		MinioEvent,
		EmitterEvent,
		NSQEvent,
		PulsarEvent,
		RedisEvent,
		ResourceEvent,
		HDFSEvent,
		FileEvent,
		GenericEvent,
	}
)

// TriggerType is the type of trigger
type TriggerType string

// possible trigger types
var (
	OpenWhiskTrigger      TriggerType = "OpenWhisk"
	ArgoWorkflowTrigger   TriggerType = "ArgoWorkflow"
	LambdaTrigger         TriggerType = "Lambda"
	CustomTrigger         TriggerType = "Custom"
	HTTPTrigger           TriggerType = "HTTP"
	KafkaTrigger          TriggerType = "Kafka"
	LogTrigger            TriggerType = "Log"
	NATSTrigger           TriggerType = "NATS"
	SlackTrigger          TriggerType = "Slack"
	K8sTrigger            TriggerType = "Kubernetes"
	AzureEventHubsTrigger TriggerType = "AzureEventHubs"
)

// EventBusType is the type of event bus
type EventBusType string

// possible event bus types
var (
	EventBusNATS EventBusType = "nats"
)

// BasicAuth contains the reference to K8s secrets that holds the username and password
type BasicAuth struct {
	// Username refers to the Kubernetes secret that holds the username required for basic auth.
	Username *corev1.SecretKeySelector `json:"username,omitempty" protobuf:"bytes,1,opt,name=username"`
	// Password refers to the Kubernetes secret that holds the password required for basic auth.
	Password *corev1.SecretKeySelector `json:"password,omitempty" protobuf:"bytes,2,opt,name=password"`
}

// TLSConfig refers to TLS configuration for a client.
type TLSConfig struct {
	// CACertSecret refers to the secret that contains the CA cert
	CACertSecret *corev1.SecretKeySelector `json:"caCertSecret,omitempty" protobuf:"bytes,1,opt,name=caCertSecret"`
	// ClientCertSecret refers to the secret that contains the client cert
	ClientCertSecret *corev1.SecretKeySelector `json:"clientCertSecret,omitempty" protobuf:"bytes,2,opt,name=clientCertSecret"`
	// ClientKeySecret refers to the secret that contains the client key
	ClientKeySecret *corev1.SecretKeySelector `json:"clientKeySecret,omitempty" protobuf:"bytes,3,opt,name=clientKeySecret"`

	// DeprecatedCACertPath refers the file path that contains the CA cert.
	// Deprecated: will be removed in v1.5, use CACertSecret instead
	DeprecatedCACertPath string `json:"caCertPath,omitempty" protobuf:"bytes,4,opt,name=caCertPath"`
	// DeprecatedClientCertPath refers the file path that contains client cert.
	// Deprecated: will be removed in v1.5, use ClientCertSecret instead
	DeprecatedClientCertPath string `json:"clientCertPath,omitempty" protobuf:"bytes,5,opt,name=clientCertPath"`
	// DeprecatedClientKeyPath refers the file path that contains client key.
	// Deprecated: will be removed in v1.5, use ClientKeySecret instead
	DeprecatedClientKeyPath string `json:"clientKeyPath,omitempty" protobuf:"bytes,6,opt,name=clientKeyPath"`
}

type SASLConfig struct {
	// SASLUsername refers to the secret that contains sasl plaintext username
	SASLUsername *corev1.SecretKeySelector `json:"saslUsername,omitempty" protobuf:"bytes,1,opt,name=saslUsername"`
	// SASLPassword refers to the secret that contains tsasl plaintext password
	SASLPassword *corev1.SecretKeySelector `json:"saslPassword,omitempty" protobuf:"bytes,2,opt,name=saslPassword"`
}

// Backoff for an operation
type Backoff struct {
	// The initial duration in nanoseconds or strings like "1s", "3m"
	// +optional
	Duration *Int64OrString `json:"duration,omitempty" protobuf:"varint,1,opt,name=duration"`
	// Duration is multiplied by factor each iteration
	// +optional
	Factor *Amount `json:"factor,omitempty" protobuf:"bytes,2,opt,name=factor"`
	// The amount of jitter applied each iteration
	// +optional
	Jitter *Amount `json:"jitter,omitempty" protobuf:"bytes,3,opt,name=jitter"`
	// Exit with error after this many steps
	// +optional
	Steps int32 `json:"steps,omitempty" protobuf:"varint,4,opt,name=steps"`
}

func (b Backoff) GetSteps() int {
	return int(b.Steps)
}

// Metadata holds the annotations and labels of an event source pod
type Metadata struct {
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,1,rep,name=annotations"`
	Labels      map[string]string `json:"labels,omitempty" protobuf:"bytes,2,rep,name=labels"`
}
