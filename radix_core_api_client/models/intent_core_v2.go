package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IntentCoreV2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A map of the hex-encoded blob hash, to hex-encoded blob content. Only returned if enabled in `TransactionFormatOptions` on your request.
    blobs_hex IntentCoreV2_blobs_hexable
    // The children_specifiers property
    children_specifiers []string
    // The decompiled transaction manifest instructions. Only returned if enabled in `TransactionFormatOptions` on your request.
    instructions *string
    // The metadata common to both transaction intents and subintents.The `min_proposer_timestamp_inclusive` and `max_proposer_timestamp_exclusive` are both optional.
    intent_header IntentHeaderV2able
    // The message property
    message TransactionMessageable
}
// NewIntentCoreV2 instantiates a new IntentCoreV2 and sets the default values.
func NewIntentCoreV2()(*IntentCoreV2) {
    m := &IntentCoreV2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIntentCoreV2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIntentCoreV2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntentCoreV2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *IntentCoreV2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBlobsHex gets the blobs_hex property value. A map of the hex-encoded blob hash, to hex-encoded blob content. Only returned if enabled in `TransactionFormatOptions` on your request.
// returns a IntentCoreV2_blobs_hexable when successful
func (m *IntentCoreV2) GetBlobsHex()(IntentCoreV2_blobs_hexable) {
    return m.blobs_hex
}
// GetChildrenSpecifiers gets the children_specifiers property value. The children_specifiers property
// returns a []string when successful
func (m *IntentCoreV2) GetChildrenSpecifiers()([]string) {
    return m.children_specifiers
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IntentCoreV2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["blobs_hex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIntentCoreV2_blobs_hexFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlobsHex(val.(IntentCoreV2_blobs_hexable))
        }
        return nil
    }
    res["children_specifiers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetChildrenSpecifiers(res)
        }
        return nil
    }
    res["instructions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstructions(val)
        }
        return nil
    }
    res["intent_header"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIntentHeaderV2FromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntentHeader(val.(IntentHeaderV2able))
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTransactionMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val.(TransactionMessageable))
        }
        return nil
    }
    return res
}
// GetInstructions gets the instructions property value. The decompiled transaction manifest instructions. Only returned if enabled in `TransactionFormatOptions` on your request.
// returns a *string when successful
func (m *IntentCoreV2) GetInstructions()(*string) {
    return m.instructions
}
// GetIntentHeader gets the intent_header property value. The metadata common to both transaction intents and subintents.The `min_proposer_timestamp_inclusive` and `max_proposer_timestamp_exclusive` are both optional.
// returns a IntentHeaderV2able when successful
func (m *IntentCoreV2) GetIntentHeader()(IntentHeaderV2able) {
    return m.intent_header
}
// GetMessage gets the message property value. The message property
// returns a TransactionMessageable when successful
func (m *IntentCoreV2) GetMessage()(TransactionMessageable) {
    return m.message
}
// Serialize serializes information the current object
func (m *IntentCoreV2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("blobs_hex", m.GetBlobsHex())
        if err != nil {
            return err
        }
    }
    if m.GetChildrenSpecifiers() != nil {
        err := writer.WriteCollectionOfStringValues("children_specifiers", m.GetChildrenSpecifiers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("instructions", m.GetInstructions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("intent_header", m.GetIntentHeader())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("message", m.GetMessage())
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
func (m *IntentCoreV2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBlobsHex sets the blobs_hex property value. A map of the hex-encoded blob hash, to hex-encoded blob content. Only returned if enabled in `TransactionFormatOptions` on your request.
func (m *IntentCoreV2) SetBlobsHex(value IntentCoreV2_blobs_hexable)() {
    m.blobs_hex = value
}
// SetChildrenSpecifiers sets the children_specifiers property value. The children_specifiers property
func (m *IntentCoreV2) SetChildrenSpecifiers(value []string)() {
    m.children_specifiers = value
}
// SetInstructions sets the instructions property value. The decompiled transaction manifest instructions. Only returned if enabled in `TransactionFormatOptions` on your request.
func (m *IntentCoreV2) SetInstructions(value *string)() {
    m.instructions = value
}
// SetIntentHeader sets the intent_header property value. The metadata common to both transaction intents and subintents.The `min_proposer_timestamp_inclusive` and `max_proposer_timestamp_exclusive` are both optional.
func (m *IntentCoreV2) SetIntentHeader(value IntentHeaderV2able)() {
    m.intent_header = value
}
// SetMessage sets the message property value. The message property
func (m *IntentCoreV2) SetMessage(value TransactionMessageable)() {
    m.message = value
}
type IntentCoreV2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlobsHex()(IntentCoreV2_blobs_hexable)
    GetChildrenSpecifiers()([]string)
    GetInstructions()(*string)
    GetIntentHeader()(IntentHeaderV2able)
    GetMessage()(TransactionMessageable)
    SetBlobsHex(value IntentCoreV2_blobs_hexable)()
    SetChildrenSpecifiers(value []string)()
    SetInstructions(value *string)()
    SetIntentHeader(value IntentHeaderV2able)()
    SetMessage(value TransactionMessageable)()
}
