# CSBackendTmp
Temp repo for come social backend

Use Iris for web server, Facebook ent for orm, Redis for cache and Mysql for database.

Examples:

Init a model(table in database): go run -mod=mod entgo.io/ent/cmd/ent init User

Gen code: go generate

Gen openapi.json: go run -mod=mod entc.go

Check database tables: go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema

Also can gen graph for relationship between models: go run -mod=mod github.com/hedwigz/entviz/cmd/entviz ./schema/


