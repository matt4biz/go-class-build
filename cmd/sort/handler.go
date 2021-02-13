package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"sortbuild"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func getLoop(r *http.Request) int {
	loop := 1

	if i, err := strconv.Atoi(r.FormValue("loop")); err == nil {
		loop = i - 1
	}

	return loop
}

func getDelay(r *http.Request) int {
	delay := 8

	if i, err := strconv.Atoi(r.FormValue("delay")); err == nil {
		delay = i
	}

	return delay
}

func qsortHigh(w http.ResponseWriter, r *http.Request) {
	q := sortbuild.QSort{Part: sortbuild.PartHigh}
	sortbuild.Animate(w, getLoop(r), getDelay(r), q.QStep)
}

func qsortMiddle(w http.ResponseWriter, r *http.Request) {
	q := sortbuild.QSort{Part: sortbuild.PartMiddle}
	sortbuild.Animate(w, getLoop(r), getDelay(r), q.QStep)
}

func qsortMedian(w http.ResponseWriter, r *http.Request) {
	q := sortbuild.QSort{Part: sortbuild.PartMedian}
	sortbuild.Animate(w, getLoop(r), getDelay(r), q.QStep)
}

func qsortInsert(w http.ResponseWriter, r *http.Request) {
	q := sortbuild.QSort{Part: sortbuild.PartInsert}
	sortbuild.Animate(w, getLoop(r), getDelay(r), q.QStep)
}

func qsortFlag(w http.ResponseWriter, r *http.Request) {
	q := sortbuild.QSort{Part: sortbuild.PartFlag}
	sortbuild.Animate(w, getLoop(r), getDelay(r), q.QStepFlag)
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	sortbuild.Animate(w, getLoop(r), getDelay(r), sortbuild.InsertionStep)
}

func showVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, version)
}
