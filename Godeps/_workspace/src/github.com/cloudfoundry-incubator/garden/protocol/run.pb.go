// Code generated by protoc-gen-gogo.
// source: run.proto
// DO NOT EDIT!

package warden

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type RunRequest struct {
	Handle           *string                `protobuf:"bytes,1,req,name=handle" json:"handle,omitempty"`
	Script           *string                `protobuf:"bytes,2,req,name=script" json:"script,omitempty"`
	Privileged       *bool                  `protobuf:"varint,3,opt,name=privileged,def=0" json:"privileged,omitempty"`
	Rlimits          *ResourceLimits        `protobuf:"bytes,4,opt,name=rlimits" json:"rlimits,omitempty"`
	Env              []*EnvironmentVariable `protobuf:"bytes,5,rep,name=env" json:"env,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *RunRequest) Reset()         { *m = RunRequest{} }
func (m *RunRequest) String() string { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()    {}

const Default_RunRequest_Privileged bool = false

func (m *RunRequest) GetHandle() string {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return ""
}

func (m *RunRequest) GetScript() string {
	if m != nil && m.Script != nil {
		return *m.Script
	}
	return ""
}

func (m *RunRequest) GetPrivileged() bool {
	if m != nil && m.Privileged != nil {
		return *m.Privileged
	}
	return Default_RunRequest_Privileged
}

func (m *RunRequest) GetRlimits() *ResourceLimits {
	if m != nil {
		return m.Rlimits
	}
	return nil
}

func (m *RunRequest) GetEnv() []*EnvironmentVariable {
	if m != nil {
		return m.Env
	}
	return nil
}

func init() {
}