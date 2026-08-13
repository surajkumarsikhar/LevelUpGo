package main

import "fmt"

// Define Library as a named map type (map albums by title)
type Library map[string]Album

// Write NewLibrary here
func NewLibrary() Library {
	return make(Library)
}

// Write AddAlbum method here
func (lib Library) AddAlbum(a Album) {
	lib[a.Title] = a
}

// Write FindAlbum method here (return Album and bool)
func (lib Library) FindAlbum(title string) (a Album, ok bool) {
	a, ok = lib[title]
	return
}

// Write RemoveAlbum method here (return bool)
func (lib Library) RemoveAlbum(title string) (existed bool) {
	_, existed = lib[title]
	delete(lib, title)
	return
}

// Write AllSongs method here (return every song from every album)
func (lib Library) AllSongs() []Song {
	songs := []Song{}
	for t, _ := range lib {
		songs = append(songs, lib[t].Songs...)
	}
	return songs
}

type Song struct {
	Title    string
	Artist   string
	Duration int
	Genre    string
}

type Album struct {
	Title  string
	Artist string
	Year   int
	Songs  []Song
}

func NewSong(title, artist string, duration int, genre string) Song {
	return Song{Title: title, Artist: artist, Duration: duration, Genre: genre}
}

func NewAlbum(title, artist string, year int) Album {
	return Album{Title: title, Artist: artist, Year: year, Songs: []Song{}}
}

func (a *Album) AddSong(s Song) {
	a.Songs = append(a.Songs, s)
}

func (a Album) TotalDuration() int {
	total := 0
	for _, s := range a.Songs {
		total += s.Duration
	}
	return total
}

func main() {
	lib := NewLibrary()

	abbey := NewAlbum("Abbey Road", "The Beatles", 1969)
	abbey.AddSong(NewSong("Come Together", "The Beatles", 259, "rock"))
	abbey.AddSong(NewSong("Something", "The Beatles", 182, "rock"))
	lib.AddAlbum(abbey)

	album, found := lib.FindAlbum("Abbey Road")
	fmt.Printf("Found: %v, %s by %s\n", found, album.Title, album.Artist)
	fmt.Printf("All songs: %d\n", len(lib.AllSongs()))
}
