package swim

import (
	"context"

	"github.com/mdlfop/fun-with-gossip/proto/rpc"
)

type State uint64

const (
	ALIVE State = iota
	DEAD
)

type Node struct {
	id         uint64
	port       uint16
	state      State
	memberList MemberList
}

func (node *Node) Ping(context.Context, *PingRequest) (*rpc.PingResponse, error) {

}
func (node *Node) PingIndirect(context.Context, *IndirectPingRequest) (*IndirectPingResponse, error) {

}

func (node *Node) PingRandom()         {}
func (node *Node) IndirectPing()       {}
func (node *Node) HandlePingRandom()   {}
func (node *Node) HandleIndirectPing() {}

func NewNode() {}

type Member struct {
	id    uint64
	port  uint16
	state State
}

type MemberList struct {
	Members []Member
}

func NewMemberList() {}

func (ml *MemberList) PingRandom()         {}
func (ml *MemberList) IndirectPing()       {}
func (ml *MemberList) HandlePingRandom()   {}
func (ml *MemberList) HandleIndirectPing() {}
