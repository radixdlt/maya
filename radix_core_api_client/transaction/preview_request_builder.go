package transaction

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e "github.com/radixdlt/maya/radix_core_api_client/models"
)

// PreviewRequestBuilder builds and executes requests for operations under \transaction\preview
type PreviewRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PreviewRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PreviewRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPreviewRequestBuilderInternal instantiates a new PreviewRequestBuilder and sets the default values.
func NewPreviewRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PreviewRequestBuilder) {
    m := &PreviewRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/transaction/preview", pathParameters),
    }
    return m
}
// NewPreviewRequestBuilder instantiates a new PreviewRequestBuilder and sets the default values.
func NewPreviewRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PreviewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPreviewRequestBuilderInternal(urlParams, requestAdapter)
}
// Post preview a transaction against the latest network state, and returns the preview receipt.If the node has enabled it, you may be able to also preview against recent network state.For V2 transactions (and beyond) the `/preview-v2` endpoint should be used instead.
// returns a TransactionPreviewResponseable when successful
// returns a BasicErrorResponse error when the service returns a 400 status code
// returns a BasicErrorResponse error when the service returns a 500 status code
func (m *PreviewRequestBuilder) Post(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.TransactionPreviewRequestable, requestConfiguration *PreviewRequestBuilderPostRequestConfiguration)(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.TransactionPreviewResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateBasicErrorResponseFromDiscriminatorValue,
        "500": i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateBasicErrorResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.CreateTransactionPreviewResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.TransactionPreviewResponseable), nil
}
// ToPostRequestInformation preview a transaction against the latest network state, and returns the preview receipt.If the node has enabled it, you may be able to also preview against recent network state.For V2 transactions (and beyond) the `/preview-v2` endpoint should be used instead.
// returns a *RequestInformation when successful
func (m *PreviewRequestBuilder) ToPostRequestInformation(ctx context.Context, body i950fd50751836b31c29b35851f9bd0af5f3c32c63b41836407b7c026d8db159e.TransactionPreviewRequestable, requestConfiguration *PreviewRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PreviewRequestBuilder when successful
func (m *PreviewRequestBuilder) WithUrl(rawUrl string)(*PreviewRequestBuilder) {
    return NewPreviewRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
