// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by model-api-gen. DO NOT EDIT.

package apis

import (
	time "time"
)

// SAdminSharableVirtualResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SAdminSharableVirtualResourceBase.
type SAdminSharableVirtualResourceBase struct {
	SSharableVirtualResourceBase
	Records string `json:"records"`
}

// SAutoDeleteResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SAutoDeleteResourceBase.
type SAutoDeleteResourceBase struct {
	// 是否跟随资源自动删除
	// example: false
	AutoDelete bool `json:"auto_delete"`
}

// SCertificateResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SCertificateResourceBase.
type SCertificateResourceBase struct {
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
	// derived attributes
	PublicKeyAlgorithm      string    `json:"public_key_algorithm"`
	PublicKeyBitLen         int       `json:"public_key_bit_len"`
	SignatureAlgorithm      string    `json:"signature_algorithm"`
	Fingerprint             string    `json:"fingerprint"`
	NotBefore               time.Time `json:"not_before"`
	NotAfter                time.Time `json:"not_after"`
	CommonName              string    `json:"common_name"`
	SubjectAlternativeNames string    `json:"subject_alternative_names"`
}

// SDistinctField is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SDistinctField.
type SDistinctField struct {
	// 资源类型
	// example: network
	ObjType string `json:"obj_type"`
	// 资源组合ID
	// example: obj_type::key::value
	Id string `json:"id"`
	// Distinct Field
	// exmaple: 部门
	Key string `json:"key"`
	// Distinct Value
	// example: 技术部
	Value string `json:"value"`
}

// SDomainLevelResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SDomainLevelResourceBase.
type SDomainLevelResourceBase struct {
	SStandaloneResourceBase
	SDomainizedResourceBase
	// 归属Domain信息的来源, local: 本地设置, cloud: 从云上同步过来
	// example: local
	DomainSrc string `json:"domain_src"`
}

// SDomainizedResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SDomainizedResourceBase.
type SDomainizedResourceBase struct {
	// 域Id
	DomainId string `json:"domain_id"`
}

// SEnabledResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SEnabledResourceBase.
type SEnabledResourceBase struct {
	// 资源是否启用
	Enabled *bool `json:"enabled,omitempty"`
}

// SEnabledStatusDomainLevelResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SEnabledStatusDomainLevelResourceBase.
type SEnabledStatusDomainLevelResourceBase struct {
	SStatusDomainLevelResourceBase
	SEnabledResourceBase
}

// SEnabledStatusInfrasResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SEnabledStatusInfrasResourceBase.
type SEnabledStatusInfrasResourceBase struct {
	SStatusInfrasResourceBase
	SEnabledResourceBase
}

// SEnabledStatusStandaloneResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SEnabledStatusStandaloneResourceBase.
type SEnabledStatusStandaloneResourceBase struct {
	SStatusStandaloneResourceBase
	SEnabledResourceBase
}

// SEncryptedResource is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SEncryptedResource.
type SEncryptedResource struct {
	// 加密密钥ID
	EncryptKeyId string `json:"encrypt_key_id"`
}

// SExternalizedResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SExternalizedResourceBase.
type SExternalizedResourceBase struct {
	// 云上Id, 对应云上资源自身Id
	ExternalId string `json:"external_id"`
	// 资源导入时间
	ImportedAt time.Time `json:"imported_at"`
	// 资源来源, cloud: 从云上同步下来的资源, local: 从本地创建的资源或资源在本地更改过项目
	Source string `json:"source"`
}

// SI18n is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SI18n.
type SI18n struct {
	SStandaloneAnonResourceBase
	// 资源类型
	// example: network
	ObjType string `json:"obj_type"`
	// 资源ID
	// example: 87321a70-1ecb-422a-8b0c-c9aa632a46a7
	ObjId string `json:"obj_id"`
	// 资源KEY
	// exmaple: name
	KeyName string `json:"key_name"`
	// 资源原始值
	// example: 技术部
	KeyValue string `json:"key_value"`
	// KeyValue 对应中文翻译
	Cn string `json:"cn"`
	// KeyValue 对应英文翻译
	En string `json:"en"`
}

// SInfrasResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SInfrasResourceBase.
type SInfrasResourceBase struct {
	SDomainLevelResourceBase
	SSharableBaseResource
}

// SJointResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SJointResourceBase.
type SJointResourceBase struct {
	SResourceBase
	RowId int64 `json:"row_id"`
}

// SKeystoneCacheObject is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SKeystoneCacheObject.
type SKeystoneCacheObject struct {
	SStandaloneResourceBase
	DomainId string `json:"domain_id"`
	Domain   string `json:"domain"`
}

// SLogBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SLogBase.
type SLogBase struct {
	Id int64 `json:"id"`
}

// SMetadata is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SMetadata.
type SMetadata struct {
	// 资源类型
	// example: network
	ObjType string `json:"obj_type"`
	// 资源ID
	// example: 87321a70-1ecb-422a-8b0c-c9aa632a46a7
	ObjId string `json:"obj_id"`
	// 资源组合ID
	// example: network::87321a70-1ecb-422a-8b0c-c9aa632a46a7
	Id string `json:"id"`
	// 标签KEY
	// exmaple: 部门
	Key string `json:"key"`
	// 标签值
	// example: 技术部
	Value string `json:"value"`
	// 是否被删除
	Deleted bool `json:"deleted"`
}

// SMultiArchResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SMultiArchResourceBase.
type SMultiArchResourceBase struct {
	// 操作系统 CPU 架构
	// example: x86 arm
	OsArch string `json:"os_arch"`
}

// SOpsLog is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SOpsLog.
type SOpsLog struct {
	SLogBase
	ObjType         string    `json:"obj_type"`
	ObjId           string    `json:"obj_id"`
	ObjName         string    `json:"obj_name"`
	Action          string    `json:"action"`
	Notes           string    `json:"notes"`
	ProjectId       string    `json:"tenant_id"`
	Project         string    `json:"tenant"`
	ProjectDomainId string    `json:"project_domain_id"`
	ProjectDomain   string    `json:"project_domain"`
	UserId          string    `json:"user_id"`
	User            string    `json:"user"`
	DomainId        string    `json:"domain_id"`
	Domain          string    `json:"domain"`
	Roles           string    `json:"roles"`
	OpsTime         time.Time `json:"ops_time"`
	OwnerDomainId   string    `json:"owner_domain_id"`
	OwnerProjectId  string    `json:"owner_tenant_id"`
}

// SProjectizedResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SProjectizedResourceBase.
type SProjectizedResourceBase struct {
	SDomainizedResourceBase
	// 项目Id
	ProjectId string `json:"tenant_id"`
}

// SRecordChecksumResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SRecordChecksumResourceBase.
type SRecordChecksumResourceBase struct {
	RecordChecksum string `json:"record_checksum"`
}

// SResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SResourceBase.
type SResourceBase struct {
	// 资源创建时间
	CreatedAt time.Time `json:"created_at"`
	// 资源更新时间
	UpdatedAt time.Time `json:"updated_at"`
	// 资源被更新次数
	UpdateVersion int `json:"update_version"`
	// 资源删除时间
	DeletedAt time.Time `json:"deleted_at"`
	// 资源是否被删除
	Deleted bool `json:"deleted"`
}

// SRole is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SRole.
type SRole struct {
	SKeystoneCacheObject
}

// SScopedResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SScopedResourceBase.
type SScopedResourceBase struct {
	SProjectizedResourceBase
}

// SSharableBaseResource is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SSharableBaseResource.
type SSharableBaseResource struct {
	// 是否共享
	IsPublic bool `json:"is_public"`
	// 默认共享范围
	PublicScope string `json:"public_scope"`
	// 共享设置的来源, local: 本地设置, cloud: 从云上同步过来
	// example: local
	PublicSrc string `json:"public_src"`
}

// SSharableVirtualResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SSharableVirtualResourceBase.
type SSharableVirtualResourceBase struct {
	SVirtualResourceBase
	SSharableBaseResource
}

// SSharedResource is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SSharedResource.
type SSharedResource struct {
	SResourceBase
	Id           int64  `json:"id"`
	ResourceType string `json:"resource_type"`
	ResourceId   string `json:"resource_id"`
	// OwnerProjectId  string `width:"128" charset:"ascii" nullable:"false" index:"true" json:"owner_project_id"`
	TargetProjectId string `json:"target_project_id"`
	TargetType      string `json:"target_type"`
}

// SStandaloneAnonResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SStandaloneAnonResourceBase.
type SStandaloneAnonResourceBase struct {
	SResourceBase
	// 资源UUID
	Id string `json:"id"`
	// 资源描述信息
	Description string `json:"description"`
	// 是否是模拟资源, 部分从公有云上同步的资源并不真实存在, 例如宿主机
	// list 接口默认不会返回这类资源，除非显示指定 is_emulate=true 过滤参数
	IsEmulated bool `json:"is_emulated"`
}

// SStandaloneResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SStandaloneResourceBase.
type SStandaloneResourceBase struct {
	SStandaloneAnonResourceBase
	// 资源名称
	Name string `json:"name"`
}

// SStatusDomainLevelResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SStatusDomainLevelResourceBase.
type SStatusDomainLevelResourceBase struct {
	SDomainLevelResourceBase
	SStatusResourceBase
}

// SStatusDomainLevelUserResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SStatusDomainLevelUserResourceBase.
type SStatusDomainLevelUserResourceBase struct {
	SStatusDomainLevelResourceBase
	// 本地用户Id
	OwnerId string `json:"owner_id"`
}

// SStatusInfrasResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SStatusInfrasResourceBase.
type SStatusInfrasResourceBase struct {
	SInfrasResourceBase
	SStatusResourceBase
}

// SStatusResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SStatusResourceBase.
type SStatusResourceBase struct {
	// 资源状态
	Status string `json:"status"`
	// 操作进度0-100
	Progress float32 `json:"progress"`
}

// SStatusStandaloneResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SStatusStandaloneResourceBase.
type SStatusStandaloneResourceBase struct {
	SStandaloneResourceBase
	SStatusResourceBase
}

// STenant is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.STenant.
type STenant struct {
	SKeystoneCacheObject
}

// SUser is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SUser.
type SUser struct {
	SKeystoneCacheObject
}

// SUserResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SUserResourceBase.
type SUserResourceBase struct {
	SStandaloneResourceBase
	// 本地用户Id
	OwnerId string `json:"owner_id"`
}

// SVirtualJointResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SVirtualJointResourceBase.
type SVirtualJointResourceBase struct {
	SJointResourceBase
}

// SVirtualResourceBase is an autogenerated struct via yunion.io/x/onecloud/pkg/cloudcommon/db.SVirtualResourceBase.
type SVirtualResourceBase struct {
	SStatusStandaloneResourceBase
	SProjectizedResourceBase
	// 云上同步资源是否在本地被更改过配置, local: 更改过, cloud: 未更改过
	// example: local
	ProjectSrc string `json:"project_src"`
	// 是否是系统资源
	IsSystem bool `json:"is_system"`
	// 资源放入回收站时间
	PendingDeletedAt time.Time `json:"pending_deleted_at"`
	// 资源是否处于回收站中
	PendingDeleted bool `json:"pending_deleted"`
	// 资源是否被冻结
	Freezed bool `json:"freezed"`
}
