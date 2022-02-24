dapr_name="image-api-go"

go build --tags "tensorflow image" &&
dapr run --app-id $dapr_name \
         --app-protocol grpc \
         --app-port 9003 \
         --dapr-grpc-port 3501 \
         --log-level debug \
         --components-path ../config/components \
         --config ../config/config.yaml \
         ./$dapr_name
