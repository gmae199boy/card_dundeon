package card

type CardType int

const (
	CHARACTER CardType = iota + 1
	PASSIVE
	MAGIC
	EQUIP
)

var (
	cardList []Card
)

type Card struct {
	name     string
	desc     string
	atk      int
	hp       int
	sd       int
	cardType CardType
}

func Init(cardType CardType, data []Card) {

}
