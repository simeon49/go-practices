package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app      = kingpin.New("myApp", "第三方命令输入解析库")
	debug    = app.Flag("debug", "enable debug mode").Default("false").Bool()
	serverIP = app.Flag("server", "server address").Default("127.0.0.1").IP()

	register     = app.Command("register", "Register a new user.")
	registerNick = register.Arg("nick", "nickname for user").Required().String()
	registerName = register.Arg("name", "name of user").Required().String()

	post        = app.Command("post", "Post a message to a channel.")
	postImage   = post.Flag("image", "image to post").ExistingFile()
	postChannel = post.Arg("channel", "channel to post to").Required().String()
	postText    = post.Arg("text", "text to post").String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Register user
	case "register":
		println(*registerNick, *registerName)

	// Post message
	case "post":
		if *postImage != "" {
		}
		if *postText != "" {
		}
	}
}
