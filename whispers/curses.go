package whispers

const (
	CURSE_TYPE_ERROR  int = iota
	CURSE_TYPE_HELLO
	CURSE_TYPE_TAXES
	CURSE_TYPE_TAXES_ADDED
	CURSE_TYPE_TAXES_PAID
)

type Curse struct {
	Type    int
	Content string
}

func NewCurse(t int) Curse {
	return Curse{
		Type: 		t,
		Content: 	"Boo!",
	}
}