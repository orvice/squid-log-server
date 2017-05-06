package main

import (
	"github.com/catpie/musdk-go"
)

var (
	client *musdk.Client
)

func InitMu() {
	cfg := Conf.Mu
	client = musdk.NewClient(cfg.ApiUri, cfg.Token, cfg.NodeId, cfg.SType)
}
