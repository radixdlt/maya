package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ConsensusManagerFieldValidatorRewardsValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The proposer_rewards property
    proposer_rewards []ProposerRewardable
    // The rewards_vault property
    rewards_vault EntityReferenceable
}
// NewConsensusManagerFieldValidatorRewardsValue instantiates a new ConsensusManagerFieldValidatorRewardsValue and sets the default values.
func NewConsensusManagerFieldValidatorRewardsValue()(*ConsensusManagerFieldValidatorRewardsValue) {
    m := &ConsensusManagerFieldValidatorRewardsValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConsensusManagerFieldValidatorRewardsValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConsensusManagerFieldValidatorRewardsValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConsensusManagerFieldValidatorRewardsValue(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConsensusManagerFieldValidatorRewardsValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConsensusManagerFieldValidatorRewardsValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["proposer_rewards"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProposerRewardFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProposerRewardable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProposerRewardable)
                }
            }
            m.SetProposerRewards(res)
        }
        return nil
    }
    res["rewards_vault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRewardsVault(val.(EntityReferenceable))
        }
        return nil
    }
    return res
}
// GetProposerRewards gets the proposer_rewards property value. The proposer_rewards property
// returns a []ProposerRewardable when successful
func (m *ConsensusManagerFieldValidatorRewardsValue) GetProposerRewards()([]ProposerRewardable) {
    return m.proposer_rewards
}
// GetRewardsVault gets the rewards_vault property value. The rewards_vault property
// returns a EntityReferenceable when successful
func (m *ConsensusManagerFieldValidatorRewardsValue) GetRewardsVault()(EntityReferenceable) {
    return m.rewards_vault
}
// Serialize serializes information the current object
func (m *ConsensusManagerFieldValidatorRewardsValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetProposerRewards() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProposerRewards()))
        for i, v := range m.GetProposerRewards() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("proposer_rewards", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("rewards_vault", m.GetRewardsVault())
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
func (m *ConsensusManagerFieldValidatorRewardsValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetProposerRewards sets the proposer_rewards property value. The proposer_rewards property
func (m *ConsensusManagerFieldValidatorRewardsValue) SetProposerRewards(value []ProposerRewardable)() {
    m.proposer_rewards = value
}
// SetRewardsVault sets the rewards_vault property value. The rewards_vault property
func (m *ConsensusManagerFieldValidatorRewardsValue) SetRewardsVault(value EntityReferenceable)() {
    m.rewards_vault = value
}
type ConsensusManagerFieldValidatorRewardsValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetProposerRewards()([]ProposerRewardable)
    GetRewardsVault()(EntityReferenceable)
    SetProposerRewards(value []ProposerRewardable)()
    SetRewardsVault(value EntityReferenceable)()
}
