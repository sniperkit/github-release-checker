package gitlab

type Config struct {
	User   string
	Token  string
	Tokens []string
	Stats  bool
	Cache  bool
}
