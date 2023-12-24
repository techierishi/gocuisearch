package main

import (
	"log/slog"
	"strconv"
)

type SearchApp struct {
	IsRaw bool

	Content string

	Time float64
	Idx  int
}

func NewSearchAppFromCmdLine(cmdLine string) SearchApp {
	return SearchApp{
		IsRaw:   true,
		Content: cmdLine,
	}
}

func NewSearchApp(r *RowItem) SearchApp {
	time, err := strconv.ParseFloat(r.Time, 64)
	if err != nil {
		slog.Error("Error while parsing time as float %v", err)
	}
	return SearchApp{
		IsRaw:   false,
		Idx:     r.Idx,
		Content: r.Content,
		Time:    time,
	}
}
