package model

func Models() []interface{} {
	return []interface{}{
		new(Amway),
		new(User),
		new(Relation),
	}
}
