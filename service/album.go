package service

import (
	"context"
)

type Album struct {
	ID     *uint    `json:"id"`
	Title  *string  `json:"title"`
	Artist *string  `json:"artist"`
	Price  *float64 `json:"price"`
}

type AlbumService struct {
}

func (s *AlbumService) ListAlbums(ctx context.Context, filter *Album, result *[]Album) error {
	*result = make([]Album, 0)
	for _, v := range albums {
		isSatisfy := true
		if filter.ID != nil {
			isSatisfy = isSatisfy && *filter.ID == *v.ID
		}
		if filter.Title != nil {
			isSatisfy = isSatisfy && *filter.Title == *v.Title
		}
		if filter.Artist != nil {
			isSatisfy = isSatisfy && *filter.Artist == *v.Artist
		}
		if filter.Price != nil {
			isSatisfy = isSatisfy && *filter.Price == *v.Price
		}
		if isSatisfy {
			*result = append(*result, v)
		}
	}
	return nil
}

func newUint(v uint) *uint {
	return &v
}
func newString(v string) *string {
	return &v
}
func newFloat64(v float64) *float64 {
	return &v
}

var albums = []Album{
	{ID: newUint(1), Title: newString("Blue Train"), Artist: newString("John Coltrane"), Price: newFloat64(56.99)},
	{ID: newUint(2), Title: newString("Jeru"), Artist: newString("Gerry Mulligan"), Price: newFloat64(17.99)},
	{ID: newUint(3), Title: newString("Sarah Vaughan and Clifford Brown"), Artist: newString("Sarah Vaughan"), Price: newFloat64(39.99)},
}
