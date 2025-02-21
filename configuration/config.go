package configuration

import (
	"database/sql"

	"github.com/Bayan2019/go-ozinshe/repositories"
)

// var MapCfg *MapConfiguration
var ApiCfg *ApiConfiguration

// type MapConfiguration struct{}

type ApiConfiguration struct {
	Conn      *sql.DB
	DB        *repositories.Queries
	DirImages string
	DirVideos string
}
