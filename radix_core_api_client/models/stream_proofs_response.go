package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StreamProofsResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// A continuation token is returned if and only if there are further non-empty pages of items currently available.The token can be provided in a following request to fetch the next page of results.The filter and sort should not be changed when re-using the continuation token.
	continuation_token *string
	// A page of ledger proofs stored by this node.
	page []LedgerProofable
}

// NewStreamProofsResponse instantiates a new StreamProofsResponse and sets the default values.
func NewStreamProofsResponse() *StreamProofsResponse {
	m := &StreamProofsResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateStreamProofsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStreamProofsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewStreamProofsResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StreamProofsResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContinuationToken gets the continuation_token property value. A continuation token is returned if and only if there are further non-empty pages of items currently available.The token can be provided in a following request to fetch the next page of results.The filter and sort should not be changed when re-using the continuation token.
// returns a *string when successful
func (m *StreamProofsResponse) GetContinuationToken() *string {
	return m.continuation_token
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StreamProofsResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["continuation_token"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContinuationToken(val)
		}
		return nil
	}
	res["page"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateLedgerProofFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]LedgerProofable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(LedgerProofable)
				}
			}
			m.SetPage(res)
		}
		return nil
	}
	return res
}

// GetPage gets the page property value. A page of ledger proofs stored by this node.
// returns a []LedgerProofable when successful
func (m *StreamProofsResponse) GetPage() []LedgerProofable {
	return m.page
}

// Serialize serializes information the current object
func (m *StreamProofsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("continuation_token", m.GetContinuationToken())
		if err != nil {
			return err
		}
	}
	if m.GetPage() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPage()))
		for i, v := range m.GetPage() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("page", cast)
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
func (m *StreamProofsResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContinuationToken sets the continuation_token property value. A continuation token is returned if and only if there are further non-empty pages of items currently available.The token can be provided in a following request to fetch the next page of results.The filter and sort should not be changed when re-using the continuation token.
func (m *StreamProofsResponse) SetContinuationToken(value *string) {
	m.continuation_token = value
}

// SetPage sets the page property value. A page of ledger proofs stored by this node.
func (m *StreamProofsResponse) SetPage(value []LedgerProofable) {
	m.page = value
}

type StreamProofsResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContinuationToken() *string
	GetPage() []LedgerProofable
	SetContinuationToken(value *string)
	SetPage(value []LedgerProofable)
}
