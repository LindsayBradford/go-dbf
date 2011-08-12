package godbf

var encoding map[int]byte

func init() {
	encoding = make(map[int]byte)
	encoding[437] = 0x01   //DOS USA code page 437 
	encoding[850] = 0x02   // DOS Multilingual code page 850 
	encoding[1252] = 0x03  // Windows ANSI code page 1252 
	encoding[10000] = 0x04 // Standard Macintosh 
	encoding[865] = 0x08   // Danish OEM
	encoding[437] = 0x09   // Dutch OEM
	encoding[850] = 0x0A   // Dutch OEM Secondary codepage
	encoding[437] = 0x0B   // Finnish OEM
	encoding[437] = 0x0D   // French OEM
	encoding[850] = 0x0E   // French OEM Secondary codepage
	encoding[437] = 0x0F   // German OEM
	encoding[850] = 0x10   // German OEM Secondary codepage
	encoding[437] = 0x11   // Italian OEM
	encoding[850] = 0x12   // Italian OEM Secondary codepage
	encoding[932] = 0x13   // Japanese Shift-JIS
	encoding[850] = 0x14   // Spanish OEM secondary codepage
	encoding[437] = 0x15   // Swedish OEM
	encoding[850] = 0x16   // Swedish OEM secondary codepage
	encoding[865] = 0x17   // Norwegian OEM
	encoding[437] = 0x18   // Spanish OEM
	encoding[437] = 0x19   // English OEM (Britain)
	encoding[850] = 0x1A   // English OEM (Britain) secondary codepage
	encoding[437] = 0x1B   // English OEM (U.S.)
	encoding[863] = 0x1C   // French OEM (Canada)
	encoding[850] = 0x1D   // French OEM secondary codepage
	encoding[852] = 0x1F   // Czech OEM
	encoding[852] = 0x22   // Hungarian OEM
	encoding[852] = 0x23   // Polish OEM
	encoding[860] = 0x24   // Portuguese OEM
	encoding[850] = 0x25   // Portuguese OEM secondary codepage
	encoding[866] = 0x26   // Russian OEM
	encoding[850] = 0x37   // English OEM (U.S.) secondary codepage
	encoding[852] = 0x40   // Romanian OEM
	encoding[936] = 0x4D   // Chinese GBK (PRC)
	encoding[949] = 0x4E   // Korean (ANSI/OEM)
	encoding[950] = 0x4F   // Chinese Big5 (Taiwan)
	encoding[874] = 0x50   // Thai (ANSI/OEM)
	encoding[1252] = 0x57  // ANSI
	encoding[1252] = 0x58  // Western European ANSI
	encoding[1252] = 0x59  // Spanish ANSI
	encoding[852] = 0x64   // Eastern European MSDOS
	encoding[866] = 0x65   // Russian MSDOS
	encoding[865] = 0x66   // Nordic MSDOS
	encoding[861] = 0x67   // Icelandic MSDOS
	encoding[895] = 0x68   // Kamenicky (Czech) MS-DOS 
	encoding[620] = 0x69   // Mazovia (Polish) MS-DOS 
	encoding[737] = 0x6A   // Greek MSDOS (437G)
	encoding[857] = 0x6B   // Turkish MSDOS
	encoding[863] = 0x6C   // FrenchCanadian MSDOS
	encoding[950] = 0x78   // Taiwan Big 5
	encoding[949] = 0x79   // Hangul (Wansung)
	encoding[936] = 0x7A   // PRC GBK
	encoding[932] = 0x7B   // Japanese Shift-JIS
	encoding[874] = 0x7C   // Thai Windows/MSDOS
	encoding[1255] = 0x7D  // Hebrew Windows 
	encoding[1256] = 0x7E  // Arabic Windows 
	encoding[737] = 0x86   // Greek OEM
	encoding[852] = 0x87   // Slovenian OEM
	encoding[857] = 0x88   // Turkish OEM
	encoding[10007] = 0x96 // Russian Macintosh 
	encoding[10029] = 0x97 // Eastern European Macintosh 
	encoding[10006] = 0x98 // Greek Macintosh 
	encoding[1250] = 0xC8  // Eastern European Windows
	encoding[1251] = 0xC9  // Russian Windows
	encoding[1254] = 0xCA  // Turkish Windows
	encoding[1253] = 0xCB  // Greek Windows
	encoding[1257] = 0xCC  // Baltic Windows
}

func GetEncodingCode(codePage int) (code byte) {
	if val, ok := encoding[codePage]; ok {
		code = val
	} else {
		code = encoding[1252]
	}
	return
}
