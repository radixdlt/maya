package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StreamProofsRequest a request to retrieve a sublist of proofs.
type StreamProofsRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A continuation token is returned if and only if there are further non-empty pages of items currently available.The token can be provided in a following request to fetch the next page of results.The filter and sort should not be changed when re-using the continuation token.
    continuation_token *string
    // If not provided, defaults to "Any".
    filter StreamProofsFilterable
    // If specified, the maximum number of proofs that will be returned.
    max_page_size *int32
    // The logical name of the network
    network *string
}
// NewStreamProofsRequest instantiates a new StreamProofsRequest and sets the default values.
func NewStreamProofsRequest()(*StreamProofsRequest) {
    m := &StreamProofsRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStreamProofsRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStreamProofsRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStreamProofsRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StreamProofsRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContinuationToken gets the continuation_token property value. A continuation token is returned if and only if there are further non-empty pages of items currently available.The token can be provided in a following request to fetch the next page of results.The filter and sort should not be changed when re-using the continuation token.
// returns a *string when successful
func (m *StreamProofsRequest) GetContinuationToken()(*string) {
    return m.continuation_token
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StreamProofsRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["continuation_token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContinuationToken(val)
        }
        return nil
    }
    res["filter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStreamProofsFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilter(val.(StreamProofsFilterable))
        }
        return nil
    }
    res["max_page_size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxPageSize(val)
        }
        return nil
    }
    res["network"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetwork(val)
        }
        return nil
    }
    return res
}
// GetFilter gets the filter property value. If not provided, defaults to "Any".
// returns a StreamProofsFilterable when successful
func (m *StreamProofsRequest) GetFilter()(StreamProofsFilterable) {
    return m.filter
}
// GetMaxPageSize gets the max_page_size property value. If specified, the maximum number of proofs that will be returned.
// returns a *int32 when successful
func (m *StreamProofsRequest) GetMaxPageSize()(*int32) {
    return m.max_page_size
}
// GetNetwork gets the network property value. The logical name of the network
// returns a *string when successful
func (m *StreamProofsRequest) GetNetwork()(*string) {
    return m.network
}
// Serialize serializes information the current object
func (m *StreamProofsRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("continuation_token", m.GetContinuationToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("max_page_size", m.GetMaxPageSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("network", m.GetNetwork())
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
func (m *StreamProofsRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContinuationToken sets the continuation_token property value. A continuation token is returned if and only if there are further non-empty pages of items currently available.The token can be provided in a following request to fetch the next page of results.The filter and sort should not be changed when re-using the continuation token.
func (m *StreamProofsRequest) SetContinuationToken(value *string)() {
    m.continuation_token = value
}
// SetFilter sets the filter property value. If not provided, defaults to "Any".
func (m *StreamProofsRequest) SetFilter(value StreamProofsFilterable)() {
    m.filter = value
}
// SetMaxPageSize sets the max_page_size property value. If specified, the maximum number of proofs that will be returned.
func (m *StreamProofsRequest) SetMaxPageSize(value *int32)() {
    m.max_page_size = value
}
// SetNetwork sets the network property value. The logical name of the network
func (m *StreamProofsRequest) SetNetwork(value *string)() {
    m.network = value
}
type StreamProofsRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContinuationToken()(*string)
    GetFilter()(StreamProofsFilterable)
    GetMaxPageSize()(*int32)
    GetNetwork()(*string)
    SetContinuationToken(value *string)()
    SetFilter(value StreamProofsFilterable)()
    SetMaxPageSize(value *int32)()
    SetNetwork(value *string)()
}
