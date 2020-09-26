package main

var DNAFull = ScoringFunc{
	'A': map[uint8]int{
		'A': 5,
		'T': -4,
		'G': -4,
		'C': -4,
	},
	'T': map[uint8]int{
		'A': -4,
		'T': 5,
		'G': -4,
		'C': -4,
	},
	'G': map[uint8]int{
		'A': -4,
		'T': -4,
		'G': 5,
		'C': -4,
	},
	'C': map[uint8]int{
		'A': -4,
		'T': -4,
		'G': -4,
		'C': 5,
	},
}

var SimpleFunc = ScoringFunc{
	'A': map[uint8]int{
		'A': 1,
		'T': -1,
		'G': -1,
		'C': -1,
	},
	'T': map[uint8]int{
		'A': -1,
		'T': 1,
		'G': -1,
		'C': -1,
	},
	'G': map[uint8]int{
		'A': -1,
		'T': -1,
		'G': 1,
		'C': -1,
	},
	'C': map[uint8]int{
		'A': -1,
		'T': -1,
		'G': -1,
		'C': 1,
	},
}
