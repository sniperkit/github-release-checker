package user

type Chan <-chan *User
type chanW chan<- *User
type chanRW chan *User

// TODO(leon): This is ugly
func onlyReadable(in chanRW) <-chan *User {
	return in
}

// TODO(leon): This is ugly
func onlyWritable(in chanRW) chan<- *User {
	return in
}
