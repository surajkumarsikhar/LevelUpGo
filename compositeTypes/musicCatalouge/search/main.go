package main

import "fmt"

// Write SearchByArtist here (return all albums by the given artist)
func (lib Library) SearchByArtist(artist string) []Album {
	result := []Album{}
	for _, a := range lib {
		if a.Artist == artist {
			result = append(result, a)
		}
	}
	return result
}

// Write GenreCount here (count songs per genre across all albums)
func (lib Library) GenreCount() map[string]int {
	result := make(map[string]int)
	for _, a := range lib {
		for _, s := range a.Songs {
			result[s.Genre]++
		}
	}
	return result
}

// Write TotalSongs here (total number of songs across all albums)
func (lib Library) TotalSongs() int {
	result := 0
	for _, a := range lib {
		result += len(a.Songs)
	}
	return result
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

type Library map[string]Album

func NewLibrary() Library {
	return make(Library)
}

func (lib Library) AddAlbum(a Album) {
	lib[a.Title] = a
}

func (lib Library) FindAlbum(title string) (Album, bool) {
	album, ok := lib[title]
	return album, ok
}

func (lib Library) RemoveAlbum(title string) bool {
	_, ok := lib[title]
	if ok {
		delete(lib, title)
	}
	return ok
}

func (lib Library) AllSongs() []Song {
	var songs []Song
	for _, album := range lib {
		for _, song := range album.Songs {
			songs = append(songs, song)
		}
	}
	return songs
}

func main() {
	lib := NewLibrary()

	abbey := NewAlbum("Abbey Road", "The Beatles", 1969)
	abbey.AddSong(NewSong("Come Together", "The Beatles", 259, "rock"))
	abbey.AddSong(NewSong("Something", "The Beatles", 182, "rock"))
	lib.AddAlbum(abbey)

	thriller := NewAlbum("Thriller", "Michael Jackson", 1982)
	thriller.AddSong(NewSong("Thriller", "Michael Jackson", 357, "pop"))
	thriller.AddSong(NewSong("Beat It", "Michael Jackson", 258, "rock"))
	lib.AddAlbum(thriller)

	fmt.Printf("Beatles albums: %d\n", len(lib.SearchByArtist("The Beatles")))
	fmt.Printf("Genre counts: %v\n", lib.GenreCount())
	fmt.Printf("Total songs: %d\n", lib.TotalSongs())
}
