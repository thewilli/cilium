// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v4.23.4
// source: envoy/extensions/http/cache/file_system_http_cache/v3/file_system_http_cache.proto

package file_system_http_cachev3

import (
	v3 "github.com/cilium/proxy/go/envoy/extensions/common/async_files/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configuration for a cache implementation that caches in the local file system.
//
// By default this cache uses a least-recently-used eviction strategy.
//
// For implementation details, see `DESIGN.md <https://github.com/envoyproxy/envoy/blob/main/source/extensions/http/cache/file_system_http_cache/DESIGN.md>`_.
// [#next-free-field: 11]
type FileSystemHttpCacheConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration of a manager for how the file system is used asynchronously.
	ManagerConfig *v3.AsyncFileManagerConfig `protobuf:"bytes,1,opt,name=manager_config,json=managerConfig,proto3" json:"manager_config,omitempty"`
	// Path at which the cache files will be stored.
	//
	// This also doubles as the unique identifier for a cache, so a cache can be shared
	// between different routes, or separate paths can be used to specify separate caches.
	//
	// If the same “cache_path“ is used in more than one “CacheConfig“, the rest of the
	// “FileSystemHttpCacheConfig“ must also match, and will refer to the same cache
	// instance.
	CachePath string `protobuf:"bytes,2,opt,name=cache_path,json=cachePath,proto3" json:"cache_path,omitempty"`
	// The maximum size of the cache in bytes - when reached, cache eviction is triggered.
	//
	// This is measured as the sum of file sizes, such that it includes headers, trailers,
	// and metadata, but does not include e.g. file system overhead and block size padding.
	//
	// If unset there is no limit except file system failure.
	MaxCacheSizeBytes *wrapperspb.UInt64Value `protobuf:"bytes,3,opt,name=max_cache_size_bytes,json=maxCacheSizeBytes,proto3" json:"max_cache_size_bytes,omitempty"`
	// The maximum size of a cache entry in bytes - larger responses will not be cached.
	//
	// This is measured as the file size for the cache entry, such that it includes
	// headers, trailers, and metadata.
	//
	// If unset there is no limit.
	//
	// [#not-implemented-hide:]
	MaxIndividualCacheEntrySizeBytes *wrapperspb.UInt64Value `protobuf:"bytes,4,opt,name=max_individual_cache_entry_size_bytes,json=maxIndividualCacheEntrySizeBytes,proto3" json:"max_individual_cache_entry_size_bytes,omitempty"`
	// The maximum number of cache entries - when reached, cache eviction is triggered.
	//
	// If unset there is no limit.
	MaxCacheEntryCount *wrapperspb.UInt64Value `protobuf:"bytes,5,opt,name=max_cache_entry_count,json=maxCacheEntryCount,proto3" json:"max_cache_entry_count,omitempty"`
	// A number of folders into which to subdivide the cache.
	//
	// Setting this can help with performance in file systems where a large number of inodes
	// in a single branch degrades performance. The optimal value in that case would be
	// “sqrt(expected_cache_entry_count)“.
	//
	// On file systems that perform well with many inodes, the default value of 1 should be used.
	//
	// [#not-implemented-hide:]
	CacheSubdivisions uint32 `protobuf:"varint,6,opt,name=cache_subdivisions,json=cacheSubdivisions,proto3" json:"cache_subdivisions,omitempty"`
	// The amount of the maximum cache size or count to evict when cache eviction is
	// triggered. For example, if “max_cache_size_bytes“ is 10000000 and “evict_fraction“
	// is 0.2, then when the cache exceeds 10MB, entries will be evicted until the cache size is
	// less than or equal to 8MB.
	//
	// The default value of 0 means when the cache exceeds 10MB, entries will be evicted only
	// until the cache is less than or equal to 10MB.
	//
	// Evicting a larger fraction will mean the eviction thread will run less often (sparing
	// CPU load) at the cost of more cache misses due to the extra evicted entries.
	//
	// [#not-implemented-hide:]
	EvictFraction float32 `protobuf:"fixed32,7,opt,name=evict_fraction,json=evictFraction,proto3" json:"evict_fraction,omitempty"`
	// The longest amount of time to wait before running a cache eviction pass. An eviction
	// pass may not necessarily remove any files, but it will update the cache state to match
	// the on-disk state. This can be important if multiple instances are accessing the same
	// cache in parallel. (e.g. if two instances each independently added non-overlapping 10MB
	// of content to a cache with a 15MB limit, neither instance would be aware that the limit
	// was exceeded without this synchronizing pass.)
	//
	// If an eviction pass has not happened within this duration, the eviction thread will
	// be awoken and perform an eviction pass.
	//
	// If unset, there will be no eviction passes except those triggered by cache limits.
	//
	// [#not-implemented-hide:]
	MaxEvictionPeriod *durationpb.Duration `protobuf:"bytes,8,opt,name=max_eviction_period,json=maxEvictionPeriod,proto3" json:"max_eviction_period,omitempty"`
	// The shortest amount of time between cache eviction passes. This can be used to reduce
	// eviction churn, if your cache max size can be flexible. If a cache eviction pass already
	// occurred more recently than this period when another would be triggered, that new
	// pass is cancelled.
	//
	// This means the cache can potentially grow beyond “max_cache_size_bytes“ by as much as
	// can be written within the duration specified.
	//
	// Generally you would use *either* “min_eviction_period“ *or* “evict_fraction“ to
	// reduce churn. Both together will work but since they're both aiming for the same goal,
	// it's simpler not to.
	//
	// [#not-implemented-hide:]
	MinEvictionPeriod *durationpb.Duration `protobuf:"bytes,9,opt,name=min_eviction_period,json=minEvictionPeriod,proto3" json:"min_eviction_period,omitempty"`
	// If true, and the cache path does not exist, attempt to create the cache path, including
	// any missing directories leading up to it. On failure, the config is rejected.
	//
	// If false, and the cache path does not exist, the config is rejected.
	//
	// [#not-implemented-hide:]
	CreateCachePath bool `protobuf:"varint,10,opt,name=create_cache_path,json=createCachePath,proto3" json:"create_cache_path,omitempty"`
}

func (x *FileSystemHttpCacheConfig) Reset() {
	*x = FileSystemHttpCacheConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSystemHttpCacheConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSystemHttpCacheConfig) ProtoMessage() {}

func (x *FileSystemHttpCacheConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSystemHttpCacheConfig.ProtoReflect.Descriptor instead.
func (*FileSystemHttpCacheConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescGZIP(), []int{0}
}

func (x *FileSystemHttpCacheConfig) GetManagerConfig() *v3.AsyncFileManagerConfig {
	if x != nil {
		return x.ManagerConfig
	}
	return nil
}

func (x *FileSystemHttpCacheConfig) GetCachePath() string {
	if x != nil {
		return x.CachePath
	}
	return ""
}

func (x *FileSystemHttpCacheConfig) GetMaxCacheSizeBytes() *wrapperspb.UInt64Value {
	if x != nil {
		return x.MaxCacheSizeBytes
	}
	return nil
}

func (x *FileSystemHttpCacheConfig) GetMaxIndividualCacheEntrySizeBytes() *wrapperspb.UInt64Value {
	if x != nil {
		return x.MaxIndividualCacheEntrySizeBytes
	}
	return nil
}

func (x *FileSystemHttpCacheConfig) GetMaxCacheEntryCount() *wrapperspb.UInt64Value {
	if x != nil {
		return x.MaxCacheEntryCount
	}
	return nil
}

func (x *FileSystemHttpCacheConfig) GetCacheSubdivisions() uint32 {
	if x != nil {
		return x.CacheSubdivisions
	}
	return 0
}

func (x *FileSystemHttpCacheConfig) GetEvictFraction() float32 {
	if x != nil {
		return x.EvictFraction
	}
	return 0
}

func (x *FileSystemHttpCacheConfig) GetMaxEvictionPeriod() *durationpb.Duration {
	if x != nil {
		return x.MaxEvictionPeriod
	}
	return nil
}

func (x *FileSystemHttpCacheConfig) GetMinEvictionPeriod() *durationpb.Duration {
	if x != nil {
		return x.MinEvictionPeriod
	}
	return nil
}

func (x *FileSystemHttpCacheConfig) GetCreateCachePath() bool {
	if x != nil {
		return x.CreateCachePath
	}
	return false
}

var File_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto protoreflect.FileDescriptor

var file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDesc = []byte{
	0x0a, 0x52, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x3f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2f, 0x76, 0x33, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x78, 0x64,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x05, 0x0a, 0x19, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x48, 0x74, 0x74, 0x70, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x6f, 0x0a, 0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x26, 0x0a, 0x0a, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x4d, 0x0a, 0x14,
	0x6d, 0x61, 0x78, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x6d, 0x0a, 0x25, 0x6d,
	0x61, 0x78, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x20, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x64,
	0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x43, 0x61, 0x63, 0x68, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x15, 0x6d, 0x61,
	0x78, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x63, 0x61, 0x63, 0x68, 0x65, 0x53, 0x75,
	0x62, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x76,
	0x69, 0x63, 0x74, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0d, 0x65, 0x76, 0x69, 0x63, 0x74, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x49, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x76, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x45, 0x76,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x49, 0x0a, 0x13,
	0x6d, 0x69, 0x6e, 0x5f, 0x65, 0x76, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6d, 0x69, 0x6e, 0x45, 0x76, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x42, 0xe8, 0x01, 0x0a, 0x43, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x18, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x48, 0x74, 0x74, 0x70, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x75, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67,
	0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x2f, 0x76, 0x33, 0x3b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x76, 0x33, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescOnce sync.Once
	file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescData = file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDesc
)

func file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescGZIP() []byte {
	file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescData)
	})
	return file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDescData
}

var file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_goTypes = []interface{}{
	(*FileSystemHttpCacheConfig)(nil), // 0: envoy.extensions.http.cache.file_system_http_cache.v3.FileSystemHttpCacheConfig
	(*v3.AsyncFileManagerConfig)(nil), // 1: envoy.extensions.common.async_files.v3.AsyncFileManagerConfig
	(*wrapperspb.UInt64Value)(nil),    // 2: google.protobuf.UInt64Value
	(*durationpb.Duration)(nil),       // 3: google.protobuf.Duration
}
var file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.http.cache.file_system_http_cache.v3.FileSystemHttpCacheConfig.manager_config:type_name -> envoy.extensions.common.async_files.v3.AsyncFileManagerConfig
	2, // 1: envoy.extensions.http.cache.file_system_http_cache.v3.FileSystemHttpCacheConfig.max_cache_size_bytes:type_name -> google.protobuf.UInt64Value
	2, // 2: envoy.extensions.http.cache.file_system_http_cache.v3.FileSystemHttpCacheConfig.max_individual_cache_entry_size_bytes:type_name -> google.protobuf.UInt64Value
	2, // 3: envoy.extensions.http.cache.file_system_http_cache.v3.FileSystemHttpCacheConfig.max_cache_entry_count:type_name -> google.protobuf.UInt64Value
	3, // 4: envoy.extensions.http.cache.file_system_http_cache.v3.FileSystemHttpCacheConfig.max_eviction_period:type_name -> google.protobuf.Duration
	3, // 5: envoy.extensions.http.cache.file_system_http_cache.v3.FileSystemHttpCacheConfig.min_eviction_period:type_name -> google.protobuf.Duration
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_init()
}
func file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_init() {
	if File_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileSystemHttpCacheConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_msgTypes,
	}.Build()
	File_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto = out.File
	file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_rawDesc = nil
	file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_goTypes = nil
	file_envoy_extensions_http_cache_file_system_http_cache_v3_file_system_http_cache_proto_depIdxs = nil
}
