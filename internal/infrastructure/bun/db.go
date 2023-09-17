package bun

import (
	"database/sql"

	"github.com/caarlos0/env"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	"github.com/go-sql-driver/mysql"
)

type DBEnv struct {
	Name      string `env:"MYSQL_DATABASE"`
	User      string `env:"MYSQL_USER"`
	Password  string `env:"MYSQL_PASSWORD"`
	Address   string `env:"MYSQL_HOST"`
	Collation string `env:"MYSQL_COLLATION"`
}

func NewDB() *bun.DB {
	cfg := DBEnv{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	c := mysql.Config{
		User:                 cfg.User,
		Passwd:               cfg.Password,
		Net:                  "tcp",
		Addr:                 cfg.Address,
		DBName:               cfg.Name,
		Collation:            cfg.Collation,
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	_db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		panic(err)
	}
	if err = _db.Ping(); err != nil {
		panic(err)
	}

	db := bun.NewDB(_db, mysqldialect.New())
	return db
}
