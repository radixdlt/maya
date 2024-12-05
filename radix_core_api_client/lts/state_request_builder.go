package lts

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StateRequestBuilder builds and executes requests for operations under \lts\state
type StateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccountAllFungibleResourceBalances the accountAllFungibleResourceBalances property
// returns a *StateAccountAllFungibleResourceBalancesRequestBuilder when successful
func (m *StateRequestBuilder) AccountAllFungibleResourceBalances()(*StateAccountAllFungibleResourceBalancesRequestBuilder) {
    return NewStateAccountAllFungibleResourceBalancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccountDepositBehaviour the accountDepositBehaviour property
// returns a *StateAccountDepositBehaviourRequestBuilder when successful
func (m *StateRequestBuilder) AccountDepositBehaviour()(*StateAccountDepositBehaviourRequestBuilder) {
    return NewStateAccountDepositBehaviourRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccountFungibleResourceBalance the accountFungibleResourceBalance property
// returns a *StateAccountFungibleResourceBalanceRequestBuilder when successful
func (m *StateRequestBuilder) AccountFungibleResourceBalance()(*StateAccountFungibleResourceBalanceRequestBuilder) {
    return NewStateAccountFungibleResourceBalanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewStateRequestBuilderInternal instantiates a new StateRequestBuilder and sets the default values.
func NewStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StateRequestBuilder) {
    m := &StateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/lts/state", pathParameters),
    }
    return m
}
// NewStateRequestBuilder instantiates a new StateRequestBuilder and sets the default values.
func NewStateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStateRequestBuilderInternal(urlParams, requestAdapter)
}
