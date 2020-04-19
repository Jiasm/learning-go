package main

import (
	"os"
	"fmt"
	"sort"
	"time"
	"text/tabwriter"
)

type Track struct {
		Title		string
		Artist	string
		Album		string
		Year		int
		Length	time.Duration
}

var tracks = []*Track{
		{ "Go", "Delilah", "From the Roots Up", 2012, length("3m38s") },
		{ "Go", "Moby", "Moby", 1922, length("3m37s") },
		{ "Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s") },
		{ "Ready 2 Go", "Martin Solveing", "Smash", 2011, length("4m24s") },
}

type byArtist []*Track

type customSort struct {
	t []*Track
	less func(x, y * Track) bool
}

func (x customSort) Len() int { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func (x byArtist) Len() int						{ return len(x) }
func (x byArtist) Less(i, j int) bool	{ return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)			{ x[i], x[j] = x[j], x[i] }

func length(s string) time.Duration {
		d, err := time.ParseDuration(s)

		if err != nil {
				panic(s)
		}

		return d
}

func printTracks(tracks []*Track) {
		const format = "%v\t%v\t%v\t%v\t%v\t\n"

		// sorting and reverse sorting
		sort.Sort(byArtist(tracks))

		tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)

		fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
		fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")

		for _, t := range tracks {
				fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
		}

		tw.Flush()
}

func printTracksReverse(tracks []*Track) {
		const format = "%v\t%v\t%v\t%v\t%v\t\n"

		// sorting and reverse sorting
		sort.Sort(sort.Reverse(byArtist(tracks)))

		tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)

		fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
		fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")

		for _, t := range tracks {
				fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
		}

		tw.Flush()
}

func printTracksMultiple(tracks []*Track) {
		const format = "%v\t%v\t%v\t%v\t%v\t\n"

		// sorting and reverse sorting
		sort.Sort(customSort{ tracks, func(x, y *Track) bool {
			if x.Title != y.Title {
				return x.Title < y.Title
			}

			if x.Year != y.Year {
				return x.Year < y.Year
			}

			if x.Length != y.Length {
				return x.Length < y.Length
			}

			return false
		}})

		tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)

		fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
		fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")

		for _, t := range tracks {
				fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
		}

		tw.Flush()
}

func printTracksMultipleReverse(tracks []*Track) {
		const format = "%v\t%v\t%v\t%v\t%v\t\n"

		// sorting and reverse sorting
		sort.Sort(sort.Reverse(customSort{ tracks, func(x, y *Track) bool {
			if x.Title != y.Title {
				return x.Title < y.Title
			}

			if x.Year != y.Year {
				return x.Year < y.Year
			}

			if x.Length != y.Length {
				return x.Length < y.Length
			}

			return false
		}}))

		tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)

		fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
		fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")

		for _, t := range tracks {
				fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
		}

		tw.Flush()
}

func main() {
	printTracks(tracks)
	printTracksReverse(tracks)
	printTracksMultiple(tracks)
	printTracksMultipleReverse(tracks)
}