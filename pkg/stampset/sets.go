package stampset

import "github.com/wtks/qbot/pkg/stamp"

var Numbers = []string{
	stamp.Zero,
	stamp.One,
	stamp.Two,
	stamp.Three,
	stamp.Four,
	stamp.Five,
	stamp.Six,
	stamp.Seven,
	stamp.Eight,
	stamp.Nine,
}

func NumberToInt(stamp string) int {
	for i, s := range Numbers {
		if s == stamp {
			return i
		}
	}
	return -1
}
