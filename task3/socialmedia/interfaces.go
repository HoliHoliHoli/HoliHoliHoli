package socialmedia

//SocialMedia blah
type SocialMedia interface {
	Feed() []string
	Fame() int
}

type xmlstruct struct{
	SocialMedia
	feed []string
	fame int
}