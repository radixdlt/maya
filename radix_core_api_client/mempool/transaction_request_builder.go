package mempool

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e "github.com/radixdlt/maya/radix_core_api_client/models"
)

// TransactionRequestBuilder builds and executes requests for operations under \mempool\transaction
type TransactionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TransactionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TransactionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTransactionRequestBuilderInternal instantiates a new TransactionRequestBuilder and sets the default values.
func NewTransactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransactionRequestBuilder) {
    m := &TransactionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/mempool/transaction", pathParameters),
    }
    return m
}
// NewTransactionRequestBuilder instantiates a new TransactionRequestBuilder and sets the default values.
func NewTransactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTransactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post returns the payload of a transaction currently in the mempool
// returns a MempoolTransactionResponseable when successful
// returns a BasicErrorResponse error when the service returns a 404 status code
// returns a BasicErrorResponse error when the service returns a 500 status code
func (m *TransactionRequestBuilder) Post(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.MempoolTransactionRequestable, requestConfiguration *TransactionRequestBuilderPostRequestConfiguration)(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.MempoolTransactionResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateBasicErrorResponseFromDiscriminatorValue,
        "500": i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateBasicErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateMempoolTransactionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.MempoolTransactionResponseable), nil
}
// ToPostRequestInformation returns the payload of a transaction currently in the mempool
// returns a *RequestInformation when successful
func (m *TransactionRequestBuilder) ToPostRequestInformation(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.MempoolTransactionRequestable, requestConfiguration *TransactionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TransactionRequestBuilder when successful
func (m *TransactionRequestBuilder) WithUrl(rawUrl string)(*TransactionRequestBuilder) {
    return NewTransactionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
