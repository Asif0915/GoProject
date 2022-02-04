package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Age       int                `json:"age,omitempty" bson:"age,omitempty"`
}
type Books struct {
	B_ID   string `json:"b_id,omitempty" bson:"b_id,omitempty"`
	BName  string `json:"bname,omitempty" bson:"bname,omitempty"`
	Bprice int    `json:"bprice,omitempty" bson:"price,omitempty"`
}
type Borrow struct {
	Person primitive.ObjectID `json:"p_id,omitempty" bson:"p_id,omitempty"`
	B_ID   string             `json:"b_id,omitempty" bson:"b_id,omitempty"`
	B_date string             `json:"date,omitempty" bson:"price,omitempty"`
}

// type Borrowed struct {
// 	B_ID        string `json:"b_id,omitempty" bson:"b_id,omitempty"`
// 	Date_borrow string `json:"price,omitempty" bson:"price,omitempty"`
// 	BName       string `json:"bname,omitempty" bson:"bname,omitempty"`
// 	Person primitive.ObjectID `json:"p_id,omitempty" bson:"p_id,omitempty"`
// }

// type Podcast struct {
// 	ID     primitive.ObjectID `bson:"_id,omitempty"`
// 	Title  string             `bson:"title,omitempty"`
// 	Author string             `bson:"author,omitempty"`
// 	Tags   []string           `bson:"tags,omitempty"`
// }
// type Episode struct {
// 	ID          primitive.ObjectID `bson:"_id,omitempty"`
// 	Podcast     primitive.ObjectID `bson:"podcast,omitempty"`
// 	Title       string             `bson:"title,omitempty"`
// 	Description string             `bson:"description,omitempty"`
// 	Duration    int32              `bson:"duration,omitempty"`
// }
// type PodcastEpisode struct {
// 	ID          primitive.ObjectID `bson:"_id,omitempty"`
// 	Podcast     Podcast            `bson:"podcast,omitempty"`
// 	Title       string             `bson:"title,omitempty"`
// 	Description string             `bson:"description,omitempty"`
// 	Duration    int32              `bson:"duration,omitempty"`
// }
