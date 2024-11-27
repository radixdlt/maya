package state

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StateRequestBuilder builds and executes requests for operations under \state
type StateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessController the accessController property
// returns a *AccessControllerRequestBuilder when successful
func (m *StateRequestBuilder) AccessController()(*AccessControllerRequestBuilder) {
    return NewAccessControllerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Account the account property
// returns a *AccountRequestBuilder when successful
func (m *StateRequestBuilder) Account()(*AccountRequestBuilder) {
    return NewAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Component the component property
// returns a *ComponentRequestBuilder when successful
func (m *StateRequestBuilder) Component()(*ComponentRequestBuilder) {
    return NewComponentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConsensusManager the consensusManager property
// returns a *ConsensusManagerRequestBuilder when successful
func (m *StateRequestBuilder) ConsensusManager()(*ConsensusManagerRequestBuilder) {
    return NewConsensusManagerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewStateRequestBuilderInternal instantiates a new StateRequestBuilder and sets the default values.
func NewStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StateRequestBuilder) {
    m := &StateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/state", pathParameters),
    }
    return m
}
// NewStateRequestBuilder instantiates a new StateRequestBuilder and sets the default values.
func NewStateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStateRequestBuilderInternal(urlParams, requestAdapter)
}
// NonFungible the nonFungible property
// returns a *NonFungibleRequestBuilder when successful
func (m *StateRequestBuilder) NonFungible()(*NonFungibleRequestBuilder) {
    return NewNonFungibleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PackageEscaped the package property
// returns a *PackageRequestBuilder when successful
func (m *StateRequestBuilder) PackageEscaped()(*PackageRequestBuilder) {
    return NewPackageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Resource the resource property
// returns a *ResourceRequestBuilder when successful
func (m *StateRequestBuilder) Resource()(*ResourceRequestBuilder) {
    return NewResourceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Validator the validator property
// returns a *ValidatorRequestBuilder when successful
func (m *StateRequestBuilder) Validator()(*ValidatorRequestBuilder) {
    return NewValidatorRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
