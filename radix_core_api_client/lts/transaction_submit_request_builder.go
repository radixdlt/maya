package lts

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455 "github.com/radixdlt/maya/radix_core_api_client/models"
)

// TransactionSubmitRequestBuilder builds and executes requests for operations under \lts\transaction\submit
type TransactionSubmitRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TransactionSubmitRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TransactionSubmitRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewTransactionSubmitRequestBuilderInternal instantiates a new TransactionSubmitRequestBuilder and sets the default values.
func NewTransactionSubmitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TransactionSubmitRequestBuilder {
	m := &TransactionSubmitRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/lts/transaction/submit", pathParameters),
	}
	return m
}

// NewTransactionSubmitRequestBuilder instantiates a new TransactionSubmitRequestBuilder and sets the default values.
func NewTransactionSubmitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TransactionSubmitRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTransactionSubmitRequestBuilderInternal(urlParams, requestAdapter)
}

// Post submits a notarized transaction to the network. Returns whether the transaction submission was already included in the node's mempool.
// returns a LtsTransactionSubmitResponseable when successful
// returns a LtsTransactionSubmitErrorResponse error when the service returns a 400 status code
// returns a LtsTransactionSubmitErrorResponse error when the service returns a 500 status code
func (m *TransactionSubmitRequestBuilder) Post(ctx context.Context, body ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.LtsTransactionSubmitRequestable, requestConfiguration *TransactionSubmitRequestBuilderPostRequestConfiguration) (ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.LtsTransactionSubmitResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"400": ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.CreateLtsTransactionSubmitErrorResponseFromDiscriminatorValue,
		"500": ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.CreateLtsTransactionSubmitErrorResponseFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.CreateLtsTransactionSubmitResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.LtsTransactionSubmitResponseable), nil
}

// ToPostRequestInformation submits a notarized transaction to the network. Returns whether the transaction submission was already included in the node's mempool.
// returns a *RequestInformation when successful
func (m *TransactionSubmitRequestBuilder) ToPostRequestInformation(ctx context.Context, body ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.LtsTransactionSubmitRequestable, requestConfiguration *TransactionSubmitRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TransactionSubmitRequestBuilder when successful
func (m *TransactionSubmitRequestBuilder) WithUrl(rawUrl string) *TransactionSubmitRequestBuilder {
	return NewTransactionSubmitRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
