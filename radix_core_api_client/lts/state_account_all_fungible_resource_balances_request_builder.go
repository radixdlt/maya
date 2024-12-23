package lts

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e "github.com/radixdlt/maya/radix_core_api_client/models"
)

// StateAccountAllFungibleResourceBalancesRequestBuilder builds and executes requests for operations under \lts\state\account-all-fungible-resource-balances
type StateAccountAllFungibleResourceBalancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// StateAccountAllFungibleResourceBalancesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type StateAccountAllFungibleResourceBalancesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewStateAccountAllFungibleResourceBalancesRequestBuilderInternal instantiates a new StateAccountAllFungibleResourceBalancesRequestBuilder and sets the default values.
func NewStateAccountAllFungibleResourceBalancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StateAccountAllFungibleResourceBalancesRequestBuilder) {
    m := &StateAccountAllFungibleResourceBalancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/lts/state/account-all-fungible-resource-balances", pathParameters),
    }
    return m
}
// NewStateAccountAllFungibleResourceBalancesRequestBuilder instantiates a new StateAccountAllFungibleResourceBalancesRequestBuilder and sets the default values.
func NewStateAccountAllFungibleResourceBalancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StateAccountAllFungibleResourceBalancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStateAccountAllFungibleResourceBalancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post returns balances for all resources associated with an account
// returns a LtsStateAccountAllFungibleResourceBalancesResponseable when successful
// returns a BasicErrorResponse error when the service returns a 500 status code
func (m *StateAccountAllFungibleResourceBalancesRequestBuilder) Post(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.LtsStateAccountAllFungibleResourceBalancesRequestable, requestConfiguration *StateAccountAllFungibleResourceBalancesRequestBuilderPostRequestConfiguration)(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.LtsStateAccountAllFungibleResourceBalancesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "500": i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateBasicErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateLtsStateAccountAllFungibleResourceBalancesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.LtsStateAccountAllFungibleResourceBalancesResponseable), nil
}
// ToPostRequestInformation returns balances for all resources associated with an account
// returns a *RequestInformation when successful
func (m *StateAccountAllFungibleResourceBalancesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.LtsStateAccountAllFungibleResourceBalancesRequestable, requestConfiguration *StateAccountAllFungibleResourceBalancesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *StateAccountAllFungibleResourceBalancesRequestBuilder when successful
func (m *StateAccountAllFungibleResourceBalancesRequestBuilder) WithUrl(rawUrl string)(*StateAccountAllFungibleResourceBalancesRequestBuilder) {
    return NewStateAccountAllFungibleResourceBalancesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
