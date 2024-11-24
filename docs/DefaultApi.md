# {{classname}}

All URIs are relative to *https://quote-api.jup.ag/v6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IndexedRouteMapGet**](DefaultApi.md#IndexedRouteMapGet) | **Get** /indexed-route-map | GET /indexed-route-map
[**ProgramIdToLabelGet**](DefaultApi.md#ProgramIdToLabelGet) | **Get** /program-id-to-label | GET /program-id-to-label
[**QuoteGet**](DefaultApi.md#QuoteGet) | **Get** /quote | GET /quote
[**SwapInstructionsPost**](DefaultApi.md#SwapInstructionsPost) | **Post** /swap-instructions | POST /swap-instructions
[**SwapPost**](DefaultApi.md#SwapPost) | **Post** /swap | POST /swap
[**TokensGet**](DefaultApi.md#TokensGet) | **Get** /tokens | GET /tokens

# **IndexedRouteMapGet**
> IndexedRouteMapResponse IndexedRouteMapGet(ctx, optional)
GET /indexed-route-map

DEPRECATED, please use /tokens for tradable mints. Returns a hash map, input mint as key and an array of valid output mint as values, token mints are indexed to reduce the file size

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiIndexedRouteMapGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiIndexedRouteMapGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onlyDirectRoutes** | **optional.Bool**| Default is false. Direct Routes limits Jupiter routing to single hop routes only. | 

### Return type

[**IndexedRouteMapResponse**](IndexedRouteMapResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProgramIdToLabelGet**
> map[string]string ProgramIdToLabelGet(ctx, )
GET /program-id-to-label

Returns a hash, which key is the program id and value is the label. This is used to help map error from transaction by identifying the fault program id. With that, we can use the `excludeDexes` or `dexes` parameter.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuoteGet**
> QuoteResponse QuoteGet(ctx, inputMint, outputMint, amount, optional)
GET /quote

Sends a GET request to the Jupiter API to get the best priced quote.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputMint** | **string**| Input token mint address | 
  **outputMint** | **string**| Output token mint address | 
  **amount** | **int32**| The amount to swap, have to factor in the token decimals. | 
 **optional** | ***DefaultApiQuoteGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiQuoteGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **slippageBps** | **optional.Int32**| The slippage in basis points, 1 basis point is 0.01%. If the output token amount exceeds the slippage then the swap transaction will fail. | 
 **autoSlippage** | **optional.Bool**| Automatically calculate the slippage based on pairs. | 
 **autoSlippageCollisionUsdValue** | **optional.Int32**| Automatic slippage collision value. | 
 **computeAutoSlippage** | **optional.Bool**| Compute auto slippage value without using it. | 
 **maxAutoSlippageBps** | **optional.Int32**| Max slippage in basis points for auto slippage calculation. Default is 400. | 
 **swapMode** | **optional.String**| (ExactIn or ExactOut) Defaults to ExactIn. ExactOut is for supporting use cases where you need an exact token amount, like payments. In this case the slippage is on the input token. | 
 **dexes** | [**optional.Interface of []string**](string.md)| Default is that all DEXes are included. You can pass in the DEXes that you want to include only and separate them by &#x60;,&#x60;. You can check out the full list [here](https://quote-api.jup.ag/v6/program-id-to-label). | 
 **excludeDexes** | [**optional.Interface of []string**](string.md)| Default is that all DEXes are included. You can pass in the DEXes that you want to exclude and separate them by &#x60;,&#x60;. You can check out the full list [here](https://quote-api.jup.ag/v6/program-id-to-label). | 
 **restrictIntermediateTokens** | **optional.Bool**| Restrict intermediate tokens to a top token set that has stable liquidity. This will help to ease potential high slippage error rate when swapping with minimal impact on pricing. | 
 **onlyDirectRoutes** | **optional.Bool**| Default is false. Direct Routes limits Jupiter routing to single hop routes only. | 
 **asLegacyTransaction** | **optional.Bool**| Default is false. Instead of using versioned transaction, this will use the legacy transaction. | 
 **platformFeeBps** | **optional.Int32**| If you want to charge the user a fee, you can specify the fee in BPS. Fee % is taken out of the output token. | 
 **maxAccounts** | **optional.Int32**| Rough estimate of the max accounts to be used for the quote, so that you can compose with your own accounts | 
 **minimizeSlippage** | **optional.Bool**| Default is false. Miminize slippage attempts to find routes with lower slippage. | 
 **preferLiquidDexes** | **optional.Bool**| Default is false. Enabling it would only consider markets with high liquidity to reduce slippage. | 

### Return type

[**QuoteResponse**](QuoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwapInstructionsPost**
> SwapInstructionsResponse SwapInstructionsPost(ctx, body)
POST /swap-instructions

Returns instructions that you can use from the quote you get from `/quote`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SwapRequest**](SwapRequest.md)|  | 

### Return type

[**SwapInstructionsResponse**](SwapInstructionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwapPost**
> SwapResponse SwapPost(ctx, body)
POST /swap

Returns a transaction that you can use from the quote you get from `/quote`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SwapRequest**](SwapRequest.md)|  | 

### Return type

[**SwapResponse**](SwapResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TokensGet**
> []string TokensGet(ctx, )
GET /tokens

Returns a list of all the tradable mints

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

