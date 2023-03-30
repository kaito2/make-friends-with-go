# DB Transaction

## Migration

Check current tables.

```
$ go run github.com/k0kubun/sqldef/cmd/mysqldef -uroot -proot -P33306 db --export
-- No table exists --
```

Save it to edit.

```
$ go run github.com/k0kubun/sqldef/cmd/mysqldef -uroot -proot -P33306 db --export > schema.sql
```

Dry run.

```
$ go run github.com/k0kubun/sqldef/cmd/mysqldef -uroot -proot -P33306 db --dry-run < schema.sql
-- dry run --
CREATE TABLE user (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(128) DEFAULT 'kaito2'
) Engine = InnoDB DEFAULT CHARSET = utf8mb4;
```

Apply it.

```
$ go run github.com/k0kubun/sqldef/cmd/mysqldef -uroot -proot -P33306 db < schema.sql
-- Apply --
CREATE TABLE user (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(128) DEFAULT 'kaito2'
) Engine = InnoDB DEFAULT CHARSET = utf8mb4;
```
