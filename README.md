# HLF-Transaction-Receipt
In Fabric Network After A Successful Write Transaction It Only Returns A  Status Code (If Don't have hyperledger blockexplorer to check block and transaction). But A Transaction Metadata Is Needed For Track And Trace.

# Conclusions

it was shown that without custom transaction receipt what we get from the network as response and what we get as response from the with custom transaction receipt. 

- In the contract methods are `Init_Asset()`,`Read_Asset()`,`Create_Asset()`,`Update_Asset()` and helper method is `Has_Asset()`(It's Checks If The Asset Exists In Ledger or Not)

- On the `Create_Asset()` & `Update_Asset()` methods custom transaction receipt logic is applied 
- On the `Init_Asset()`,`Read_Asset()` methods custom transaction receipt logic is applied is not applied

Two Screenshot is from after transation one is returning transction receipt and another one is not

