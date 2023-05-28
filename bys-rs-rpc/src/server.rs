use std::time::{SystemTime, UNIX_EPOCH};

use proto::hoststatus::{
    host_status_service_server::{HostStatusService, HostStatusServiceServer},
    GetStatusResponse,
};
use tonic::transport::Server;

mod proto;

#[derive(Debug)]
pub struct HostStatusServer {}

#[tonic::async_trait]
impl HostStatusService for HostStatusServer {
    async fn get_status(
        &self,
        request: tonic::Request<proto::hoststatus::GetStatusRequest>,
    ) -> Result<tonic::Response<GetStatusResponse>, tonic::Status> {
        match request.remote_addr() {
            Some(addr) => {
                println!("recv a request from {:?}", addr.to_string());
            },
            None => {
                println!("recv a request from UNKOWN");
            }
        }

        let ts_ms = SystemTime::now().duration_since(UNIX_EPOCH).unwrap();
        let resp = GetStatusResponse {
            hostname: "rust.host_status_service".to_string(),
            timestamp_ms: ts_ms.as_millis() as i64,
        };

        Ok(tonic::Response::new(resp))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "127.0.0.1:10086";
    let hs = HostStatusServer {};

    println!("host status server listening on {}", addr);

    let svc = HostStatusServiceServer::new(hs);

    Server::builder()
        .add_service(svc)
        .serve(addr.parse().unwrap())
        .await?;
    Ok(())
}
