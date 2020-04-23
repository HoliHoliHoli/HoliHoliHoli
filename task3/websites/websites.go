package websites

type Twitter struct {
	URL       string
	Username  string
	Followers int
}

// Feed returns the latest Twitter posts
func (t *Twitter) Feed() []string {
	return []string{
		"Twitter feeds",
		"Coding is basically copying other people's work",
	}
}

// Fame tell how famous a user is. The higher
// the number of followers the more famous
func (t *Twitter) Fame() int {
	return t.Followers
}




type Facebook struct {
	URL     string
	Name    string
	Friends int
}

// Feed returns the latest Facebook posts
func (f *Facebook) Feed() []string {
	return []string{
		"Facebook feeds",
		"Hey, here's my cool new selfie",
		"What is code?",
	}
}

// Fame tell how famous a user is. The higher
// the number of friends the more famous
func (f *Facebook) Fame() int {
	return f.Friends
}




type Linkedin struct {
	URL         string
	Name        string
	Connections int
}

// Feed returns the latest Linkedin posts
func (l *Linkedin) Feed() []string {
	return []string{
		"LinkedIn feeds",
		"Hey, I just started a new position at Hotels.ng",
	}
}

// Fame tell how famous a user is. The higher
// the number of Connections the more famous
func (l *Linkedin) Fame() int {
	return l.Connections
}
