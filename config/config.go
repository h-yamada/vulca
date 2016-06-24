package config

var Conf Config

type Config struct {
	VulsDBPath string
	CveDBPath  string
}
