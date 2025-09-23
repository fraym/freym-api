# vNext

# v0.49.0

- (feature) Add `types` option to deployment command to force type deployments
- (feature) Add `dry-run` option to deployment command
- (feature) Add `doNotUseSnapshots` option to streams client
- (feature) Add `TryLock` and `TryRLock` methods to sync client
- (feature) Add `returnEmptyDataIfNotFound` option to single view entry queries in crud and projections clients

# v0.48.3

- (bug) Fix `StreamHandlerError` constructor in js streams package

# v0.48.2

- (bug) Add retry number to subscription event in go streams package
- (improvement) Update all dependencies

# v0.48.1

- (bug) Fix `onOutdated` and `onStatusCode` options to graphql client

# v0.48.0

- (feature) Add `onOutdated` and `onStatusCode` options to graphql client

# v0.47.0

- (feature) Add `deployment-id` and `next-deployment-id` commands to deployment command

# v0.46.0

- (bc) Update all dependencies

# v0.45.0

- (feature) streams: list events that had errors
- (feature) streams: requeue event that had errors
- (feature) streams: Improve retry mechanisms by adding a retry number
- (feature) streams: Improve retry mechanisms by returning a `doRetry` flag on error
- (feature) streams: Add transactional consistency by correlation id
- (feature) crud + projections: Strong consistency by id

# v0.44.1

- (bug) Fix `wait` param in crud and projections to contain a list of conditions

# v0.44.0

- (feature) Add upsert rpc to crud
- (feature) Add `wait` param to list data rpc
- (feature) Add `wait` param to view data rpc
- (feature) Add `wait` param to view data list rpc

# v0.43.0

- (feature) Add `current` and `set-current` commands to deployment command
- (improvement) Update all dependencies

# v0.42.0

- (feature) Add `ci` arg to deployment command for better ci integration

# v0.41.1

- (bug) Read streams group id from env variables

# v0.41.0

- (feature) Add rpc to get the last handled event of a consumer group for a topic

# v0.40.0

- (feature) Add crud backchannel

# v0.39.0

- (feature) Add projections backchannel

# v0.38.4

- (bug) Remove mutex from sync client

# v0.38.3

- (bug) Fix env variable namings

# v0.38.2

- (bug) Fix `@baseView` support

# v0.38.1

- (bug) Add missing baseView to deployment

# v0.38.0

- (feature) Add `@view` and `@baseView` to crud
- (feature) Add `@baseView` to projections
- (bc) Cleanup client function naming

# v0.37.2

- (bug) Fix output dir of `@fraym/react` package

# v0.37.1

- (bug) Fix paths
- (improvement) Read `deploymentTarget` from `CRUD_CLIENT_DEPLOYMENT_TARGET` env var if not provided in config
- (improvement) Read `deploymentTarget` from `PROJECTIONS_CLIENT_DEPLOYMENT_TARGET` env var if not provided in config
- (improvement) Read `deploymentId` from `STREAMS_CLIENT_DEPLOYMENT_ID` env var if not provided in config

# v0.37.0

- (feature) Add `@fraym/graphql`

# v0.36.0

- (feature) Add `@fraym/react`

# v0.35.1

- (bug) Fix projections go client

# v0.35.0

- (feature) Add `RenameEventType` rpc to streams
- (feature) Allow multi field indices by allowing `@index` on object level
- (feature) Add `revalidate` param to `@view`
- (feature) Add `skip=##` flag to deployment command
- (bug) Fix `force` deployment command flag
- (feature) Add `@renamed` to crud
- (internal) Remove unused dangerously remove gdpr check flag
- (feature) Update crud data by filter
- (feature) Add `@projectedAt`

# v0.34.4

- (bug) Fix `jose` version

# v0.34.3

- (feature) Add `--force` flag to create deployment command

# v0.34.2

- (bug) Remove `user_id` from `EventMetadata` in favor of `AuthData`

# v0.34.1

- (improvement) Add user ID to auth data

# v0.34.0

- (feature) Add pagination to auth users

# v0.33.0

- (improvement) Allow all supported types in js auth packages `generateJwt` for `expirationTime`
- (improvement) Update all dependencies
- (feature) Add JSON functions to projections go client
- (improvement) Improve `RUnlock` performance in go sync client

# v0.32.2

- (bug) Fix exports in sync JS package

# v0.32.1

- (bug) Fix proto JS package

# v0.32.0

- (feature) Add read locks to sync
- (feature) JS sync client

# v0.31.9

- (bug) Add `userId` and `deploymentId` to subscription event in go streams package
- (internal) Switch go uuid package

# v0.31.8

- (bug) Add missing deployment target to namespace rollback apis

# v0.31.7

- (bug) Add missing parameters in migration apis

# v0.31.6

- (bug) Use target instead of deploymentId in projections and crud

# v0.31.5

- (bug) Add activation deployment step to all services

# v0.31.4

- (bug) Add activation deployment step

# v0.31.3

- (bug) Fix deployment apis (creator defines the target)

# v0.31.2

- (bug) Fix streams backchannel message data (only send tenant and topic again)

# v0.31.1

- (bug) Cleanup deployment renaming

# v0.31.0

- (bug) Fix order of arguments (go)
- (bc) Use variadic args in sync client for resource definitions (go)
- (bc) Rename migrations to deployments

# v0.30.0

- (bug) Fix stream pagination
- (improvement) Allow relative paths for sql file names for view projections
- (feature) Move `@freym/proto` and all api clients into this package
