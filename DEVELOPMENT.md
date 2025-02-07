# Development

You'll need the following apps for a smooth development experience:

- minikube
- lens
- okteto

## Running the dev environment

Start minikube if not already done:

```shell
minikube start
```

Install dev image if not already done (https://github.com/fraym/go-dev-container). See [README.md](https://github.com/fraym/go-dev-container) for documentation.

Start dev container (go):

```shell
make dev-go
```

Start dev container (js):

```shell
make dev-js
```

Connect your IDE to that dev container.

Stop the dev container after development (go):

```shell
make dev-go-stop
```

Stop the dev container after development (js):

```shell
make dev-js-stop
```
