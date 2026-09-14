# migratex

Utility wrapper over `golang-migrate` to fit my workflow needs.

## Install

```sh
go install github.com/joaomagfreitas/migratex/cmd/migratex@latest
```

After install `migratex` should be available as a command, if Go bin userspace is known in PATH.

## Usage

Upgrade database by running next migration.

```sh
migratex --conn=...
```

Downgrade database by running current migration cleanup script.

```sh
migratex --down --conn=...
```

If in need to run all migrations (up or down).

```sh
migratex --all --conn=...
```

---

If something went wrong when running the migration script, you will need to force the older version-

```sh
migratex --force=... --conn=...
```

### Connection string and config files

You can either use a connection string or a yaml config file to specify the database connection configuration. Config files are a better alternative to isolate connection details from the workflow scripts:

```sh
migratex --cfg=...
```
