# Readme

At times while doing CI/CD with Github Actions, you need to read environment variables stored in `Github Secrets`. This sample program demonstrates how you can access `Github Secrets` stored in Github while doing CI/CD.

## Steps

1. Go to **Github Repository > Settings**. From the left menu, selection **Secrets > Actions**

![Actions link](https://i.imgur.com/sFnQ2Zk.png)

2. Select `New repository secret` and add a key-value pair. For our example, we have used key name as `T1` and Value as `Hello world`. Then save it

![Save actions secret](https://i.imgur.com/MCPcOOU.png)

3. Once you save it, it should be visible on the screen

![Secrets list](https://i.imgur.com/kcz1K0h.png)

4. Now you need to add the secret in your Github Actions Workflow file. You need to call the secret using

```
env:
    T1: ${{secrets.T1}}
```
Value from `secrets.T1` is being assigned to Runner environment variable `T1`

Here is my [Github Actions workflow](https://github.com/chilarai/go-env-github-actions/blob/master/.github/workflows/go.yml)

```
name: Go

on:
    push:
        branches: ["master"]
    pull_request:
        branches: ["master"]

env:
    T1: ${{secrets.T1}}

jobs:
    build:
        runs-on: ubuntu-latest
        steps:
            - uses: actions/checkout@v3

            - name: Set up Go
              uses: actions/setup-go@v3
              with:
                  go-version: 1.18

            - name: Build
              run: go build -v ./...

            - name: Test
              run: go test -v ./...

            - name: Environment list
              run: env

```

Once you do that, you should access this in your unit tests and code as

```
env1, err1 := os.LookupEnv("T1")
log.Println(env1, err1)
```
