package xorm_type_assist

type BoolType = string

const (
	TRUE  BoolType = "T"
	FALSE BoolType = "F"
)

type GenderType = string

const (
	Feminine  GenderType = "f"
	Masculine GenderType = "m"
	UniGender GenderType = "u"
)
