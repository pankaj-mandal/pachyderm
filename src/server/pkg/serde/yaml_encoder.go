package serde

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"gopkg.in/pachyderm/yaml.v3"
)

type YAMLEncoder struct {
	e *yaml.Encoder

	// OrigName sets whether this YAMLEncoder uses the original (.proto) name of
	// fields when marshalling to protos
	origName bool
}

func EncodeYAML(v interface{}, options ...EncoderOption) ([]byte, error) {
	var buf bytes.Buffer
	e := NewYAMLEncoder(&buf, options...)
	if err := e.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func NewYAMLEncoder(w io.Writer, options ...EncoderOption) *YAMLEncoder {
	e := &YAMLEncoder{e: yaml.NewEncoder(w)}
	for _, o := range options {
		o(e)
	}
	return e
}

func (e *YAMLEncoder) Encode(v interface{}) error {
	return e.EncodeTransform(v, nil)
}

func (e *YAMLEncoder) EncodeTransform(v interface{}, f func(map[string]interface{}) error) error {
	// Encode to JSON first
	var buf bytes.Buffer
	j := json.NewEncoder(&buf)
	if err := j.Encode(v); err != nil {
		return fmt.Errorf("serialization error while canonicalizing output: %v", err)
	}

	return e.jsonToYAMLTransform(buf.Bytes(), f)
}

func (e *YAMLEncoder) EncodeProto(v proto.Message) error {
	return e.EncodeProtoTransform(v, nil)
}

func (e *YAMLEncoder) EncodeProtoTransform(v proto.Message, f func(map[string]interface{}) error) error {
	// Encode to JSON first
	var buf bytes.Buffer
	m := jsonpb.Marshaler{
		OrigName: e.origName,
	}
	if err := m.Marshal(&buf, v); err != nil {
		return fmt.Errorf("serialization error while canonicalizing output: %v", err)
	}

	return e.jsonToYAMLTransform(buf.Bytes(), f)
}

func (e *YAMLEncoder) jsonToYAMLTransform(intermediateJSON []byte,
	f func(map[string]interface{}) error) error {
	// Unmarshal from JSON to intermediate map ('holder')
	holder := map[string]interface{}{}
	if err := json.Unmarshal(intermediateJSON, &holder); err != nil {
		return fmt.Errorf("deserialization error while canonicalizing output: %v", err)
	}

	// transform 'holder' (e.g. de-stringifying TFJob)
	if f != nil {
		if err := f(holder); err != nil {
			return err
		}
	}

	// Encode 'holder' to YAML
	if err := e.e.Encode(holder); err != nil {
		return fmt.Errorf("serialization error while canonicalizing yaml: %v", err)
	}
	return nil
}
