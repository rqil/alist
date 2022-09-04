package ftp

import (
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/op"
)

type Addition struct {
	Address  string `json:"address" required:"true"`
	Username string `json:"username" required:"true"`
	Password string `json:"password" required:"true"`
	driver.RootFolderPath
}

var config = driver.Config{
	Name:        "FTP",
	LocalSort:   true,
	OnlyLocal:   true,
	DefaultRoot: "/",
}

func New() driver.Driver {
	return &FTP{}
}

func init() {
	op.RegisterDriver(config, New)
}
