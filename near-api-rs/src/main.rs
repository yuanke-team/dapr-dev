use near_jsonrpc_client::{methods, JsonRpcClient};
use near_jsonrpc_primitives::types::transactions::TransactionInfo;
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    
    let mainnet_client = JsonRpcClient::connect("https://archival-rpc.mainnet.near.org");
    
    let tx_status_request = methods::tx::RpcTransactionStatusRequest {
        transaction_info: TransactionInfo::TransactionId {
            hash: "9FtHUFBQsZ2MG77K3x3MJ9wjX3UT8zE1TczCrhZEcG8U".parse()?,
            account_id: "miraclx.near".parse()?,
        },
    };
    
    // call a method on the server via the connected client
    let tx_status = mainnet_client.call(tx_status_request).await?;
    
    println!("{:?}", tx_status);
}
