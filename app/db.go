package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/revel/revel"
)

var (
	DB *gorm.DB
)

func InitDB() *gorm.DB {
	var (
		driver string
		spec   string
		found  bool
	)

	// Read configuration
	if driver, found = revel.Config.String("gorm.driver"); !found {
		revel.ERROR.Fatal("No gorm.driver found.")
	}
	if spec, found = revel.Config.String("gorm.spec"); !found {
		revel.ERROR.Fatal("No gorm.spec found.")
	}

	maxIdleConns := revel.Config.IntDefault("gorm.max_idle_conns", 10)
	maxOpenConns := revel.Config.IntDefault("gorm.max_open_conns", 100)
	singularTable := revel.Config.BoolDefault("gorm.singular_table", false)
	logMode := revel.Config.BoolDefault("gorm.log_mode", false)

	// Initialize `gorm`
	dbm, err := gorm.Open(driver, spec)
	if err != nil {
		revel.ERROR.Fatal(err)
	}

	DB = &dbm

	dbm.DB().Ping()
	dbm.DB().SetMaxIdleConns(maxIdleConns)
	dbm.DB().SetMaxOpenConns(maxOpenConns)
	dbm.SingularTable(singularTable)
	dbm.LogMode(logMode)

	return &dbm
}
