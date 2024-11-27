package transaction

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TransactionRequestBuilder builds and executes requests for operations under \transaction
type TransactionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallPreview the callPreview property
// returns a *CallPreviewRequestBuilder when successful
func (m *TransactionRequestBuilder) CallPreview()(*CallPreviewRequestBuilder) {
    return NewCallPreviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewTransactionRequestBuilderInternal instantiates a new TransactionRequestBuilder and sets the default values.
func NewTransactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransactionRequestBuilder) {
    m := &TransactionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/transaction", pathParameters),
    }
    return m
}
// NewTransactionRequestBuilder instantiates a new TransactionRequestBuilder and sets the default values.
func NewTransactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTransactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Parse the parse property
// returns a *ParseRequestBuilder when successful
func (m *TransactionRequestBuilder) Parse()(*ParseRequestBuilder) {
    return NewParseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Preview the preview property
// returns a *PreviewRequestBuilder when successful
func (m *TransactionRequestBuilder) Preview()(*PreviewRequestBuilder) {
    return NewPreviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PreviewV2 the previewV2 property
// returns a *PreviewV2RequestBuilder when successful
func (m *TransactionRequestBuilder) PreviewV2()(*PreviewV2RequestBuilder) {
    return NewPreviewV2RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Receipt the receipt property
// returns a *ReceiptRequestBuilder when successful
func (m *TransactionRequestBuilder) Receipt()(*ReceiptRequestBuilder) {
    return NewReceiptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Status the status property
// returns a *StatusRequestBuilder when successful
func (m *TransactionRequestBuilder) Status()(*StatusRequestBuilder) {
    return NewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Submit the submit property
// returns a *SubmitRequestBuilder when successful
func (m *TransactionRequestBuilder) Submit()(*SubmitRequestBuilder) {
    return NewSubmitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
