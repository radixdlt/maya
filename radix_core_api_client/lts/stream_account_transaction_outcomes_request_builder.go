package lts

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e "github.com/radixdlt/maya/radix_core_api_client/models"
)

// StreamAccountTransactionOutcomesRequestBuilder builds and executes requests for operations under \lts\stream\account-transaction-outcomes
type StreamAccountTransactionOutcomesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// StreamAccountTransactionOutcomesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type StreamAccountTransactionOutcomesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewStreamAccountTransactionOutcomesRequestBuilderInternal instantiates a new StreamAccountTransactionOutcomesRequestBuilder and sets the default values.
func NewStreamAccountTransactionOutcomesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StreamAccountTransactionOutcomesRequestBuilder) {
    m := &StreamAccountTransactionOutcomesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/lts/stream/account-transaction-outcomes", pathParameters),
    }
    return m
}
// NewStreamAccountTransactionOutcomesRequestBuilder instantiates a new StreamAccountTransactionOutcomesRequestBuilder and sets the default values.
func NewStreamAccountTransactionOutcomesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StreamAccountTransactionOutcomesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStreamAccountTransactionOutcomesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post returns a list of committed transaction outcomes (containing balance changes) from a given state version, filtered to only transactions which involved the given account.
// returns a LtsStreamAccountTransactionOutcomesResponseable when successful
// returns a BasicErrorResponse error when the service returns a 400 status code
// returns a BasicErrorResponse error when the service returns a 500 status code
func (m *StreamAccountTransactionOutcomesRequestBuilder) Post(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.LtsStreamAccountTransactionOutcomesRequestable, requestConfiguration *StreamAccountTransactionOutcomesRequestBuilderPostRequestConfiguration)(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.LtsStreamAccountTransactionOutcomesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateBasicErrorResponseFromDiscriminatorValue,
        "500": i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateBasicErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateLtsStreamAccountTransactionOutcomesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.LtsStreamAccountTransactionOutcomesResponseable), nil
}
// ToPostRequestInformation returns a list of committed transaction outcomes (containing balance changes) from a given state version, filtered to only transactions which involved the given account.
// returns a *RequestInformation when successful
func (m *StreamAccountTransactionOutcomesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.LtsStreamAccountTransactionOutcomesRequestable, requestConfiguration *StreamAccountTransactionOutcomesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *StreamAccountTransactionOutcomesRequestBuilder when successful
func (m *StreamAccountTransactionOutcomesRequestBuilder) WithUrl(rawUrl string)(*StreamAccountTransactionOutcomesRequestBuilder) {
    return NewStreamAccountTransactionOutcomesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
