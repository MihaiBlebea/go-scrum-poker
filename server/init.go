package server

import "embed"

//go:embed webapp/*
var static embed.FS
