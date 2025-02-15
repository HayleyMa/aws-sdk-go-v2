// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type OrderStatus string

// Enum values for OrderStatus
const (
	OrderStatusReceived   OrderStatus = "RECEIVED"
	OrderStatusPending    OrderStatus = "PENDING"
	OrderStatusProcessing OrderStatus = "PROCESSING"
	OrderStatusInstalling OrderStatus = "INSTALLING"
	OrderStatusFulfilled  OrderStatus = "FULFILLED"
	OrderStatusCancelled  OrderStatus = "CANCELLED"
)

// Values returns all known values for OrderStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OrderStatus) Values() []OrderStatus {
	return []OrderStatus{
		"RECEIVED",
		"PENDING",
		"PROCESSING",
		"INSTALLING",
		"FULFILLED",
		"CANCELLED",
	}
}

type PaymentOption string

// Enum values for PaymentOption
const (
	PaymentOptionAllUpfront     PaymentOption = "ALL_UPFRONT"
	PaymentOptionNoUpfront      PaymentOption = "NO_UPFRONT"
	PaymentOptionPartialUpfront PaymentOption = "PARTIAL_UPFRONT"
)

// Values returns all known values for PaymentOption. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PaymentOption) Values() []PaymentOption {
	return []PaymentOption{
		"ALL_UPFRONT",
		"NO_UPFRONT",
		"PARTIAL_UPFRONT",
	}
}

type PaymentTerm string

// Enum values for PaymentTerm
const (
	PaymentTermThreeYears PaymentTerm = "THREE_YEARS"
)

// Values returns all known values for PaymentTerm. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PaymentTerm) Values() []PaymentTerm {
	return []PaymentTerm{
		"THREE_YEARS",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeOutpost ResourceType = "OUTPOST"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"OUTPOST",
	}
}
