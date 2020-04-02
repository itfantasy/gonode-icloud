package main

import (
	"fmt"

	"github.com/itfantasy/gonode"
	"github.com/itfantasy/gonode/behaviors/gen_server"
	"github.com/itfantasy/gonode/utils/ini"
	"github.com/itfantasy/gonode/utils/io"

	"github.com/itfantasy/gonode-icloud/icloud/logics/game"
	"github.com/itfantasy/gonode-toolkit/toolkit/gen_room"
)

type RoomServer struct {
}

func (r *RoomServer) Setup() *gen_server.NodeInfo {
	conf, err := ini.Load(io.CurDir() + "conf.ini")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	info := new(gen_room.RoomServerInfo)
	info.Id = conf.Get("node", "id")
	info.Url = conf.Get("node", "url")
	info.LogLevel = conf.Get("log", "loglevel")
	info.LogComp = conf.Get("log", "logcomp")
	info.RegComp = conf.Get("reg", "regcomp")
	info.PubDomain = conf.Get("node", "pubdomain")
	if err := gen_room.InitGameDB(conf.Get("gamedb", "comp")); err != nil {
		return nil
	}
	return info.ExpandToNodeInfo()
}
func (r *RoomServer) Start() {

}
func (r *RoomServer) OnConn(id string) {
	fmt.Println("new conn !! " + id)
	if gonode.IsPeer(id) {
		game.HandleConn(id)
	}
}
func (r *RoomServer) OnMsg(id string, msg []byte) {
	if gonode.IsPeer(id) {
		game.HandleMsg(id, msg)
	}
}
func (r *RoomServer) OnClose(id string, reason error) {
	fmt.Println("conn closed !! " + id + " -- reason:" + reason.Error())
	if gonode.IsPeer(id) {
		game.HandleClose(id)
	}
}

func main() {
	gonode.Bind(new(RoomServer))
	gonode.Launch()
}
