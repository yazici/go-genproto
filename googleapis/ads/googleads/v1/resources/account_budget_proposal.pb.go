// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/account_budget_proposal.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An account-level budget proposal.
//
// All fields prefixed with 'proposed' may not necessarily be applied directly.
// For example, proposed spending limits may be adjusted before their
// application.  This is true if the 'proposed' field has an 'approved'
// counterpart, e.g. spending limits.
//
// Please note that the proposal type (proposal_type) changes which fields are
// required and which must remain empty.
type AccountBudgetProposal struct {
	// The resource name of the proposal.
	// AccountBudgetProposal resource names have the form:
	//
	//
	// `customers/{customer_id}/accountBudgetProposals/{account_budget_proposal_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the proposal.
	Id *wrappers.Int64Value `protobuf:"bytes,14,opt,name=id,proto3" json:"id,omitempty"`
	// The resource name of the billing setup associated with this proposal.
	BillingSetup *wrappers.StringValue `protobuf:"bytes,2,opt,name=billing_setup,json=billingSetup,proto3" json:"billing_setup,omitempty"`
	// The resource name of the account-level budget associated with this
	// proposal.
	AccountBudget *wrappers.StringValue `protobuf:"bytes,3,opt,name=account_budget,json=accountBudget,proto3" json:"account_budget,omitempty"`
	// The type of this proposal, e.g. END to end the budget associated with this
	// proposal.
	ProposalType enums.AccountBudgetProposalTypeEnum_AccountBudgetProposalType `protobuf:"varint,4,opt,name=proposal_type,json=proposalType,proto3,enum=google.ads.googleads.v1.enums.AccountBudgetProposalTypeEnum_AccountBudgetProposalType" json:"proposal_type,omitempty"`
	// The status of this proposal.
	// When a new proposal is created, the status defaults to PENDING.
	Status enums.AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus `protobuf:"varint,15,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus" json:"status,omitempty"`
	// The name to assign to the account-level budget.
	ProposedName *wrappers.StringValue `protobuf:"bytes,5,opt,name=proposed_name,json=proposedName,proto3" json:"proposed_name,omitempty"`
	// The approved start date time in yyyy-mm-dd hh:mm:ss format.
	ApprovedStartDateTime *wrappers.StringValue `protobuf:"bytes,20,opt,name=approved_start_date_time,json=approvedStartDateTime,proto3" json:"approved_start_date_time,omitempty"`
	// A purchase order number is a value that enables the user to help them
	// reference this budget in their monthly invoices.
	ProposedPurchaseOrderNumber *wrappers.StringValue `protobuf:"bytes,12,opt,name=proposed_purchase_order_number,json=proposedPurchaseOrderNumber,proto3" json:"proposed_purchase_order_number,omitempty"`
	// Notes associated with this budget.
	ProposedNotes *wrappers.StringValue `protobuf:"bytes,13,opt,name=proposed_notes,json=proposedNotes,proto3" json:"proposed_notes,omitempty"`
	// The date time when this account-level budget proposal was created, which is
	// not the same as its approval date time, if applicable.
	CreationDateTime *wrappers.StringValue `protobuf:"bytes,16,opt,name=creation_date_time,json=creationDateTime,proto3" json:"creation_date_time,omitempty"`
	// The date time when this account-level budget was approved, if applicable.
	ApprovalDateTime *wrappers.StringValue `protobuf:"bytes,17,opt,name=approval_date_time,json=approvalDateTime,proto3" json:"approval_date_time,omitempty"`
	// The proposed start date time of the account-level budget, which cannot be
	// in the past.
	//
	// Types that are valid to be assigned to ProposedStartTime:
	//	*AccountBudgetProposal_ProposedStartDateTime
	//	*AccountBudgetProposal_ProposedStartTimeType
	ProposedStartTime isAccountBudgetProposal_ProposedStartTime `protobuf_oneof:"proposed_start_time"`
	// The proposed end date time of the account-level budget, which cannot be in
	// the past.
	//
	// Types that are valid to be assigned to ProposedEndTime:
	//	*AccountBudgetProposal_ProposedEndDateTime
	//	*AccountBudgetProposal_ProposedEndTimeType
	ProposedEndTime isAccountBudgetProposal_ProposedEndTime `protobuf_oneof:"proposed_end_time"`
	// The approved end date time of the account-level budget.
	//
	// Types that are valid to be assigned to ApprovedEndTime:
	//	*AccountBudgetProposal_ApprovedEndDateTime
	//	*AccountBudgetProposal_ApprovedEndTimeType
	ApprovedEndTime isAccountBudgetProposal_ApprovedEndTime `protobuf_oneof:"approved_end_time"`
	// The proposed spending limit.
	//
	// Types that are valid to be assigned to ProposedSpendingLimit:
	//	*AccountBudgetProposal_ProposedSpendingLimitMicros
	//	*AccountBudgetProposal_ProposedSpendingLimitType
	ProposedSpendingLimit isAccountBudgetProposal_ProposedSpendingLimit `protobuf_oneof:"proposed_spending_limit"`
	// The approved spending limit.
	//
	// Types that are valid to be assigned to ApprovedSpendingLimit:
	//	*AccountBudgetProposal_ApprovedSpendingLimitMicros
	//	*AccountBudgetProposal_ApprovedSpendingLimitType
	ApprovedSpendingLimit isAccountBudgetProposal_ApprovedSpendingLimit `protobuf_oneof:"approved_spending_limit"`
	XXX_NoUnkeyedLiteral  struct{}                                      `json:"-"`
	XXX_unrecognized      []byte                                        `json:"-"`
	XXX_sizecache         int32                                         `json:"-"`
}

func (m *AccountBudgetProposal) Reset()         { *m = AccountBudgetProposal{} }
func (m *AccountBudgetProposal) String() string { return proto.CompactTextString(m) }
func (*AccountBudgetProposal) ProtoMessage()    {}
func (*AccountBudgetProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_budget_proposal_522c3e292dda2d1a, []int{0}
}
func (m *AccountBudgetProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBudgetProposal.Unmarshal(m, b)
}
func (m *AccountBudgetProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBudgetProposal.Marshal(b, m, deterministic)
}
func (dst *AccountBudgetProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBudgetProposal.Merge(dst, src)
}
func (m *AccountBudgetProposal) XXX_Size() int {
	return xxx_messageInfo_AccountBudgetProposal.Size(m)
}
func (m *AccountBudgetProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBudgetProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBudgetProposal proto.InternalMessageInfo

func (m *AccountBudgetProposal) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AccountBudgetProposal) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *AccountBudgetProposal) GetBillingSetup() *wrappers.StringValue {
	if m != nil {
		return m.BillingSetup
	}
	return nil
}

func (m *AccountBudgetProposal) GetAccountBudget() *wrappers.StringValue {
	if m != nil {
		return m.AccountBudget
	}
	return nil
}

func (m *AccountBudgetProposal) GetProposalType() enums.AccountBudgetProposalTypeEnum_AccountBudgetProposalType {
	if m != nil {
		return m.ProposalType
	}
	return enums.AccountBudgetProposalTypeEnum_UNSPECIFIED
}

func (m *AccountBudgetProposal) GetStatus() enums.AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus {
	if m != nil {
		return m.Status
	}
	return enums.AccountBudgetProposalStatusEnum_UNSPECIFIED
}

func (m *AccountBudgetProposal) GetProposedName() *wrappers.StringValue {
	if m != nil {
		return m.ProposedName
	}
	return nil
}

func (m *AccountBudgetProposal) GetApprovedStartDateTime() *wrappers.StringValue {
	if m != nil {
		return m.ApprovedStartDateTime
	}
	return nil
}

func (m *AccountBudgetProposal) GetProposedPurchaseOrderNumber() *wrappers.StringValue {
	if m != nil {
		return m.ProposedPurchaseOrderNumber
	}
	return nil
}

func (m *AccountBudgetProposal) GetProposedNotes() *wrappers.StringValue {
	if m != nil {
		return m.ProposedNotes
	}
	return nil
}

func (m *AccountBudgetProposal) GetCreationDateTime() *wrappers.StringValue {
	if m != nil {
		return m.CreationDateTime
	}
	return nil
}

func (m *AccountBudgetProposal) GetApprovalDateTime() *wrappers.StringValue {
	if m != nil {
		return m.ApprovalDateTime
	}
	return nil
}

type isAccountBudgetProposal_ProposedStartTime interface {
	isAccountBudgetProposal_ProposedStartTime()
}

type AccountBudgetProposal_ProposedStartDateTime struct {
	ProposedStartDateTime *wrappers.StringValue `protobuf:"bytes,18,opt,name=proposed_start_date_time,json=proposedStartDateTime,proto3,oneof"`
}

type AccountBudgetProposal_ProposedStartTimeType struct {
	ProposedStartTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,7,opt,name=proposed_start_time_type,json=proposedStartTimeType,proto3,enum=google.ads.googleads.v1.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*AccountBudgetProposal_ProposedStartDateTime) isAccountBudgetProposal_ProposedStartTime() {}

func (*AccountBudgetProposal_ProposedStartTimeType) isAccountBudgetProposal_ProposedStartTime() {}

func (m *AccountBudgetProposal) GetProposedStartTime() isAccountBudgetProposal_ProposedStartTime {
	if m != nil {
		return m.ProposedStartTime
	}
	return nil
}

func (m *AccountBudgetProposal) GetProposedStartDateTime() *wrappers.StringValue {
	if x, ok := m.GetProposedStartTime().(*AccountBudgetProposal_ProposedStartDateTime); ok {
		return x.ProposedStartDateTime
	}
	return nil
}

func (m *AccountBudgetProposal) GetProposedStartTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := m.GetProposedStartTime().(*AccountBudgetProposal_ProposedStartTimeType); ok {
		return x.ProposedStartTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

type isAccountBudgetProposal_ProposedEndTime interface {
	isAccountBudgetProposal_ProposedEndTime()
}

type AccountBudgetProposal_ProposedEndDateTime struct {
	ProposedEndDateTime *wrappers.StringValue `protobuf:"bytes,19,opt,name=proposed_end_date_time,json=proposedEndDateTime,proto3,oneof"`
}

type AccountBudgetProposal_ProposedEndTimeType struct {
	ProposedEndTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,9,opt,name=proposed_end_time_type,json=proposedEndTimeType,proto3,enum=google.ads.googleads.v1.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*AccountBudgetProposal_ProposedEndDateTime) isAccountBudgetProposal_ProposedEndTime() {}

func (*AccountBudgetProposal_ProposedEndTimeType) isAccountBudgetProposal_ProposedEndTime() {}

func (m *AccountBudgetProposal) GetProposedEndTime() isAccountBudgetProposal_ProposedEndTime {
	if m != nil {
		return m.ProposedEndTime
	}
	return nil
}

func (m *AccountBudgetProposal) GetProposedEndDateTime() *wrappers.StringValue {
	if x, ok := m.GetProposedEndTime().(*AccountBudgetProposal_ProposedEndDateTime); ok {
		return x.ProposedEndDateTime
	}
	return nil
}

func (m *AccountBudgetProposal) GetProposedEndTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := m.GetProposedEndTime().(*AccountBudgetProposal_ProposedEndTimeType); ok {
		return x.ProposedEndTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

type isAccountBudgetProposal_ApprovedEndTime interface {
	isAccountBudgetProposal_ApprovedEndTime()
}

type AccountBudgetProposal_ApprovedEndDateTime struct {
	ApprovedEndDateTime *wrappers.StringValue `protobuf:"bytes,21,opt,name=approved_end_date_time,json=approvedEndDateTime,proto3,oneof"`
}

type AccountBudgetProposal_ApprovedEndTimeType struct {
	ApprovedEndTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,22,opt,name=approved_end_time_type,json=approvedEndTimeType,proto3,enum=google.ads.googleads.v1.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*AccountBudgetProposal_ApprovedEndDateTime) isAccountBudgetProposal_ApprovedEndTime() {}

func (*AccountBudgetProposal_ApprovedEndTimeType) isAccountBudgetProposal_ApprovedEndTime() {}

func (m *AccountBudgetProposal) GetApprovedEndTime() isAccountBudgetProposal_ApprovedEndTime {
	if m != nil {
		return m.ApprovedEndTime
	}
	return nil
}

func (m *AccountBudgetProposal) GetApprovedEndDateTime() *wrappers.StringValue {
	if x, ok := m.GetApprovedEndTime().(*AccountBudgetProposal_ApprovedEndDateTime); ok {
		return x.ApprovedEndDateTime
	}
	return nil
}

func (m *AccountBudgetProposal) GetApprovedEndTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := m.GetApprovedEndTime().(*AccountBudgetProposal_ApprovedEndTimeType); ok {
		return x.ApprovedEndTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

type isAccountBudgetProposal_ProposedSpendingLimit interface {
	isAccountBudgetProposal_ProposedSpendingLimit()
}

type AccountBudgetProposal_ProposedSpendingLimitMicros struct {
	ProposedSpendingLimitMicros *wrappers.Int64Value `protobuf:"bytes,10,opt,name=proposed_spending_limit_micros,json=proposedSpendingLimitMicros,proto3,oneof"`
}

type AccountBudgetProposal_ProposedSpendingLimitType struct {
	ProposedSpendingLimitType enums.SpendingLimitTypeEnum_SpendingLimitType `protobuf:"varint,11,opt,name=proposed_spending_limit_type,json=proposedSpendingLimitType,proto3,enum=google.ads.googleads.v1.enums.SpendingLimitTypeEnum_SpendingLimitType,oneof"`
}

func (*AccountBudgetProposal_ProposedSpendingLimitMicros) isAccountBudgetProposal_ProposedSpendingLimit() {
}

func (*AccountBudgetProposal_ProposedSpendingLimitType) isAccountBudgetProposal_ProposedSpendingLimit() {
}

func (m *AccountBudgetProposal) GetProposedSpendingLimit() isAccountBudgetProposal_ProposedSpendingLimit {
	if m != nil {
		return m.ProposedSpendingLimit
	}
	return nil
}

func (m *AccountBudgetProposal) GetProposedSpendingLimitMicros() *wrappers.Int64Value {
	if x, ok := m.GetProposedSpendingLimit().(*AccountBudgetProposal_ProposedSpendingLimitMicros); ok {
		return x.ProposedSpendingLimitMicros
	}
	return nil
}

func (m *AccountBudgetProposal) GetProposedSpendingLimitType() enums.SpendingLimitTypeEnum_SpendingLimitType {
	if x, ok := m.GetProposedSpendingLimit().(*AccountBudgetProposal_ProposedSpendingLimitType); ok {
		return x.ProposedSpendingLimitType
	}
	return enums.SpendingLimitTypeEnum_UNSPECIFIED
}

type isAccountBudgetProposal_ApprovedSpendingLimit interface {
	isAccountBudgetProposal_ApprovedSpendingLimit()
}

type AccountBudgetProposal_ApprovedSpendingLimitMicros struct {
	ApprovedSpendingLimitMicros *wrappers.Int64Value `protobuf:"bytes,23,opt,name=approved_spending_limit_micros,json=approvedSpendingLimitMicros,proto3,oneof"`
}

type AccountBudgetProposal_ApprovedSpendingLimitType struct {
	ApprovedSpendingLimitType enums.SpendingLimitTypeEnum_SpendingLimitType `protobuf:"varint,24,opt,name=approved_spending_limit_type,json=approvedSpendingLimitType,proto3,enum=google.ads.googleads.v1.enums.SpendingLimitTypeEnum_SpendingLimitType,oneof"`
}

func (*AccountBudgetProposal_ApprovedSpendingLimitMicros) isAccountBudgetProposal_ApprovedSpendingLimit() {
}

func (*AccountBudgetProposal_ApprovedSpendingLimitType) isAccountBudgetProposal_ApprovedSpendingLimit() {
}

func (m *AccountBudgetProposal) GetApprovedSpendingLimit() isAccountBudgetProposal_ApprovedSpendingLimit {
	if m != nil {
		return m.ApprovedSpendingLimit
	}
	return nil
}

func (m *AccountBudgetProposal) GetApprovedSpendingLimitMicros() *wrappers.Int64Value {
	if x, ok := m.GetApprovedSpendingLimit().(*AccountBudgetProposal_ApprovedSpendingLimitMicros); ok {
		return x.ApprovedSpendingLimitMicros
	}
	return nil
}

func (m *AccountBudgetProposal) GetApprovedSpendingLimitType() enums.SpendingLimitTypeEnum_SpendingLimitType {
	if x, ok := m.GetApprovedSpendingLimit().(*AccountBudgetProposal_ApprovedSpendingLimitType); ok {
		return x.ApprovedSpendingLimitType
	}
	return enums.SpendingLimitTypeEnum_UNSPECIFIED
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AccountBudgetProposal) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AccountBudgetProposal_OneofMarshaler, _AccountBudgetProposal_OneofUnmarshaler, _AccountBudgetProposal_OneofSizer, []interface{}{
		(*AccountBudgetProposal_ProposedStartDateTime)(nil),
		(*AccountBudgetProposal_ProposedStartTimeType)(nil),
		(*AccountBudgetProposal_ProposedEndDateTime)(nil),
		(*AccountBudgetProposal_ProposedEndTimeType)(nil),
		(*AccountBudgetProposal_ApprovedEndDateTime)(nil),
		(*AccountBudgetProposal_ApprovedEndTimeType)(nil),
		(*AccountBudgetProposal_ProposedSpendingLimitMicros)(nil),
		(*AccountBudgetProposal_ProposedSpendingLimitType)(nil),
		(*AccountBudgetProposal_ApprovedSpendingLimitMicros)(nil),
		(*AccountBudgetProposal_ApprovedSpendingLimitType)(nil),
	}
}

func _AccountBudgetProposal_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AccountBudgetProposal)
	// proposed_start_time
	switch x := m.ProposedStartTime.(type) {
	case *AccountBudgetProposal_ProposedStartDateTime:
		b.EncodeVarint(18<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProposedStartDateTime); err != nil {
			return err
		}
	case *AccountBudgetProposal_ProposedStartTimeType:
		b.EncodeVarint(7<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ProposedStartTimeType))
	case nil:
	default:
		return fmt.Errorf("AccountBudgetProposal.ProposedStartTime has unexpected type %T", x)
	}
	// proposed_end_time
	switch x := m.ProposedEndTime.(type) {
	case *AccountBudgetProposal_ProposedEndDateTime:
		b.EncodeVarint(19<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProposedEndDateTime); err != nil {
			return err
		}
	case *AccountBudgetProposal_ProposedEndTimeType:
		b.EncodeVarint(9<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ProposedEndTimeType))
	case nil:
	default:
		return fmt.Errorf("AccountBudgetProposal.ProposedEndTime has unexpected type %T", x)
	}
	// approved_end_time
	switch x := m.ApprovedEndTime.(type) {
	case *AccountBudgetProposal_ApprovedEndDateTime:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ApprovedEndDateTime); err != nil {
			return err
		}
	case *AccountBudgetProposal_ApprovedEndTimeType:
		b.EncodeVarint(22<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ApprovedEndTimeType))
	case nil:
	default:
		return fmt.Errorf("AccountBudgetProposal.ApprovedEndTime has unexpected type %T", x)
	}
	// proposed_spending_limit
	switch x := m.ProposedSpendingLimit.(type) {
	case *AccountBudgetProposal_ProposedSpendingLimitMicros:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProposedSpendingLimitMicros); err != nil {
			return err
		}
	case *AccountBudgetProposal_ProposedSpendingLimitType:
		b.EncodeVarint(11<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ProposedSpendingLimitType))
	case nil:
	default:
		return fmt.Errorf("AccountBudgetProposal.ProposedSpendingLimit has unexpected type %T", x)
	}
	// approved_spending_limit
	switch x := m.ApprovedSpendingLimit.(type) {
	case *AccountBudgetProposal_ApprovedSpendingLimitMicros:
		b.EncodeVarint(23<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ApprovedSpendingLimitMicros); err != nil {
			return err
		}
	case *AccountBudgetProposal_ApprovedSpendingLimitType:
		b.EncodeVarint(24<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ApprovedSpendingLimitType))
	case nil:
	default:
		return fmt.Errorf("AccountBudgetProposal.ApprovedSpendingLimit has unexpected type %T", x)
	}
	return nil
}

func _AccountBudgetProposal_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AccountBudgetProposal)
	switch tag {
	case 18: // proposed_start_time.proposed_start_date_time
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(wrappers.StringValue)
		err := b.DecodeMessage(msg)
		m.ProposedStartTime = &AccountBudgetProposal_ProposedStartDateTime{msg}
		return true, err
	case 7: // proposed_start_time.proposed_start_time_type
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ProposedStartTime = &AccountBudgetProposal_ProposedStartTimeType{enums.TimeTypeEnum_TimeType(x)}
		return true, err
	case 19: // proposed_end_time.proposed_end_date_time
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(wrappers.StringValue)
		err := b.DecodeMessage(msg)
		m.ProposedEndTime = &AccountBudgetProposal_ProposedEndDateTime{msg}
		return true, err
	case 9: // proposed_end_time.proposed_end_time_type
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ProposedEndTime = &AccountBudgetProposal_ProposedEndTimeType{enums.TimeTypeEnum_TimeType(x)}
		return true, err
	case 21: // approved_end_time.approved_end_date_time
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(wrappers.StringValue)
		err := b.DecodeMessage(msg)
		m.ApprovedEndTime = &AccountBudgetProposal_ApprovedEndDateTime{msg}
		return true, err
	case 22: // approved_end_time.approved_end_time_type
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ApprovedEndTime = &AccountBudgetProposal_ApprovedEndTimeType{enums.TimeTypeEnum_TimeType(x)}
		return true, err
	case 10: // proposed_spending_limit.proposed_spending_limit_micros
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(wrappers.Int64Value)
		err := b.DecodeMessage(msg)
		m.ProposedSpendingLimit = &AccountBudgetProposal_ProposedSpendingLimitMicros{msg}
		return true, err
	case 11: // proposed_spending_limit.proposed_spending_limit_type
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ProposedSpendingLimit = &AccountBudgetProposal_ProposedSpendingLimitType{enums.SpendingLimitTypeEnum_SpendingLimitType(x)}
		return true, err
	case 23: // approved_spending_limit.approved_spending_limit_micros
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(wrappers.Int64Value)
		err := b.DecodeMessage(msg)
		m.ApprovedSpendingLimit = &AccountBudgetProposal_ApprovedSpendingLimitMicros{msg}
		return true, err
	case 24: // approved_spending_limit.approved_spending_limit_type
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ApprovedSpendingLimit = &AccountBudgetProposal_ApprovedSpendingLimitType{enums.SpendingLimitTypeEnum_SpendingLimitType(x)}
		return true, err
	default:
		return false, nil
	}
}

func _AccountBudgetProposal_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AccountBudgetProposal)
	// proposed_start_time
	switch x := m.ProposedStartTime.(type) {
	case *AccountBudgetProposal_ProposedStartDateTime:
		s := proto.Size(x.ProposedStartDateTime)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AccountBudgetProposal_ProposedStartTimeType:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.ProposedStartTimeType))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// proposed_end_time
	switch x := m.ProposedEndTime.(type) {
	case *AccountBudgetProposal_ProposedEndDateTime:
		s := proto.Size(x.ProposedEndDateTime)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AccountBudgetProposal_ProposedEndTimeType:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.ProposedEndTimeType))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// approved_end_time
	switch x := m.ApprovedEndTime.(type) {
	case *AccountBudgetProposal_ApprovedEndDateTime:
		s := proto.Size(x.ApprovedEndDateTime)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AccountBudgetProposal_ApprovedEndTimeType:
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(x.ApprovedEndTimeType))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// proposed_spending_limit
	switch x := m.ProposedSpendingLimit.(type) {
	case *AccountBudgetProposal_ProposedSpendingLimitMicros:
		s := proto.Size(x.ProposedSpendingLimitMicros)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AccountBudgetProposal_ProposedSpendingLimitType:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.ProposedSpendingLimitType))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// approved_spending_limit
	switch x := m.ApprovedSpendingLimit.(type) {
	case *AccountBudgetProposal_ApprovedSpendingLimitMicros:
		s := proto.Size(x.ApprovedSpendingLimitMicros)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AccountBudgetProposal_ApprovedSpendingLimitType:
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(x.ApprovedSpendingLimitType))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*AccountBudgetProposal)(nil), "google.ads.googleads.v1.resources.AccountBudgetProposal")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/account_budget_proposal.proto", fileDescriptor_account_budget_proposal_522c3e292dda2d1a)
}

var fileDescriptor_account_budget_proposal_522c3e292dda2d1a = []byte{
	// 840 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xef, 0x6e, 0xeb, 0x34,
	0x18, 0xc6, 0x69, 0x3a, 0x0e, 0x3a, 0x3e, 0xed, 0xd8, 0x49, 0xe9, 0x39, 0xd9, 0x1f, 0x1d, 0x6d,
	0xa0, 0x49, 0x93, 0x10, 0xa9, 0x3a, 0x26, 0x90, 0xc2, 0x07, 0x68, 0x61, 0x6c, 0x4c, 0x30, 0xaa,
	0x74, 0x14, 0x69, 0xaa, 0x14, 0xb9, 0x8d, 0x09, 0x11, 0x89, 0x1d, 0xd9, 0xce, 0xd0, 0xc4, 0x05,
	0x20, 0x6e, 0x83, 0x8f, 0x5c, 0x0a, 0x77, 0xc1, 0x57, 0xae, 0x02, 0xd9, 0x89, 0xdd, 0xa4, 0x6d,
	0xd6, 0x68, 0xe3, 0x9b, 0x13, 0xfb, 0x79, 0x5e, 0xff, 0xfc, 0xbe, 0xaf, 0x13, 0xf0, 0x79, 0x40,
	0x48, 0x10, 0xa1, 0x1e, 0xf4, 0x59, 0x2f, 0x1b, 0x8a, 0xd1, 0x5d, 0xbf, 0x47, 0x11, 0x23, 0x29,
	0x9d, 0x23, 0xd6, 0x83, 0xf3, 0x39, 0x49, 0x31, 0xf7, 0x66, 0xa9, 0x1f, 0x20, 0xee, 0x25, 0x94,
	0x24, 0x84, 0xc1, 0xc8, 0x4e, 0x28, 0xe1, 0xc4, 0x3c, 0xca, 0x54, 0x36, 0xf4, 0x99, 0xad, 0x0d,
	0xec, 0xbb, 0xbe, 0xad, 0x0d, 0xf6, 0x86, 0x55, 0x31, 0x10, 0x4e, 0xe3, 0x4a, 0x7f, 0x8f, 0x71,
	0xc8, 0x53, 0x96, 0x85, 0xd9, 0xfb, 0xe2, 0x71, 0x1e, 0xfc, 0x3e, 0x41, 0xb9, 0xc3, 0xa7, 0x0f,
	0x3b, 0xb0, 0x04, 0x61, 0x3f, 0xc4, 0x81, 0x17, 0x85, 0x71, 0xc8, 0x8b, 0xc2, 0x8f, 0x1e, 0x16,
	0xf2, 0x30, 0x46, 0xc5, 0xe5, 0x6f, 0xf2, 0xe5, 0xf2, 0x69, 0x96, 0xfe, 0xd4, 0xfb, 0x95, 0xc2,
	0x24, 0x41, 0x54, 0x91, 0x1c, 0x28, 0xbb, 0x24, 0xec, 0x41, 0x8c, 0x09, 0x87, 0x3c, 0x24, 0x38,
	0x9f, 0x7d, 0xff, 0x9f, 0x1d, 0xd0, 0x1d, 0x64, 0x30, 0x43, 0xc9, 0x32, 0xca, 0x51, 0xcc, 0x0f,
	0x40, 0x5b, 0x1d, 0xa9, 0x87, 0x61, 0x8c, 0xac, 0xc6, 0x61, 0xe3, 0xe4, 0xb9, 0xdb, 0x52, 0x2f,
	0xaf, 0x61, 0x8c, 0xcc, 0x0f, 0x81, 0x11, 0xfa, 0xd6, 0xf6, 0x61, 0xe3, 0xe4, 0xc5, 0xe9, 0x7e,
	0x9e, 0x0f, 0x5b, 0xed, 0xc4, 0xfe, 0x06, 0xf3, 0x4f, 0xce, 0x26, 0x30, 0x4a, 0x91, 0x6b, 0x84,
	0xbe, 0x39, 0x00, 0xed, 0x59, 0x18, 0x45, 0x02, 0x9a, 0x21, 0x9e, 0x26, 0x96, 0x21, 0x75, 0x07,
	0x2b, 0xba, 0x31, 0xa7, 0x21, 0x0e, 0x32, 0x61, 0x2b, 0x97, 0x8c, 0x85, 0xc2, 0xfc, 0x12, 0x6c,
	0x97, 0x8f, 0xde, 0x6a, 0xd6, 0xf0, 0x68, 0xc3, 0x22, 0xa1, 0xf9, 0x1b, 0x68, 0x97, 0x12, 0x66,
	0x6d, 0x1d, 0x36, 0x4e, 0xb6, 0x4f, 0x27, 0x76, 0x55, 0x69, 0xc9, 0x83, 0xb7, 0xd7, 0x1e, 0xd3,
	0xcd, 0x7d, 0x82, 0xce, 0x71, 0x1a, 0x57, 0xcf, 0xba, 0xad, 0xa4, 0xf0, 0x64, 0x52, 0xf0, 0x2c,
	0x2b, 0x34, 0xeb, 0x5d, 0x19, 0xf5, 0xf6, 0x31, 0x51, 0xc7, 0xd2, 0xa1, 0x3a, 0x6e, 0x36, 0xef,
	0xe6, 0x91, 0xc4, 0xc1, 0x67, 0x7b, 0x40, 0x7e, 0x96, 0xca, 0xb7, 0xeb, 0x1c, 0xbc, 0x92, 0xc8,
	0x44, 0xff, 0x00, 0x2c, 0x98, 0x24, 0x94, 0xdc, 0x21, 0x5f, 0x34, 0x0a, 0xe5, 0x9e, 0x0f, 0x39,
	0xf2, 0x44, 0x31, 0x5a, 0xef, 0xd5, 0x70, 0xeb, 0x2a, 0xf5, 0x58, 0x88, 0xbf, 0x82, 0x1c, 0xdd,
	0x84, 0x31, 0x32, 0x21, 0x78, 0xa3, 0x77, 0x96, 0xa4, 0x74, 0xfe, 0x33, 0x64, 0xc8, 0x23, 0xd4,
	0x47, 0xd4, 0xc3, 0x69, 0x3c, 0x43, 0xd4, 0x6a, 0xd5, 0x30, 0xdf, 0x57, 0x1e, 0xa3, 0xdc, 0xe2,
	0x7b, 0xe1, 0x70, 0x2d, 0x0d, 0x44, 0xc9, 0x2c, 0xe0, 0x09, 0x47, 0xcc, 0x6a, 0xd7, 0x29, 0x19,
	0x4d, 0x2f, 0x24, 0xe6, 0x15, 0x30, 0xe7, 0x14, 0xc9, 0xce, 0x29, 0x80, 0xef, 0xd4, 0x30, 0xda,
	0x51, 0x3a, 0xcd, 0x7c, 0x05, 0xcc, 0xec, 0x30, 0x60, 0x54, 0xf0, 0x7a, 0x59, 0xc7, 0x4b, 0xe9,
	0xb4, 0xd7, 0x8f, 0xc0, 0xd2, 0x70, 0xcb, 0x69, 0x31, 0x37, 0x3b, 0x5e, 0xbe, 0xe5, 0x76, 0x95,
	0xbe, 0x9c, 0x18, 0xb2, 0x62, 0xac, 0xef, 0x1d, 0xeb, 0x1d, 0x59, 0xb8, 0x67, 0x1b, 0x0a, 0x57,
	0xd8, 0xe8, 0xee, 0x50, 0x0f, 0x2b, 0x01, 0xd5, 0x84, 0x39, 0x06, 0xaf, 0x74, 0x40, 0x84, 0xfd,
	0x02, 0x47, 0xa7, 0x06, 0x47, 0xc3, 0xed, 0x28, 0xf5, 0x39, 0xf6, 0x35, 0xc5, 0x2f, 0x4b, 0xa6,
	0x0b, 0x86, 0xe7, 0x4f, 0x60, 0x28, 0x07, 0x2b, 0x12, 0xe8, 0x16, 0x29, 0x13, 0x74, 0x6b, 0x10,
	0x18, 0x6e, 0x47, 0xa9, 0x97, 0x08, 0x4a, 0xa6, 0x0b, 0x82, 0x57, 0x4f, 0x20, 0x28, 0x07, 0xd3,
	0x04, 0xb3, 0x42, 0x37, 0x2e, 0x7d, 0x9f, 0xe2, 0x70, 0x4e, 0x09, 0xb3, 0xc0, 0xc6, 0x9b, 0xfe,
	0xb2, 0xb9, 0x68, 0xc7, 0x71, 0xee, 0xf1, 0xad, 0xb0, 0xf8, 0x4e, 0x3a, 0x98, 0x7f, 0x34, 0xc0,
	0x41, 0x55, 0x10, 0xc9, 0xf5, 0x42, 0x72, 0x7d, 0xbd, 0x81, 0xab, 0x64, 0xad, 0x01, 0x57, 0xde,
	0x5e, 0x36, 0xdd, 0xdd, 0xb5, 0xbb, 0x51, 0xbc, 0x8b, 0x4b, 0x6d, 0x2d, 0xef, 0xeb, 0xcd, 0xbc,
	0x5b, 0xee, 0xbe, 0xbe, 0xdb, 0x2a, 0x78, 0xab, 0x82, 0x48, 0x5e, 0xeb, 0x7f, 0xe5, 0xdd, 0x72,
	0x77, 0xd7, 0xee, 0x46, 0x4c, 0x0e, 0xbb, 0xa0, 0xb3, 0xa6, 0xa9, 0x87, 0x1d, 0xf0, 0x72, 0xa5,
	0x4b, 0xc4, 0xcb, 0x95, 0xc2, 0x1b, 0xee, 0x82, 0xd7, 0x15, 0xb9, 0x13, 0x53, 0x15, 0x98, 0xc3,
	0xdf, 0x0d, 0x70, 0x3c, 0x27, 0xb1, 0xbd, 0xf1, 0xcf, 0x6d, 0xb8, 0xb7, 0xf6, 0x6b, 0x36, 0x12,
	0xa7, 0x3e, 0x6a, 0xdc, 0x5e, 0xe5, 0x06, 0x01, 0x89, 0x20, 0x0e, 0x6c, 0x42, 0x83, 0x5e, 0x80,
	0xb0, 0xcc, 0x89, 0xfa, 0x51, 0x4a, 0x42, 0xf6, 0xc0, 0xaf, 0xe5, 0x67, 0x7a, 0xf4, 0xa7, 0xd1,
	0xbc, 0x18, 0x0c, 0xfe, 0x32, 0x8e, 0x2e, 0x32, 0xcb, 0x81, 0xcf, 0xec, 0x6c, 0x28, 0x46, 0x93,
	0xbe, 0xed, 0xaa, 0x95, 0x7f, 0xab, 0x35, 0xd3, 0x81, 0xcf, 0xa6, 0x7a, 0xcd, 0x74, 0xd2, 0x9f,
	0xea, 0x35, 0xff, 0x1a, 0xc7, 0xd9, 0x84, 0xe3, 0x0c, 0x7c, 0xe6, 0x38, 0x7a, 0x95, 0xe3, 0x4c,
	0xfa, 0x8e, 0xa3, 0xd7, 0xcd, 0x9e, 0xc9, 0xcd, 0x7e, 0xfc, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5d, 0xf5, 0x49, 0xe3, 0x06, 0x0b, 0x00, 0x00,
}
