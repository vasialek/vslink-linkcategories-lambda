package models

import "time"

// TypeOfResource use to set status and type of resources, like LinkCategory
type TypeOfResource int

const (
	// ActivePublic active and available to all users
	ActivePublic TypeOfResource = iota + 1
	// ActivePrivate active, but only for owner
	ActivePrivate
	// DeletedPublic deleted
	DeletedPublic
	// DeletedPrivate deleted
	DeletedPrivate
)

// LinkCategory to store information about categories
type LinkCategory struct {
	LinkCategoryID string         `json:"link_category_id"`
	StatusID       TypeOfResource `json:"status_id"`
	UserID         string         `json:"user_id"`
	Name           string         `json:"name"`
	IconClass      string         `json:"icon_class"`
	CreatedAt      time.Time      `json:"created_at"`
}
