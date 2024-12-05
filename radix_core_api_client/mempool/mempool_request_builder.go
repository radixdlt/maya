package mempool

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MempoolRequestBuilder builds and executes requests for operations under \mempool
type MempoolRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewMempoolRequestBuilderInternal instantiates a new MempoolRequestBuilder and sets the default values.
func NewMempoolRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MempoolRequestBuilder) {
    m := &MempoolRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/mempool", pathParameters),
    }
    return m
}
// NewMempoolRequestBuilder instantiates a new MempoolRequestBuilder and sets the default values.
func NewMempoolRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MempoolRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMempoolRequestBuilderInternal(urlParams, requestAdapter)
}
// List the list property
// returns a *ListRequestBuilder when successful
func (m *MempoolRequestBuilder) List()(*ListRequestBuilder) {
    return NewListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Transaction the transaction property
// returns a *TransactionRequestBuilder when successful
func (m *MempoolRequestBuilder) Transaction()(*TransactionRequestBuilder) {
    return NewTransactionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
