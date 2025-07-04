# vNext

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
