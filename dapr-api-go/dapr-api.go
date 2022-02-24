package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	// daprd "github.com/dapr/go-sdk/service/http"
	daprd "github.com/dapr/go-sdk/service/grpc"
)

func sayhelloHandler(_ context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	fmt.Println(string(in.Data))
	out = &common.Content{
		Data:        []byte("grcp go api test"),
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return out, nil
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName:%s, Topic:%s, ID:%s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	return true, nil
}

func getIpHandler(_ context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	out = &common.Content{
		Data:        []byte(getIp()),
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return out, nil
}

func getIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	s, err := daprd.NewService(":9003")
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
	if err := s.AddServiceInvocationHandler("grcp", sayhelloHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

	if err := s.AddServiceInvocationHandler("ip", getIpHandler); err != nil {
		log.Fatalf("error getIpHandler invocation handler: %v", err)
	}

	// add some topic subscriptions
	var sub = &common.Subscription{
		PubsubName: "pubsub-messages",
		Topic:      "neworder",
		Route:      "/orders",
	}
	if err := s.AddTopicEventHandler(sub, eventHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}
