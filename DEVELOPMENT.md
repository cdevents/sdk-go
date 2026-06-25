# Developing

## Setting up a development environment

### Setup a GitHub account accessible via SSH

GitHub is used for project Source Code Management (SCM) using the SSH protocol for authentication.

1. Create [a GitHub account](https://github.com/join) if you do not already have one.
1. Setup
[GitHub access via SSH](https://help.github.com/articles/connecting-to-github-with-ssh/)

### Install tools

You must install these tools:

1. [`git`](https://help.github.com/articles/set-up-git/): For source control

1. [`go`](https://golang.org/doc/install): The language this SDK is built in.
    Check the [.go-version](.go-version) file for the minimum required version.

1. [`make`](https://www.gnu.org/software/make/): not strictly required but handy to run
   tests with a single command.

### Setup a fork

The sdk-go project requires that you develop (commit) code changes to branches that belong to a fork of the `cdevents/sdk-go` repository in your GitHub account before submitting them as Pull Requests (PRs) to the actual project repository.

1. [Create a fork](https://help.github.com/articles/fork-a-repo/) of the `cdevents/sdk-go` repository in your GitHub account.

1. Create a clone of your fork on your local machine, including submodules:

    ```shell
    git clone --recurse-submodules git@github.com:${YOUR_GITHUB_USERNAME}/sdk-go.git
    ```

    If you have already cloned the repository without `--recurse-submodules`, initialise the submodules separately:

    ```shell
    git submodule update --init --recursive
    ```

    The repository uses git submodules to vendor the [CDEvents spec](https://github.com/cdevents/spec)
    schemas. These live under `pkg/api/spec-v*` (one per supported spec major.minor version).
    The submodules must be present for code generation (`make generate`) and tests to work.

1. Configure `git` remote repositories

    Adding `cdevents/sdk-go` as the `upstream` and your fork as the `origin` remote repositories to your `.git/config` sets you up nicely for regularly [syncing your fork](https://help.github.com/articles/syncing-a-fork/) and submitting pull requests.

    1. Change into the project directory

        ```shell
        cd sdk-go
        ```

    1. Configure sdk-go as the `upstream` repository

        ```shell
        git remote add upstream git@github.com:cdevents/sdk-go.git

        # Optional: Prevent accidental pushing of commits by changing the upstream URL to `no_push`
        git remote set-url --push upstream no_push
        ```

    1. Configure your fork as the `origin` repository

        ```shell
        git remote add origin git@github.com:${YOUR_GITHUB_USERNAME}/sdk-go.git
        ```

## Developing, building and testing

Make target all defined to run unit tests, format imports, format go code and run the linter.

To format the go code and imports:

```shell
make fmt
```

To run the go linter:

```shell
make lint
```

To run unit tests:

```shell
make test
```

To run all targets, before creating a commit:

```shell
make all
```

## Creating a release

### Patch version (e.g. v0.5.0 → v0.5.1)

A patch release updates the SDK to a new patch version of the CDEvents spec
(or fixes SDK bugs) without adding new event types or changing the API surface.

1. **Update the spec submodule** to point to the new spec tag:

    ```shell
    cd pkg/api/spec-v0.5          # use the appropriate major.minor folder
    git fetch origin
    git checkout v0.5.1            # the new spec tag
    cd -
    git add pkg/api/spec-v0.5
    ```

    If the spec schemas have bugs that need workarounds, add patch files under
    `hack/patches/`. The generator applies them automatically before code generation
    and reverts them from the submodule afterwards.

1. **Update `SPEC_VERSIONS` in `tools/generator.go`** — change the version string
   for the affected spec slot (e.g. `"0.5.0"` → `"0.5.1"`). The generator uses
   `semver.MajorMinor` to locate the `spec-v0.5` folder, so the folder name does
   not change, but the full version is embedded in the generated code.

1. **Regenerate the SDK**:

    ```shell
    make generate
    ```

    This deletes all `zz_*` files, re-runs the generator, and checks that the
    working tree is clean. Review the diff to confirm only the expected version
    strings changed.

1. **Run tests**:

    ```shell
    make all
    ```

1. **Open a PR, get it reviewed and merged.**

1. **Tag the release** from `main`:

    ```shell
    git tag -a v0.5.1 -m "Release v0.5.1"
    git push upstream v0.5.1
    ```

1. **Create a GitHub Release** from the new tag with release notes summarising
   spec changes and any SDK-level fixes.

### New spec major.minor version (e.g. v0.5.x → v0.6.0)

A major.minor release adds support for a new version of the CDEvents spec,
which may introduce new event types, new fields, or structural changes.

1. **Add a new spec submodule** for the new major.minor:

    ```shell
    git submodule add -b spec-v0.6 https://github.com/cdevents/spec pkg/api/spec-v0.6
    ```

1. **Update the generator** in `tools/generator.go`:
   - Append the new version to `SPEC_VERSIONS` (e.g. add `"0.6.0"`).
   - If the new spec introduces structural changes to the schema (e.g. new context
     fields), update `DataFromSchema` and the templates under `tools/templates/`
     accordingly.

1. **Regenerate the SDK**:

    ```shell
    make generate
    ```

    This will create a new `pkg/api/v06/` package with type aliases, factory
    functions, and the `SpecVersion` constant for the new version.

1. **Add conformance and doc tests** for the new version under `pkg/api/v06/`
   (see existing versions for the pattern: `conformance_test.go`, `docs_test.go`,
   `factory_test.go`).

1. **Run tests**:

    ```shell
    make all
    ```

1. **Open a PR, get it reviewed and merged.**

1. **Tag the release** from `main`:

    ```shell
    git tag -a v0.6.0 -m "Release v0.6.0"
    git push upstream v0.6.0
    ```

1. **Create a GitHub Release** from the new tag with release notes covering the
   new spec version, new event types, and any breaking changes.
