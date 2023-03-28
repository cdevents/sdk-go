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
    > **Note** Golang [version v1.18](https://golang.org/dl/) or higher is required.

1. [`make`](https://www.gnu.org/software/make/): not stricly required but handy to run
   tests with a single command.

### Setup a fork

The sdk-go project requires that you develop (commit) code changes to branches that belong to a fork of the `cdevents/sdk-go` repository in your GitHub account before submitting them as Pull Requests (PRs) to the actual project repository.

1. [Create a fork](https://help.github.com/articles/fork-a-repo/) of the `cdevents/sdk-go` repository in your GitHub account.

1. Create a clone of your fork on your local machine:

    ```shell
    git clone git@github.com:${YOUR_GITHUB_USERNAME}/sdk-go.git
    ```

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
$ make fmt
```

To run the go linter:

```shell
$ make lint
```

To run unit tests:
```shell
$ make test
```

To refresh generated code:
```shell
$ make generate
```

To run all targets, before creating a commit:

```shell
make all
```

### Updating the SDK with a new version of CDEvents

When a new version of the CDEvents spec is released, the
SDK needs to be updated and released accordingly. To do so,
perform the following steps.

Update the spec submodule:
```shell
CDEVENTS_TAG=<git tag of the new version>
pushd pgk/api/spec
git fetch --all
git checkout ${CDEVENTS_TAG}
popd
```

Update the spec version in the [code](https://github.com/cdevents/sdk-go/blob/5fdad604dc7c15f82e285f9ec6307a81f42e9612/pkg/api/types.go#L33)
```golang
	CDEventsSpecVersion       = "<CDEVENTS_TAG>"
```

Regenerate the code and tests from the schemas:
```shell
make generate
```

The SDK includes one test [`examples_test.go`](https://github.com/cdevents/sdk-go/blob/main/pkg/api/examples_test.go),
that verifies it's possible, through the SDK, to create the example
events defined in the spec.
New events, changes to the schemas of existing events and changes to the
examples may cause this test to fail. After updating the code,
run all tests and fix any failure as required.

```shell
make test
```

Once all tests are passing, create a new commit and verify that
`generate` does not trigger any change to the code. 

```shell
make generate  # or make all, to also check formatting at tests
```