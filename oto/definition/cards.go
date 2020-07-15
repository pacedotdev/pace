package public

// CardsService is used to work with cards.
type CardsService interface {
	// GetCard gets a card.
	GetCard(GetCardRequest) GetCardResponse
	// CreateCard creates a new Card.
	CreateCard(CreateCardRequest) CreateCardResponse
}

type GetCardRequest struct {
	OrgID  string
	CardID string
}

type GetCardResponse struct {
	Card Card
}

type CreateCardRequest struct {
	// OrgID is the org ID in which to create the card.
	OrgID string
	// TeamID is the team ID in which to create the card.
	TeamID string
	// Title is the title of the card.
	Title string
	// ParentTargetKind is the kind of target to relate this card to (e.g. card or message)
	ParentTargetKind string
	// ParentTargetID is the ID of the item to relate this new card to.
	ParentTargetID string
}

type CreateCardResponse struct {
	// Card is the card that was just created.
	Card Card
}

type Card struct {
	// ID is the unique ID of the card within the org.
	ID       string
	CTime    string
	MTime    string
	Order    float64
	TeamID   string
	Slug     string
	Title    string
	Status   string
	Author   Person
	Body     string
	BodyHTML string
	Tags     []string
	// TakenByCurrentUser indicates whether the current user has
	// taken this card or not.
	TakenByCurrentUser bool
	// TakenByPeople is a list of people who have taken responsibility
	// for this Card.
	TakenByPeople []Person
	// Files are the list of files that are attached to
	// this Card.
	Files               []File
	RelatedCardsSummary RelatedCardsSummary
}

type RelatedCardsSummary struct {
	Total    int
	Done     int
	Progress int
}
