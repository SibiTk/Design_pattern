package main

import "fmt"


type Song struct {
	Title string
}

type Playlist struct {
	songs []Song
}

func NewPlaylist() *Playlist {
	return &Playlist{}
}

func (p *Playlist) AddSong(title string) {
	p.songs = append(p.songs, Song{Title: title})
}

func (p *Playlist) CreateIterator() *PlaylistIterator {
	return &PlaylistIterator{
		songs: p.songs,
		index: 0,
	}
}


type PlaylistIterator struct {
	songs []Song
	index int
}

func (it *PlaylistIterator) HasNext() bool {
	return it.index < len(it.songs)
}

func (it *PlaylistIterator) Next() Song {
	song := it.songs[it.index]
	it.index++
	return song
}


func main() {
	playlist := NewPlaylist()
	playlist.AddSong("hcsbvjsnnv")
	playlist.AddSong("snjnv,n ,sn ")
	playlist.AddSong(" nmcs svmnm,nv")

	iterator := playlist.CreateIterator()

	fmt.Println(" Playing playlist:")
	for iterator.HasNext() {
		song := iterator.Next()
		fmt.Println("Now Playing:", song.Title)
	}
}
