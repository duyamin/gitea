// Code generated by protoc-gen-go.
// source: Snapshot.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = math.Inf

type SnapshotFileInfo_Type int32

const (
	SnapshotFileInfo_HFILE SnapshotFileInfo_Type = 1
	SnapshotFileInfo_WAL   SnapshotFileInfo_Type = 2
)

var SnapshotFileInfo_Type_name = map[int32]string{
	1: "HFILE",
	2: "WAL",
}
var SnapshotFileInfo_Type_value = map[string]int32{
	"HFILE": 1,
	"WAL":   2,
}

func (x SnapshotFileInfo_Type) Enum() *SnapshotFileInfo_Type {
	p := new(SnapshotFileInfo_Type)
	*p = x
	return p
}
func (x SnapshotFileInfo_Type) String() string {
	return proto1.EnumName(SnapshotFileInfo_Type_name, int32(x))
}
func (x *SnapshotFileInfo_Type) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(SnapshotFileInfo_Type_value, data, "SnapshotFileInfo_Type")
	if err != nil {
		return err
	}
	*x = SnapshotFileInfo_Type(value)
	return nil
}

type SnapshotFileInfo struct {
	Type             *SnapshotFileInfo_Type `protobuf:"varint,1,req,name=type,enum=proto.SnapshotFileInfo_Type" json:"type,omitempty"`
	Hfile            *string                `protobuf:"bytes,3,opt,name=hfile" json:"hfile,omitempty"`
	WalServer        *string                `protobuf:"bytes,4,opt,name=wal_server" json:"wal_server,omitempty"`
	WalName          *string                `protobuf:"bytes,5,opt,name=wal_name" json:"wal_name,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *SnapshotFileInfo) Reset()         { *m = SnapshotFileInfo{} }
func (m *SnapshotFileInfo) String() string { return proto1.CompactTextString(m) }
func (*SnapshotFileInfo) ProtoMessage()    {}

func (m *SnapshotFileInfo) GetType() SnapshotFileInfo_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return SnapshotFileInfo_HFILE
}

func (m *SnapshotFileInfo) GetHfile() string {
	if m != nil && m.Hfile != nil {
		return *m.Hfile
	}
	return ""
}

func (m *SnapshotFileInfo) GetWalServer() string {
	if m != nil && m.WalServer != nil {
		return *m.WalServer
	}
	return ""
}

func (m *SnapshotFileInfo) GetWalName() string {
	if m != nil && m.WalName != nil {
		return *m.WalName
	}
	return ""
}

type SnapshotRegionManifest struct {
	Version          *int32                                `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	RegionInfo       *RegionInfo                           `protobuf:"bytes,2,req,name=region_info" json:"region_info,omitempty"`
	FamilyFiles      []*SnapshotRegionManifest_FamilyFiles `protobuf:"bytes,3,rep,name=family_files" json:"family_files,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *SnapshotRegionManifest) Reset()         { *m = SnapshotRegionManifest{} }
func (m *SnapshotRegionManifest) String() string { return proto1.CompactTextString(m) }
func (*SnapshotRegionManifest) ProtoMessage()    {}

func (m *SnapshotRegionManifest) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *SnapshotRegionManifest) GetRegionInfo() *RegionInfo {
	if m != nil {
		return m.RegionInfo
	}
	return nil
}

func (m *SnapshotRegionManifest) GetFamilyFiles() []*SnapshotRegionManifest_FamilyFiles {
	if m != nil {
		return m.FamilyFiles
	}
	return nil
}

type SnapshotRegionManifest_StoreFile struct {
	Name      *string    `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Reference *Reference `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
	// TODO: Add checksums or other fields to verify the file
	FileSize         *uint64 `protobuf:"varint,3,opt,name=file_size" json:"file_size,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SnapshotRegionManifest_StoreFile) Reset()         { *m = SnapshotRegionManifest_StoreFile{} }
func (m *SnapshotRegionManifest_StoreFile) String() string { return proto1.CompactTextString(m) }
func (*SnapshotRegionManifest_StoreFile) ProtoMessage()    {}

func (m *SnapshotRegionManifest_StoreFile) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SnapshotRegionManifest_StoreFile) GetReference() *Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *SnapshotRegionManifest_StoreFile) GetFileSize() uint64 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

type SnapshotRegionManifest_FamilyFiles struct {
	FamilyName       []byte                              `protobuf:"bytes,1,req,name=family_name" json:"family_name,omitempty"`
	StoreFiles       []*SnapshotRegionManifest_StoreFile `protobuf:"bytes,2,rep,name=store_files" json:"store_files,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *SnapshotRegionManifest_FamilyFiles) Reset()         { *m = SnapshotRegionManifest_FamilyFiles{} }
func (m *SnapshotRegionManifest_FamilyFiles) String() string { return proto1.CompactTextString(m) }
func (*SnapshotRegionManifest_FamilyFiles) ProtoMessage()    {}

func (m *SnapshotRegionManifest_FamilyFiles) GetFamilyName() []byte {
	if m != nil {
		return m.FamilyName
	}
	return nil
}

func (m *SnapshotRegionManifest_FamilyFiles) GetStoreFiles() []*SnapshotRegionManifest_StoreFile {
	if m != nil {
		return m.StoreFiles
	}
	return nil
}

type SnapshotDataManifest struct {
	TableSchema      *TableSchema              `protobuf:"bytes,1,req,name=table_schema" json:"table_schema,omitempty"`
	RegionManifests  []*SnapshotRegionManifest `protobuf:"bytes,2,rep,name=region_manifests" json:"region_manifests,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *SnapshotDataManifest) Reset()         { *m = SnapshotDataManifest{} }
func (m *SnapshotDataManifest) String() string { return proto1.CompactTextString(m) }
func (*SnapshotDataManifest) ProtoMessage()    {}

func (m *SnapshotDataManifest) GetTableSchema() *TableSchema {
	if m != nil {
		return m.TableSchema
	}
	return nil
}

func (m *SnapshotDataManifest) GetRegionManifests() []*SnapshotRegionManifest {
	if m != nil {
		return m.RegionManifests
	}
	return nil
}

func init() {
	proto1.RegisterEnum("proto.SnapshotFileInfo_Type", SnapshotFileInfo_Type_name, SnapshotFileInfo_Type_value)
}
