package cli

import (
	"fmt"

	"github.com/kschamplin/gotelem"
	"github.com/kschamplin/gotelem/mprpc"
	"github.com/urfave/cli/v2"
	imgui "github.com/AllenDang/cimgui-go"
)


func init() {
	subCmds = append(subCmds, clientCmd)
}


var clientCmd = &cli.Command{
	Name: "client",
	Aliases: []string{"c"},
	Usage: "interact with a gotelem server",
	ArgsUsage: "[server url]",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name: "gui",
			Aliases: []string{"g"},
			Usage: "start a local TUI",
		},
	},
	Description: `
Connects to a gotelem server or relay. Can be used to 
	`,
	Action: client,
}


func loop() {
	imgui.ShowDemoWindow()	
}

func client(ctx *cli.Context) error {
	backend := imgui.CreateBackend()
	backend.CreateWindow("hi there", 1200, 900, 0)
	backend.Run(loop)
	return nil
}


// the client should connect to a TCP server and listen to packets.
func CANFrameHandler(f *gotelem.Frame) (*mprpc.RPCEmpty, error){
	fmt.Printf("got frame, %v\n", f)
	return nil, nil
}

var initialRPCHandlers = map[string]mprpc.ServiceFunc{
	"can": mprpc.MakeService(CANFrameHandler),
}
