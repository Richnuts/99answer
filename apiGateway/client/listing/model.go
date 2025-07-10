package listing

type CreateListingInput struct {
	UserID      int    `json:"user_id"`
	ListingType string `json:"listing_type"`
	Price       int    `json:"price"`
}
