package main

import (
	"context"
	"fmt"
	"os"

	goplugin "github.com/hashicorp/go-plugin"
	"github.com/manifoldco/promptui"

	"aspect.build/cli/bazel/buildeventstream"
	"aspect.build/cli/bazel/command_line"
	"aspect.build/cli/pkg/bazel"
	"aspect.build/cli/pkg/ioutils"
	"aspect.build/cli/pkg/plugin/sdk/v1alpha3/config"
	aspectplugin "aspect.build/cli/pkg/plugin/sdk/v1alpha3/plugin"
)

// main starts up the plugin as a child process of the CLI and connects the gRPC communication.
func main() {
	goplugin.Serve(config.NewConfigFor(&HelloWorldPlugin{}))
}

// HelloWorldPlugin declares the fields on an instance of the plugin.
type HelloWorldPlugin struct {
	// Base gives default implementations of the plugin methods, so implementing them below is optional.
	// See the definition of aspectplugin.Base for more methods that can be implemented by the plugin.
	aspectplugin.Base
	// This plugin will store some state from the Build Events for use at the end of the build.
	command_line.CommandLine
}

// CustomCommands contributes a new 'hello-world' command alongside the built-in ones like 'build' and 'test'.
func (plugin *HelloWorldPlugin) CustomCommands() ([]*aspectplugin.Command, error) {
	return []*aspectplugin.Command{
		aspectplugin.NewCommand(
			"hello-world",
			"Print 'Hello World!' to the command line.",
			"Print 'Hello World!' to the command line. Echo any given argument. Then run a 'bazel help'",
			func(ctx context.Context, args []string, bzl bazel.Bazel) error {
				fmt.Println("Hello World!")
				fmt.Print("Arguments passed to command: ")
				fmt.Println(args)
				fmt.Println("Going to run: 'bazel help'")

				bzl.RunCommand(ioutils.DefaultStreams, "help")

				return nil
			},
		),
	}, nil
}
// BEPEventCallback subscribes to all Build Events, and lets our logic react to ones we care about.
func (plugin *HelloWorldPlugin) BEPEventCallback(event *buildeventstream.BuildEvent) error {
	switch event.Payload.(type) {
		case *buildeventstream.BuildEvent_StructuredCommandLine:
			commandLine := *event.GetStructuredCommandLine()
			if commandLine.CommandLineLabel == "canonical" {
				plugin.CommandLine = commandLine
			}
	}
	return nil
}

// PostBuildHook will be called at the end of an `aspect build` execution, after Bazel completes.
func (plugin *HelloWorldPlugin) PostBuildHook(
	isInteractiveMode bool,
	promptRunner ioutils.PromptRunner,
) error {
	// We condition prompting on whether there's an interactive user to engage with.
	if isInteractiveMode {
		// The manifoldco/promptui library creates many styles of interactive prompts.
		// Check out the examples: https://github.com/manifoldco/promptui/tree/master/_examples
		prompt := promptui.Prompt{
			Label:     "Thanks for trying the hello-world plugin! Would you like to see the command that was run",
			IsConfirm: true,
		}
		// Since the prompt is a boolean, any non-nil error should represent a NO.
		if _, err := promptRunner.Run(prompt); err == nil {
			plugin.printTargetPattern()
		}
	}
	return nil
}

// printTargetPattern is just representative of some logic a plugin might want to perform on the data collected.
func (plugin *HelloWorldPlugin) printTargetPattern() {
	for _, section := range plugin.CommandLine.Sections {
		fmt.Fprintf(os.Stdout, "%s\n", section.SectionLabel)
		if section.SectionLabel == "residual" {
			switch f := section.SectionType.(type) {
			case *command_line.CommandLineSection_ChunkList:
				fmt.Fprintf(os.Stdout, "target pattern was %s\n", f.ChunkList.Chunk[0])
			}
		}
	}
}
