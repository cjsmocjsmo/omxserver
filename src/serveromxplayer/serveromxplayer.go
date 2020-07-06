///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// LICENSE: GNU General Public License, version 2 (GPLv2)
// Copyright 2016, Charlie J. Smotherman
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License v2
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
package main

import (
	"log"
	"html/template"
	"net/http"
	"encoding/json"
	"net/url"
	omxplayer "serveromxplayer/lib"
	"github.com/gorilla/mux"
)

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func showOmx(w http.ResponseWriter, r *http.Request) {
	tmppath := "./static/omx.template"
	tmpl := template.Must(template.ParseFiles(tmppath))
	tmpl.Execute(w, tmpl)
}

func omxplayerPlayMediaHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(err)
	}
	q := u.Query()
	movie := q.Get("movie")
	log.Println("this is movie")
	log.Println(movie)
	omxplayer.Play(movie)

	omxplayer.ParseOutput()
	log.Printf("Omxlayer should be playing %s", movie)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Should be playing")
}

func omxplayerPlayMediaReactHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("this is url string %s", r.URL.String())
	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Println(err)
	}
	log.Printf("this is u %s", u)
	q := u.Query()
	log.Printf("this is q %s", q)
	log.Printf("this is q00 %v", q["medPath"])
	movie2 := q.Get("medPath")
	log.Printf("this is movie2 %s", movie2)

	omxplayer.Play(movie2)
	omxplayer.ParseOutput()
	log.Printf("Omxlayer should be playing %s", movie2)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Should be playing")
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	omxplayer.Resume()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Should be playing")
}

func pauseHandler(w http.ResponseWriter, r *http.Request) {
	omxplayer.Pause()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Should be playing")
}

func stopHandler(w http.ResponseWriter, r *http.Request) {
	omxplayer.Stop()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Should be playing")
}

//Seek forward
func nextHandler(w http.ResponseWriter, r *http.Request) {
	omxplayer.Fwd()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Should be playing")
}

//Seek backward
func previousHandler(w http.ResponseWriter, r *http.Request) {
	omxplayer.Bwd()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Should be playing")
}

//Next Chapter
func nextChapterHandler(w http.ResponseWriter, r *http.Request) {
	omxplayer.Next()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Should be playing")
}

//Previous Chapter
func previousChapterHandler(w http.ResponseWriter, r *http.Request) {
	omxplayer.Prev()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Should be playing")
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/static").Subrouter()
	r.HandleFunc("/Test", showOmx)
	r.HandleFunc("/OmxplayerPlayMedia", omxplayerPlayMediaHandler)
	r.HandleFunc("/OmxplayerPlayMediaReact", omxplayerPlayMediaReactHandler)
	r.HandleFunc("/Play", playHandler)
	r.HandleFunc("/Pause", pauseHandler)
	r.HandleFunc("/Stop", stopHandler)
	r.HandleFunc("/Next", nextHandler)
	r.HandleFunc("/Previous", previousHandler)
	r.HandleFunc("/NextChapter", nextChapterHandler)
	r.HandleFunc("/PreviousChapter", previousChapterHandler)
	s.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/media/"))))
	http.ListenAndServe(":8181", (r))
}
