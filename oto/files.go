package public

// File represents an attached file.
type File struct {
	// ID is the identifier for this file.
	// example: "5f19afce3979fb39"
	ID string
	// CTime is the time the file was uploaded.
	// example: "2020-07-23T15:42:06.597897724Z"
	CTime string
	// Name is the name of the file.
	// example: "filename.jpg"
	Name string
	// Path is the path of the file.
	// example: "/path/to/filename.jpg"
	Path string
	// ContentType is the type of the file.
	// example: "image/jpg"
	ContentType string
	// FileType is the type of file.
	// Can be "file", "video", "image", "audio" or "screenshare".
	// example: "image"
	FileType string
	// Size is the size of the file in bytes.
	// example: 65211
	Size int
	// DownloadURL URL which can be used to get the file.
	// example: "/d/path/to/filename.jpg"
	DownloadURL string
	// ThumbnailURL is an optional thumbnail URL for this file.
	// example: "/d/path/to/filename-thumbnail.jpg"
	ThumbnailURL string
	// Author is the person who uploaded the file.
	Author Person
}
