package database

import (
	"blog-api/routes"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pgPass = os.Getenv("PG_PASS")

var pgConnString = "postgres://postgres.xlvhpbcxavkjwhvvumjr:" + pgPass + "@aws-0-eu-west-2.pooler.supabase.com:5432/postgres"

func Database() {
	db, err := gorm.Open(postgres.Open(pgConnString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	db.AutoMigrate(&routes.Blog{})

}
