package utils

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func InitSnowflake(nodeID int64) {
	var err error
	node, err = snowflake.NewNode(nodeID)
	if err != nil {
		panic("snowflake init faild: " + err.Error())
	}
}

func GenID() int64 {
	return node.Generate().Int64()
}
