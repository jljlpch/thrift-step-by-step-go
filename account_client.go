package main

import (
	"fmt"
	"github.com/LC2010/thrift-step-by-step-go/gen-go/account"
	"github.com/apache/thrift/lib/go/thrift"
	"net"
	"os"
)

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "19090"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := account.NewAccountClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:19090", "", err)
		os.Exit(1)
	}
	defer transport.Close()

	// 第一个请求，登录账号
	request := &account.Request{
		Name:     "super",
		Password: "123",
		Op:       account.Operation_LOGIN,
	}

	r, err := client.DoAction(request)
	fmt.Println(r)

	// 第二个请求，注册账号
	request.Op = account.Operation_REGISTER
	r, err = client.DoAction(request)
	fmt.Println(r)

	// 第三个请求，name为空的请求
	request.Name = ""
	r, err = client.DoAction(request)

}
