package public

type CommentsService interface {
	// AddComment adds a comment.
	AddComment(AddCommentRequest) AddCommentResponse
}

type AddCommentRequest struct {
	OrgID      string
	TargetKind string
	TargetID   string
	Body       string
}

type AddCommentResponse struct {
	Comment Comment
}

type Comment struct {
	ID       string
	CTime    string
	MTime    string
	Body     string
	BodyHTML string
	Author   Person
	Files    []File
}
