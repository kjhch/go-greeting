package service

import (
	"context"
	"fmt"
)

type Album struct {
	ID     *uint    `json:"id"`
	Title  *string  `json:"title"`
	Artist *string  `json:"artist"`
	Price  *float64 `json:"price"`
}

func (a Album) String() string {
	return fmt.Sprintf("{%v %v %v %v}", *a.ID, *a.Title, *a.Artist, *a.Price)
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

func NewUint(v uint) *uint {
	return &v
}
func NewString(v string) *string {
	return &v
}
func NewFloat64(v float64) *float64 {
	return &v
}

var albums = []Album{
	{ID: NewUint(1), Title: NewString("Blue Train"), Artist: NewString("John Coltrane"), Price: NewFloat64(56.99)},
	{ID: NewUint(2), Title: NewString("Jeru"), Artist: NewString("Gerry Mulligan"), Price: NewFloat64(17.99)},
	{ID: NewUint(3), Title: NewString("Sarah Vaughan and Clifford Brown"), Artist: NewString("Sarah Vaughan"), Price: NewFloat64(39.99)},
}
