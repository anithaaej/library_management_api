package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name            string             `json:"name,omitempty"`
	AuthorName      string             `json:"authorname,omitempty"`
	BookCategory    string             `json:"bookcategory,omitempty"`
	ISBN            string             `json:"isbn,omitempty"`
	CoverImage      string             `json:"coverimage,omitempty"`
	Rating          float32            `json:"rating,omitempty"`
	CopiesAvailable int                `json:"copiesavailable"`
	RentedCopies    int                `json:"rentedcopies"`
}
