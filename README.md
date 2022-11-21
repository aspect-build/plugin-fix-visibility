# Template for aspect-cli plugins

This repo provides the fastest way to make a plugin for the [aspect cli].

It contains a plugin written in Go, with a GitHub actions CI/CD pipeline to release it.

More details about aspect cli plugins is on the [plugin documentation].

## Instructions

Create a new repo with the green "Use this template" button above.
Then in your repo...

1. Find-and-replace `hello_world` with your plugin name.
1. Find-and-replace `github.com/aspect-build/aspect-cli-plugin-template` with the name of your Go module. See <https://go.dev/doc/modules/developing>
1. Delete everything above the SNIP line below, and start coding on your features!

---------- %<  SNIP %< ------------

# My Plugin

This is a plugin for the Aspect CLI.

## Developing

To try the plugin, first check that you have the most recent [aspect cli release] installed.

First build the plugin from source:

```bash
% bazel build ...
```

Note that the `.aspect/cli/plugins.yaml` file has a reference to the path under `bazel-bin` where the plugin binary was just written.
On the first build, you'll see a warning printed that the plugin doesn't exist at this path.
This is just the development flow for working on plugins; users will reference the plugin's releases which are downloaded for them automatically.

Now just run `aspect`. You should see that `hello-world` appears in the help output. This shows that our plugin was loaded and contributed a custom command to the CLI.

```
Usage:
  aspect [command]

Custom Commands from Plugins:
  hello-world        Print 'Hello World!' to the command line.
```

## Releasing

Just push a tag to your GitHub repo.
The actions integration will create a release.

[bazelisk]: https://bazel.build/install/bazelisk
[aspect cli]: https://aspect.build/cli
[plugin documentation]: https://docs.aspect.build/aspect-build/aspect-cli/5.0.1/docs/help/topics/plugins.html
[aspect cli release]: https://github.com/aspect-build/aspect-cli/releases
