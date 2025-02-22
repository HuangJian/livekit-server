package selector

import (
	"github.com/livekit/protocol/livekit"
	"github.com/thoas/go-funk"
)

// RandomSelector selects an available node at random
type RandomSelector struct {
}

func (s *RandomSelector) SelectNode(nodes []*livekit.Node) (*livekit.Node, error) {
	nodes = GetAvailableNodes(nodes)
	if len(nodes) == 0 {
		return nil, ErrNoAvailableNodes
	}

	idx := funk.RandomInt(0, len(nodes))
	return nodes[idx], nil
}
