package public

// File represents an attached file.
type File struct {
	// ID is the identifier for this file.
	ID string
	// CTime is the time the file was uploaded.
	CTime string
	// Name is the name of the file.
	Name string
	// Path is the path of the file.
	Path string
	// ContentType is the type of the file.
	ContentType string
	// FileType is the type of file.
	// Can be "file", "video", "image", "audio" or "screenshare".
	FileType string
	// Size is the size of the file in bytes.
	Size int
	// DownloadURL URL which can be used to get the file.
	DownloadURL string
	// ThumbnailURL is an optional thumbnail URL for this file.
	ThumbnailURL string
	// Author is the person who uploaded the file.
	Author Person
}
