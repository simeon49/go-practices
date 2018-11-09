package demochain

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./core/demochain"
)

var bc *demochain.Blockchain

func startServer() {
	fmt.Println("启动服务: http://localhost:8080")
	http.HandleFunc("/", blockchainsGetHandle)
	http.HandleFunc("/blockchains/get", blockchainsGetHandle)
	http.HandleFunc("/blockchains/write", blockchainsWriteHandle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func blockchainsGetHandle(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(bc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(bytes))
}

func blockchainsWriteHandle(w http.ResponseWriter, r *http.Request) {
	bc.SendData(r.URL.Query().Get("data"))
	blockchainsGetHandle(w, r)
}

func Run() {
	fmt.Println("初始区块链...")
	bc = demochain.NewBlockchain()
	bc.SendData("abcd")
	bc.SendData("efg")
	bc.Print()

	startServer()
}
