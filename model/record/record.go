package model

type Record struct {
	Date        string   `json:"date"`
	Artist      []string `json:"artist"`
	Album       string   `json:"album"`
	Rating      string   `json:"rating"`
	SavedTracks string   `json:"saved_tracks"`
	BestSongs   []string `json:"best_songs"`
	WorstSongs  []string `json:"worst_songs"`
	Thoughts    string   `json:"thoughts"`
}
