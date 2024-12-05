package lts

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e "github.com/radixdlt/maya/radix_core_api_client/models"
)

// TransactionConstructionRequestBuilder builds and executes requests for operations under \lts\transaction\construction
type TransactionConstructionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TransactionConstructionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TransactionConstructionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTransactionConstructionRequestBuilderInternal instantiates a new TransactionConstructionRequestBuilder and sets the default values.
func NewTransactionConstructionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransactionConstructionRequestBuilder) {
    m := &TransactionConstructionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/lts/transaction/construction", pathParameters),
    }
    return m
}
// NewTransactionConstructionRequestBuilder instantiates a new TransactionConstructionRequestBuilder and sets the default values.
func NewTransactionConstructionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransactionConstructionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTransactionConstructionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post returns information necessary to build a transaction
// returns a LtsTransactionConstructionResponseable when successful
// returns a BasicErrorResponse error when the service returns a 500 status code
func (m *TransactionConstructionRequestBuilder) Post(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.LtsTransactionConstructionRequestable, requestConfiguration *TransactionConstructionRequestBuilderPostRequestConfiguration)(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.LtsTransactionConstructionResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "500": i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateBasicErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateLtsTransactionConstructionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.LtsTransactionConstructionResponseable), nil
}
// ToPostRequestInformation returns information necessary to build a transaction
// returns a *RequestInformation when successful
func (m *TransactionConstructionRequestBuilder) ToPostRequestInformation(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.LtsTransactionConstructionRequestable, requestConfiguration *TransactionConstructionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TransactionConstructionRequestBuilder when successful
func (m *TransactionConstructionRequestBuilder) WithUrl(rawUrl string)(*TransactionConstructionRequestBuilder) {
    return NewTransactionConstructionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
