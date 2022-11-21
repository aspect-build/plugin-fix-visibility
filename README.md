# Fix-visibility plugin

This is a plugin for the [Aspect CLI].

It listens to Bazel's build event protocol and filters for the Aborted_ANALYSIS_FAILURE event.
If the Analysis Failure is due to a message containing `is not visible from target` this indicates
that the `visibility` attribute of a target doesn't include the package where our target is defined.

After the build completes, the plugin offers to repair the problem by adding the missing `visibility` entry.

In this demo, we uncomment the `alias` target from `example/BUILD.bazel` and run `bazel build example` to see the failure.
The plugin offers to automatically add the missing entry to `visibility` of the `//:dev` target,
then we run the same build again, and it passes.
The demo then runs `git diff` so you can see what edit was made.

[![asciicast](https://asciinema.org/a/1IRPgMQmhJC3L8RM1XTwRYUfa.svg)](https://asciinema.org/a/1IRPgMQmhJC3L8RM1XTwRYUfa)

## Developing

To try the plugin, first check that you have Aspect installed, by running `bazel version` and checking for
`Aspect CLI version` in the output.

First build the plugin from source:

```bash
% bazel build ...
```

Note that the `.aspect/cli/plugins.yaml` file has a reference to the path under `bazel-bin` where the plugin binary was just written.
On the first build, you'll see a warning printed that the plugin doesn't exist at this path.
This is just the development flow for working on plugins; users will reference the plugin's releases which are downloaded for them automatically.

## Releasing

Just push a tag to your GitHub repo.
The actions integration will create a release.

[bazelisk]: https://bazel.build/install/bazelisk
[aspect cli]: https://aspect.build/cli
[plugin documentation]: https://docs.aspect.build/aspect-build/aspect-cli/5.0.1/docs/help/topics/plugins.html
[aspect cli release]: https://github.com/aspect-build/aspect-cli/releases
