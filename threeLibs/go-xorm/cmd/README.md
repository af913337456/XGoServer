
[中文](https://github.com/go-xorm/cmd/blob/master/README_CN.md)

# xorm tools

xorm tools is a set of  tools for database operation. 

## Source Install

`go get github.com/go-xorm/cmd/xorm`

and you should install the depends below:

* github.com/go-xorm/xorm

* Mysql: [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

* MyMysql: [github.com/ziutek/mymysql/godrv](https://github.com/ziutek/mymysql/godrv)

* Postgres: [github.com/lib/pq](https://github.com/lib/pq)

* SQLite: [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) 

* MSSQL: [github.com/denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb)

** For sqlite support, you need build via `go build -tags sqlite3` because of this driver ask cgo.

## Commands

All the commands below.

* **reverse**     reverse a db to codes
* **shell**       a general shell to operate all kinds of database
* **dump**        dump database all table struct's and data to standard output
* **source**      execute a sql from std in
* **driver**      list all supported drivers

## Reverse

Reverse command is a tool to convert your database struct to all kinds languages of structs or classes. After you installed the tool, you can type 

`xorm help reverse`

to get help

example:

`cd $GOPATH/github.com/go-xorm/cmd/xorm`

sqlite:
`xorm reverse sqite3 test.db templates/goxorm`

mysql:
`xorm reverse mysql root:@/xorm_test?charset=utf8 templates/goxorm`

mymysql:
`xorm reverse mymysql xorm_test2/root/ templates/goxorm`

postgres:
`xorm reverse postgres "dbname=xorm_test sslmode=disable" templates/goxorm`

will generated go files in `./model` directory

### Template and Config

Now, xorm tool supports go and c++ two languages and have go, goxorm, c++ three of default templates. In template directory, we can put a config file to control how to generating.

```
lang=go
genJson=1
```

lang must be go or c++ now.
genJson can be 1 or 0, if 1 then the struct will have json tag.

## Shell

Shell command provides a tool to operate database. For example, you can create table, alter table, insert data, delete data and etc.

`xorm shell sqlite3 test.db` will connect to the sqlite3 database and you can type `help` to list all the shell commands.

## Dump

Dump command provides a tool to dump all database structs and data as SQL to your standard output.

`xorm dump sqlite3 test.db` could dump sqlite3 database test.db to standard output. If you want to save to file, just
type `xorm dump sqlite3 test.db > test.sql`.

## Source

`xorm source sqlite3 test.db < test.sql` will execute sql file on the test.db.

## Driver

List all supported drivers since default build will not include sqlite3.

## LICENSE

 BSD License
 [http://creativecommons.org/licenses/BSD/](http://creativecommons.org/licenses/BSD/)
