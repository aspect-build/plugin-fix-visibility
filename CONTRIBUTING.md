## Developing

To try the plugin, first check that you have Aspect installed, by running `bazel version` and checking for
`Aspect CLI version` in the output.

First build the plugin from source:

```bash
% bazel build ...
```

Note that the `.aspect/cli/config.yaml` file has a reference to the path under `bazel-bin` where the plugin binary was just written.
On the first build, you'll see a warning printed that the plugin doesn't exist at this path.
This is just the development flow for working on plugins; users will reference the plugin's releases which are downloaded for them automatically.

## Releasing

Just push a tag to your GitHub repo.
The actions integration will create a release.
