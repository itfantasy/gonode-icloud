package lobby

import (
	"fmt"
	"strings"

	"github.com/itfantasy/gonode"
	"github.com/itfantasy/gonode/behaviors/gen_server"
	"github.com/itfantasy/gonode/utils/ini"
	"github.com/itfantasy/gonode/utils/io"
)

type LobbyBoot struct {
	server LobbyServer
}

func (this *LobbyBoot) Setup() *gen_server.NodeInfo {
	conf, err := ini.Load(io.CurDir() + "conf.ini")
	if err != nil {
		return nil
	}
	nodeInfo := new(gen_server.NodeInfo)

	nodeInfo.Id = conf.Get("node", "id")
	nodeInfo.Url = conf.Get("node", "url")
	nodeInfo.PubUrl = conf.Get("node", "puburl")
	nodeInfo.BackEnds = conf.Get("node", "backends")

	nodeInfo.LogLevel = conf.Get("log", "loglevel")
	nodeInfo.LogComp = conf.Get("log", "logcomp")

	nodeInfo.RegComp = conf.Get("reg", "regcomp")

	return nodeInfo
}

func (this *LobbyBoot) Start() {
	fmt.Println("node starting...")
	this.server.Start()
}
func (this *LobbyBoot) OnConn(id string) {
	fmt.Println("new conn !! " + id)
}
func (this *LobbyBoot) OnMsg(id string, msg []byte) {
	if strings.Contains(id, "room") {
		this.server.OnServerMsg(id, msg)
	} else {
		this.server.OnMsg(id, msg)
	}
}
func (this *LobbyBoot) OnClose(id string) {
	fmt.Println("conn closed !! " + id)
}
func (this *LobbyBoot) Initialize(server LobbyServer) {
	this.server = server
	gonode.Node().Initialize(this)
}
