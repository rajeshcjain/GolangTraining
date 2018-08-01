package main

import (
	"github.com/alecthomas/kingpin"
	"fmt"
)

/*

THIS IS A PROGRAM TO UNDERSTAND THE "kingpin"

*/


var (
	/*
debug and serverIP are related with all the commands.
--------------------------------------------------
$ go build && ./my-app.exe --help
usage: my-app.exe [<flags>] <command> [<args> ...]

Flags:
  --help              Show context-sensitive help (also try --help-long and
                      --help-man).
  --debug             enable debug mode
  --server=127.0.0.1  server address

Commands:
  help [<command>...]
    Show help.

  register [<flags>] <nick> <name>
    Register a new user.

  post [<flags>] <channel> [<text>]
    Post a message to a channel.

               AND


   when we press
----------------------------------------------------------------
   $ ./my-app.exe register --help
usage: my-app.exe register [<flags>] <nick> <name>

Register a new user.

Flags:
  --help              Show context-sensitive help (also try --help-long and
                      --help-man).
  --debug             enable debug mode
  --server=127.0.0.1  server address
  --ln=LN             enter the last name

Args:
  <nick>  nickname for user
  <name>  name of user

So when we run the above command we can see debug and server along with
nick and name ARG.we have specifically attached it with register.So these
will not appear with post command

which is
-----------------------------------------------------------------------------
$ ./my-app.exe post --help
usage: my-app.exe post [<flags>] <channel> [<text>]

Post a message to a channel.

Flags:
  --help              Show context-sensitive help (also try --help-long and
                      --help-man).
  --debug             enable debug mode
  --server=127.0.0.1  server address
  --image=IMAGE       image to post

Args:
  <channel>  channel to post to
  [<text>]   text to post
----------------------
OUTPUT of the program
----------------------
$ ./my-app.exe register raj chand --ln="jain"
raj
chand
jain
*/
	debug    = kingpin.Flag("debug", "enable debug mode").Default("false").Bool()
	serverIP = kingpin.Flag("server", "server address").Default("127.0.0.1").IP()

	register     = kingpin.Command("register", "Register a new user.")
	registerNick = register.Arg("nick", "nickname for user").Required().String()
	registerName = register.Arg("name", "name of user").Required().String()
        registerFlag = register.Flag("ln", "enter the last name").String()

	post        = kingpin.Command("post", "Post a message to a channel.")
	postImage   = post.Flag("image", "image to post").ExistingFile()
	postChannel = post.Arg("channel", "channel to post to").Required().String()
	postText    = post.Arg("text", "text to post").String()
)

func main() {
	switch kingpin.Parse() {
	// Register user
	case "register":{
		fmt.Println(*registerNick)
		fmt.Println(*registerName)
		fmt.Println(*registerFlag)
	}



		// Post message
	case "post":
		if postImage != nil {
		}
		if *postText != "" {
		}
	}
}
