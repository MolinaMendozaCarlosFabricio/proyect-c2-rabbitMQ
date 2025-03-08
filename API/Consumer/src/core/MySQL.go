package core

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Database DatabaseConfig `toml:"database"`
	Server   ServerConfig   `toml:"server"`
}

type DatabaseConfig struct {
	DNS string `toml:"dns"`
}

type ServerConfig struct {
	Port  int  `toml:"port"`
	Debug bool `toml:"debug"`
}

type ConectionMySQL struct {
	DB  *sql.DB
	Err string
}

func MySQLConection() *ConectionMySQL {
	var config Config
	errorMsg := ""

	if _, err := toml.DecodeFile("config_db.toml", &config); err != nil {
		errorMsg = fmt.Sprintf("Error al leer el archivo toml: %v", err)
		return &ConectionMySQL{DB: nil, Err: errorMsg}
	}

	if config.Database.DNS == "" {
		errorMsg = "Error: DNS de la base de datos vacío"
		return &ConectionMySQL{DB: nil, Err: errorMsg}
	}

	db, err := sql.Open("mysql", config.Database.DNS)
	if err != nil {
		errorMsg = fmt.Sprintf("Error al establecer la conexión con la BD: %v", err)
		return &ConectionMySQL{DB: nil, Err: errorMsg}
	}

	if err := db.Ping(); err != nil {
		db.Close()
		errorMsg = fmt.Sprintf("Error al hacer ping en la BD: %v", err)
		return &ConectionMySQL{DB: nil, Err: errorMsg}
	}

	db.SetMaxOpenConns(10)

	fmt.Println("Conexión exitosa a la base de datos")
	return &ConectionMySQL{DB: db, Err: ""}
}

func (conection *ConectionMySQL) ExecPreparedQuerys(query string, values ...interface{}) (sql.Result, error) {
	stmt, err := conection.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("Error al preparar la consulta", err)
	}
	defer stmt.Close()

	results, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("Error al realizar la consulta", err)
	}
	return results, nil
}

func (conection ConectionMySQL) FetchRows(query string, values ...interface{}) (*sql.Rows, error) {
	rows, err := conection.DB.Query(query, values...)
	if err != nil {
		return nil, fmt.Errorf("Erro al conseguir las filas afectadas")
	}
	return rows, nil
}