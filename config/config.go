package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type (
	DBConfig struct {
		HOST          string `envconfig:"DB_HOST"`
		PORT          string `envconfig:"DB_PORT"`
		USER          string `envconfig:"DB_USER"`
		NAME          string `envconfig:"DB_NAME"`
		PASS          string `envconfig:"DB_PASSWORD"`
		MAX_OPEN_CONN int    `envconfig:"DB_MAX_OPEN_CONN"`
		MAX_IDLE_CONN int    `envconfig:"DB_MAX_IDDLE_CONN"`
	}

	application struct {
		DBConfig DBConfig
		DB       *sqlx.DB
		HOST     string `envconfig:"HOST"`
		PORT     string `envconfig:"PORT"`
		NAME     string `envconfig:"NAME"`
		VERSION  string `envconfig:"VERSION"`
	}
)

var Application *application

func (a *application) InitConfig() error {
	Application = &application{}

	var err error
	Application.initENV()
	if err = Application.initDB(); err != nil {
		log.Println("Database connection Error", err)
		return err
	}

	return nil
}

func (a *application) initDB() error {
	log.Println("connecting to database", a.DBConfig.HOST)
	var connString string = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		a.DBConfig.HOST, a.DBConfig.PORT, a.DBConfig.USER, a.DBConfig.PASS, a.DBConfig.NAME,
	)
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		log.Println("Error connecting to database host:", a.DBConfig.HOST, a.DBConfig.PORT, err)
		return err
	}
	if err = db.Ping(); err != nil {
		log.Println("Error connecting to database host:", a.DBConfig.HOST, a.DBConfig.PORT, err)
		return err
	}
	db.SetMaxOpenConns(a.DBConfig.MAX_OPEN_CONN)
	db.SetMaxIdleConns(a.DBConfig.MAX_IDLE_CONN)
	a.DB = db
	log.Println("Database connection success")
	return nil
}

func (a *application) initENV() error {
	var err error
	if err = godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env, file not found")
	}

	if err = envconfig.Process("", a); err != nil {
		return err
	}
	return err
}
