package models

import (
	"encoding/json"
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

type AgentTypeV1Alpha struct {
	ApiVersion string                   `json:"apiVersion,omitempty" yaml:"apiVersion"`
	Kind       string                   `json:"kind,omitempty" yaml:"kind"`
	Metadata   AgentTypeV1AlphaMetadata `json:"metadata" yaml:"metadata"`
	Status     AgentTypeV1AlphaStatus   `json:"status" yaml:"status"`
}

type AgentTypeV1AlphaMetadata struct {
	Name       string      `json:"name,omitempty" yaml:"name,omitempty"`
	CreateTime json.Number `json:"create_time,omitempty" yaml:"create_time,omitempty"`
	UpdateTime json.Number `json:"update_time,omitempty" yaml:"update_time,omitempty"`
}

type AgentTypeV1AlphaStatus struct {
	TotalAgentCount   int    `json:"total_agent_count,omitempty" yaml:"total_agent_count,omitempty"`
	RegistrationToken string `json:"registration_token,omitempty" yaml:"registration_token,omitempty"`
}

func NewAgentTypeV1Alpha(name string) AgentTypeV1Alpha {
	a := AgentTypeV1Alpha{}
	a.Metadata.Name = name
	a.setApiVersionAndKind()
	return a
}

func NewAgentTypeV1AlphaFromJson(data []byte) (*AgentTypeV1Alpha, error) {
	a := AgentTypeV1Alpha{}

	err := json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}

	a.setApiVersionAndKind()
	return &a, nil
}

func NewAgentTypeV1AlphaFromYaml(data []byte) (*AgentTypeV1Alpha, error) {
	a := AgentTypeV1Alpha{}

	err := yaml.UnmarshalStrict(data, &a)
	if err != nil {
		return nil, err
	}

	a.setApiVersionAndKind()
	return &a, nil
}

func (s *AgentTypeV1Alpha) setApiVersionAndKind() {
	s.ApiVersion = "v1alpha"
	s.Kind = "SelfHostedAgentType"
}

func (s *AgentTypeV1Alpha) ObjectName() string {
	return fmt.Sprintf("SelfHostedAgentType/%s", s.Metadata.Name)
}

func (s *AgentTypeV1Alpha) ToJson() ([]byte, error) {
	return json.Marshal(s)
}

func (s *AgentTypeV1Alpha) ToYaml() ([]byte, error) {
	return yaml.Marshal(s)
}
