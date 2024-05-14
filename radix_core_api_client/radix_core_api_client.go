package radix_core_api_client

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
	i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
	i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
	i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
	id97d97a0f61eb7aba03f6dfe0b81d915dc4592e0ad0b2b0a35f86898fa3a1281 "gitlab.com/mayachain/mayanode/bifrost/pkg/chainclients/radix/coreapi/client/lts"
	i8960c5fe7653f0efc6e7f2bc06a2ec6d1c2357a8f02decb6e4cbe18f9e0eaf52 "gitlab.com/mayachain/mayanode/bifrost/pkg/chainclients/radix/coreapi/client/mempool"
	i78093a4e4a0b144ec97604aab95c3df07bd4e66fb0282f14b7e3b2b4674004d4 "gitlab.com/mayachain/mayanode/bifrost/pkg/chainclients/radix/coreapi/client/state"
	iea54b86da45c623cf5fd94330c1536d128bebca6af259285a97e377b603dd494 "gitlab.com/mayachain/mayanode/bifrost/pkg/chainclients/radix/coreapi/client/status"
	i8f4072d456c1d74b5adfef98d8684acded835b6cb52e24c169a318cb684aeb45 "gitlab.com/mayachain/mayanode/bifrost/pkg/chainclients/radix/coreapi/client/stream"
	i59c61699e638c094c5d0fa5d1bedbe6d028f8ca21aa088bd92877c88d3693cb5 "gitlab.com/mayachain/mayanode/bifrost/pkg/chainclients/radix/coreapi/client/transaction"
)

// RadixCoreApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type RadixCoreApiClient struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewRadixCoreApiClient instantiates a new RadixCoreApiClient and sets the default values.
func NewRadixCoreApiClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *RadixCoreApiClient {
	m := &RadixCoreApiClient{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
	}
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory {
		return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory {
		return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory {
		return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory {
		return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory {
		return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory {
		return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory {
		return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory()
	})
	if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
		m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("localhost:3333/core")
	}
	m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
	return m
}

// Lts the lts property
// returns a *LtsRequestBuilder when successful
func (m *RadixCoreApiClient) Lts() *id97d97a0f61eb7aba03f6dfe0b81d915dc4592e0ad0b2b0a35f86898fa3a1281.LtsRequestBuilder {
	return id97d97a0f61eb7aba03f6dfe0b81d915dc4592e0ad0b2b0a35f86898fa3a1281.NewLtsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Mempool the mempool property
// returns a *MempoolRequestBuilder when successful
func (m *RadixCoreApiClient) Mempool() *i8960c5fe7653f0efc6e7f2bc06a2ec6d1c2357a8f02decb6e4cbe18f9e0eaf52.MempoolRequestBuilder {
	return i8960c5fe7653f0efc6e7f2bc06a2ec6d1c2357a8f02decb6e4cbe18f9e0eaf52.NewMempoolRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// State the state property
// returns a *StateRequestBuilder when successful
func (m *RadixCoreApiClient) State() *i78093a4e4a0b144ec97604aab95c3df07bd4e66fb0282f14b7e3b2b4674004d4.StateRequestBuilder {
	return i78093a4e4a0b144ec97604aab95c3df07bd4e66fb0282f14b7e3b2b4674004d4.NewStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Status the status property
// returns a *StatusRequestBuilder when successful
func (m *RadixCoreApiClient) Status() *iea54b86da45c623cf5fd94330c1536d128bebca6af259285a97e377b603dd494.StatusRequestBuilder {
	return iea54b86da45c623cf5fd94330c1536d128bebca6af259285a97e377b603dd494.NewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Stream the stream property
// returns a *StreamRequestBuilder when successful
func (m *RadixCoreApiClient) Stream() *i8f4072d456c1d74b5adfef98d8684acded835b6cb52e24c169a318cb684aeb45.StreamRequestBuilder {
	return i8f4072d456c1d74b5adfef98d8684acded835b6cb52e24c169a318cb684aeb45.NewStreamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Transaction the transaction property
// returns a *TransactionRequestBuilder when successful
func (m *RadixCoreApiClient) Transaction() *i59c61699e638c094c5d0fa5d1bedbe6d028f8ca21aa088bd92877c88d3693cb5.TransactionRequestBuilder {
	return i59c61699e638c094c5d0fa5d1bedbe6d028f8ca21aa088bd92877c88d3693cb5.NewTransactionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
