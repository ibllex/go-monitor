package main

type Config struct {
	Slug    string `ini:"slug"`
	Port    int    `ini:"port"`
	History int    `ini:"history"`
}
