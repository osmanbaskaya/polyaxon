// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/base.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type KV struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *any.Any `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cadb02d0b0a23f3, []int{0}
}

func (m *KV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KV.Unmarshal(m, b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KV.Marshal(b, m, deterministic)
}
func (m *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(m, src)
}
func (m *KV) XXX_Size() int {
	return xxx_messageInfo_KV.Size(m)
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KV) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

// Request data with user namespace
type UserResourceRequest struct {
	// User
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResourceRequest) Reset()         { *m = UserResourceRequest{} }
func (m *UserResourceRequest) String() string { return proto.CompactTextString(m) }
func (*UserResourceRequest) ProtoMessage()    {}
func (*UserResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cadb02d0b0a23f3, []int{1}
}

func (m *UserResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResourceRequest.Unmarshal(m, b)
}
func (m *UserResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResourceRequest.Marshal(b, m, deterministic)
}
func (m *UserResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResourceRequest.Merge(m, src)
}
func (m *UserResourceRequest) XXX_Size() int {
	return xxx_messageInfo_UserResourceRequest.Size(m)
}
func (m *UserResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserResourceRequest proto.InternalMessageInfo

func (m *UserResourceRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

// Request data with owner namespace
type OwnerResourceRequest struct {
	// Owner of the namespace
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OwnerResourceRequest) Reset()         { *m = OwnerResourceRequest{} }
func (m *OwnerResourceRequest) String() string { return proto.CompactTextString(m) }
func (*OwnerResourceRequest) ProtoMessage()    {}
func (*OwnerResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cadb02d0b0a23f3, []int{2}
}

func (m *OwnerResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnerResourceRequest.Unmarshal(m, b)
}
func (m *OwnerResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnerResourceRequest.Marshal(b, m, deterministic)
}
func (m *OwnerResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerResourceRequest.Merge(m, src)
}
func (m *OwnerResourceRequest) XXX_Size() int {
	return xxx_messageInfo_OwnerResourceRequest.Size(m)
}
func (m *OwnerResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerResourceRequest proto.InternalMessageInfo

func (m *OwnerResourceRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

// Request data to get/delete entity
type OwnerEntityResourceRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Unique integer identifier of the entity
	Uuid                 string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OwnerEntityResourceRequest) Reset()         { *m = OwnerEntityResourceRequest{} }
func (m *OwnerEntityResourceRequest) String() string { return proto.CompactTextString(m) }
func (*OwnerEntityResourceRequest) ProtoMessage()    {}
func (*OwnerEntityResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cadb02d0b0a23f3, []int{3}
}

func (m *OwnerEntityResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnerEntityResourceRequest.Unmarshal(m, b)
}
func (m *OwnerEntityResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnerEntityResourceRequest.Marshal(b, m, deterministic)
}
func (m *OwnerEntityResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerEntityResourceRequest.Merge(m, src)
}
func (m *OwnerEntityResourceRequest) XXX_Size() int {
	return xxx_messageInfo_OwnerEntityResourceRequest.Size(m)
}
func (m *OwnerEntityResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerEntityResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerEntityResourceRequest proto.InternalMessageInfo

func (m *OwnerEntityResourceRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *OwnerEntityResourceRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

// Request data with owner/project namespace
type ProjectResourceRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Project under namesapce
	Project              string   `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectResourceRequest) Reset()         { *m = ProjectResourceRequest{} }
func (m *ProjectResourceRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectResourceRequest) ProtoMessage()    {}
func (*ProjectResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cadb02d0b0a23f3, []int{4}
}

func (m *ProjectResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectResourceRequest.Unmarshal(m, b)
}
func (m *ProjectResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectResourceRequest.Marshal(b, m, deterministic)
}
func (m *ProjectResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectResourceRequest.Merge(m, src)
}
func (m *ProjectResourceRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectResourceRequest.Size(m)
}
func (m *ProjectResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectResourceRequest proto.InternalMessageInfo

func (m *ProjectResourceRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ProjectResourceRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

type Uuids struct {
	Uuids                []string `protobuf:"bytes,1,rep,name=uuids,proto3" json:"uuids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Uuids) Reset()         { *m = Uuids{} }
func (m *Uuids) String() string { return proto.CompactTextString(m) }
func (*Uuids) ProtoMessage()    {}
func (*Uuids) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cadb02d0b0a23f3, []int{5}
}

func (m *Uuids) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Uuids.Unmarshal(m, b)
}
func (m *Uuids) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Uuids.Marshal(b, m, deterministic)
}
func (m *Uuids) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Uuids.Merge(m, src)
}
func (m *Uuids) XXX_Size() int {
	return xxx_messageInfo_Uuids.Size(m)
}
func (m *Uuids) XXX_DiscardUnknown() {
	xxx_messageInfo_Uuids.DiscardUnknown(m)
}

var xxx_messageInfo_Uuids proto.InternalMessageInfo

func (m *Uuids) GetUuids() []string {
	if m != nil {
		return m.Uuids
	}
	return nil
}

// Request to act on multiple entities under project
type ProjectResourceUuidsBodyRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Project under namesapce
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Uuids of the entities
	Uuids                *Uuids   `protobuf:"bytes,3,opt,name=uuids,proto3" json:"uuids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectResourceUuidsBodyRequest) Reset()         { *m = ProjectResourceUuidsBodyRequest{} }
func (m *ProjectResourceUuidsBodyRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectResourceUuidsBodyRequest) ProtoMessage()    {}
func (*ProjectResourceUuidsBodyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cadb02d0b0a23f3, []int{6}
}

func (m *ProjectResourceUuidsBodyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectResourceUuidsBodyRequest.Unmarshal(m, b)
}
func (m *ProjectResourceUuidsBodyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectResourceUuidsBodyRequest.Marshal(b, m, deterministic)
}
func (m *ProjectResourceUuidsBodyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectResourceUuidsBodyRequest.Merge(m, src)
}
func (m *ProjectResourceUuidsBodyRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectResourceUuidsBodyRequest.Size(m)
}
func (m *ProjectResourceUuidsBodyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectResourceUuidsBodyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectResourceUuidsBodyRequest proto.InternalMessageInfo

func (m *ProjectResourceUuidsBodyRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ProjectResourceUuidsBodyRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ProjectResourceUuidsBodyRequest) GetUuids() *Uuids {
	if m != nil {
		return m.Uuids
	}
	return nil
}

// Request data to get/delete entity
type ProjectEntityResourceRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Project where the experiement will be assigned
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Unique integer identifier of the entity
	Uuid                 string   `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectEntityResourceRequest) Reset()         { *m = ProjectEntityResourceRequest{} }
func (m *ProjectEntityResourceRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectEntityResourceRequest) ProtoMessage()    {}
func (*ProjectEntityResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cadb02d0b0a23f3, []int{7}
}

func (m *ProjectEntityResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectEntityResourceRequest.Unmarshal(m, b)
}
func (m *ProjectEntityResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectEntityResourceRequest.Marshal(b, m, deterministic)
}
func (m *ProjectEntityResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectEntityResourceRequest.Merge(m, src)
}
func (m *ProjectEntityResourceRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectEntityResourceRequest.Size(m)
}
func (m *ProjectEntityResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectEntityResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectEntityResourceRequest proto.InternalMessageInfo

func (m *ProjectEntityResourceRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ProjectEntityResourceRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ProjectEntityResourceRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

// Request list data with user namespace
type UserResouceListRequest struct {
	// User
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Pagination offset
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Limit size
	Limit int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// Sort to order the search
	Sort string `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	// Query filter the search search
	Query                string   `protobuf:"bytes,5,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResouceListRequest) Reset()         { *m = UserResouceListRequest{} }
func (m *UserResouceListRequest) String() string { return proto.CompactTextString(m) }
func (*UserResouceListRequest) ProtoMessage()    {}
func (*UserResouceListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cadb02d0b0a23f3, []int{8}
}

func (m *UserResouceListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResouceListRequest.Unmarshal(m, b)
}
func (m *UserResouceListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResouceListRequest.Marshal(b, m, deterministic)
}
func (m *UserResouceListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResouceListRequest.Merge(m, src)
}
func (m *UserResouceListRequest) XXX_Size() int {
	return xxx_messageInfo_UserResouceListRequest.Size(m)
}
func (m *UserResouceListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResouceListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserResouceListRequest proto.InternalMessageInfo

func (m *UserResouceListRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserResouceListRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *UserResouceListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *UserResouceListRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *UserResouceListRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

// Request list data with owner namespace
type OwnerResouceListRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Pagination offset
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Limit size
	Limit int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// Sort to order the search
	Sort string `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	// Query filter the search search
	Query                string   `protobuf:"bytes,5,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OwnerResouceListRequest) Reset()         { *m = OwnerResouceListRequest{} }
func (m *OwnerResouceListRequest) String() string { return proto.CompactTextString(m) }
func (*OwnerResouceListRequest) ProtoMessage()    {}
func (*OwnerResouceListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cadb02d0b0a23f3, []int{9}
}

func (m *OwnerResouceListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnerResouceListRequest.Unmarshal(m, b)
}
func (m *OwnerResouceListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnerResouceListRequest.Marshal(b, m, deterministic)
}
func (m *OwnerResouceListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerResouceListRequest.Merge(m, src)
}
func (m *OwnerResouceListRequest) XXX_Size() int {
	return xxx_messageInfo_OwnerResouceListRequest.Size(m)
}
func (m *OwnerResouceListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerResouceListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerResouceListRequest proto.InternalMessageInfo

func (m *OwnerResouceListRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *OwnerResouceListRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *OwnerResouceListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *OwnerResouceListRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *OwnerResouceListRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

// Request list data with owner/project namespace
type ProjectResourceListRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Project under namesapce
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Pagination offset
	Offset int32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	// Limit size
	Limit int32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// Sort to order the search
	Sort string `protobuf:"bytes,5,opt,name=sort,proto3" json:"sort,omitempty"`
	// Query filter the search search
	Query                string   `protobuf:"bytes,6,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectResourceListRequest) Reset()         { *m = ProjectResourceListRequest{} }
func (m *ProjectResourceListRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectResourceListRequest) ProtoMessage()    {}
func (*ProjectResourceListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cadb02d0b0a23f3, []int{10}
}

func (m *ProjectResourceListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectResourceListRequest.Unmarshal(m, b)
}
func (m *ProjectResourceListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectResourceListRequest.Marshal(b, m, deterministic)
}
func (m *ProjectResourceListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectResourceListRequest.Merge(m, src)
}
func (m *ProjectResourceListRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectResourceListRequest.Size(m)
}
func (m *ProjectResourceListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectResourceListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectResourceListRequest proto.InternalMessageInfo

func (m *ProjectResourceListRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ProjectResourceListRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ProjectResourceListRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ProjectResourceListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ProjectResourceListRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *ProjectResourceListRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

// Params for requesting an artifact
type RunArtifactRequest struct {
	// Owner of the namespace
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// Project where the experiement will be assigned
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Unique integer identifier of the entity
	Uuid string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Artifact filepath
	Path string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	// Artifact logging step
	Step int32 `protobuf:"varint,5,opt,name=step,proto3" json:"step,omitempty"`
	// Artifact type
	Type                 string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunArtifactRequest) Reset()         { *m = RunArtifactRequest{} }
func (m *RunArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*RunArtifactRequest) ProtoMessage()    {}
func (*RunArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cadb02d0b0a23f3, []int{11}
}

func (m *RunArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunArtifactRequest.Unmarshal(m, b)
}
func (m *RunArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunArtifactRequest.Marshal(b, m, deterministic)
}
func (m *RunArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunArtifactRequest.Merge(m, src)
}
func (m *RunArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_RunArtifactRequest.Size(m)
}
func (m *RunArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunArtifactRequest proto.InternalMessageInfo

func (m *RunArtifactRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RunArtifactRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *RunArtifactRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *RunArtifactRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *RunArtifactRequest) GetStep() int32 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *RunArtifactRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*KV)(nil), "v1.KV")
	proto.RegisterType((*UserResourceRequest)(nil), "v1.UserResourceRequest")
	proto.RegisterType((*OwnerResourceRequest)(nil), "v1.OwnerResourceRequest")
	proto.RegisterType((*OwnerEntityResourceRequest)(nil), "v1.OwnerEntityResourceRequest")
	proto.RegisterType((*ProjectResourceRequest)(nil), "v1.ProjectResourceRequest")
	proto.RegisterType((*Uuids)(nil), "v1.Uuids")
	proto.RegisterType((*ProjectResourceUuidsBodyRequest)(nil), "v1.ProjectResourceUuidsBodyRequest")
	proto.RegisterType((*ProjectEntityResourceRequest)(nil), "v1.ProjectEntityResourceRequest")
	proto.RegisterType((*UserResouceListRequest)(nil), "v1.UserResouceListRequest")
	proto.RegisterType((*OwnerResouceListRequest)(nil), "v1.OwnerResouceListRequest")
	proto.RegisterType((*ProjectResourceListRequest)(nil), "v1.ProjectResourceListRequest")
	proto.RegisterType((*RunArtifactRequest)(nil), "v1.RunArtifactRequest")
}

func init() { proto.RegisterFile("v1/base.proto", fileDescriptor_7cadb02d0b0a23f3) }

var fileDescriptor_7cadb02d0b0a23f3 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xdf, 0x6a, 0x13, 0x41,
	0x14, 0xc6, 0x99, 0x6c, 0x26, 0xd2, 0x23, 0x82, 0x8c, 0x21, 0xae, 0x41, 0x69, 0x98, 0xab, 0x2a,
	0xb2, 0x21, 0xfa, 0x04, 0x2d, 0x28, 0x82, 0x82, 0x32, 0x50, 0xef, 0x37, 0xc9, 0x49, 0x1d, 0x8d,
	0x3b, 0xdb, 0xf9, 0x13, 0xd9, 0x7b, 0xc1, 0x17, 0xf0, 0x11, 0x7c, 0x50, 0x99, 0x33, 0xbb, 0xb6,
	0xd6, 0xad, 0x56, 0xb1, 0x77, 0xe7, 0x4c, 0xbe, 0xf9, 0xbe, 0x5f, 0xce, 0x19, 0x16, 0x6e, 0xed,
	0x16, 0xf3, 0x65, 0xe9, 0xb0, 0xa8, 0xad, 0xf1, 0x46, 0x0c, 0x76, 0x8b, 0xe9, 0xbd, 0x13, 0x63,
	0x4e, 0xb6, 0x38, 0xa7, 0x93, 0x65, 0xd8, 0xcc, 0xcb, 0xaa, 0x49, 0x3f, 0xcb, 0x23, 0x18, 0xbc,
	0x7c, 0x2b, 0x6e, 0x43, 0xf6, 0x01, 0x9b, 0x9c, 0xcd, 0xd8, 0xc1, 0x9e, 0x8a, 0xa5, 0x78, 0x04,
	0x7c, 0x57, 0x6e, 0x03, 0xe6, 0x83, 0x19, 0x3b, 0xb8, 0xf9, 0x64, 0x5c, 0x24, 0x8b, 0xa2, 0xb3,
	0x28, 0x0e, 0xab, 0x46, 0x25, 0x89, 0x7c, 0x08, 0x77, 0x8e, 0x1d, 0x5a, 0x85, 0xce, 0x04, 0xbb,
	0x42, 0x85, 0xa7, 0x01, 0x9d, 0x17, 0x02, 0x86, 0xc1, 0xa1, 0x6d, 0x5d, 0xa9, 0x96, 0x8f, 0x61,
	0xfc, 0xfa, 0x53, 0xf5, 0xab, 0x76, 0x0c, 0xdc, 0xc4, 0xf3, 0x56, 0x9c, 0x1a, 0xf9, 0x1c, 0xa6,
	0xa4, 0x7e, 0x56, 0x79, 0xed, 0x9b, 0x2b, 0xdd, 0xa1, 0xd4, 0xa0, 0xd7, 0xc4, 0x1d, 0x53, 0x83,
	0x5e, 0xcb, 0x17, 0x30, 0x79, 0x63, 0xcd, 0x7b, 0x5c, 0xf9, 0xab, 0x79, 0xe4, 0x70, 0xa3, 0x4e,
	0xfa, 0xd6, 0xa6, 0x6b, 0xe5, 0x03, 0xe0, 0xc7, 0x41, 0xaf, 0x5d, 0xbc, 0x18, 0xad, 0x5d, 0xce,
	0x66, 0x59, 0xbc, 0x48, 0x8d, 0xb4, 0xb0, 0x7f, 0x21, 0x88, 0xd4, 0x47, 0x66, 0xdd, 0xfc, 0x63,
	0xa2, 0xd8, 0xef, 0x82, 0x32, 0x5a, 0xc4, 0x5e, 0xb1, 0x5b, 0x14, 0x64, 0xda, 0x65, 0x2e, 0xe1,
	0x7e, 0x9b, 0xf9, 0x37, 0x63, 0xba, 0x3c, 0xb0, 0x1b, 0x60, 0x76, 0x6e, 0x80, 0x9f, 0x19, 0x4c,
	0x7e, 0xac, 0x78, 0x85, 0xaf, 0xb4, 0xf3, 0xbf, 0xd9, 0xb2, 0x98, 0xc0, 0xc8, 0x6c, 0x36, 0x0e,
	0x93, 0x37, 0x57, 0x6d, 0x17, 0x51, 0xb6, 0xfa, 0xa3, 0xf6, 0xe4, 0xcd, 0x55, 0x6a, 0xa2, 0x83,
	0x33, 0xd6, 0xe7, 0xc3, 0xe4, 0x10, 0xeb, 0xa8, 0x3c, 0x0d, 0x68, 0x9b, 0x9c, 0x27, 0x68, 0x6a,
	0xe4, 0x17, 0x06, 0x77, 0xcf, 0x9e, 0xcf, 0xcf, 0x1c, 0xfd, 0x7f, 0xf3, 0xba, 0x48, 0xbe, 0x31,
	0x98, 0x5e, 0xd8, 0xf4, 0x9f, 0x61, 0x2e, 0x9f, 0xf9, 0x19, 0x66, 0xd6, 0x8f, 0x39, 0xec, 0xc3,
	0xe4, 0x7d, 0x98, 0xa3, 0xf3, 0x98, 0x5f, 0x19, 0x08, 0x15, 0xaa, 0x43, 0xeb, 0xf5, 0xa6, 0x5c,
	0xf9, 0xff, 0xf8, 0x24, 0xe2, 0x59, 0x5d, 0xfa, 0x77, 0xdd, 0xac, 0x62, 0x4d, 0x60, 0x1e, 0x6b,
	0x02, 0xe3, 0x8a, 0xea, 0x78, 0xe6, 0x9b, 0x1a, 0x5b, 0x2e, 0xaa, 0x97, 0x23, 0xfa, 0x8a, 0x3c,
	0xfd, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x83, 0x1e, 0x01, 0xab, 0x04, 0x00, 0x00,
}