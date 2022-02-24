dapr_name="web-port"

go build --tags $dapr_name &&

dapr run --app-id $dapr_name \
         --app-protocol http \
         --app-port 9080 \
         --dapr-http-port 3500 \
         --log-level debug \
         --components-path ../config/components \
         --config ../config/config.yaml \
         ./$dapr_name