package data

import "vslink-linkcategories-lambda/models"

type LinkCategoryRepository struct {
}

func NewLinkCategoryRepository() *LinkCategoryRepository {
	return &LinkCategoryRepository{}
}

// GetAllActive returns list of categories for links
func (cr *LinkCategoryRepository) GetAllActive() ([]models.LinkCategory, error) {

	var categories = []models.LinkCategory{
		{
			LinkCategoryID: "38f4f9d5f12a416193a6c22a56d9c142",
			IconClass:      "fa fa-plus text-primary",
			Name:           "Important links",
		},
		{
			LinkCategoryID: "6ea6ac0db8114f17ba717145e0990ce7",
			IconClass:      "fa fa-graduation-cap text-warning",
			Name:           "Learning",
		},
		{
			LinkCategoryID: "74bbe5264219493688c75f4cd90b2cf4",
			IconClass:      "fa fa-smile-o text-success",
			Name:           "Funny",
		},
	}

	return categories, nil
}
