package main

import (
	"fmt"

	"math/rand"
	"time"

	"gioui.org/layout"
)

var (
	two_bubi   float32 = 2
	three_bubi float32 = 3
	four_bubi  float32 = 4
	five_bubi  float32 = 5
	six_bubi   float32 = 6
	seven_bubi float32 = 7
	eight_bubi float32 = 8
	nine_bubi  float32 = 9
	ten_bubi   float32 = 10
	jack_bubi  float32 = 10
	queen_bubi float32 = 10
	king_bubi  float32 = 10
	Ace_bubi   float32 = 11

	two_Krest   float32 = 2
	three_Krest float32 = 3
	four_Krest  float32 = 4
	five_Krest  float32 = 5
	six_Krest   float32 = 6
	seven_Krest float32 = 7
	eight_Krest float32 = 8
	nine_Krest  float32 = 9
	ten_Krest   float32 = 10
	jack_Krest  float32 = 10
	queen_Krest float32 = 10
	king_Krest  float32 = 10
	Ace_Krest   float32 = 11

	two_Piki   float32 = 2
	three_Piki float32 = 3
	four_Piki  float32 = 4
	five_Piki  float32 = 5
	six_Piki   float32 = 6
	seven_Piki float32 = 7
	eight_Piki float32 = 8
	nine_Piki  float32 = 9
	ten_Piki   float32 = 10
	jack_Piki  float32 = 10
	queen_Piki float32 = 10
	king_Piki  float32 = 10
	Ace_Piki   float32 = 11

	two_Chervi   float32 = 2
	three_Chervi float32 = 3
	four_Chervi  float32 = 4
	five_Chervi  float32 = 5
	six_Chervi   float32 = 6
	seven_Chervi float32 = 7
	eight_Chervi float32 = 8
	nine_Chervi  float32 = 9
	ten_Chervi   float32 = 10
	jack_Chervi  float32 = 10
	queen_Chervi float32 = 10
	king_Chervi  float32 = 10
	Ace_Chervi   float32 = 11
)
var deck = []Card{
	{"двойка бубей", two_bubi},
	{"тройка бубей", three_bubi},
	{"четверка бубей", four_bubi},
	{"пятерка бубей", five_bubi},
	{"шестерка бубей", six_bubi},
	{"семерка бубей", seven_bubi},
	{"восьмерка бубей", eight_bubi},
	{"девятка бубей", nine_bubi},
	{"десятка бубей", ten_bubi},
	{"валет бубей", jack_bubi},
	{"дама бубей", queen_bubi},
	{"король бубей", king_bubi},
	{"туз бубей", Ace_bubi},
	{"двойка крестей", two_Krest},
	{"тройка крестей", three_Krest},
	{"четверка крестей", four_Krest},
	{"пятерка крестей", five_Krest},
	{"шестерка крестей", six_Krest},
	{"семерка крестей", seven_Krest},
	{"восьмерка крестей", eight_Krest},
	{"девятка крестей", nine_Krest},
	{"десятка крестей", ten_Krest},
	{"валет крестей", jack_Krest},
	{"дама крестей", queen_Krest},
	{"король крестей", king_Krest},
	{"туз крестей", Ace_Krest},
	{"двойка пик", two_Piki},
	{"тройка пик", three_Piki},
	{"четверка пик", four_Piki},
	{"пятерка пик", five_Piki},
	{"шестерка пик", six_Piki},
	{"семерка пик", seven_Piki},
	{"восьмерка пик", eight_Piki},
	{"девятка пик", nine_Piki},
	{"десятка пик", ten_Piki},
	{"валет пик", jack_Piki},
	{"дама пик", queen_Piki},
	{"король пик", king_Piki},
	{"туз пик", Ace_Piki},
	{"двойка червей", two_Chervi},
	{"тройка червей", three_Chervi},
	{"четверка червей", four_Chervi},
	{"пятерка червей", five_Chervi},
	{"шестерка червей", six_Chervi},
	{"семерка червей", seven_Chervi},
	{"восьмерка червей", eight_Chervi},
	{"девятка червей", nine_Chervi},
	{"десятка червей", ten_Chervi},
	{"валет червей", jack_Chervi},
	{"дама червей", queen_Chervi},
	{"король червей", king_Chervi},
	{"туз червей", Ace_Chervi},
}

type C = layout.Context
type D = layout.Dimensions
type Card struct {
	Name  string
	Value float32
}

func main() {

	var NowBalance float32
	NowBalance = 1000
	ActualBalance := &NowBalance
	fmt.Printf("Добро пожаловать в казино! На вашем балансе %v очков.", NowBalance)
	fmt.Println(" Начнём игру!")
	fmt.Println("Управление происходит с помощью цифр")

	for {
		var UserChoise int
		time.Sleep(1100 * time.Millisecond)
		fmt.Println("Что бы начать новую раздачу - введите 1")
		fmt.Println("Посмотреть баланс - введите 2")
		fmt.Scan(&UserChoise)
		switch UserChoise {
		case 1:
			GameStart(deck, ActualBalance)
		case 2:
			SeeBalance(ActualBalance)

		}
	}

}

func GameStart(deck []Card, Balance *float32) {
	var UserStake float32
	fmt.Println("Сделайте вашу ставку: ")
	fmt.Scan(&UserStake)
	for {
		if UserStake > *Balance {
			fmt.Println("Не хватает очков для ставки, введите корректную сумму")
			fmt.Scan(&UserStake)
		} else {
			break
		}
	}
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Карта дилера:")
	DillerCard1 := getRandomCard(deck)
	DillerHandSclice := []Card{DillerCard1}
	checkBlackjack(DillerHandSclice)
	adjustAceValue(&DillerHandSclice)
	time.Sleep(1500 * time.Millisecond)
	fmt.Println(DillerCard1.Name)
	time.Sleep(1600 * time.Millisecond)
	PlayerCard1 := getRandomCard(deck)
	PlayerHand := []Card{PlayerCard1}
	adjustAceValue(&PlayerHand)
	checkBlackjack(PlayerHand)
	fmt.Println("Ваша карта:")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println(PlayerCard1.Name)
	time.Sleep(1600 * time.Millisecond)
	fmt.Println("Диллер получает вторую карту")
	DillerCard2 := getRandomCard(deck)
	DillerHand2 := []Card{DillerCard2}
	adjustAceValue(&DillerHand2)
	checkBlackjack(DillerHand2)
	if checkBlackjack(DillerHand2) == true {
		time.Sleep(1600 * time.Millisecond)
		fmt.Println(DillerCard2.Name)
		time.Sleep(1300 * time.Millisecond)
		fmt.Println("У диллера блэкджек, вы проиграли!")
		*Balance -= *Balance
		return
	}
	time.Sleep(1100 * time.Millisecond)
	fmt.Println("Ваша вторая карта:")
	PlayerCard2 := getRandomCard(deck)
	PlayerHand2 := []Card{PlayerCard2}
	adjustAceValue(&PlayerHand2)
	checkBlackjack(PlayerHand2)
	fmt.Println(PlayerCard2.Name)
	if checkBlackjack(PlayerHand2) == true {
		time.Sleep(1600 * time.Millisecond)
		fmt.Println("БЛЭКДЖЕК, вы выиграли!")
		fmt.Printf(" Ваш выигрыш: %v", UserStake*1, 5)
		*Balance = *Balance * 1.5
		return
	}
	PlayerValue := PlayerCard1.Value + PlayerCard2.Value
	time.Sleep(1100 * time.Millisecond)
	fmt.Printf("Ваша рука: %v, каковы дальнейшие действия?\n", PlayerValue)
	time.Sleep(1100 * time.Millisecond)
	fmt.Println("1. Ещё одну")
	fmt.Println("2. Достаточно")
	//Тут основной цикл управления игрока
	for {
		var SecondChoice int
		fmt.Scan(&SecondChoice)

		if SecondChoice == 1 {
			PlayerCard := getRandomCard(deck)
			PlayerHand3 := []Card{PlayerCard}
			adjustAceValue(&PlayerHand3)
			time.Sleep(1100 * time.Millisecond)
			fmt.Println("Ваша следующая карта:", PlayerCard.Name)
			PlayerValue += PlayerCard.Value

			if PlayerValue > 21 {
				time.Sleep(1100 * time.Millisecond)
				fmt.Println("У Вас больше 21, вы проиграли!")
				*Balance -= UserStake
				break
			}
			time.Sleep(1100 * time.Millisecond)
			fmt.Printf("Ваша рука: %v, каковы дальнейшие действия?\n", PlayerValue)
			fmt.Println("1. Ещё одну")
			fmt.Println("2. Достаточно")
		} else if SecondChoice == 2 {
			DealerValue := DillerCard1.Value + DillerCard2.Value
			PlayerPressedEnough(DealerValue, PlayerValue, UserStake, Balance, DillerCard2.Name)
			break
		} else {
			fmt.Println("Некорректный ввод, попробуйте снова.")
		}
	}
}

func PlayerPressedEnough(dealerResult, playerResult, stake float32, balance *float32, DillerCard2Name string) {

	time.Sleep(2550 * time.Millisecond)
	fmt.Println("Диллер открывает свою карту:")
	time.Sleep(1550 * time.Millisecond)
	fmt.Println(DillerCard2Name)
	for {
		if dealerResult < 17 {
			fmt.Println("Дилер берёт карту:")
			time.Sleep(1100 * time.Millisecond)
			newdealercard := getRandomCard(deck)
			DealerHand3 := []Card{newdealercard}
			adjustAceValue(&DealerHand3)
			if dealerResult > 11 {

			}
			fmt.Println(newdealercard.Name)
			dealerResult += newdealercard.Value
		} else if dealerResult >= 17 {
			fmt.Println("Диллер перестал брать карты")
			break
		}
	}

	switch {
	case dealerResult >= 22:
		time.Sleep(1100 * time.Millisecond)
		fmt.Println("У диллера больше 21, вы выиграли!")
		*balance += stake
	case dealerResult < 22:

		mathResult := dealerResult - playerResult
		if mathResult > 0 {
			fmt.Printf("Вы проиграли, у диллера %v, у вас - %v, на %v меньше\n", dealerResult, playerResult, mathResult)
			*balance -= stake
		} else if mathResult < 0 {
			fmt.Printf("Вы выиграли: ваши %v против %v дилера, вы выиграли %v!!!", playerResult, dealerResult, stake*2)
			*balance += stake
		} else {
			fmt.Println("Ничья!")
		}
	}
}

func getRandomCard(deck []Card) Card {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	return deck[rnd.Intn(len(deck))]
}
func SeeBalance(balance *float32) {
	time.Sleep(1100 * time.Millisecond)
	fmt.Println("Ваш баланс:")
	fmt.Println(*balance)
}

// Функция для корректировки значения туза
func adjustAceValue(hand *[]Card) {
	for i := range *hand {
		if (*hand)[i].Value == 11 {
			// Calculate the total value of the hand
			totalValue := calculateHandValue(*hand)
			// If the total value exceeds 21, change the Ace's value to 1
			if totalValue > 21 {
				(*hand)[i].Value = 1
				// Recalculate the total value after adjustment
				totalValue = calculateHandValue(*hand)
				// If the adjusted total value is still over 21, continue adjusting Aces
				if totalValue > 21 {
					continue
				} else {
					// If the adjusted total value is 21 or under, stop adjusting
					break
				}
			}
		}
	}
}
func checkBlackjack(hand []Card) bool {
	// Check for an Ace and a ten-value card in the hand
	hasAce := false
	hasTenValueCard := false

	for _, card := range hand {
		if card.Value == 11 {
			hasAce = true
		}
		if card.Value == 10 {
			hasTenValueCard = true
		}
	}

	// If the hand has both an Ace and a ten-value card, and the total value is 21, it's a blackjack
	return hasAce && hasTenValueCard && calculateHandValue(hand) == 21
}

// Функция для подсчёта общего значения карт в руке
func calculateHandValue(hand []Card) float32 {
	var totalValue float32 = 0
	for _, card := range hand {
		totalValue += card.Value
	}
	return totalValue
}
