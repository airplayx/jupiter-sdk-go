# Go API client for Jupiter

***Initially generated by: [Swagger Codegen](https://github.com/swagger-api/swagger-codegen.git)***

***Last updated by @Halimao***

The core of [jup.ag](https://jup.ag). Easily get a quote and swap through Jupiter API.  

- Rate Limit We update our rate limit from time to time depending on the load of our servers. We recommend running your own instance of the API if you want to have high rate limit, here to learn how to run the [self-hosted API](https://station.jup.ag/docs/apis/self-hosted).  

- API Wrapper - Typescript [@jup-ag/api](https://github.com/jup-ag/jupiter-quote-api-node)  

- Data types - Public keys are base58 encoded strings - raw data such as Vec<u8\> are base64 encoded strings 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 6.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
```golang
import "github.com/airplayx/jupiter-sdk-go"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.jup.ag/swap/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**IndexedRouteMapGet**](docs/DefaultApi.md#indexedroutemapget) | **Get** /indexed-route-map | GET /indexed-route-map
*DefaultApi* | [**ProgramIdToLabelGet**](docs/DefaultApi.md#programidtolabelget) | **Get** /program-id-to-label | GET /program-id-to-label
*DefaultApi* | [**QuoteGet**](docs/DefaultApi.md#quoteget) | **Get** /quote | GET /quote
*DefaultApi* | [**SwapInstructionsPost**](docs/DefaultApi.md#swapinstructionspost) | **Post** /swap-instructions | POST /swap-instructions
*DefaultApi* | [**SwapPost**](docs/DefaultApi.md#swappost) | **Post** /swap | POST /swap
*DefaultApi* | [**TokensGet**](docs/DefaultApi.md#tokensget) | **Get** /tokens | GET /tokens

## Documentation For Models

 - [AccountMeta](docs/AccountMeta.md)
 - [AnyOfSwapRequestComputeUnitPriceMicroLamports](docs/AnyOfSwapRequestComputeUnitPriceMicroLamports.md)
 - [AnyOfSwapRequestPrioritizationFeeLamports](docs/AnyOfSwapRequestPrioritizationFeeLamports.md)
 - [IndexedRouteMapResponse](docs/IndexedRouteMapResponse.md)
 - [Instruction](docs/Instruction.md)
 - [PlatformFee](docs/PlatformFee.md)
 - [PriorityFeeWithMaxLamports](docs/PriorityFeeWithMaxLamports.md)
 - [QuoteResponse](docs/QuoteResponse.md)
 - [RoutePlanStep](docs/RoutePlanStep.md)
 - [SwapInfo](docs/SwapInfo.md)
 - [SwapInstructionsResponse](docs/SwapInstructionsResponse.md)
 - [SwapMode](docs/SwapMode.md)
 - [SwapRequest](docs/SwapRequest.md)
 - [SwapRequestDynamicSlippage](docs/SwapRequestDynamicSlippage.md)
 - [SwapResponse](docs/SwapResponse.md)
 - [SwapResponseDynamicSlippageReport](docs/SwapResponseDynamicSlippageReport.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author
@Halimao

