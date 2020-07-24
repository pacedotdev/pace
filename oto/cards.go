package public

// CardsService allows you to programmatically manage cards in Pace.
type CardsService interface {
	// GetCard gets a card.
	GetCard(GetCardRequest) GetCardResponse
	// CreateCard creates a new Card.
	CreateCard(CreateCardRequest) CreateCardResponse
	// UpdateCard updates the title and body of the card.
	UpdateCard(UpdateCardRequest) UpdateCardResponse
	// UpdateCardStatus updates a card's status.
	UpdateCardStatus(UpdateCardStatusRequest) UpdateCardStatusResponse
	// TakeCard takes responsibility for a card.
	// Can be undone with PutBackCard.
	TakeCard(TakeCardRequest) TakeCardResponse
	// PutBackCard removes a user from the list of responsbile users.
	// Undoes TakeCard.
	PutBackCard(PutBackCardRequest) PutBackCardResponse
	// DeleteCard deletes a card.
	DeleteCard(DeleteCardRequest) DeleteCardResponse
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
	// ParentTargetKind is the kind of target to relate this card to (e.g. "card", "message", or "showcase")
	// example: "card"
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
	// example: "your-org-id"
	OrgID string
	// CardID is the ID number of the card.
	// example: "123"
	CardID string
	// Status is the new status of the card.
	// Valid strings are "future", "next", "progress", "done".
	// example: "progress"
	Status string
}

// UpdateCardStatusResponse is the output object for UpdateCardService.
type UpdateCardStatusResponse struct {
	// Card is the card that was updated.
	Card Card
}

// UpdateCardRequest is the input object for UpdateCard.
type UpdateCardRequest struct {
	// OrgID is the ID of the org.
	// example: "your-org-id"
	OrgID string
	// CardID is the ID of the card to update.
	// example: "123"
	CardID string
	// Title is the new title for the card.
	// example: "New title"
	Title string
	// Body is the new markdown body for the card.
	// example: "I have been **updated**"
	Body string
}

// UpdateCardResponse is the output object for UpdateCard.
type UpdateCardResponse struct {
	// Card is the recently updated Card.
	Card Card
}

// TakeCardRequest is the input object for TakeCard.
type TakeCardRequest struct {
	// OrgID is the ID of your org.
	// example: "your-org-id"
	OrgID string
	// CardID is the ID of the card to assign yourself to.
	// example: "123"
	CardID string
}

// TakeCardResponse is the output object for TakeCard.
type TakeCardResponse struct {
	// Card is the newly updated Card.
	Card Card
}

// PutBackCardRequest is the input object for PutBackCard.
type PutBackCardRequest struct {
	// OrgID is the ID of your org.
	// example: "your-org-id"
	OrgID string
	// CardID is the ID of the card to unassign yourself from.
	// example: "123"
	CardID string
}

// PutBackCardResponse is the output object for PutBackCard.
type PutBackCardResponse struct {
	// Card is the newly updated Card.
	Card Card
}

// DeleteCardRequest is the input object for DeleteCard.
type DeleteCardRequest struct {
	// OrgID is the ID of your org.
	// example: "your-org-id"
	OrgID string
	// CardID is the ID of the card to delete.
	// example: "123"
	CardID string
}

// DeleteCardResponse is the output object for DeleteCard.
type DeleteCardResponse struct{}
