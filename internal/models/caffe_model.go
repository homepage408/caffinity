package models

type CafeResponse struct {
	ID             int32        `json:"id"`
	Name           string       `json:"name"`
	TagLine        string       `json:"tag_line,omitempty"`
	Address        string       `json:"address,omitempty"`
	Lat            string       `json:"lat,omitempty"`
	Lng            string       `json:"lng,omitempty"`
	Hours          string       `json:"hours,omitempty"`
	Phone          string       `json:"phone,omitempty"`
	Instagram      string       `json:"instagram,omitempty"`
	Rating         string       `json:"rating,omitempty"`
	Reviews        int32        `json:"reviews,omitempty"`
	PriceLevel     int32        `json:"price_level,omitempty"`
	Approved       bool         `json:"approved"`
	ApprovedReason string       `json:"approved_reason,omitempty"`
	HeroImage      string       `json:"hero_image,omitempty"`
	VibeEmoji      string       `json:"vibe_emoji,omitempty"`
	Tags           []Tags       `json:"tags,omitempty"`
	Menus          []Menus      `json:"menus,omitempty"`
	Facilities     []Facilities `json:"facilities,omitempty"`
}

type Menus struct {
	ID          int32  `json:"id"`
	CafeID      int32  `json:"cafe_id"`
	Name        string `json:"name"`
	Price       int32  `json:"price,omitempty"`
	Strength    int32  `json:"strength,omitempty"`
	IsSafe      bool   `json:"is_safe"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
}

type Facilities struct {
	ID     int32  `json:"id"`
	CafeID int32  `json:"cafe_id"`
	Name   string `json:"name"`
}

type Tags struct {
	ID     int32  `json:"id"`
	CafeID int32  `json:"cafe_id"`
	Name   string `json:"name"`
}
