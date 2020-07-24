package public

// CommentsService allows you to create comments in Pace.
type CommentsService interface {
	// AddComment adds a comment.
	AddComment(AddCommentRequest) AddCommentResponse
}

// AddCommentRequest is the input object for AddComment.
type AddCommentRequest struct {
	// OrgID is the ID of the org.
	// example: "your-org-id"
	OrgID string
	// TargetKind is the kind of item this comment is for.
	// Can be "card", "message", "showcase".
	// example: "card"
	TargetKind string
	// TargetID is the ID of the target.
	TargetID string
	// Body is the markdown body of the comment.
	Body string
}

// AddCommentResponse is the output object for AddComment.
type AddCommentResponse struct {
	// Comment is the comment that was created.
	Comment Comment
}

// Comment is a single comment in Pace.
type Comment struct {
	// ID is the ID of the comment.
	// example: "5f19afce3979fb39"
	ID string
	// CTime is the time this was created.
	// example: "2020-07-23T15:42:06.597897724Z"
	CTime string
	// MTime is the time this comment was last modified.
	// example: "2020-07-23T15:42:06.597897724Z"
	MTime string
	// Body is the markdown body of the comment.
	// example: "Hello **Pace**"
	Body string
	// BodyHTML is the HTML formatted body of the comment.
	// example: "<p>Hello <strong>Pace</strong></p>"
	BodyHTML string
	// Author is the person who posted this comment.
	Author Person
}
