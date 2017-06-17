package auth

//Googleauth struct reperesnts authentication struct for google.
type Googleauth struct {
	ClientID, SecretKey string
}



//AWSauth struct reperesnts authentication struct for google.
type AWSauthbase struct {
	ClientID, SecretKey string
}



type Auth struct {
	AccessKey string
	SecretKey string
}
