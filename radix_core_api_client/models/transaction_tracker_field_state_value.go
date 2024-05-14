package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionTrackerFieldStateValue struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The epochs_per_partition property
	epochs_per_partition *int64
	// The partition_range_end_inclusive property
	partition_range_end_inclusive *int32
	// The partition_range_start_inclusive property
	partition_range_start_inclusive *int32
	// The start_epoch property
	start_epoch *int64
	// The start_partition property
	start_partition *int32
}

// NewTransactionTrackerFieldStateValue instantiates a new TransactionTrackerFieldStateValue and sets the default values.
func NewTransactionTrackerFieldStateValue() *TransactionTrackerFieldStateValue {
	m := &TransactionTrackerFieldStateValue{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTransactionTrackerFieldStateValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionTrackerFieldStateValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionTrackerFieldStateValue(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionTrackerFieldStateValue) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEpochsPerPartition gets the epochs_per_partition property value. The epochs_per_partition property
// returns a *int64 when successful
func (m *TransactionTrackerFieldStateValue) GetEpochsPerPartition() *int64 {
	return m.epochs_per_partition
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionTrackerFieldStateValue) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["epochs_per_partition"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEpochsPerPartition(val)
		}
		return nil
	}
	res["partition_range_end_inclusive"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPartitionRangeEndInclusive(val)
		}
		return nil
	}
	res["partition_range_start_inclusive"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPartitionRangeStartInclusive(val)
		}
		return nil
	}
	res["start_epoch"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStartEpoch(val)
		}
		return nil
	}
	res["start_partition"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStartPartition(val)
		}
		return nil
	}
	return res
}

// GetPartitionRangeEndInclusive gets the partition_range_end_inclusive property value. The partition_range_end_inclusive property
// returns a *int32 when successful
func (m *TransactionTrackerFieldStateValue) GetPartitionRangeEndInclusive() *int32 {
	return m.partition_range_end_inclusive
}

// GetPartitionRangeStartInclusive gets the partition_range_start_inclusive property value. The partition_range_start_inclusive property
// returns a *int32 when successful
func (m *TransactionTrackerFieldStateValue) GetPartitionRangeStartInclusive() *int32 {
	return m.partition_range_start_inclusive
}

// GetStartEpoch gets the start_epoch property value. The start_epoch property
// returns a *int64 when successful
func (m *TransactionTrackerFieldStateValue) GetStartEpoch() *int64 {
	return m.start_epoch
}

// GetStartPartition gets the start_partition property value. The start_partition property
// returns a *int32 when successful
func (m *TransactionTrackerFieldStateValue) GetStartPartition() *int32 {
	return m.start_partition
}

// Serialize serializes information the current object
func (m *TransactionTrackerFieldStateValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("epochs_per_partition", m.GetEpochsPerPartition())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("partition_range_end_inclusive", m.GetPartitionRangeEndInclusive())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("partition_range_start_inclusive", m.GetPartitionRangeStartInclusive())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("start_epoch", m.GetStartEpoch())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("start_partition", m.GetStartPartition())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TransactionTrackerFieldStateValue) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEpochsPerPartition sets the epochs_per_partition property value. The epochs_per_partition property
func (m *TransactionTrackerFieldStateValue) SetEpochsPerPartition(value *int64) {
	m.epochs_per_partition = value
}

// SetPartitionRangeEndInclusive sets the partition_range_end_inclusive property value. The partition_range_end_inclusive property
func (m *TransactionTrackerFieldStateValue) SetPartitionRangeEndInclusive(value *int32) {
	m.partition_range_end_inclusive = value
}

// SetPartitionRangeStartInclusive sets the partition_range_start_inclusive property value. The partition_range_start_inclusive property
func (m *TransactionTrackerFieldStateValue) SetPartitionRangeStartInclusive(value *int32) {
	m.partition_range_start_inclusive = value
}

// SetStartEpoch sets the start_epoch property value. The start_epoch property
func (m *TransactionTrackerFieldStateValue) SetStartEpoch(value *int64) {
	m.start_epoch = value
}

// SetStartPartition sets the start_partition property value. The start_partition property
func (m *TransactionTrackerFieldStateValue) SetStartPartition(value *int32) {
	m.start_partition = value
}

type TransactionTrackerFieldStateValueable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEpochsPerPartition() *int64
	GetPartitionRangeEndInclusive() *int32
	GetPartitionRangeStartInclusive() *int32
	GetStartEpoch() *int64
	GetStartPartition() *int32
	SetEpochsPerPartition(value *int64)
	SetPartitionRangeEndInclusive(value *int32)
	SetPartitionRangeStartInclusive(value *int32)
	SetStartEpoch(value *int64)
	SetStartPartition(value *int32)
}
