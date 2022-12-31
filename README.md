# Fix-visibility plugin

This is a plugin for the [Aspect CLI](https://aspect.build/cli).

It listens to Bazel's build event protocol and filters for the Aborted_ANALYSIS_FAILURE event.
If the Analysis Failure is due to a message containing `is not visible from target` this indicates
that the `visibility` attribute of a target doesn't include the package where our target is defined.

After the build completes, the plugin offers to repair the problem by adding the missing `visibility` entry.

## Demo

In this demo, we uncomment the `alias` target from `example/BUILD.bazel` and run `bazel build example` to see the failure.
The plugin offers to automatically add the missing entry to `visibility` of the `//:dev` target,
then we run the same build again, and it passes.
The demo then runs `git diff` so you can see what edit was made.

[![asciicast](https://asciinema.org/a/1IRPgMQmhJC3L8RM1XTwRYUfa.svg)](https://asciinema.org/a/1IRPgMQmhJC3L8RM1XTwRYUfa)
