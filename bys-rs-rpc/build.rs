fn main() {
    tonic_build::configure()
        .out_dir("src/proto") // 生成代码的存放目录
        .compile(
            &["../proto/host_status_service.proto"], // 欲生成的 proto 文件列表
            &["../proto"],                           // proto 依赖所在的根目录
        )
        .unwrap_or_else(|e| panic!("failed to compile protos {:?}", e));
}
