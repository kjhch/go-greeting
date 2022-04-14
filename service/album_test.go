package service

import "testing"

func TestListAlbums(t *testing.T) {
	s := AlbumService{}
	var result []Album
	filter := &Album{
		ID:     NewUint(1),
		Title:  nil,
		Artist: nil,
		Price:  NewFloat64(56.99),
	}
	err := s.ListAlbums(nil, filter, &result)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
