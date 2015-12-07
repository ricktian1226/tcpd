package main

import (
	"flag"
	"fmt"
	//"time"

	//proto "code.google.com/p/goprotobuf/proto"

	//"feiyu.com/protocol"
	"github.com/funny/link"
	"github.com/funny/link/example/codec"

	//"feiyu.com/xiaoyao/error"
	//"feiyu.com/xiaoyao/log"
)

func main() {
	var address string
	flag.StringVar(&address, "addr", "192.168.93.129:10003", "server address")
	flag.Parse()

	session, err := link.Connect("tcp", address, codec.Bytes(codec.Uint16BE))
	//session, err := link.Connect("tcp", address, codec.ProtoBuf(codec.Uint16BE))
	if err != nil {
		panic(err)
	}

	go func() {
		var msg []byte
		//msg := tcpd_proto.TcpdTest{}

		for {
			if err := session.Receive(&msg); err != nil {
				break
			}
			fmt.Printf("receive (%v)\n", string(msg))
			//fmt.Printf("uid(%d) msg(%s)\n", msg.GetUid(), msg.GetMsg())
			//xylog.DebugNoId("shit here")
		}
	}()

	for {

		var str string
		if _, err := fmt.Scanf("%s\n", &str); err != nil {
			fmt.Printf("session.Send failed : %v", err)
			break
		}

		msg := []byte(str)
		//msg := tcpd_proto.TcpdTest{
		//	Uid: proto.Uint64(1),
		//	Msg: proto.String(str),
		//}

		//xylog.DebugNoId("i will send : %v", msg)
		if err = session.Send(msg); err != nil {
			//xylog.DebugNoId(" session.Send failed : %v", err)
			fmt.Printf("session.Send failed : %v", err)
			break
		}

		//time.Sleep(time.Second * 10)

	}

	session.Close()
	println("bye bitch")

}
