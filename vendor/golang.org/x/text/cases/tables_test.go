// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package cases

var (
	caseIgnorable = map[rune]bool{
		0x0027: true,
		0x002e: true,
		0x003a: true,
		0x00b7: true,
		0x0387: true,
		0x05f4: true,
		0x2018: true,
		0x2019: true,
		0x2024: true,
		0x2027: true,
		0xfe13: true,
		0xfe52: true,
		0xfe55: true,
		0xff07: true,
		0xff0e: true,
		0xff1a: true,
	}

	special = map[rune]struct{ toLower, toTitle, toUpper string }{
		0x00df: {"ß", "Ss", "SS"},
		0x0130: {"i̇", "İ", "İ"},
		0xfb00: {"ﬀ", "Ff", "FF"},
		0xfb01: {"ﬁ", "Fi", "FI"},
		0xfb02: {"ﬂ", "Fl", "FL"},
		0xfb03: {"ﬃ", "Ffi", "FFI"},
		0xfb04: {"ﬄ", "Ffl", "FFL"},
		0xfb05: {"ﬅ", "St", "ST"},
		0xfb06: {"ﬆ", "St", "ST"},
		0x0587: {"և", "Եւ", "ԵՒ"},
		0xfb13: {"ﬓ", "Մն", "ՄՆ"},
		0xfb14: {"ﬔ", "Մե", "ՄԵ"},
		0xfb15: {"ﬕ", "Մի", "ՄԻ"},
		0xfb16: {"ﬖ", "Վն", "ՎՆ"},
		0xfb17: {"ﬗ", "Մխ", "ՄԽ"},
		0x0149: {"ŉ", "ʼN", "ʼN"},
		0x0390: {"ΐ", "Ϊ́", "Ϊ́"},
		0x03b0: {"ΰ", "Ϋ́", "Ϋ́"},
		0x01f0: {"ǰ", "J̌", "J̌"},
		0x1e96: {"ẖ", "H̱", "H̱"},
		0x1e97: {"ẗ", "T̈", "T̈"},
		0x1e98: {"ẘ", "W̊", "W̊"},
		0x1e99: {"ẙ", "Y̊", "Y̊"},
		0x1e9a: {"ẚ", "Aʾ", "Aʾ"},
		0x1f50: {"ὐ", "Υ̓", "Υ̓"},
		0x1f52: {"ὒ", "Υ̓̀", "Υ̓̀"},
		0x1f54: {"ὔ", "Υ̓́", "Υ̓́"},
		0x1f56: {"ὖ", "Υ̓͂", "Υ̓͂"},
		0x1fb6: {"ᾶ", "Α͂", "Α͂"},
		0x1fc6: {"ῆ", "Η͂", "Η͂"},
		0x1fd2: {"ῒ", "Ϊ̀", "Ϊ̀"},
		0x1fd3: {"ΐ", "Ϊ́", "Ϊ́"},
		0x1fd6: {"ῖ", "Ι͂", "Ι͂"},
		0x1fd7: {"ῗ", "Ϊ͂", "Ϊ͂"},
		0x1fe2: {"ῢ", "Ϋ̀", "Ϋ̀"},
		0x1fe3: {"ΰ", "Ϋ́", "Ϋ́"},
		0x1fe4: {"ῤ", "Ρ̓", "Ρ̓"},
		0x1fe6: {"ῦ", "Υ͂", "Υ͂"},
		0x1fe7: {"ῧ", "Ϋ͂", "Ϋ͂"},
		0x1ff6: {"ῶ", "Ω͂", "Ω͂"},
		0x1f80: {"ᾀ", "ᾈ", "ἈΙ"},
		0x1f81: {"ᾁ", "ᾉ", "ἉΙ"},
		0x1f82: {"ᾂ", "ᾊ", "ἊΙ"},
		0x1f83: {"ᾃ", "ᾋ", "ἋΙ"},
		0x1f84: {"ᾄ", "ᾌ", "ἌΙ"},
		0x1f85: {"ᾅ", "ᾍ", "ἍΙ"},
		0x1f86: {"ᾆ", "ᾎ", "ἎΙ"},
		0x1f87: {"ᾇ", "ᾏ", "ἏΙ"},
		0x1f88: {"ᾀ", "ᾈ", "ἈΙ"},
		0x1f89: {"ᾁ", "ᾉ", "ἉΙ"},
		0x1f8a: {"ᾂ", "ᾊ", "ἊΙ"},
		0x1f8b: {"ᾃ", "ᾋ", "ἋΙ"},
		0x1f8c: {"ᾄ", "ᾌ", "ἌΙ"},
		0x1f8d: {"ᾅ", "ᾍ", "ἍΙ"},
		0x1f8e: {"ᾆ", "ᾎ", "ἎΙ"},
		0x1f8f: {"ᾇ", "ᾏ", "ἏΙ"},
		0x1f90: {"ᾐ", "ᾘ", "ἨΙ"},
		0x1f91: {"ᾑ", "ᾙ", "ἩΙ"},
		0x1f92: {"ᾒ", "ᾚ", "ἪΙ"},
		0x1f93: {"ᾓ", "ᾛ", "ἫΙ"},
		0x1f94: {"ᾔ", "ᾜ", "ἬΙ"},
		0x1f95: {"ᾕ", "ᾝ", "ἭΙ"},
		0x1f96: {"ᾖ", "ᾞ", "ἮΙ"},
		0x1f97: {"ᾗ", "ᾟ", "ἯΙ"},
		0x1f98: {"ᾐ", "ᾘ", "ἨΙ"},
		0x1f99: {"ᾑ", "ᾙ", "ἩΙ"},
		0x1f9a: {"ᾒ", "ᾚ", "ἪΙ"},
		0x1f9b: {"ᾓ", "ᾛ", "ἫΙ"},
		0x1f9c: {"ᾔ", "ᾜ", "ἬΙ"},
		0x1f9d: {"ᾕ", "ᾝ", "ἭΙ"},
		0x1f9e: {"ᾖ", "ᾞ", "ἮΙ"},
		0x1f9f: {"ᾗ", "ᾟ", "ἯΙ"},
		0x1fa0: {"ᾠ", "ᾨ", "ὨΙ"},
		0x1fa1: {"ᾡ", "ᾩ", "ὩΙ"},
		0x1fa2: {"ᾢ", "ᾪ", "ὪΙ"},
		0x1fa3: {"ᾣ", "ᾫ", "ὫΙ"},
		0x1fa4: {"ᾤ", "ᾬ", "ὬΙ"},
		0x1fa5: {"ᾥ", "ᾭ", "ὭΙ"},
		0x1fa6: {"ᾦ", "ᾮ", "ὮΙ"},
		0x1fa7: {"ᾧ", "ᾯ", "ὯΙ"},
		0x1fa8: {"ᾠ", "ᾨ", "ὨΙ"},
		0x1fa9: {"ᾡ", "ᾩ", "ὩΙ"},
		0x1faa: {"ᾢ", "ᾪ", "ὪΙ"},
		0x1fab: {"ᾣ", "ᾫ", "ὫΙ"},
		0x1fac: {"ᾤ", "ᾬ", "ὬΙ"},
		0x1fad: {"ᾥ", "ᾭ", "ὭΙ"},
		0x1fae: {"ᾦ", "ᾮ", "ὮΙ"},
		0x1faf: {"ᾧ", "ᾯ", "ὯΙ"},
		0x1fb3: {"ᾳ", "ᾼ", "ΑΙ"},
		0x1fbc: {"ᾳ", "ᾼ", "ΑΙ"},
		0x1fc3: {"ῃ", "ῌ", "ΗΙ"},
		0x1fcc: {"ῃ", "ῌ", "ΗΙ"},
		0x1ff3: {"ῳ", "ῼ", "ΩΙ"},
		0x1ffc: {"ῳ", "ῼ", "ΩΙ"},
		0x1fb2: {"ᾲ", "Ὰͅ", "ᾺΙ"},
		0x1fb4: {"ᾴ", "Άͅ", "ΆΙ"},
		0x1fc2: {"ῂ", "Ὴͅ", "ῊΙ"},
		0x1fc4: {"ῄ", "Ήͅ", "ΉΙ"},
		0x1ff2: {"ῲ", "Ὼͅ", "ῺΙ"},
		0x1ff4: {"ῴ", "Ώͅ", "ΏΙ"},
		0x1fb7: {"ᾷ", "ᾼ͂", "Α͂Ι"},
		0x1fc7: {"ῇ", "ῌ͂", "Η͂Ι"},
		0x1ff7: {"ῷ", "ῼ͂", "Ω͂Ι"},
	}

	foldMap = map[rune]struct{ simple, full, special string }{
		0x0049: {"", "", "ı"},
		0x00b5: {"μ", "μ", ""},
		0x00df: {"", "ss", ""},
		0x0130: {"", "i̇", "i"},
		0x0149: {"", "ʼn", ""},
		0x017f: {"s", "s", ""},
		0x01f0: {"", "ǰ", ""},
		0x0345: {"ι", "ι", ""},
		0x0390: {"", "ΐ", ""},
		0x03b0: {"", "ΰ", ""},
		0x03c2: {"σ", "σ", ""},
		0x03d0: {"β", "β", ""},
		0x03d1: {"θ", "θ", ""},
		0x03d5: {"φ", "φ", ""},
		0x03d6: {"π", "π", ""},
		0x03f0: {"κ", "κ", ""},
		0x03f1: {"ρ", "ρ", ""},
		0x03f5: {"ε", "ε", ""},
		0x0587: {"", "եւ", ""},
		0x13f8: {"Ᏸ", "Ᏸ", ""},
		0x13f9: {"Ᏹ", "Ᏹ", ""},
		0x13fa: {"Ᏺ", "Ᏺ", ""},
		0x13fb: {"Ᏻ", "Ᏻ", ""},
		0x13fc: {"Ᏼ", "Ᏼ", ""},
		0x13fd: {"Ᏽ", "Ᏽ", ""},
		0x1c80: {"в", "в", ""},
		0x1c81: {"д", "д", ""},
		0x1c82: {"о", "о", ""},
		0x1c83: {"с", "с", ""},
		0x1c84: {"т", "т", ""},
		0x1c85: {"т", "т", ""},
		0x1c86: {"ъ", "ъ", ""},
		0x1c87: {"ѣ", "ѣ", ""},
		0x1c88: {"ꙋ", "ꙋ", ""},
		0x1e96: {"", "ẖ", ""},
		0x1e97: {"", "ẗ", ""},
		0x1e98: {"", "ẘ", ""},
		0x1e99: {"", "ẙ", ""},
		0x1e9a: {"", "aʾ", ""},
		0x1e9b: {"ṡ", "ṡ", ""},
		0x1e9e: {"", "ss", ""},
		0x1f50: {"", "ὐ", ""},
		0x1f52: {"", "ὒ", ""},
		0x1f54: {"", "ὔ", ""},
		0x1f56: {"", "ὖ", ""},
		0x1f80: {"", "ἀι", ""},
		0x1f81: {"", "ἁι", ""},
		0x1f82: {"", "ἂι", ""},
		0x1f83: {"", "ἃι", ""},
		0x1f84: {"", "ἄι", ""},
		0x1f85: {"", "ἅι", ""},
		0x1f86: {"", "ἆι", ""},
		0x1f87: {"", "ἇι", ""},
		0x1f88: {"", "ἀι", ""},
		0x1f89: {"", "ἁι", ""},
		0x1f8a: {"", "ἂι", ""},
		0x1f8b: {"", "ἃι", ""},
		0x1f8c: {"", "ἄι", ""},
		0x1f8d: {"", "ἅι", ""},
		0x1f8e: {"", "ἆι", ""},
		0x1f8f: {"", "ἇι", ""},
		0x1f90: {"", "ἠι", ""},
		0x1f91: {"", "ἡι", ""},
		0x1f92: {"", "ἢι", ""},
		0x1f93: {"", "ἣι", ""},
		0x1f94: {"", "ἤι", ""},
		0x1f95: {"", "ἥι", ""},
		0x1f96: {"", "ἦι", ""},
		0x1f97: {"", "ἧι", ""},
		0x1f98: {"", "ἠι", ""},
		0x1f99: {"", "ἡι", ""},
		0x1f9a: {"", "ἢι", ""},
		0x1f9b: {"", "ἣι", ""},
		0x1f9c: {"", "ἤι", ""},
		0x1f9d: {"", "ἥι", ""},
		0x1f9e: {"", "ἦι", ""},
		0x1f9f: {"", "ἧι", ""},
		0x1fa0: {"", "ὠι", ""},
		0x1fa1: {"", "ὡι", ""},
		0x1fa2: {"", "ὢι", ""},
		0x1fa3: {"", "ὣι", ""},
		0x1fa4: {"", "ὤι", ""},
		0x1fa5: {"", "ὥι", ""},
		0x1fa6: {"", "ὦι", ""},
		0x1fa7: {"", "ὧι", ""},
		0x1fa8: {"", "ὠι", ""},
		0x1fa9: {"", "ὡι", ""},
		0x1faa: {"", "ὢι", ""},
		0x1fab: {"", "ὣι", ""},
		0x1fac: {"", "ὤι", ""},
		0x1fad: {"", "ὥι", ""},
		0x1fae: {"", "ὦι", ""},
		0x1faf: {"", "ὧι", ""},
		0x1fb2: {"", "ὰι", ""},
		0x1fb3: {"", "αι", ""},
		0x1fb4: {"", "άι", ""},
		0x1fb6: {"", "ᾶ", ""},
		0x1fb7: {"", "ᾶι", ""},
		0x1fbc: {"", "αι", ""},
		0x1fbe: {"ι", "ι", ""},
		0x1fc2: {"", "ὴι", ""},
		0x1fc3: {"", "ηι", ""},
		0x1fc4: {"", "ήι", ""},
		0x1fc6: {"", "ῆ", ""},
		0x1fc7: {"", "ῆι", ""},
		0x1fcc: {"", "ηι", ""},
		0x1fd2: {"", "ῒ", ""},
		0x1fd3: {"", "ΐ", ""},
		0x1fd6: {"", "ῖ", ""},
		0x1fd7: {"", "ῗ", ""},
		0x1fe2: {"", "ῢ", ""},
		0x1fe3: {"", "ΰ", ""},
		0x1fe4: {"", "ῤ", ""},
		0x1fe6: {"", "ῦ", ""},
		0x1fe7: {"", "ῧ", ""},
		0x1ff2: {"", "ὼι", ""},
		0x1ff3: {"", "ωι", ""},
		0x1ff4: {"", "ώι", ""},
		0x1ff6: {"", "ῶ", ""},
		0x1ff7: {"", "ῶι", ""},
		0x1ffc: {"", "ωι", ""},
		0xab70: {"Ꭰ", "Ꭰ", ""},
		0xab71: {"Ꭱ", "Ꭱ", ""},
		0xab72: {"Ꭲ", "Ꭲ", ""},
		0xab73: {"Ꭳ", "Ꭳ", ""},
		0xab74: {"Ꭴ", "Ꭴ", ""},
		0xab75: {"Ꭵ", "Ꭵ", ""},
		0xab76: {"Ꭶ", "Ꭶ", ""},
		0xab77: {"Ꭷ", "Ꭷ", ""},
		0xab78: {"Ꭸ", "Ꭸ", ""},
		0xab79: {"Ꭹ", "Ꭹ", ""},
		0xab7a: {"Ꭺ", "Ꭺ", ""},
		0xab7b: {"Ꭻ", "Ꭻ", ""},
		0xab7c: {"Ꭼ", "Ꭼ", ""},
		0xab7d: {"Ꭽ", "Ꭽ", ""},
		0xab7e: {"Ꭾ", "Ꭾ", ""},
		0xab7f: {"Ꭿ", "Ꭿ", ""},
		0xab80: {"Ꮀ", "Ꮀ", ""},
		0xab81: {"Ꮁ", "Ꮁ", ""},
		0xab82: {"Ꮂ", "Ꮂ", ""},
		0xab83: {"Ꮃ", "Ꮃ", ""},
		0xab84: {"Ꮄ", "Ꮄ", ""},
		0xab85: {"Ꮅ", "Ꮅ", ""},
		0xab86: {"Ꮆ", "Ꮆ", ""},
		0xab87: {"Ꮇ", "Ꮇ", ""},
		0xab88: {"Ꮈ", "Ꮈ", ""},
		0xab89: {"Ꮉ", "Ꮉ", ""},
		0xab8a: {"Ꮊ", "Ꮊ", ""},
		0xab8b: {"Ꮋ", "Ꮋ", ""},
		0xab8c: {"Ꮌ", "Ꮌ", ""},
		0xab8d: {"Ꮍ", "Ꮍ", ""},
		0xab8e: {"Ꮎ", "Ꮎ", ""},
		0xab8f: {"Ꮏ", "Ꮏ", ""},
		0xab90: {"Ꮐ", "Ꮐ", ""},
		0xab91: {"Ꮑ", "Ꮑ", ""},
		0xab92: {"Ꮒ", "Ꮒ", ""},
		0xab93: {"Ꮓ", "Ꮓ", ""},
		0xab94: {"Ꮔ", "Ꮔ", ""},
		0xab95: {"Ꮕ", "Ꮕ", ""},
		0xab96: {"Ꮖ", "Ꮖ", ""},
		0xab97: {"Ꮗ", "Ꮗ", ""},
		0xab98: {"Ꮘ", "Ꮘ", ""},
		0xab99: {"Ꮙ", "Ꮙ", ""},
		0xab9a: {"Ꮚ", "Ꮚ", ""},
		0xab9b: {"Ꮛ", "Ꮛ", ""},
		0xab9c: {"Ꮜ", "Ꮜ", ""},
		0xab9d: {"Ꮝ", "Ꮝ", ""},
		0xab9e: {"Ꮞ", "Ꮞ", ""},
		0xab9f: {"Ꮟ", "Ꮟ", ""},
		0xaba0: {"Ꮠ", "Ꮠ", ""},
		0xaba1: {"Ꮡ", "Ꮡ", ""},
		0xaba2: {"Ꮢ", "Ꮢ", ""},
		0xaba3: {"Ꮣ", "Ꮣ", ""},
		0xaba4: {"Ꮤ", "Ꮤ", ""},
		0xaba5: {"Ꮥ", "Ꮥ", ""},
		0xaba6: {"Ꮦ", "Ꮦ", ""},
		0xaba7: {"Ꮧ", "Ꮧ", ""},
		0xaba8: {"Ꮨ", "Ꮨ", ""},
		0xaba9: {"Ꮩ", "Ꮩ", ""},
		0xabaa: {"Ꮪ", "Ꮪ", ""},
		0xabab: {"Ꮫ", "Ꮫ", ""},
		0xabac: {"Ꮬ", "Ꮬ", ""},
		0xabad: {"Ꮭ", "Ꮭ", ""},
		0xabae: {"Ꮮ", "Ꮮ", ""},
		0xabaf: {"Ꮯ", "Ꮯ", ""},
		0xabb0: {"Ꮰ", "Ꮰ", ""},
		0xabb1: {"Ꮱ", "Ꮱ", ""},
		0xabb2: {"Ꮲ", "Ꮲ", ""},
		0xabb3: {"Ꮳ", "Ꮳ", ""},
		0xabb4: {"Ꮴ", "Ꮴ", ""},
		0xabb5: {"Ꮵ", "Ꮵ", ""},
		0xabb6: {"Ꮶ", "Ꮶ", ""},
		0xabb7: {"Ꮷ", "Ꮷ", ""},
		0xabb8: {"Ꮸ", "Ꮸ", ""},
		0xabb9: {"Ꮹ", "Ꮹ", ""},
		0xabba: {"Ꮺ", "Ꮺ", ""},
		0xabbb: {"Ꮻ", "Ꮻ", ""},
		0xabbc: {"Ꮼ", "Ꮼ", ""},
		0xabbd: {"Ꮽ", "Ꮽ", ""},
		0xabbe: {"Ꮾ", "Ꮾ", ""},
		0xabbf: {"Ꮿ", "Ꮿ", ""},
		0xfb00: {"", "ff", ""},
		0xfb01: {"", "fi", ""},
		0xfb02: {"", "fl", ""},
		0xfb03: {"", "ffi", ""},
		0xfb04: {"", "ffl", ""},
		0xfb05: {"", "st", ""},
		0xfb06: {"", "st", ""},
		0xfb13: {"", "մն", ""},
		0xfb14: {"", "մե", ""},
		0xfb15: {"", "մի", ""},
		0xfb16: {"", "վն", ""},
		0xfb17: {"", "մխ", ""},
	}

	breakProp = []struct{ lo, hi rune }{
		{0x0, 0x26},
		{0x28, 0x2d},
		{0x2f, 0x2f},
		{0x3b, 0x40},
		{0x5b, 0x5e},
		{0x60, 0x60},
		{0x7b, 0xa9},
		{0xab, 0xac},
		{0xae, 0xb4},
		{0xb6, 0xb6},
		{0xb8, 0xb9},
		{0xbb, 0xbf},
		{0xd7, 0xd7},
		{0xf7, 0xf7},
		{0x2c2, 0x2c5},
		{0x2d2, 0x2d6},
		{0x2d8, 0x2df},
		{0x2e5, 0x2eb},
		{0x2ed, 0x2ed},
		{0x2ef, 0x2ff},
		{0x375, 0x375},
		{0x378, 0x379},
		{0x37e, 0x37e},
		{0x380, 0x385},
		{0x38b, 0x38b},
		{0x38d, 0x38d},
		{0x3a2, 0x3a2},
		{0x3f6, 0x3f6},
		{0x482, 0x482},
		{0x530, 0x530},
		{0x557, 0x558},
		{0x55a, 0x560},
		{0x588, 0x590},
		{0x5be, 0x5be},
		{0x5c0, 0x5c0},
		{0x5c3, 0x5c3},
		{0x5c6, 0x5c6},
		{0x5c8, 0x5cf},
		{0x5eb, 0x5ef},
		{0x5f5, 0x5ff},
		{0x606, 0x60f},
		{0x61b, 0x61b},
		{0x61d, 0x61f},
		{0x66a, 0x66a},
		{0x66c, 0x66d},
		{0x6d4, 0x6d4},
		{0x6de, 0x6de},
		{0x6e9, 0x6e9},
		{0x6fd, 0x6fe},
		{0x700, 0x70e},
		{0x74b, 0x74c},
		{0x7b2, 0x7bf},
		{0x7f6, 0x7f9},
		{0x7fb, 0x7ff},
		{0x82e, 0x83f},
		{0x85c, 0x89f},
		{0x8b5, 0x8b5},
		{0x8be, 0x8d3},
		{0x964, 0x965},
		{0x970, 0x970},
		{0x984, 0x984},
		{0x98d, 0x98e},
		{0x991, 0x992},
		{0x9a9, 0x9a9},
		{0x9b1, 0x9b1},
		{0x9b3, 0x9b5},
		{0x9ba, 0x9bb},
		{0x9c5, 0x9c6},
		{0x9c9, 0x9ca},
		{0x9cf, 0x9d6},
		{0x9d8, 0x9db},
		{0x9de, 0x9de},
		{0x9e4, 0x9e5},
		{0x9f2, 0xa00},
		{0xa04, 0xa04},
		{0xa0b, 0xa0e},
		{0xa11, 0xa12},
		{0xa29, 0xa29},
		{0xa31, 0xa31},
		{0xa34, 0xa34},
		{0xa37, 0xa37},
		{0xa3a, 0xa3b},
		{0xa3d, 0xa3d},
		{0xa43, 0xa46},
		{0xa49, 0xa4a},
		{0xa4e, 0xa50},
		{0xa52, 0xa58},
		{0xa5d, 0xa5d},
		{0xa5f, 0xa65},
		{0xa76, 0xa80},
		{0xa84, 0xa84},
		{0xa8e, 0xa8e},
		{0xa92, 0xa92},
		{0xaa9, 0xaa9},
		{0xab1, 0xab1},
		{0xab4, 0xab4},
		{0xaba, 0xabb},
		{0xac6, 0xac6},
		{0xaca, 0xaca},
		{0xace, 0xacf},
		{0xad1, 0xadf},
		{0xae4, 0xae5},
		{0xaf0, 0xaf8},
		{0xafa, 0xb00},
		{0xb04, 0xb04},
		{0xb0d, 0xb0e},
		{0xb11, 0xb12},
		{0xb29, 0xb29},
		{0xb31, 0xb31},
		{0xb34, 0xb34},
		{0xb3a, 0xb3b},
		{0xb45, 0xb46},
		{0xb49, 0xb4a},
		{0xb4e, 0xb55},
		{0xb58, 0xb5b},
		{0xb5e, 0xb5e},
		{0xb64, 0xb65},
		{0xb70, 0xb70},
		{0xb72, 0xb81},
		{0xb84, 0xb84},
		{0xb8b, 0xb8d},
		{0xb91, 0xb91},
		{0xb96, 0xb98},
		{0xb9b, 0xb9b},
		{0xb9d, 0xb9d},
		{0xba0, 0xba2},
		{0xba5, 0xba7},
		{0xbab, 0xbad},
		{0xbba, 0xbbd},
		{0xbc3, 0xbc5},
		{0xbc9, 0xbc9},
		{0xbce, 0xbcf},
		{0xbd1, 0xbd6},
		{0xbd8, 0xbe5},
		{0xbf0, 0xbff},
		{0xc04, 0xc04},
		{0xc0d, 0xc0d},
		{0xc11, 0xc11},
		{0xc29, 0xc29},
		{0xc3a, 0xc3c},
		{0xc45, 0xc45},
		{0xc49, 0xc49},
		{0xc4e, 0xc54},
		{0xc57, 0xc57},
		{0xc5b, 0xc5f},
		{0xc64, 0xc65},
		{0xc70, 0xc7f},
		{0xc84, 0xc84},
		{0xc8d, 0xc8d},
		{0xc91, 0xc91},
		{0xca9, 0xca9},
		{0xcb4, 0xcb4},
		{0xcba, 0xcbb},
		{0xcc5, 0xcc5},
		{0xcc9, 0xcc9},
		{0xcce, 0xcd4},
		{0xcd7, 0xcdd},
		{0xcdf, 0xcdf},
		{0xce4, 0xce5},
		{0xcf0, 0xcf0},
		{0xcf3, 0xd00},
		{0xd04, 0xd04},
		{0xd0d, 0xd0d},
		{0xd11, 0xd11},
		{0xd3b, 0xd3c},
		{0xd45, 0xd45},
		{0xd49, 0xd49},
		{0xd4f, 0xd53},
		{0xd58, 0xd5e},
		{0xd64, 0xd65},
		{0xd70, 0xd79},
		{0xd80, 0xd81},
		{0xd84, 0xd84},
		{0xd97, 0xd99},
		{0xdb2, 0xdb2},
		{0xdbc, 0xdbc},
		{0xdbe, 0xdbf},
		{0xdc7, 0xdc9},
		{0xdcb, 0xdce},
		{0xdd5, 0xdd5},
		{0xdd7, 0xdd7},
		{0xde0, 0xde5},
		{0xdf0, 0xdf1},
		{0xdf4, 0xe30},
		{0xe32, 0xe33},
		{0xe3b, 0xe46},
		{0xe4f, 0xe4f},
		{0xe5a, 0xeb0},
		{0xeb2, 0xeb3},
		{0xeba, 0xeba},
		{0xebd, 0xec7},
		{0xece, 0xecf},
		{0xeda, 0xeff},
		{0xf01, 0xf17},
		{0xf1a, 0xf1f},
		{0xf2a, 0xf34},
		{0xf36, 0xf36},
		{0xf38, 0xf38},
		{0xf3a, 0xf3d},
		{0xf48, 0xf48},
		{0xf6d, 0xf70},
		{0xf85, 0xf85},
		{0xf98, 0xf98},
		{0xfbd, 0xfc5},
		{0xfc7, 0x102a},
		{0x103f, 0x103f},
		{0x104a, 0x1055},
		{0x105a, 0x105d},
		{0x1061, 0x1061},
		{0x1065, 0x1066},
		{0x106e, 0x1070},
		{0x1075, 0x1081},
		{0x108e, 0x108e},
		{0x109e, 0x109f},
		{0x10c6, 0x10c6},
		{0x10c8, 0x10cc},
		{0x10ce, 0x10cf},
		{0x10fb, 0x10fb},
		{0x1249, 0x1249},
		{0x124e, 0x124f},
		{0x1257, 0x1257},
		{0x1259, 0x1259},
		{0x125e, 0x125f},
		{0x1289, 0x1289},
		{0x128e, 0x128f},
		{0x12b1, 0x12b1},
		{0x12b6, 0x12b7},
		{0x12bf, 0x12bf},
		{0x12c1, 0x12c1},
		{0x12c6, 0x12c7},
		{0x12d7, 0x12d7},
		{0x1311, 0x1311},
		{0x1316, 0x1317},
		{0x135b, 0x135c},
		{0x1360, 0x137f},
		{0x1390, 0x139f},
		{0x13f6, 0x13f7},
		{0x13fe, 0x1400},
		{0x166d, 0x166e},
		{0x1680, 0x1680},
		{0x169b, 0x169f},
		{0x16eb, 0x16ed},
		{0x16f9, 0x16ff},
		{0x170d, 0x170d},
		{0x1715, 0x171f},
		{0x1735, 0x173f},
		{0x1754, 0x175f},
		{0x176d, 0x176d},
		{0x1771, 0x1771},
		{0x1774, 0x17b3},
		{0x17d4, 0x17dc},
		{0x17de, 0x17df},
		{0x17ea, 0x180a},
		{0x180f, 0x180f},
		{0x181a, 0x181f},
		{0x1878, 0x187f},
		{0x18ab, 0x18af},
		{0x18f6, 0x18ff},
		{0x191f, 0x191f},
		{0x192c, 0x192f},
		{0x193c, 0x1945},
		{0x1950, 0x19cf},
		{0x19da, 0x19ff},
		{0x1a1c, 0x1a54},
		{0x1a5f, 0x1a5f},
		{0x1a7d, 0x1a7e},
		{0x1a8a, 0x1a8f},
		{0x1a9a, 0x1aaf},
		{0x1abf, 0x1aff},
		{0x1b4c, 0x1b4f},
		{0x1b5a, 0x1b6a},
		{0x1b74, 0x1b7f},
		{0x1bf4, 0x1bff},
		{0x1c38, 0x1c3f},
		{0x1c4a, 0x1c4c},
		{0x1c7e, 0x1c7f},
		{0x1c89, 0x1ccf},
		{0x1cd3, 0x1cd3},
		{0x1cf7, 0x1cf7},
		{0x1cfa, 0x1cff},
		{0x1df6, 0x1dfa},
		{0x1f16, 0x1f17},
		{0x1f1e, 0x1f1f},
		{0x1f46, 0x1f47},
		{0x1f4e, 0x1f4f},
		{0x1f58, 0x1f58},
		{0x1f5a, 0x1f5a},
		{0x1f5c, 0x1f5c},
		{0x1f5e, 0x1f5e},
		{0x1f7e, 0x1f7f},
		{0x1fb5, 0x1fb5},
		{0x1fbd, 0x1fbd},
		{0x1fbf, 0x1fc1},
		{0x1fc5, 0x1fc5},
		{0x1fcd, 0x1fcf},
		{0x1fd4, 0x1fd5},
		{0x1fdc, 0x1fdf},
		{0x1fed, 0x1ff1},
		{0x1ff5, 0x1ff5},
		{0x1ffd, 0x200b},
		{0x2010, 0x2017},
		{0x201a, 0x2023},
		{0x2025, 0x2026},
		{0x2028, 0x2029},
		{0x2030, 0x203e},
		{0x2041, 0x2053},
		{0x2055, 0x205f},
		{0x2065, 0x2065},
		{0x2070, 0x2070},
		{0x2072, 0x207e},
		{0x2080, 0x208f},
		{0x209d, 0x20cf},
		{0x20f1, 0x2101},
		{0x2103, 0x2106},
		{0x2108, 0x2109},
		{0x2114, 0x2114},
		{0x2116, 0x2118},
		{0x211e, 0x2123},
		{0x2125, 0x2125},
		{0x2127, 0x2127},
		{0x2129, 0x2129},
		{0x212e, 0x212e},
		{0x213a, 0x213b},
		{0x2140, 0x2144},
		{0x214a, 0x214d},
		{0x214f, 0x215f},
		{0x2189, 0x24b5},
		{0x24ea, 0x2bff},
		{0x2c2f, 0x2c2f},
		{0x2c5f, 0x2c5f},
		{0x2ce5, 0x2cea},
		{0x2cf4, 0x2cff},
		{0x2d26, 0x2d26},
		{0x2d28, 0x2d2c},
		{0x2d2e, 0x2d2f},
		{0x2d68, 0x2d6e},
		{0x2d70, 0x2d7e},
		{0x2d97, 0x2d9f},
		{0x2da7, 0x2da7},
		{0x2daf, 0x2daf},
		{0x2db7, 0x2db7},
		{0x2dbf, 0x2dbf},
		{0x2dc7, 0x2dc7},
		{0x2dcf, 0x2dcf},
		{0x2dd7, 0x2dd7},
		{0x2ddf, 0x2ddf},
		{0x2e00, 0x2e2e},
		{0x2e30, 0x3004},
		{0x3006, 0x3029},
		{0x3030, 0x303a},
		{0x303d, 0x3098},
		{0x309b, 0x3104},
		{0x312e, 0x3130},
		{0x318f, 0x319f},
		{0x31bb, 0x9fff},
		{0xa48d, 0xa4cf},
		{0xa4fe, 0xa4ff},
		{0xa60d, 0xa60f},
		{0xa62c, 0xa63f},
		{0xa673, 0xa673},
		{0xa67e, 0xa67e},
		{0xa6f2, 0xa716},
		{0xa720, 0xa721},
		{0xa789, 0xa78a},
		{0xa7af, 0xa7af},
		{0xa7b8, 0xa7f6},
		{0xa828, 0xa83f},
		{0xa874, 0xa87f},
		{0xa8c6, 0xa8cf},
		{0xa8da, 0xa8df},
		{0xa8f8, 0xa8fa},
		{0xa8fc, 0xa8fc},
		{0xa8fe, 0xa8ff},
		{0xa92e, 0xa92f},
		{0xa954, 0xa95f},
		{0xa97d, 0xa97f},
		{0xa9c1, 0xa9ce},
		{0xa9da, 0xa9e4},
		{0xa9e6, 0xa9ef},
		{0xa9fa, 0xa9ff},
		{0xaa37, 0xaa3f},
		{0xaa4e, 0xaa4f},
		{0xaa5a, 0xaa7a},
		{0xaa7e, 0xaaaf},
		{0xaab1, 0xaab1},
		{0xaab5, 0xaab6},
		{0xaab9, 0xaabd},
		{0xaac0, 0xaac0},
		{0xaac2, 0xaadf},
		{0xaaf0, 0xaaf1},
		{0xaaf7, 0xab00},
		{0xab07, 0xab08},
		{0xab0f, 0xab10},
		{0xab17, 0xab1f},
		{0xab27, 0xab27},
		{0xab2f, 0xab2f},
		{0xab5b, 0xab5b},
		{0xab66, 0xab6f},
		{0xabeb, 0xabeb},
		{0xabee, 0xabef},
		{0xabfa, 0xabff},
		{0xd7a4, 0xd7af},
		{0xd7c7, 0xd7ca},
		{0xd7fc, 0xfaff},
		{0xfb07, 0xfb12},
		{0xfb18, 0xfb1c},
		{0xfb29, 0xfb29},
		{0xfb37, 0xfb37},
		{0xfb3d, 0xfb3d},
		{0xfb3f, 0xfb3f},
		{0xfb42, 0xfb42},
		{0xfb45, 0xfb45},
		{0xfbb2, 0xfbd2},
		{0xfd3e, 0xfd4f},
		{0xfd90, 0xfd91},
		{0xfdc8, 0xfdef},
		{0xfdfc, 0xfdff},
		{0xfe10, 0xfe12},
		{0xfe14, 0xfe1f},
		{0xfe30, 0xfe32},
		{0xfe35, 0xfe4c},
		{0xfe50, 0xfe51},
		{0xfe53, 0xfe54},
		{0xfe56, 0xfe6f},
		{0xfe75, 0xfe75},
		{0xfefd, 0xfefe},
		{0xff00, 0xff06},
		{0xff08, 0xff0d},
		{0xff0f, 0xff19},
		{0xff1b, 0xff20},
		{0xff3b, 0xff3e},
		{0xff40, 0xff40},
		{0xff5b, 0xff9d},
		{0xffbf, 0xffc1},
		{0xffc8, 0xffc9},
		{0xffd0, 0xffd1},
		{0xffd8, 0xffd9},
		{0xffdd, 0xfff8},
		{0xfffc, 0xffff},
		{0x1000c, 0x1000c},
		{0x10027, 0x10027},
		{0x1003b, 0x1003b},
		{0x1003e, 0x1003e},
		{0x1004e, 0x1004f},
		{0x1005e, 0x1007f},
		{0x100fb, 0x1013f},
		{0x10175, 0x101fc},
		{0x101fe, 0x1027f},
		{0x1029d, 0x1029f},
		{0x102d1, 0x102df},
		{0x102e1, 0x102ff},
		{0x10320, 0x1032f},
		{0x1034b, 0x1034f},
		{0x1037b, 0x1037f},
		{0x1039e, 0x1039f},
		{0x103c4, 0x103c7},
		{0x103d0, 0x103d0},
		{0x103d6, 0x103ff},
		{0x1049e, 0x1049f},
		{0x104aa, 0x104af},
		{0x104d4, 0x104d7},
		{0x104fc, 0x104ff},
		{0x10528, 0x1052f},
		{0x10564, 0x105ff},
		{0x10737, 0x1073f},
		{0x10756, 0x1075f},
		{0x10768, 0x107ff},
		{0x10806, 0x10807},
		{0x10809, 0x10809},
		{0x10836, 0x10836},
		{0x10839, 0x1083b},
		{0x1083d, 0x1083e},
		{0x10856, 0x1085f},
		{0x10877, 0x1087f},
		{0x1089f, 0x108df},
		{0x108f3, 0x108f3},
		{0x108f6, 0x108ff},
		{0x10916, 0x1091f},
		{0x1093a, 0x1097f},
		{0x109b8, 0x109bd},
		{0x109c0, 0x109ff},
		{0x10a04, 0x10a04},
		{0x10a07, 0x10a0b},
		{0x10a14, 0x10a14},
		{0x10a18, 0x10a18},
		{0x10a34, 0x10a37},
		{0x10a3b, 0x10a3e},
		{0x10a40, 0x10a5f},
		{0x10a7d, 0x10a7f},
		{0x10a9d, 0x10abf},
		{0x10ac8, 0x10ac8},
		{0x10ae7, 0x10aff},
		{0x10b36, 0x10b3f},
		{0x10b56, 0x10b5f},
		{0x10b73, 0x10b7f},
		{0x10b92, 0x10bff},
		{0x10c49, 0x10c7f},
		{0x10cb3, 0x10cbf},
		{0x10cf3, 0x10fff},
		{0x11047, 0x11065},
		{0x11070, 0x1107e},
		{0x110bb, 0x110bc},
		{0x110be, 0x110cf},
		{0x110e9, 0x110ef},
		{0x110fa, 0x110ff},
		{0x11135, 0x11135},
		{0x11140, 0x1114f},
		{0x11174, 0x11175},
		{0x11177, 0x1117f},
		{0x111c5, 0x111c9},
		{0x111cd, 0x111cf},
		{0x111db, 0x111db},
		{0x111dd, 0x111ff},
		{0x11212, 0x11212},
		{0x11238, 0x1123d},
		{0x1123f, 0x1127f},
		{0x11287, 0x11287},
		{0x11289, 0x11289},
		{0x1128e, 0x1128e},
		{0x1129e, 0x1129e},
		{0x112a9, 0x112af},
		{0x112eb, 0x112ef},
		{0x112fa, 0x112ff},
		{0x11304, 0x11304},
		{0x1130d, 0x1130e},
		{0x11311, 0x11312},
		{0x11329, 0x11329},
		{0x11331, 0x11331},
		{0x11334, 0x11334},
		{0x1133a, 0x1133b},
		{0x11345, 0x11346},
		{0x11349, 0x1134a},
		{0x1134e, 0x1134f},
		{0x11351, 0x11356},
		{0x11358, 0x1135c},
		{0x11364, 0x11365},
		{0x1136d, 0x1136f},
		{0x11375, 0x113ff},
		{0x1144b, 0x1144f},
		{0x1145a, 0x1147f},
		{0x114c6, 0x114c6},
		{0x114c8, 0x114cf},
		{0x114da, 0x1157f},
		{0x115b6, 0x115b7},
		{0x115c1, 0x115d7},
		{0x115de, 0x115ff},
		{0x11641, 0x11643},
		{0x11645, 0x1164f},
		{0x1165a, 0x1167f},
		{0x116b8, 0x116bf},
		{0x116ca, 0x1171c},
		{0x1172c, 0x1172f},
		{0x1173a, 0x1189f},
		{0x118ea, 0x118fe},
		{0x11900, 0x11abf},
		{0x11af9, 0x11bff},
		{0x11c09, 0x11c09},
		{0x11c37, 0x11c37},
		{0x11c41, 0x11c4f},
		{0x11c5a, 0x11c71},
		{0x11c90, 0x11c91},
		{0x11ca8, 0x11ca8},
		{0x11cb7, 0x11fff},
		{0x1239a, 0x123ff},
		{0x1246f, 0x1247f},
		{0x12544, 0x12fff},
		{0x1342f, 0x143ff},
		{0x14647, 0x167ff},
		{0x16a39, 0x16a3f},
		{0x16a5f, 0x16a5f},
		{0x16a6a, 0x16acf},
		{0x16aee, 0x16aef},
		{0x16af5, 0x16aff},
		{0x16b37, 0x16b3f},
		{0x16b44, 0x16b4f},
		{0x16b5a, 0x16b62},
		{0x16b78, 0x16b7c},
		{0x16b90, 0x16eff},
		{0x16f45, 0x16f4f},
		{0x16f7f, 0x16f8e},
		{0x16fa0, 0x16fdf},
		{0x16fe1, 0x1bbff},
		{0x1bc6b, 0x1bc6f},
		{0x1bc7d, 0x1bc7f},
		{0x1bc89, 0x1bc8f},
		{0x1bc9a, 0x1bc9c},
		{0x1bc9f, 0x1bc9f},
		{0x1bca4, 0x1d164},
		{0x1d16a, 0x1d16c},
		{0x1d183, 0x1d184},
		{0x1d18c, 0x1d1a9},
		{0x1d1ae, 0x1d241},
		{0x1d245, 0x1d3ff},
		{0x1d455, 0x1d455},
		{0x1d49d, 0x1d49d},
		{0x1d4a0, 0x1d4a1},
		{0x1d4a3, 0x1d4a4},
		{0x1d4a7, 0x1d4a8},
		{0x1d4ad, 0x1d4ad},
		{0x1d4ba, 0x1d4ba},
		{0x1d4bc, 0x1d4bc},
		{0x1d4c4, 0x1d4c4},
		{0x1d506, 0x1d506},
		{0x1d50b, 0x1d50c},
		{0x1d515, 0x1d515},
		{0x1d51d, 0x1d51d},
		{0x1d53a, 0x1d53a},
		{0x1d53f, 0x1d53f},
		{0x1d545, 0x1d545},
		{0x1d547, 0x1d549},
		{0x1d551, 0x1d551},
		{0x1d6a6, 0x1d6a7},
		{0x1d6c1, 0x1d6c1},
		{0x1d6db, 0x1d6db},
		{0x1d6fb, 0x1d6fb},
		{0x1d715, 0x1d715},
		{0x1d735, 0x1d735},
		{0x1d74f, 0x1d74f},
		{0x1d76f, 0x1d76f},
		{0x1d789, 0x1d789},
		{0x1d7a9, 0x1d7a9},
		{0x1d7c3, 0x1d7c3},
		{0x1d7cc, 0x1d7cd},
		{0x1d800, 0x1d9ff},
		{0x1da37, 0x1da3a},
		{0x1da6d, 0x1da74},
		{0x1da76, 0x1da83},
		{0x1da85, 0x1da9a},
		{0x1daa0, 0x1daa0},
		{0x1dab0, 0x1dfff},
		{0x1e007, 0x1e007},
		{0x1e019, 0x1e01a},
		{0x1e022, 0x1e022},
		{0x1e025, 0x1e025},
		{0x1e02b, 0x1e7ff},
		{0x1e8c5, 0x1e8cf},
		{0x1e8d7, 0x1e8ff},
		{0x1e94b, 0x1e94f},
		{0x1e95a, 0x1edff},
		{0x1ee04, 0x1ee04},
		{0x1ee20, 0x1ee20},
		{0x1ee23, 0x1ee23},
		{0x1ee25, 0x1ee26},
		{0x1ee28, 0x1ee28},
		{0x1ee33, 0x1ee33},
		{0x1ee38, 0x1ee38},
		{0x1ee3a, 0x1ee3a},
		{0x1ee3c, 0x1ee41},
		{0x1ee43, 0x1ee46},
		{0x1ee48, 0x1ee48},
		{0x1ee4a, 0x1ee4a},
		{0x1ee4c, 0x1ee4c},
		{0x1ee50, 0x1ee50},
		{0x1ee53, 0x1ee53},
		{0x1ee55, 0x1ee56},
		{0x1ee58, 0x1ee58},
		{0x1ee5a, 0x1ee5a},
		{0x1ee5c, 0x1ee5c},
		{0x1ee5e, 0x1ee5e},
		{0x1ee60, 0x1ee60},
		{0x1ee63, 0x1ee63},
		{0x1ee65, 0x1ee66},
		{0x1ee6b, 0x1ee6b},
		{0x1ee73, 0x1ee73},
		{0x1ee78, 0x1ee78},
		{0x1ee7d, 0x1ee7d},
		{0x1ee7f, 0x1ee7f},
		{0x1ee8a, 0x1ee8a},
		{0x1ee9c, 0x1eea0},
		{0x1eea4, 0x1eea4},
		{0x1eeaa, 0x1eeaa},
		{0x1eebc, 0x1f12f},
		{0x1f14a, 0x1f14f},
		{0x1f16a, 0x1f16f},
		{0x1f18a, 0x1ffff},
	}

	breakTest = []string{
		"AA",
		"ÄA",
		"Aa\u2060",
		"Äa\u2060",
		"Aa|:",
		"Äa|:",
		"Aa|'",
		"Äa|'",
		"Aa|'\u2060",
		"Äa|'\u2060",
		"Aa|,",
		"Äa|,",
		"a\u2060A",
		"a\u2060̈A",
		"a\u2060a\u2060",
		"a\u2060̈a\u2060",
		"a\u2060a|:",
		"a\u2060̈a|:",
		"a\u2060a|'",
		"a\u2060̈a|'",
		"a\u2060a|'\u2060",
		"a\u2060̈a|'\u2060",
		"a\u2060a|,",
		"a\u2060̈a|,",
		"a:A",
		"a:̈A",
		"a:a\u2060",
		"a:̈a\u2060",
		"a:a|:",
		"a:̈a|:",
		"a:a|'",
		"a:̈a|'",
		"a:a|'\u2060",
		"a:̈a|'\u2060",
		"a:a|,",
		"a:̈a|,",
		"a'A",
		"a'̈A",
		"a'a\u2060",
		"a'̈a\u2060",
		"a'a|:",
		"a'̈a|:",
		"a'a|'",
		"a'̈a|'",
		"a'a|'\u2060",
		"a'̈a|'\u2060",
		"a'a|,",
		"a'̈a|,",
		"a'\u2060A",
		"a'\u2060̈A",
		"a'\u2060a\u2060",
		"a'\u2060̈a\u2060",
		"a'\u2060a|:",
		"a'\u2060̈a|:",
		"a'\u2060a|'",
		"a'\u2060̈a|'",
		"a'\u2060a|'\u2060",
		"a'\u2060̈a|'\u2060",
		"a'\u2060a|,",
		"a'\u2060̈a|,",
		"a|,|A",
		"a|,̈|A",
		"a|,|a\u2060",
		"a|,̈|a\u2060",
		"a|,|a|:",
		"a|,̈|a|:",
		"a|,|a|'",
		"a|,̈|a|'",
		"a|,|a|'\u2060",
		"a|,̈|a|'\u2060",
		"a|,|a|,",
		"a|,̈|a|,",
		"AAA",
		"A:A",
		"A|:|:|A",
		"A00A",
		"A__A",
		"a|🇦🇧|🇨|b",
		"a|🇦🇧\u200d|🇨|b",
		"a|🇦\u200d🇧|🇨|b",
		"a|🇦🇧|🇨🇩|b",
		"ä\u200d̈b",
		"1_a|:|:|a",
		"1_a|:|.|a",
		"1_a|:|,|a",
		"1_a|.|:|a",
		"1_a|.|.|a",
		"1_a|.|,|a",
		"1_a|,|:|a",
		"1_a|,|.|a",
		"1_a|,|,|a",
		"a_a|:|:|1",
		"a|:|:|a",
		"a_1|:|:|a",
		"a_a|:|:|a",
		"a_a|:|.|1",
		"a|:|.|a",
		"a_1|:|.|a",
		"a_a|:|.|a",
		"a_a|:|,|1",
		"a|:|,|a",
		"a_1|:|,|a",
		"a_a|:|,|a",
		"a_a|.|:|1",
		"a|.|:|a",
		"a_1|.|:|a",
		"a_a|.|:|a",
		"a_a|.|.|1",
		"a|.|.|a",
		"a_1|.|.|a",
		"a_a|.|.|a",
		"a_a|.|,|1",
		"a|.|,|a",
		"a_1|.|,|a",
		"a_a|.|,|a",
		"a_a|,|:|1",
		"a|,|:|a",
		"a_1|,|:|a",
		"a_a|,|:|a",
		"a_a|,|.|1",
		"a|,|.|a",
		"a_1|,|.|a",
		"a_a|,|.|a",
		"a_a|,|,|1",
		"a|,|,|a",
		"a_1|,|,|a",
		"a_a|,|,|a",
	}
)
