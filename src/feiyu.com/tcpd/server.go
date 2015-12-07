package main

import (
	"flag"
	//"io"
	"fmt"
	"os"
	"time"

	beegoconf "github.com/beego/config"
	"github.com/funny/link"
	"github.com/funny/link/example/codec"

	//"feiyu.com/protocol"
	"feiyu.com/xiaoyao/error"
	"feiyu.com/xiaoyao/log"
)

var (
	configFile   string                    //配置文件名称
	defIniConfig beegoconf.ConfigContainer //配置信息管理器
)

func main() {

	//初始化服务配置信息
	err := parseIniConfig()
	if err != xyerror.ErrOK {
		fmt.Printf("parseIniConfig failed : %v", err)
		panic(err)
	}

	//初始化服务
	var server *link.Server
	//server, err = link.Serve("tcp", defServerConfig.host, codec.ProtoBuf(codec.Uint16BE))
	server, err = link.Serve("tcp", defServerConfig.host, codec.Bytes(codec.Uint16BE))
	if err != xyerror.ErrOK {
		fmt.Printf("link.Serve failed : %v", err)
		panic(err)
	}

	channel := link.NewChannel()

	for {
		session, err := server.Accept()
		if err != nil {
			fmt.Printf("server.Accept failed : %v\n", err)
			break
		}
		session.EnableAsyncSend(1024)
		channel.Join(session)

		go func() {
			for { //死循环接受消息
				var msg []byte
				err = session.Receive(&msg)
				//msg := tcpd_proto.TcpdTest{}
				//err = session.Receive(&msg)
				if err != nil {
					//fmt.Printf("session.Receive failed : %v\n", err)
					xylog.ErrorNoId("session.Receive failed : %v", err)
					break
				}

				//fmt.Printf("session(%d).Receive  : %v\n", session.Id(), msg.String())
				fmt.Printf("session(%d).Receive  : %v\n", session.Id(), string(msg))

				//err = session.Send(&msg)
				//message := fmt.Sprintf("session#%d said : %v", session.Id(), msg)
				//err = channel.Broadcast(&msg)
				err = channel.Broadcast(msg)
				if err != xyerror.ErrOK {
					xylog.ErrorNoId("channel.Broadcast failed : %v", err)
				}
			}

			channel.Exit(session)
			session.Close()
		}()

		go func() {

			for {
				time.Sleep(time.Second * 30)
				server.PrintSessions()
			}

		}()

		//go io.Copy(session.Conn(), session.Conn())

	}
}

const (
	INI_CONFIG_ITEM_SERVER_ADDRESS = "Server::address"
)

func parseIniConfig() (err error) {
	//读取配置文件名
	flag.StringVar(&configFile, "config", os.Args[0]+".ini", "ini file")
	xylog.ProcessCmd()
	flag.Parse()

	//加载ini配置文件
	defIniConfig, err = beegoconf.NewConfig("ini", configFile)
	if err != xyerror.ErrOK {
		fmt.Printf("load config file %s failed : %v\n", configFile, err)
		return
	}

	err = xylog.ProcessIniConfig(defIniConfig)
	if err != xyerror.ErrOK {
		fmt.Printf("ProcessIniConfig failed : %v\n", err)
		return
	}

	//获取服务器ip配置信息
	defServerConfig.host = defIniConfig.String(INI_CONFIG_ITEM_SERVER_ADDRESS)

	return
}
