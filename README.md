# Freym API

[![JS Checks](https://github.com/fraym/freym-api/actions/workflows/js-checks.yaml/badge.svg)](https://github.com/fraym/freym-api/actions/workflows/js-checks.yaml)
[![Go Checks](https://github.com/fraym/freym-api/actions/workflows/go-checks.yaml/badge.svg)](https://github.com/fraym/freym-api/actions/workflows/go-checks.yaml)

## Overview

The Freym API provides a set of clients to interact with the Freym microservices.

## Clients

| Service                                 | Language   |
| --------------------------------------- | ---------- |
| [Sync](go/sync/README.md)               | go         |
| [Sync](js/sync/README.md)               | JavaScript |
| [Streams](go/streams/README.md)         | go         |
| [Streams](js/streams/README.md)         | JavaScript |
| [Projections](go/projections/README.md) | go         |
| [Projections](js/projections/README.md) | JavaScript |
| [CRUD](go/crud/README.md)               | go         |
| [CRUD](js/crud/README.md)               | JavaScript |
| [Auth](go/auth/README.md)               | go         |
| [Auth](js/auth/README.md)               | JavaScript |
| [Deployments](go/deployments/README.md) | go         |
| [Deployments](js/deployments/README.md) | JavaScript |

## Docs

Please have a look at our [documentation](https://docs.freym.becklyn.app/docs).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Releasing new versions

- update versions in all `package.json` files (using the `make version-js` command)
  - package versions
  - dependency version in case a package references another package in this monorepo
- add a git tag in the form: `v#.#.#` or `v#.#.#-alpha.#`
- Update `RELEASE.md` to contain the new release
- GitHub Actions will automatically publish all packages once everything is pushed
