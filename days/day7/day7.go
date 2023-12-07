package day7

import (
	"aoc2023/internal"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var CardCharRank = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}
var JokerCardCharRank = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

type ComboRank int

const (
	fiveOfAKind  ComboRank = 7
	fourOfAKind  ComboRank = 6
	fullHouse    ComboRank = 5
	threeOfAKind ComboRank = 4
	twoPairs     ComboRank = 3
	onePair      ComboRank = 2
	highCard     ComboRank = 1
)

type Hand struct {
	bid       int
	cards     []string
	comboRank ComboRank
}

func CreateHand(cards []string, bid int) (Hand, error) {
	if len(cards) != 5 {
		return Hand{}, errors.New("bad input value")
	}

	hand := Hand{cards: cards, bid: bid}
	cardMap := make(map[string]int)

	for _, card := range cards {
		_, hasCard := CardCharRank[card]
		if !hasCard {
			return hand, errors.New("bad input value")
		}

		_, hasCard = cardMap[card]
		if !hasCard {
			cardMap[card] = 1
			continue
		}

		cardMap[card] += 1
	}

	highestCount := 0
	for _, count := range cardMap {
		if count > highestCount {
			highestCount = count
		}
	}

	switch highestCount {
	case 5:
		hand.comboRank = fiveOfAKind
	case 4:
		hand.comboRank = fourOfAKind
	case 3:
		hand.comboRank = threeOfAKind
		for _, count := range cardMap {
			if count == 2 {
				hand.comboRank = fullHouse
				break
			}
		}
	case 2:
		// check if there is second pair
		pairCount := 0
		for _, count := range cardMap {
			if count == 2 {
				pairCount++
			}
		}
		switch pairCount {
		case 2:
			hand.comboRank = twoPairs
		case 1:
			hand.comboRank = onePair
		default:
			return hand, errors.New("bad input values")
		}
	case 1:
		hand.comboRank = highCard
	default:
		return hand, errors.New("bad input value")
	}

	return hand, nil
}

func CreateJokerHand(cards []string, bid int) (Hand, error) {
	if len(cards) != 5 {
		return Hand{}, errors.New("bad input value")
	}

	hand := Hand{cards: cards, bid: bid}
	cardMap := make(map[string]int)

	jokerCount := 0
	for _, card := range cards {
		if card == "J" {
			jokerCount++
			continue
		}

		_, hasCard := CardCharRank[card]
		if !hasCard {
			return hand, errors.New("bad input value")
		}

		_, hasCard = cardMap[card]
		if !hasCard {
			cardMap[card] = 1
			continue
		}

		cardMap[card] += 1
	}

	highestCount := 0
	for _, count := range cardMap {
		if count > highestCount {
			highestCount = count
		}
	}

	switch highestCount {
	case 5:
		hand.comboRank = fiveOfAKind
	case 4:
		hand.comboRank = fourOfAKind
		if jokerCount == 1 {
			hand.comboRank = fiveOfAKind
		}
	case 3:
		hand.comboRank = threeOfAKind
		switch jokerCount {
		case 2:
			hand.comboRank = fiveOfAKind
		case 1:
			hand.comboRank = fourOfAKind
		case 0:
			for _, count := range cardMap {
				if count == 2 {
					hand.comboRank = fullHouse
					break
				}
			}
		default:
			return hand, errors.New("bad logic")
		}

	case 2:
		// check if there is second pair
		pairCount := 0
		for _, count := range cardMap {
			if count == 2 {
				pairCount++
			}
		}
		switch pairCount {
		case 2:
			hand.comboRank = twoPairs
			if jokerCount == 1 {
				hand.comboRank = fullHouse
			}
		case 1:
			switch jokerCount {
			case 3:
				hand.comboRank = fiveOfAKind
			case 2:
				hand.comboRank = fourOfAKind
			case 1:
				hand.comboRank = threeOfAKind
			default:
				hand.comboRank = onePair
			}
		default:
			return hand, errors.New("bad logic")
		}
	case 1:
		hand.comboRank = highCard
		switch jokerCount {
		case 4:
			hand.comboRank = fiveOfAKind
		case 3:
			hand.comboRank = fourOfAKind
		case 2:
			hand.comboRank = threeOfAKind
		case 1:
			hand.comboRank = onePair
		}
	case 0:
		if jokerCount == 5 {
			hand.comboRank = fiveOfAKind
		}
	default:
		return hand, errors.New("bad input value")
	}

	return hand, nil
}

func (ha Hand) isBetter(hb Hand) bool {
	if ha.comboRank == hb.comboRank {
		for cardIndex := range ha.cards {
			cardRankA := CardCharRank[ha.cards[cardIndex]]
			cardRankB := CardCharRank[hb.cards[cardIndex]]

			if cardRankA != cardRankB {
				return cardRankA > cardRankB
			}
		}

		return false
	}

	return ha.comboRank > hb.comboRank
}

func (ha Hand) isJokerBetter(hb Hand) bool {
	if ha.comboRank == hb.comboRank {
		for cardIndex := range ha.cards {
			cardRankA := JokerCardCharRank[ha.cards[cardIndex]]
			cardRankB := JokerCardCharRank[hb.cards[cardIndex]]

			if cardRankA != cardRankB {
				return cardRankA > cardRankB
			}
		}

		return false
	}

	return ha.comboRank > hb.comboRank
}

func Exec() error {
	lines, err := internal.Read("days/day7/input.txt")
	if err != nil {
		return errors.Join(errors.New("bad read file"), err)
	}

	numRegex, err := internal.NumberRegex()
	if err != nil {
		return errors.Join(errors.New("unable to init regex"), err)
	}

	hands := make([]Hand, 0)
	jokerHands := make([]Hand, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}

		s := strings.Split(line, " ")
		cards := strings.Split(s[0], "")

		bidStr := numRegex.FindString(s[1])
		if bidStr == "" {
			return errors.New("bad input value")
		}
		bid, err := strconv.Atoi(bidStr)
		if err != nil {
			return errors.New("bad input value")
		}

		hand, err := CreateHand(cards, bid)
		if err != nil {
			return err
		}

		jokerHand, err := CreateJokerHand(cards, bid)
		if err != nil {
			return err
		}

		hands = append(hands, hand)
		jokerHands = append(jokerHands, jokerHand)
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].isBetter(hands[j])
	})

	sort.SliceStable(jokerHands, func(i, j int) bool {
		return jokerHands[i].isJokerBetter(jokerHands[j])
	})

	totalWin := 0
	totalJokerWin := 0
	multiplier := len(hands)
	for i := range hands {
		totalWin += hands[i].bid * multiplier
		totalJokerWin += jokerHands[i].bid * multiplier
		multiplier--
	}

	fmt.Println(totalWin)
	fmt.Println(totalJokerWin)

	return nil
}
