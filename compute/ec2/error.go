package ec2

// Error represents error in EC2 instance creation.
type Error struct {
	StatusCode int

	Code string

	Message string

	RequestID string `xml:"RequestID"`
}
