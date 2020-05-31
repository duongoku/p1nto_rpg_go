package p1nto

type player struct {
	Money int
	Hp int
	Atk int
	Armor int
}

func CheckPlayer(userID string) {
	tmp, ok := users[userID]
	if !ok {
		users[userID] := player{0, 50, 10, 1}
	}
	return
}