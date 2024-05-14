package lts

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StreamRequestBuilder builds and executes requests for operations under \lts\stream
type StreamRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AccountTransactionOutcomes the accountTransactionOutcomes property
// returns a *StreamAccountTransactionOutcomesRequestBuilder when successful
func (m *StreamRequestBuilder) AccountTransactionOutcomes() *StreamAccountTransactionOutcomesRequestBuilder {
	return NewStreamAccountTransactionOutcomesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewStreamRequestBuilderInternal instantiates a new StreamRequestBuilder and sets the default values.
func NewStreamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *StreamRequestBuilder {
	m := &StreamRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/lts/stream", pathParameters),
	}
	return m
}

// NewStreamRequestBuilder instantiates a new StreamRequestBuilder and sets the default values.
func NewStreamRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *StreamRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewStreamRequestBuilderInternal(urlParams, requestAdapter)
}

// TransactionOutcomes the transactionOutcomes property
// returns a *StreamTransactionOutcomesRequestBuilder when successful
func (m *StreamRequestBuilder) TransactionOutcomes() *StreamTransactionOutcomesRequestBuilder {
	return NewStreamTransactionOutcomesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
