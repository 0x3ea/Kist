package main

import (
	"path/filepath"
	"strings"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
)

var node *snowflake.Node

func snowflakeInit() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err) // 如果初始化失败，直接让程序崩掉
	}
}

func generateUID() string {
	return node.Generate().String()
}

func generateUUID() string {
	u := uuid.New().String()
	u = u[0:8] + u[9:13] + u[14:18] + u[19:23] + u[24:]
	return u
}


