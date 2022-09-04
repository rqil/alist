package sftp

import (
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/op"
)

type Addition struct {
	Address    string `json:"address" required:"true"`
	Username   string `json:"username" required:"true"`
	PrivateKey string `json:"private_key" type:"text"`
	Password   string `json:"password"`
	driver.RootFolderPath
}

var config = driver.Config{
	Name:        "SFTP",
	LocalSort:   true,
	OnlyLocal:   true,
	DefaultRoot: "/",
	CheckStatus: true,
}

func New() driver.Driver {
	return &SFTP{}
}

func init() {
	op.RegisterDriver(config, New)
}
