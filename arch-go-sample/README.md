# Arch-Go Sample

これは [fdaines/arch-go](https://github.com/fdaines/arch-go) を使って Web API の設計をバリデーションするサンプルです。

## Run on Local

```
$ go run cmd/main.go
```

### Sample access

Create user.

```
$ curl --location 'http://localhost:8080/api/v1/users/' \
       --form 'name="yo"'
```

Get user.

```
$ curl --location 'http://localhost:8080/api/v1/users/cgocu9smpa35dov1fn60'
```


## Check Rules

```
$ go run github.com/fdaines/arch-go
Looking for packages.
8 packages found...
[PASS] - Packages matching pattern 'github.com/kaito2/make-friends-with-go/arch-go-sample/internal/adapter/handler' should ['only depend on internal packages that matches [[github.com/kaito2/make-friends-with-go/arch-go-sample/internal/usecase github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/repository]]']
	Package 'github.com/kaito2/make-friends-with-go/arch-go-sample/internal/adapter/handler' passes
[PASS] - Packages matching pattern 'github.com/kaito2/make-friends-with-go/arch-go-sample/internal/adapter/implement/factory' should ['only depend on internal packages that matches [[github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/factory github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/model]]']
	Package 'github.com/kaito2/make-friends-with-go/arch-go-sample/internal/adapter/implement/factory' passes
[PASS] - Packages matching pattern 'github.com/kaito2/make-friends-with-go/arch-go-sample/internal/adapter/implement/repository' should ['only depend on internal packages that matches [[github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/repository github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/model]]']
	Package 'github.com/kaito2/make-friends-with-go/arch-go-sample/internal/adapter/implement/repository' passes
[PASS] - Packages matching pattern 'github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain' should ['only depend on internal packages that matches [[github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain]]']
	Package 'github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/factory' passes
	Package 'github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/model' passes
	Package 'github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain/repository' passes
[PASS] - Packages matching pattern 'github.com/kaito2/make-friends-with-go/arch-go-sample/internal/usecase' should ['only depend on internal packages that matches [[github.com/kaito2/make-friends-with-go/arch-go-sample/internal/domain]]']
	Package 'github.com/kaito2/make-friends-with-go/arch-go-sample/internal/usecase' passes
+---+-----------------+-------------+-------------+-------------+
| # | RULE TYPE       |       TOTAL |   SUCCEEDED |      FAILED |
+---+-----------------+-------------+-------------+-------------+
| 1 | Content Rule    |           0 |           0 |           0 |
| 2 | Dependency Rule |           5 |           5 |           0 |
| 3 | Function Rule   |           0 |           0 |           0 |
| 4 | Naming Rule     |           0 |           0 |           0 |
+---+-----------------+-------------+-------------+-------------+
|   | COMPLIANCE RATE | 100% [PASS]                             |
|   | COVERAGE RATE   |  87% [PASS]                             |
+---+-----------------+-----------------------------------------+
--------------------------------------
	Execution Summary
--------------------------------------
Total Rules: 	5
Succeeded: 	5
Failed: 	0
--------------------------------------
Compliance:      100% (Pass)
Coverage:         87% (Pass)
Time: 0.163 seconds
```
