package main

import "fmt"

// Define your Song struct here
type Song struct {
	Title    string
	Artist   string
	Duration int
	Genre    string
}

// Define your Album struct here
type Album struct {
	Title  string
	Artist string
	Year   int
	Songs  []Song
}

// Write NewSong here
func NewSong(title, artist string, duration int, genre string) Song {
	return Song{
		Title:    title,
		Artist:   artist,
		Duration: duration,
		Genre:    genre,
	}
}

// Write NewAlbum here (start with empty Songs)
func NewAlbum(title, artist string, year int) Album {
	return Album{
		Title:  title,
		Artist: artist,
		Year:   year,
	}
}

// Write AddSong method here (pointer receiver)
func (a *Album) AddSong(s Song) {
	a.Songs = append(a.Songs, s)
}

// Write TotalDuration method here (value receiver)
func (a Album) TotalDuration() int {
	total := 0
	for _, song := range a.Songs {
		total += song.Duration
	}
	return total
}

func main() {
	a := NewAlbum("Abbey Road", "The Beatles", 1969)
	a.AddSong(NewSong("Come Together", "The Beatles", 259, "rock"))
	a.AddSong(NewSong("Something", "The Beatles", 182, "rock"))
	fmt.Printf("%s (%d) - %d songs, %ds total\n", a.Title, a.Year, len(a.Songs), a.TotalDuration())
}
