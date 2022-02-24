dapr_name="dapr-api-go"

go build --tags $dapr_name &&

dapr run --app-id $dapr_name \
         --app-protocol grpc \
         --app-port 9003 \
         --dapr-grpc-port 4501 \
         --log-level debug \
         --components-path ../config/components \
         --config ../config/config.yaml \
         ./$dapr_name
