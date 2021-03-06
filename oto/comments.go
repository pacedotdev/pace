package public

// CommentsService allows you to programmatically manage comments in Pace.
type CommentsService interface {
	// AddComment adds a comment.
	AddComment(AddCommentRequest) AddCommentResponse
	// DeleteComment deletes a Comment.
	DeleteComment(DeleteCommentRequest) DeleteCommentResponse
}

// AddCommentRequest is the input object for AddComment.
type AddCommentRequest struct {
	// OrgID is the ID of the org.
	// example: "your-org-id"
	OrgID string
	// TargetKind is the kind of item this comment is for.
	// Can be "card", "message", or "showcase".
	// example: "card"
	TargetKind string
	// TargetID is the ID of the target.
	// example: "123"
	TargetID string
	// Body is the markdown body of the comment.
	// example: "This is my **comment**"
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

// DeleteCommentRequest is the input object for DeleteComment.
type DeleteCommentRequest struct {
	// ID is the ID of the comment to delete.
	// example: "5f19afce3979fb39"
	ID string
	// OrgID is the ID of your org.
	// example: "your-org-id"
	OrgID string
	// TargetKind is the kind of target this Comment was made on.
	// Can be "card", "message", or "showcase".
	// Used to help identify the Comment.
	// example: "card"
	TargetKind string
	// TargetID is the ID of the target.
	// Used to help identify the Comment.
	// example: "123"
	TargetID string
}

// DeleteCommentResponse is the output object for DeleteComment.
type DeleteCommentResponse struct{}
