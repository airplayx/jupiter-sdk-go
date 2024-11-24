# SwapInstructionsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenLedgerInstruction** | [***Instruction**](Instruction.md) |  | [optional] [default to null]
**ComputeBudgetInstructions** | [**[]Instruction**](Instruction.md) | The necessary instructions to setup the compute budget. | [default to null]
**SetupInstructions** | [**[]Instruction**](Instruction.md) | Setup missing ATA for the users. | [default to null]
**SwapInstruction** | [***Instruction**](Instruction.md) |  | [default to null]
**CleanupInstruction** | [***Instruction**](Instruction.md) |  | [optional] [default to null]
**AddressLookupTableAddresses** | **[]string** | The lookup table addresses that you can use if you are using versioned transaction. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

