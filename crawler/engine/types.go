package engine

import "go_project/crawler/model"

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []model.Profile
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}