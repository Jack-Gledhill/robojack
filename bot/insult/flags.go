package insult

import "github.com/Jack-Gledhill/robojack/utils"

const (
	FlagBG3 utils.Bitmask = 1 << iota
	FlagDnD
	FlagGeneral
	FlagProgramming
	FlagSCS
	FlagUoS
	DefaultFlags = FlagBG3 | FlagDnD | FlagGeneral | FlagProgramming | FlagSCS | FlagUoS
)

func FlaggedList(b utils.Bitmask) []string {
	var flags []string

	if b.HasFlag(FlagBG3) {
		flags = append(flags, MasterList.BG3...)
	}

	if b.HasFlag(FlagDnD) {
		flags = append(flags, MasterList.DnD...)
	}

	if b.HasFlag(FlagGeneral) {
		flags = append(flags, MasterList.General...)
	}

	if b.HasFlag(FlagProgramming) {
		flags = append(flags, MasterList.Programming...)
	}

	if b.HasFlag(FlagSCS) {
		flags = append(flags, MasterList.SCS...)
	}

	if b.HasFlag(FlagUoS) {
		flags = append(flags, MasterList.UoS...)
	}

	return flags
}

func DefaultList() []string {
	return FlaggedList(DefaultFlags)
}
