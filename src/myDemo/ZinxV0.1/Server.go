package main

import "zin/src/zinx/znet"

func main() {
	server := znet.NewServer("[zinx V0.1]")
	server.Serve()
}
