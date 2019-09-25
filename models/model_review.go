package models

type BodyReview struct {
	entity_type string
	id          string
	variables   Variables
}

type Variables struct {
	comment            string
	participants       []string
	validationOrReview string
}
