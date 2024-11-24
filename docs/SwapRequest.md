# SwapRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserPublicKey** | **string** | The user public key. | [default to null]
**WrapAndUnwrapSol** | **bool** | Default is true. If true, will automatically wrap/unwrap SOL. If false, it will use wSOL token account.  Will be ignored if &#x60;destinationTokenAccount&#x60; is set because the &#x60;destinationTokenAccount&#x60; may belong to a different user that we have no authority to close. | [optional] [default to true]
**UseSharedAccounts** | **bool** | Default is true. This enables the usage of shared program accountns. That means no intermediate token accounts or open orders accounts need to be created for the users. But it also means that the likelihood of hot accounts is higher. | [optional] [default to true]
**FeeAccount** | **string** | Fee token account, same as the output token for ExactIn and as the input token for ExactOut, it is derived using the seeds &#x3D; [\&quot;referral_ata\&quot;, referral_account, mint] and the &#x60;REFER4ZgmyYx9c6He5XfaTMiGfdLwRnkV4RPp9t9iF3&#x60; referral contract (only pass in if you set a feeBps and make sure that the feeAccount has been created). | [optional] [default to null]
**ComputeUnitPriceMicroLamports** | [***AnyOfSwapRequestComputeUnitPriceMicroLamports**](AnyOfSwapRequestComputeUnitPriceMicroLamports.md) | The compute unit price to prioritize the transaction, the additional fee will be &#x60;computeUnitLimit (1400000) * computeUnitPriceMicroLamports&#x60;. If &#x60;auto&#x60; is used, Jupiter will automatically set a priority fee and it will be capped at 5,000,000 lamports / 0.005 SOL. | [optional] [default to null]
**PrioritizationFeeLamports** | [***AnyOfSwapRequestPrioritizationFeeLamports**](AnyOfSwapRequestPrioritizationFeeLamports.md) | \\* PriorityFeeWithMaxLamports is impossible to be typed. Prioritization fee lamports paid for the transaction in addition to the signatures fee. Mutually exclusive with compute_unit_price_micro_lamports. If &#x60;auto&#x60; is used, Jupiter will automatically set a priority fee and it will be capped at 5,000,000 lamports / 0.005 SOL. | [optional] [default to null]
**AsLegacyTransaction** | **bool** | Default is false. Request a legacy transaction rather than the default versioned transaction, needs to be paired with a quote using asLegacyTransaction otherwise the transaction might be too large. | [optional] [default to false]
**UseTokenLedger** | **bool** | Default is false. This is useful when the instruction before the swap has a transfer that increases the input token amount. Then, the swap will just use the difference between the token ledger token amount and post token amount. | [optional] [default to false]
**DestinationTokenAccount** | **string** | Public key of the token account that will be used to receive the token out of the swap. If not provided, the user&#x27;s ATA will be used. If provided, we assume that the token account is already initialized. | [optional] [default to null]
**DynamicComputeUnitLimit** | **bool** | When enabled, it will do a swap simulation to get the compute unit used and set it in ComputeBudget&#x27;s compute unit limit. This will increase latency slightly since there will be one extra RPC call to simulate this. Default is &#x60;false&#x60;. | [optional] [default to null]
**SkipUserAccountsRpcCalls** | **bool** | When enabled, it will not do any rpc calls check on user&#x27;s accounts. Enable it only when you already setup all the accounts needed for the trasaction, like wrapping or unwrapping sol, destination account is already created. | [optional] [default to null]
**ProgramAuthorityId** | **int32** | The program authority id [0;7], load balanced across the available set by default | [optional] [default to null]
**AllowOptimizedWrappedSolTokenAccount** | **bool** | Default is false. Enabling it would reduce use an optimized way to open WSOL that reduce compute unit. | [optional] [default to false]
**QuoteResponse** | [***QuoteResponse**](QuoteResponse.md) |  | [default to null]
**DynamicSlippage** | [***SwapRequestDynamicSlippage**](SwapRequest_dynamicSlippage.md) |  | [optional] [default to null]
**BlockhashSlotsToExpiry** | **float64** | Optional. When passed in, Swap object will be returned with your desired slots to epxiry. | [optional] [default to null]
**CorrectLastValidBlockHeight** | **bool** | Optional. Default to false. Request Swap object to be returned with the correct blockhash prior to Agave 2.0. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

