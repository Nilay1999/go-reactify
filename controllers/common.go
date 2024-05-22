package controllers

import (
	"github.com/Nilay1999/gin-gonic-server/types"
	"github.com/golodash/galidator"
)

var (
	g          = galidator.New()
	customizer = g.Validator(types.UserType{})
)
