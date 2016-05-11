package main

import (
	p "github.com/d-smith/go-examples/plugin"
	"github.com/hashicorp/go-plugin"
	"os/exec"
	"log"
	"fmt"
	//"io/ioutil"

	"time"
)

func main() {
	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"greeter": new(p.GreeterPlugin),
	}

	//log.SetOutput(ioutil.Discard)


	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("../server/server"),
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("greeter")
	if err != nil {
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	greeter := raw.(p.Greeter)

	for i:=0; i < 10; i++ {
		fmt.Println(greeter.Greet())
		time.Sleep(time.Second)
	}
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}


