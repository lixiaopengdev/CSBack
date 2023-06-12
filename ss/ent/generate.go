package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent  generate --feature sql/upsert,sql/execquery ./schema
// if want gen schema-viz.html , just go run -mod=mod github.com/hedwigz/entviz/cmd/entviz ./schema/
