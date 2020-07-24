package public

// CardsService is used to work with cards.
type CardsService interface {
	// GetCard gets a card.
	GetCard(GetCardRequest) GetCardResponse
	// CreateCard creates a new Card.
	CreateCard(CreateCardRequest) CreateCardResponse
	// UpdateCardStatus updates a card's status.
	UpdateCardStatus(UpdateCardStatusRequest) UpdateCardStatusResponse
}

// GetCardRequest is the input object for GetCard.
type GetCardRequest struct {
	// OrgID is the ID of the org.
	// example: "your-org-id"
	OrgID string
	// CardID is the ID of the card to get.
	// example: "123"
	CardID string
}

// GetCardResponse is the output object for GetCard.
type GetCardResponse struct {
	// Card is the card.
	Card Card
}

// CreateCardRequest is the input object for CreateCard.
type CreateCardRequest struct {
	// OrgID is the org ID in which to create the card.
	// example: "your-org-id"
	OrgID string
	// TeamID is the team ID in which to create the card.
	// example: "your-team-id"
	TeamID string
	// Title is the title of the card.
	// example: "This is my new card"
	Title string
	// ParentTargetKind is the kind of target to relate this card to (e.g. card or message)
	// example: "cards"
	ParentTargetKind string
	// ParentTargetID is the ID of the item to relate this new card to.
	// example: "123"
	ParentTargetID string
}

// CreateCardResponse is the output object for CreateCard.
type CreateCardResponse struct {
	// Card is the card that was just created.
	Card Card
}

// Card is a card in Pace.
type Card struct {
	// ID is the unique ID of the card within the org.
	// example: "123"
	ID string
	// CTime is the time this was created.
	// example: "2020-07-23T15:42:06.597897724Z"
	CTime string
	// MTime is the time this comment was last modified.
	// example: "2020-07-23T15:42:06.597897724Z"
	MTime string
	// TeamID is the ID of the team that this card belongs to.
	// example: "your-team-id"
	TeamID string
	// Slug is the URL slug for this card.
	// example: "this-is-my-card"
	Slug string
	// Title is the title of the card.
	// example: "This is my card"
	Title string
	// Status is the current status of the card.
	// example: "done"
	Status string
	// Author is the Person who created this card.
	Author Person
	// Body is the markdown body of this card.
	// example: "Hello from **Pace**."
	Body string
	// BodyHTML is the HTML rendering of the body of this card.
	// example: "<p>Hello from <strong>Pace</strong>."
	BodyHTML string
	// Tags is a list of tags associated with this card.
	// example: ["bug", "question"]
	Tags []string
	// TakenByCurrentUser indicates whether the current user has
	// taken this card or not.
	// example: true
	TakenByCurrentUser bool
	// TakenByPeople is a list of people who have taken responsibility
	// for this Card.
	TakenByPeople []Person
	// Files are the list of files that are attached to
	// this Card.
	Files []File
}

// UpdateCardStatusRequest is the input object UpdateCardStatus.
type UpdateCardStatusRequest struct {
	// OrgID is the ID of the org.
	OrgID string
	// CardID is the ID number of the card.
	CardID string
	// Status is the new status of the card.
	// Valid strings are "future", "next", "progress", "done".
	Status string
}

// UpdateCardStatusResponse is the output object for UpdateCardService.
type UpdateCardStatusResponse struct {
	// Card is the card that was updated.
	Card Card
}
