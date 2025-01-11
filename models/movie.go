package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`    // ID único gerado pelo MongoDB
    Title    string             `bson:"title" json:"title"`         // Título do filme
    Director string             `bson:"director" json:"director"`   // Nome do diretor
    Year     int                `bson:"year" json:"year"`           // Ano de lançamento
    Genre    string             `bson:"genre" json:"genre"`         // Gênero do filme
    Rating   string             `bson:"rating" json:"rating"`       // Avaliação do filme
}

var validGenres = map[string]bool{
	"Action":                 true,
	"Adventure":              true,
	"Animation":              true,
	"Biography":              true,
	"Comedy":                 true,
	"Crime":                  true,
	"Documentary":            true,
	"Drama":                  true,   
	"Family":                 true,
	"Fantasy":                true,
	"History":                true,
	"Horror":                 true,
	"Mystery":                true,
	"Musical":                true,
	"Romance":                true,
	"Science Fiction":        true,
	"Sport":                  true,
	"Thriller":               true,
	"War":                    true,
	"Western":                true,
	"Action Comedy":          true,
	"Superhero":              true,
	"Dark Comedy":            true,
	"Psychological Thriller": true,
	"Epic":                   true,
	"Adventure Comedy":       true,
	"Science Fantasy":        true,
	"Cyberpunk":              true,
	"Steampunk":              true,
	"Martial Arts":           true,
	"Spy":                    true,
	"Political Thriller":     true,
	"Historical Fiction":     true,
	"Noir":                   true,
	"Slice of Life":          true,
	"Survival":               true,
	"Disaster":               true,
	"Parody":                 true,
	"Satire":                 true,
	"Romantic Comedy":        true,
	"Coming of Age":          true,
	"Crime Drama":            true,
	"Legal Drama":            true,
	"Fantasy Adventure":      true,
	"Urban Fantasy":          true,
	"Mythology":              true,
	"Western Comedy":         true,
	"Experimental":           true,
	"Silent Film":            true,
	"Road Movie":             true,
}