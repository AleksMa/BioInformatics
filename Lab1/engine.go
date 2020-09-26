package main

import "fmt"

const GapSymbol = '-'

type NeedlemanWunsch struct {
	FirstSequence  *Sequence
	SecondSequence *Sequence
	Table          Matrix
	SF             ScoringFunc
	GapValue       int
}

func NewNeedlemanWunsch(first, second *Sequence, sf ScoringFunc, GapValue int) *NeedlemanWunsch {
	nw := &NeedlemanWunsch{
		FirstSequence:  first,
		SecondSequence: second,
		Table:          make(Matrix, len(second.Value)+1),
		SF:             sf,
		GapValue:       GapValue,
	}
	// Аллоцируем первую строку
	nw.Table[0] = make(Line, len(first.Value)+1)
	// Обнуляем (0, 0)
	nw.Table[0][0] = &Cell{
		Distance: 0,
		Dir:      NullDirection,
	}
	// Обнуляем первую строку
	for i := range first.Value {
		nw.Table[0][i+1] = &Cell{
			Distance: GapValue * (i + 1),
			Dir:      LeftDirection,
		}
	}
	// Аллоцируем оставшиеся строки, зануляем первый столбец
	for i := range second.Value {
		nw.Table[i+1] = make(Line, len(first.Value)+1)
		nw.Table[i+1][0] = &Cell{
			Distance: GapValue * (i + 1),
			Dir:      TopDirection,
		}
	}
	return nw
}

func (nw *NeedlemanWunsch) Print() {
	for i := 0; i <= len(nw.SecondSequence.Value); i++ {
		for j := 0; j <= len(nw.FirstSequence.Value); j++ {
			fmt.Print(nw.Table[i][j].Distance, ", ", nw.Table[i][j].Dir)
			fmt.Print("   | ")
		}
		fmt.Println()
	}
}

func (nw *NeedlemanWunsch) Solve() (string, string) {
	nw.determine(len(nw.SecondSequence.Value), len(nw.FirstSequence.Value))

	cell := nw.Table[len(nw.SecondSequence.Value)][len(nw.FirstSequence.Value)]

	firstRes, secondRes := "", ""
	fp, sp := len(nw.FirstSequence.Value)-1, len(nw.SecondSequence.Value)-1

	for cell.Dir != NullDirection {
		if cell.Dir == DiagonalDirection {
			firstRes = string(rune(nw.FirstSequence.Value[fp])) + firstRes
			secondRes = string(rune(nw.SecondSequence.Value[sp])) + secondRes
			fp--
			sp--
		} else if cell.Dir == LeftDirection {
			firstRes = string(rune(nw.FirstSequence.Value[fp])) + firstRes
			secondRes = "-" + secondRes
			fp--
		} else if cell.Dir == TopDirection {
			firstRes = "-" + firstRes
			secondRes = string(rune(nw.SecondSequence.Value[sp])) + secondRes
			sp--
		}
		cell = nw.Table[sp+1][fp+1]
	}

	return firstRes, secondRes
}

func (nw *NeedlemanWunsch) determine(i, j int) {
	if nw.Table[i][j] != nil {
		return
	}
	leftCell, topCell, diagCell := nw.Table[i][j-1], nw.Table[i-1][j], nw.Table[i-1][j-1]
	if leftCell == nil {
		nw.determine(i, j-1)
		leftCell = nw.Table[i][j-1]
	}
	if diagCell == nil {
		nw.determine(i-1, j-1)
		diagCell = nw.Table[i-1][j-1]
	}
	if topCell == nil {
		nw.determine(i-1, j)
		topCell = nw.Table[i-1][j]
	}

	maxVal, maxNum := max3(
		nw.Table[i-1][j-1].Distance+nw.SF[nw.SecondSequence.Value[i-1]][nw.FirstSequence.Value[j-1]],
		nw.Table[i][j-1].Distance+nw.GapValue,
		nw.Table[i-1][j].Distance+nw.GapValue,
	)

	nw.Table[i][j] = &Cell{
		Distance: maxVal,
	}
	curCell := nw.Table[i][j]

	switch maxNum {
	case 1:
		curCell.Dir = DiagonalDirection
	case 2:
		curCell.Dir = LeftDirection
	case 3:
		curCell.Dir = TopDirection
	}
}

func max3(a, b, c int) (int, int) {
	if a >= b {
		if a >= c {
			return a, 1
		}
		return c, 3
	}
	if b >= c {
		return b, 2
	}
	return c, 3
}

func max2(a, b int) (int, bool) {
	if a >= b {
		return a, true
	}
	return b, false
}
