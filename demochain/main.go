package demochain

import (
	demochain "./core/demochain"
)

func Run() {
	bc := demochain.NewBlockchain()
	bc.SendData("abcd")
	bc.SendData("efg")
	bc.Print()
}
