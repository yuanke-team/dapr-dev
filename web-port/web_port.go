package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	dapr "github.com/dapr/go-sdk/client"

	"google.golang.org/grpc"
)

func httpClientSend(image []byte, w http.ResponseWriter) {
	client := &http.Client{}
	println("httpClientSend ....")

	// Dapr api format: http://localhost:<daprPort>/v1.0/invoke/<appId>/method/<method-name>
	var uri string
	// if api == "rust" {
	uri = "http://localhost:3502/v1.0/invoke/image-api-rs/method/api/image"
	// } else {
	// 	uri = "http://localhost:3503/v1.0/invoke/image-api-wasi-socket-rs/method/image"
	// }
	println("uri: ", uri)
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(image))

	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	println(resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	res := string(body)
	println("res: ", res)
	if strings.Contains(res, "Max bytes limit exceeded") {
		res = "ImageTooLarge"
	}
	w.Header().Set("Content-Type", "image/png")
	fmt.Fprintf(w, "%s", res)
}

func daprHttpClientSend(server, dir string, body []byte, w http.ResponseWriter) {
	ctx := context.Background()

	// create the client
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}

	content := &dapr.DataContent{
		ContentType: "application/json",
		Data:        body,
	}

	resp, err := client.InvokeMethodWithContent(ctx, server, dir, "post", content)
	if err != nil {
		panic(err)
	}
	log.Printf("dapr-wasmedge-go method api/image has invoked, response: %s", string(resp))
	fmt.Printf("Image classify result: %q\n", resp)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", string(resp))
}

func daprGrpcClientSend(server, dir string, body []byte, w http.ResponseWriter) {
	ctx := context.Background()

	// create the client
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}

	content := &dapr.DataContent{
		ContentType: "application/json",
		Data:        body,
	}

	resp, err := client.InvokeMethodWithContent(ctx, server, dir, "post", content)
	if err != nil {
		panic(err)
	}
	log.Printf("dapr-wasmedge-go method api/image has invoked, response: %s", string(resp))
	fmt.Printf("Image classify result: %q\n", resp)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", string(resp))
}

func grpcHandler(w http.ResponseWriter, r *http.Request) {
	println("imageHandler ....")
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		println("error: ", err.Error())
		panic(err)
	}
	daprGrpcClientSend("dapr-api-go", "grcp", body, w)
}

func daprGrpcClientPublishEvent(body []byte, w http.ResponseWriter) {

	pubsubName := "pubsub-messages"
	topicName := "neworder"

	ctx := context.Background()
	// data := []byte("{ \"message\": \"hello\" }")

	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	// defer client.Close()

	if err := client.PublishEvent(ctx, pubsubName, topicName, body); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", string("Publish Success"))
}

func pubshHandler(w http.ResponseWriter, r *http.Request) {
	println("pubshHandler ....")
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		println("error: ", err.Error())
		panic(err)
	}
	daprGrpcClientPublishEvent(body, w)
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	println("ipHandler ....")
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		println("error: ", err.Error())
		panic(err)
	}
	daprGrpcClientSend("dapr-api-go", "ip", body, w)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	println("helloWorldHandler ....")
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		println("error: ", err.Error())
		panic(err)
	}
	daprGrpcClientSend("nodejs-server", "bsc-hello-world", body, w)
}

func bscStorageHandler(w http.ResponseWriter, r *http.Request) {
	println("bscStorageHandler ....")
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		println("error: ", err.Error())
		panic(err)
	}
	daprGrpcClientSend("nodejs-server", "bsc-storage", body, w)
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	println("imgHandler ....")
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		println("error: ", err.Error())
		panic(err)
	}
	daprGrpcClientSend("image-api-go", "/api/image", body, w)
}
func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/go/api/image", imageUploadFileHandler)
	// http.HandleFunc("/api/hello", imgHandler)

	http.HandleFunc("/grpc", grpcHandler)
	http.HandleFunc("/pubsh", pubshHandler)
	http.HandleFunc("/ip", ipHandler)
	http.HandleFunc("/hello-world", helloWorldHandler)
	http.HandleFunc("/bsc-storage", bscStorageHandler)

	println("listen to 9080 ...")
	log.Fatal(http.ListenAndServe(":9080", nil))
}

// clint conn to grpc-server
func main_cliet_grpc() {
	// Testing 40 MB data exchange
	maxRequestBodySize := 40
	var opts []grpc.CallOption

	// Receive 40 MB + 1 MB (data + headers overhead) exchange
	headerBuffer := 1
	opts = append(opts, grpc.MaxCallRecvMsgSize((maxRequestBodySize+headerBuffer)*1024*1024))
	conn, err := grpc.Dial(net.JoinHostPort("127.0.0.1",
		"4501"),
		grpc.WithDefaultCallOptions(opts...), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection to 4501 ...")
	client := dapr.NewClientWithConnection(conn)

	pubsubName := "pubsub-messages"
	topicName := "neworder"

	ctx := context.Background()
	data := []byte("{ \"message\": \"hello\" }")

	if err := client.PublishEvent(ctx, pubsubName, topicName, data); err != nil {
		panic(err)
	}
	fmt.Println("data published")

	fmt.Println("Done (CTRL+C to Exit)")

	defer client.Close()
}

func GetEnvValue(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
