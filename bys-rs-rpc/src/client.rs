use proto::hoststatus::{host_status_service_client::HostStatusServiceClient, GetStatusRequest};
use tonic::Request;

mod proto;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "127.0.0.1:10086"; // TODO
    let dst = format!("http://{}", addr);
    let mut client = HostStatusServiceClient::connect(dst).await?;

    let resp = client.get_status(Request::new(GetStatusRequest {})).await?;

    println!("{:?}", resp);
    Ok(())
}
