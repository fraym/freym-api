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

Start dev container (from the directory root of the client you want to develop):

```shell
make dev
```

Connect your IDE to that dev container.
