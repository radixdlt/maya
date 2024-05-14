package stream

import (
	"context"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455 "github.com/radixdlt/maya/radix_core_api_client/models"
)

// ProofsRequestBuilder builds and executes requests for operations under \stream\proofs
type ProofsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ProofsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProofsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewProofsRequestBuilderInternal instantiates a new ProofsRequestBuilder and sets the default values.
func NewProofsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ProofsRequestBuilder {
	m := &ProofsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/stream/proofs", pathParameters),
	}
	return m
}

// NewProofsRequestBuilder instantiates a new ProofsRequestBuilder and sets the default values.
func NewProofsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ProofsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewProofsRequestBuilderInternal(urlParams, requestAdapter)
}

// Post returns a stream of proofs committed to the node's ledger.NOTE: This endpoint may return different results on different nodes:* Each node may persist different subset of signatures on a given proofs, as long as enoughof the validator set has signed.* Inside an epoch, different nodes may receive and persist / keep different proofs, subject toconstraints on gaps between proofs.Proofs during an epoch can also be garbage collected by the node after the fact. Thereforeproofs may disappear from this stream.Some proofs (such as during genesis and protocol update enactment) are created on a node and don'tinclude signatures.This stream accepts four different options in the request:* All proofs forward (from state version)* All end-of-epoch proofs (from epoch number)* All end-of-epoch proofs triggering a protocol update* All node-injected proofs enacting genesis or a protocol update (for protocol update name, from state version)The end-of-epoch proofs can be used to "trustlessly" verify the validator set for a given epoch.By tracking the fact that validators for epoch N sign the next validator set for epoch N + 1,this chain of proofs can be used to provide proof of the current validator set from a hardcodedstart.When a validator set is known for a given epoch, this can be used to verify the various transactionhash trees in the epoch, and to prove other data.NOTE: This endpoint was built after agreeing the new Radix convention for paged APIs. Its modelstherefore follow the new convention, rather than attempting to align with existing loose Core API conventions.
// returns a StreamProofsResponseable when successful
// returns a StreamProofsErrorResponse error when the service returns a 400 status code
// returns a StreamProofsErrorResponse error when the service returns a 500 status code
func (m *ProofsRequestBuilder) Post(ctx context.Context, body ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.StreamProofsRequestable, requestConfiguration *ProofsRequestBuilderPostRequestConfiguration) (ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.StreamProofsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"400": ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.CreateStreamProofsErrorResponseFromDiscriminatorValue,
		"500": ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.CreateStreamProofsErrorResponseFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.CreateStreamProofsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.StreamProofsResponseable), nil
}

// ToPostRequestInformation returns a stream of proofs committed to the node's ledger.NOTE: This endpoint may return different results on different nodes:* Each node may persist different subset of signatures on a given proofs, as long as enoughof the validator set has signed.* Inside an epoch, different nodes may receive and persist / keep different proofs, subject toconstraints on gaps between proofs.Proofs during an epoch can also be garbage collected by the node after the fact. Thereforeproofs may disappear from this stream.Some proofs (such as during genesis and protocol update enactment) are created on a node and don'tinclude signatures.This stream accepts four different options in the request:* All proofs forward (from state version)* All end-of-epoch proofs (from epoch number)* All end-of-epoch proofs triggering a protocol update* All node-injected proofs enacting genesis or a protocol update (for protocol update name, from state version)The end-of-epoch proofs can be used to "trustlessly" verify the validator set for a given epoch.By tracking the fact that validators for epoch N sign the next validator set for epoch N + 1,this chain of proofs can be used to provide proof of the current validator set from a hardcodedstart.When a validator set is known for a given epoch, this can be used to verify the various transactionhash trees in the epoch, and to prove other data.NOTE: This endpoint was built after agreeing the new Radix convention for paged APIs. Its modelstherefore follow the new convention, rather than attempting to align with existing loose Core API conventions.
// returns a *RequestInformation when successful
func (m *ProofsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ieb2700c3462f1a9e73e3b7d2d4d35314fa43b94ba0380cc504a8f60edfdf3455.StreamProofsRequestable, requestConfiguration *ProofsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ProofsRequestBuilder when successful
func (m *ProofsRequestBuilder) WithUrl(rawUrl string) *ProofsRequestBuilder {
	return NewProofsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
