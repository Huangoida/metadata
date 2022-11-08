package util

import (
	"github.com/bwmarrin/snowflake"
	"sync"
)

var node *snowflake.Node
var one sync.Once

func GenerateId() int64 {
	one.Do(func() {
		node, _ = snowflake.NewNode(1)
	})
	id := node.Generate()
	return id.Int64()
}
