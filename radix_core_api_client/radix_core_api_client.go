package radix_core_api_client

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i03e3fe4baa992461efbd1b08c75ee8cf65a59617a86136665cde9ea150281bd2 "github.com/radixdlt/maya/radix_core_api_client/stream"
    i0c9375c7e1ed432ad93cc568344579bf54a5aa028244979fc51ca86eb5a167d5 "github.com/radixdlt/maya/radix_core_api_client/mempool"
    i14eb1b912ad71031f8a2eb5d0e4d3cb5de4534d68063c91bf03e9db4027bbc4c "github.com/radixdlt/maya/radix_core_api_client/lts"
    i6fde89e38c8412255411d92856ccf93c24d2dfcc3611341b105c74540145df04 "github.com/radixdlt/maya/radix_core_api_client/transaction"
    i80ebb2a3ca695e1249923248d9e431845c131647d525fddca6ae23e5723d5fa9 "github.com/radixdlt/maya/radix_core_api_client/state"
    ibde1eaf2adaa3bf5920708cbb6daa6c37efe37b410e8eb88e07eb401c5578c8b "github.com/radixdlt/maya/radix_core_api_client/status"
)

// RadixCoreApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type RadixCoreApiClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewRadixCoreApiClient instantiates a new RadixCoreApiClient and sets the default values.
func NewRadixCoreApiClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RadixCoreApiClient) {
    m := &RadixCoreApiClient{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("localhost:3333/core")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
// Lts the lts property
// returns a *LtsRequestBuilder when successful
func (m *RadixCoreApiClient) Lts()(*i14eb1b912ad71031f8a2eb5d0e4d3cb5de4534d68063c91bf03e9db4027bbc4c.LtsRequestBuilder) {
    return i14eb1b912ad71031f8a2eb5d0e4d3cb5de4534d68063c91bf03e9db4027bbc4c.NewLtsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Mempool the mempool property
// returns a *MempoolRequestBuilder when successful
func (m *RadixCoreApiClient) Mempool()(*i0c9375c7e1ed432ad93cc568344579bf54a5aa028244979fc51ca86eb5a167d5.MempoolRequestBuilder) {
    return i0c9375c7e1ed432ad93cc568344579bf54a5aa028244979fc51ca86eb5a167d5.NewMempoolRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// State the state property
// returns a *StateRequestBuilder when successful
func (m *RadixCoreApiClient) State()(*i80ebb2a3ca695e1249923248d9e431845c131647d525fddca6ae23e5723d5fa9.StateRequestBuilder) {
    return i80ebb2a3ca695e1249923248d9e431845c131647d525fddca6ae23e5723d5fa9.NewStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Status the status property
// returns a *StatusRequestBuilder when successful
func (m *RadixCoreApiClient) Status()(*ibde1eaf2adaa3bf5920708cbb6daa6c37efe37b410e8eb88e07eb401c5578c8b.StatusRequestBuilder) {
    return ibde1eaf2adaa3bf5920708cbb6daa6c37efe37b410e8eb88e07eb401c5578c8b.NewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stream the stream property
// returns a *StreamRequestBuilder when successful
func (m *RadixCoreApiClient) Stream()(*i03e3fe4baa992461efbd1b08c75ee8cf65a59617a86136665cde9ea150281bd2.StreamRequestBuilder) {
    return i03e3fe4baa992461efbd1b08c75ee8cf65a59617a86136665cde9ea150281bd2.NewStreamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Transaction the transaction property
// returns a *TransactionRequestBuilder when successful
func (m *RadixCoreApiClient) Transaction()(*i6fde89e38c8412255411d92856ccf93c24d2dfcc3611341b105c74540145df04.TransactionRequestBuilder) {
    return i6fde89e38c8412255411d92856ccf93c24d2dfcc3611341b105c74540145df04.NewTransactionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
