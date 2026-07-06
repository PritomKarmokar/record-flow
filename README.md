# record-flow
- just want to build something don't know what to build?
- Let's start with the simple crud first then;(
- Want to break free? God knows I want to brek free?

## `Migration` Helpers
- create new sql migration
```
goose -dir migrations -s create record_table sql
```
- `up` - Apply all available migrations
```
goose up
```
- `down` - Roll back a single migration from the current version
```
goose down
```