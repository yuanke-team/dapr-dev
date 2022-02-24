dapr_name="near-api-rs"

dapr run --app-id $dapr_name \
         --app-protocol grpc \
         --app-port 50051 \
         --dapr-grpc-port 51051 \
         --log-level debug \
         --components-path ../config/components \
         --config ../config/config.yaml \
         ./target/release/$dapr_name
