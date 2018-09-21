package room

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	//"strings"

	"github.com/itfantasy/gonode"
	"github.com/itfantasy/gonode/behaviors/gen_server"
	"github.com/itfantasy/gonode/utils/ini"
	"github.com/itfantasy/gonode/utils/io"
	//	"github.com/itfantasy/gonode/utils/timer"
)

type RoomBoot struct {
	server RoomServer
}

func (this *RoomBoot) SelfInfo() (*gen_server.NodeInfo, error) {
	conf, err := ini.Load(io.CurDir() + "conf.ini")
	if err != nil {
		return nil, err
	}
	nodeInfo := new(gen_server.NodeInfo)

	nodeInfo.Id = conf.Get("node", "id")
	nodeInfo.Url = conf.Get("node", "url")
	nodeInfo.AutoDetect = conf.GetInt("node", "autodetect", 0) > 0
	nodeInfo.Public = conf.GetInt("node", "public", 0) > 0

	nodeInfo.RedUrl = conf.Get("redis", "url")
	nodeInfo.RedPool = conf.GetInt("redis", "pool", 0)
	nodeInfo.RedDB = conf.GetInt("redis", "db", 0)
	nodeInfo.RedAuth = conf.Get("redis", "auth")

	return nodeInfo, nil
}
func (this *RoomBoot) OnDetect(id string) bool {
	if id == "lobby" { // the room will auto find the lobby, and try to build a conn to the lobby
		return true
	}
	return false
}
func (this *RoomBoot) Start() {
	fmt.Println("node starting...")
}
func (this *RoomBoot) OnConn(id string) {
	this.server.OnConn(id)
}
func (this *RoomBoot) OnMsg(id string, msg []byte) {
	if strings.Contains(id, "lobby") {
		// native logic for lobbyserver
	} else {
		this.server.OnMsg(id, msg)
	}
}
func (this *RoomBoot) OnClose(id string) {
	this.server.OnClose(id)
}
func (this *RoomBoot) OnShell(id string, msg string) {

}
func (this *RoomBoot) OnRanId() string {
	return "cnt" + strconv.Itoa(rand.Intn(100000))
}
func (this *RoomBoot) Initialize(server RoomServer) {
	this.server = server
	gonode.Node().Initialize(this)
}
