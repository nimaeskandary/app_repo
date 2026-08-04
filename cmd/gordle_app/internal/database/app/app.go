package app_database

import db_core "github.com/nimaeskandary/app_repo/pkg/database/go/core"

// AppDBReader identifies the query-only App Database connection in the dependency graph.
type AppDBReader db_core.SQLDatabase

// AppDBWriter identifies the writable App Database connection in the dependency graph.
type AppDBWriter db_core.SQLDatabase

// AppDatabaseMigrator identifies the App Database migrator in the dependency graph.
type AppDatabaseMigrator db_core.Migrator

// AppDatabaseMigrateAllOnStart identifies the App Database startup migration work in the dependency graph.
type AppDatabaseMigrateAllOnStart db_core.MigrateAllOnStart
