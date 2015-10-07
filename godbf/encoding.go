package godbf

import (
	"fmt"
)

var lookup map[string]string
var encoding map[int]byte
var encodingTable map[string]byte // this map is used to convert dbase fileencodings to mahonia library encodings

func main() {
	//dizi := []int{437, 850, 1252, 10000, 865, 437, 850, 437, 437, 850, 437, 850, 437, 850, 932, 850, 437, 850, 865, 437, 437, 850, 437, 863, 850, 852, 852, 852, 860, 850, 866, 850, 852, 936, 949, 950, 874, 1252, 1252, 852, 866, 865, 861, 895, 620, 737, 857, 863, 950, 949, 936, 932, 874, 737, 852, 857, 10007, 10029, 10006, 1250, 1251, 1254, 1253, 1257}

	if code, ok := encodingTable[lookup["l5"]]; ok {
		fmt.Printf("code=%x", code)
	} else {
		panic("unsupported encoding")
	}
}

func init() {
	lookup = make(map[string]string)

	lookup["ISO-8859-2"] = "ISO-8859-2"
	lookup["ISO_8859-2:1987"] = "ISO-8859-2"
	lookup["iso-ir-101"] = "ISO-8859-2"
	lookup["latin2"] = "ISO-8859-2"
	lookup["l2"] = "ISO-8859-2"
	lookup["csISOLatin2"] = "ISO-8859-2"

	lookup["ISO-8859-3"] = "ISO-8859-3"
	lookup["ISO_8859-3:1988"] = "ISO-8859-3"
	lookup["iso-ir-109"] = "ISO-8859-3"
	lookup["latin3"] = "ISO-8859-3"
	lookup["l3"] = "ISO-8859-3"
	lookup["csISOLatin3"] = "ISO-8859-3"

	lookup["ISO-8859-4"] = "ISO-8859-4"
	lookup["ISO_8859-4:1988"] = "ISO-8859-4"
	lookup["iso-ir-110"] = "ISO-8859-4"
	lookup["latin4"] = "ISO-8859-4"
	lookup["l4"] = "ISO-8859-4"
	lookup["csISOLatin4"] = "ISO-8859-4"

	lookup["ISO-8859-5"] = "ISO-8859-5"
	lookup["ISO_8859-5:1988"] = "ISO-8859-5"
	lookup["iso-ir-144"] = "ISO-8859-5"
	lookup["cyrillic"] = "ISO-8859-5"
	lookup["csISOLatinCyrillic"] = "ISO-8859-5"

	lookup["ISO-8859-6"] = "ISO-8859-6"
	lookup["ISO_8859-6:1987"] = "ISO-8859-6"
	lookup["iso-ir-127"] = "ISO-8859-6"
	lookup["ECMA-114"] = "ISO-8859-6"
	lookup["ASMO-708"] = "ISO-8859-6"
	lookup["arabic"] = "ISO-8859-6"
	lookup["csISOLatinArabic"] = "ISO-8859-6"

	lookup["ISO-8859-7"] = "ISO-8859-7"
	lookup["ISO_8859-7:2003"] = "ISO-8859-7"
	lookup["iso-ir-126"] = "ISO-8859-7"
	lookup["ELOT_928"] = "ISO-8859-7"
	lookup["ECMA-118"] = "ISO-8859-7"
	lookup["greek"] = "ISO-8859-7"
	lookup["greek8"] = "ISO-8859-7"
	lookup["csISOLatinGreek"] = "ISO-8859-7"

	lookup["ISO-8859-8"] = "ISO-8859-8"
	lookup["ISO_8859-8:1999"] = "ISO-8859-8"
	lookup["iso-ir-138"] = "ISO-8859-8"
	lookup["hebrew"] = "ISO-8859-8"
	lookup["csISOLatinHebrew"] = "ISO-8859-8"

	lookup["ISO-8859-9"] = "ISO-8859-9"
	lookup["ISO_8859-9:1999"] = "ISO-8859-9"
	lookup["iso-ir-148"] = "ISO-8859-9"
	lookup["latin5"] = "ISO-8859-9"
	lookup["l5"] = "ISO-8859-9"
	lookup["csISOLatin5"] = "ISO-8859-9"

	lookup["ISO-8859-10"] = "ISO-8859-10"
	lookup["iso_8859-10:1992"] = "ISO-8859-10"
	lookup["l6"] = "ISO-8859-10"
	lookup["iso-ir-157"] = "ISO-8859-10"
	lookup["latin6"] = "ISO-8859-10"
	lookup["csISOLatin6"] = "ISO-8859-10"

	lookup["ISO-8859-11"] = "ISO-8859-11"
	lookup["iso_8859-11:2001"] = "ISO-8859-11"
	lookup["Latin/Thai"] = "ISO-8859-11"
	lookup["TIS-620"] = "ISO-8859-11"

	lookup["ISO-8859-13"] = "ISO-8859-13"
	lookup["latin7"] = "ISO-8859-13"
	lookup["Baltic Rim"] = "ISO-8859-13"

	lookup["ISO-8859-14"] = "ISO-8859-14"
	lookup["iso-ir-199"] = "ISO-8859-14"
	lookup["ISO_8859-14:1998"] = "ISO-8859-14"
	lookup["latin8"] = "ISO-8859-14"
	lookup["iso-celtic"] = "ISO-8859-14"
	lookup["l8"] = "ISO-8859-14"

	lookup["ISO-8859-15"] = "ISO-8859-15"
	lookup["Latin-9"] = "ISO-8859-15"

	lookup["ISO-8859-16"] = "ISO-8859-16"
	lookup["iso-ir-226"] = "ISO-8859-16"
	lookup["ISO_8859-16:2001"] = "ISO-8859-16"
	lookup["latin10"] = "ISO-8859-16"
	lookup["l10"] = "ISO-8859-16"

	lookup["macos-0_2-10.2"] = "macos-0_2-10.2"
	lookup["macos-0_2-10.2"] = "macos-0_2-10.2"
	lookup["macintosh"] = "macos-0_2-10.2"
	lookup["mac"] = "macos-0_2-10.2"
	lookup["csMacintosh"] = "macos-0_2-10.2"
	lookup["windows-10000"] = "macos-0_2-10.2"
	lookup["macroman"] = "macos-0_2-10.2"

	lookup["macos-6_2-10.4"] = "macos-6_2-10.4"
	lookup["macos-6_2-10.4"] = "macos-6_2-10.4"
	lookup["x-mac-greek"] = "macos-6_2-10.4"
	lookup["windows-10006"] = "macos-6_2-10.4"
	lookup["macgr"] = "macos-6_2-10.4"

	lookup["macos-7_3-10.2"] = "macos-7_3-10.2"
	lookup["macos-7_3-10.2"] = "macos-7_3-10.2"
	lookup["x-mac-cyrillic"] = "macos-7_3-10.2"
	lookup["windows-10007"] = "macos-7_3-10.2"
	lookup["mac-cyrillic"] = "macos-7_3-10.2"
	lookup["maccy"] = "macos-7_3-10.2"

	lookup["macos-29-10.2"] = "macos-29-10.2"
	lookup["macos-29-10.2"] = "macos-29-10.2"
	lookup["x-mac-centraleurroman"] = "macos-29-10.2"
	lookup["windows-10029"] = "macos-29-10.2"
	lookup["x-mac-ce"] = "macos-29-10.2"
	lookup["macce"] = "macos-29-10.2"
	lookup["maccentraleurope"] = "macos-29-10.2"

	lookup["macos-35-10.2"] = "macos-35-10.2"
	lookup["macos-35-10.2"] = "macos-35-10.2"
	lookup["x-mac-turkish"] = "macos-35-10.2"
	lookup["windows-10081"] = "macos-35-10.2"
	lookup["mactr"] = "macos-35-10.2"

	lookup["windows-1250"] = "windows-1250"
	lookup["1250"] = "windows-1250"

	lookup["windows-1251"] = "windows-1251"
	lookup["1251"] = "windows-1251"

	lookup["windows-1252"] = "windows-1252"
	lookup["1252"] = "windows-1252"

	lookup["windows-1253"] = "windows-1253"
	lookup["1253"] = "windows-1253"

	lookup["windows-1254"] = "windows-1254"
	lookup["1254"] = "windows-1254"

	lookup["windows-1255"] = "windows-1255"
	lookup["1255"] = "windows-1255"

	lookup["windows-1256"] = "windows-1256"
	lookup["1256"] = "windows-1256"

	lookup["windows-1257"] = "windows-1257"
	lookup["1257"] = "windows-1257"

	lookup["windows-1258"] = "windows-1258"
	lookup["1258"] = "windows-1258"

	lookup["windows-874"] = "windows-874"
	lookup["874"] = "windows-874"

	lookup["IBM037"] = "IBM037"
	lookup["cp037"] = "IBM037"
	lookup["ebcdic-cp-us"] = "IBM037"
	lookup["ebcdic-cp-ca"] = "IBM037"
	lookup["ebcdic-cp-wt"] = "IBM037"
	lookup["ebcdic-cp-nl"] = "IBM037"
	lookup["csIBM037"] = "IBM037"

	lookup["ibm-273_P100-1995"] = "ibm-273_P100-1995"
	lookup["ibm-273_P100-1995"] = "ibm-273_P100-1995"
	lookup["ibm-273"] = "ibm-273_P100-1995"
	lookup["IBM273"] = "ibm-273_P100-1995"
	lookup["CP273"] = "ibm-273_P100-1995"
	lookup["csIBM273"] = "ibm-273_P100-1995"
	lookup["ebcdic-de"] = "ibm-273_P100-1995"
	lookup["273"] = "ibm-273_P100-1995"

	lookup["ibm-277_P100-1995"] = "ibm-277_P100-1995"
	lookup["ibm-277_P100-1995"] = "ibm-277_P100-1995"
	lookup["ibm-277"] = "ibm-277_P100-1995"
	lookup["IBM277"] = "ibm-277_P100-1995"
	lookup["cp277"] = "ibm-277_P100-1995"
	lookup["EBCDIC-CP-DK"] = "ibm-277_P100-1995"
	lookup["EBCDIC-CP-NO"] = "ibm-277_P100-1995"
	lookup["csIBM277"] = "ibm-277_P100-1995"
	lookup["ebcdic-dk"] = "ibm-277_P100-1995"
	lookup["277"] = "ibm-277_P100-1995"

	lookup["ibm-278_P100-1995"] = "ibm-278_P100-1995"
	lookup["ibm-278_P100-1995"] = "ibm-278_P100-1995"
	lookup["ibm-278"] = "ibm-278_P100-1995"
	lookup["IBM278"] = "ibm-278_P100-1995"
	lookup["cp278"] = "ibm-278_P100-1995"
	lookup["ebcdic-cp-fi"] = "ibm-278_P100-1995"
	lookup["ebcdic-cp-se"] = "ibm-278_P100-1995"
	lookup["csIBM278"] = "ibm-278_P100-1995"
	lookup["ebcdic-sv"] = "ibm-278_P100-1995"
	lookup["278"] = "ibm-278_P100-1995"

	lookup["ibm-280_P100-1995"] = "ibm-280_P100-1995"
	lookup["ibm-280_P100-1995"] = "ibm-280_P100-1995"
	lookup["ibm-280"] = "ibm-280_P100-1995"
	lookup["IBM280"] = "ibm-280_P100-1995"
	lookup["CP280"] = "ibm-280_P100-1995"
	lookup["ebcdic-cp-it"] = "ibm-280_P100-1995"
	lookup["csIBM280"] = "ibm-280_P100-1995"
	lookup["280"] = "ibm-280_P100-1995"

	lookup["ibm-284_P100-1995"] = "ibm-284_P100-1995"
	lookup["ibm-284_P100-1995"] = "ibm-284_P100-1995"
	lookup["ibm-284"] = "ibm-284_P100-1995"
	lookup["IBM284"] = "ibm-284_P100-1995"
	lookup["CP284"] = "ibm-284_P100-1995"
	lookup["ebcdic-cp-es"] = "ibm-284_P100-1995"
	lookup["csIBM284"] = "ibm-284_P100-1995"
	lookup["cpibm284"] = "ibm-284_P100-1995"
	lookup["284"] = "ibm-284_P100-1995"

	lookup["ibm-285_P100-1995"] = "ibm-285_P100-1995"
	lookup["ibm-285_P100-1995"] = "ibm-285_P100-1995"
	lookup["ibm-285"] = "ibm-285_P100-1995"
	lookup["IBM285"] = "ibm-285_P100-1995"
	lookup["CP285"] = "ibm-285_P100-1995"
	lookup["ebcdic-cp-gb"] = "ibm-285_P100-1995"
	lookup["csIBM285"] = "ibm-285_P100-1995"
	lookup["cpibm285"] = "ibm-285_P100-1995"
	lookup["ebcdic-gb"] = "ibm-285_P100-1995"
	lookup["285"] = "ibm-285_P100-1995"

	lookup["ibm-290_P100-1995"] = "ibm-290_P100-1995"
	lookup["ibm-290_P100-1995"] = "ibm-290_P100-1995"
	lookup["ibm-290"] = "ibm-290_P100-1995"
	lookup["IBM290"] = "ibm-290_P100-1995"
	lookup["cp290"] = "ibm-290_P100-1995"
	lookup["EBCDIC-JP-kana"] = "ibm-290_P100-1995"
	lookup["csIBM290"] = "ibm-290_P100-1995"

	lookup["ibm-297_P100-1995"] = "ibm-297_P100-1995"
	lookup["ibm-297_P100-1995"] = "ibm-297_P100-1995"
	lookup["ibm-297"] = "ibm-297_P100-1995"
	lookup["IBM297"] = "ibm-297_P100-1995"
	lookup["cp297"] = "ibm-297_P100-1995"
	lookup["ebcdic-cp-fr"] = "ibm-297_P100-1995"
	lookup["csIBM297"] = "ibm-297_P100-1995"
	lookup["cpibm297"] = "ibm-297_P100-1995"
	lookup["297"] = "ibm-297_P100-1995"

	lookup["ibm-420_X120-1999"] = "ibm-420_X120-1999"
	lookup["ibm-420_X120-1999"] = "ibm-420_X120-1999"
	lookup["ibm-420"] = "ibm-420_X120-1999"
	lookup["IBM420"] = "ibm-420_X120-1999"
	lookup["cp420"] = "ibm-420_X120-1999"
	lookup["ebcdic-cp-ar1"] = "ibm-420_X120-1999"
	lookup["csIBM420"] = "ibm-420_X120-1999"
	lookup["420"] = "ibm-420_X120-1999"

	lookup["IBM424"] = "IBM424"
	lookup["cp424"] = "IBM424"
	lookup["ebcdic-cp-he"] = "IBM424"
	lookup["csIBM424"] = "IBM424"

	lookup["IBM437"] = "IBM437"
	lookup["cp437"] = "IBM437"
	lookup["437"] = "IBM437"
	lookup["csPC8CodePage437"] = "IBM437"

	lookup["IBM500"] = "IBM500"
	lookup["CP500"] = "IBM500"
	lookup["ebcdic-cp-be"] = "IBM500"
	lookup["ebcdic-cp-ch"] = "IBM500"
	lookup["csIBM500"] = "IBM500"

	lookup["ibm-720_P100-1997"] = "ibm-720_P100-1997"
	lookup["ibm-720_P100-1997"] = "ibm-720_P100-1997"
	lookup["ibm-720"] = "ibm-720_P100-1997"
	lookup["windows-720"] = "ibm-720_P100-1997"
	lookup["DOS-720"] = "ibm-720_P100-1997"

	lookup["IBM737"] = "IBM737"
	lookup["cp737"] = "IBM737"
	lookup["cp737_DOSGreek"] = "IBM737"

	lookup["IBM775"] = "IBM775"
	lookup["cp775"] = "IBM775"
	lookup["csPC775Baltic"] = "IBM775"

	lookup["ibm-803_P100-1999"] = "ibm-803_P100-1999"
	lookup["ibm-803_P100-1999"] = "ibm-803_P100-1999"
	lookup["ibm-803"] = "ibm-803_P100-1999"
	lookup["cp803"] = "ibm-803_P100-1999"

	lookup["ibm-838_P100-1995"] = "ibm-838_P100-1995"
	lookup["ibm-838_P100-1995"] = "ibm-838_P100-1995"
	lookup["ibm-838"] = "ibm-838_P100-1995"
	lookup["IBM838"] = "ibm-838_P100-1995"
	lookup["IBM-Thai"] = "ibm-838_P100-1995"
	lookup["csIBMThai"] = "ibm-838_P100-1995"
	lookup["cp838"] = "ibm-838_P100-1995"
	lookup["838"] = "ibm-838_P100-1995"
	lookup["ibm-9030"] = "ibm-838_P100-1995"

	lookup["IBM850"] = "IBM850"
	lookup["cp850"] = "IBM850"
	lookup["850"] = "IBM850"
	lookup["csPC850Multilingual"] = "IBM850"

	lookup["ibm-851_P100-1995"] = "ibm-851_P100-1995"
	lookup["ibm-851_P100-1995"] = "ibm-851_P100-1995"
	lookup["ibm-851"] = "ibm-851_P100-1995"
	lookup["IBM851"] = "ibm-851_P100-1995"
	lookup["cp851"] = "ibm-851_P100-1995"
	lookup["851"] = "ibm-851_P100-1995"
	lookup["csPC851"] = "ibm-851_P100-1995"

	lookup["IBM852"] = "IBM852"
	lookup["cp852"] = "IBM852"
	lookup["852"] = "IBM852"
	lookup["csPCp852"] = "IBM852"

	lookup["IBM855"] = "IBM855"
	lookup["cp855"] = "IBM855"
	lookup["855"] = "IBM855"
	lookup["csIBM855"] = "IBM855"

	lookup["IBM856"] = "IBM856"
	lookup["cp856"] = "IBM856"
	lookup["cp856_Hebrew_PC"] = "IBM856"

	lookup["ibm-857_P100-1995"] = "ibm-857_P100-1995"
	lookup["ibm-857_P100-1995"] = "ibm-857_P100-1995"
	lookup["ibm-857"] = "ibm-857_P100-1995"
	lookup["IBM857"] = "ibm-857_P100-1995"
	lookup["cp857"] = "ibm-857_P100-1995"
	lookup["857"] = "ibm-857_P100-1995"
	lookup["csIBM857"] = "ibm-857_P100-1995"
	lookup["windows-857"] = "ibm-857_P100-1995"

	lookup["ibm-858_P100-1997"] = "ibm-858_P100-1997"
	lookup["ibm-858_P100-1997"] = "ibm-858_P100-1997"
	lookup["ibm-858"] = "ibm-858_P100-1997"
	lookup["IBM00858"] = "ibm-858_P100-1997"
	lookup["CCSID00858"] = "ibm-858_P100-1997"
	lookup["CP00858"] = "ibm-858_P100-1997"
	lookup["PC-Multilingual-850+euro"] = "ibm-858_P100-1997"
	lookup["cp858"] = "ibm-858_P100-1997"
	lookup["windows-858"] = "ibm-858_P100-1997"

	lookup["ibm-860_P100-1995"] = "ibm-860_P100-1995"
	lookup["ibm-860_P100-1995"] = "ibm-860_P100-1995"
	lookup["ibm-860"] = "ibm-860_P100-1995"
	lookup["IBM860"] = "ibm-860_P100-1995"
	lookup["cp860"] = "ibm-860_P100-1995"
	lookup["860"] = "ibm-860_P100-1995"
	lookup["csIBM860"] = "ibm-860_P100-1995"

	lookup["ibm-861_P100-1995"] = "ibm-861_P100-1995"
	lookup["ibm-861_P100-1995"] = "ibm-861_P100-1995"
	lookup["ibm-861"] = "ibm-861_P100-1995"
	lookup["IBM861"] = "ibm-861_P100-1995"
	lookup["cp861"] = "ibm-861_P100-1995"
	lookup["861"] = "ibm-861_P100-1995"
	lookup["cp-is"] = "ibm-861_P100-1995"
	lookup["csIBM861"] = "ibm-861_P100-1995"
	lookup["windows-861"] = "ibm-861_P100-1995"

	lookup["ibm-862_P100-1995"] = "ibm-862_P100-1995"
	lookup["ibm-862_P100-1995"] = "ibm-862_P100-1995"
	lookup["ibm-862"] = "ibm-862_P100-1995"
	lookup["IBM862"] = "ibm-862_P100-1995"
	lookup["cp862"] = "ibm-862_P100-1995"
	lookup["862"] = "ibm-862_P100-1995"
	lookup["csPC862LatinHebrew"] = "ibm-862_P100-1995"
	lookup["DOS-862"] = "ibm-862_P100-1995"
	lookup["windows-862"] = "ibm-862_P100-1995"

	lookup["ibm-863_P100-1995"] = "ibm-863_P100-1995"
	lookup["ibm-863_P100-1995"] = "ibm-863_P100-1995"
	lookup["ibm-863"] = "ibm-863_P100-1995"
	lookup["IBM863"] = "ibm-863_P100-1995"
	lookup["cp863"] = "ibm-863_P100-1995"
	lookup["863"] = "ibm-863_P100-1995"
	lookup["csIBM863"] = "ibm-863_P100-1995"

	lookup["ibm-864_X110-1999"] = "ibm-864_X110-1999"
	lookup["ibm-864_X110-1999"] = "ibm-864_X110-1999"
	lookup["ibm-864"] = "ibm-864_X110-1999"
	lookup["IBM864"] = "ibm-864_X110-1999"
	lookup["cp864"] = "ibm-864_X110-1999"
	lookup["csIBM864"] = "ibm-864_X110-1999"

	lookup["ibm-865_P100-1995"] = "ibm-865_P100-1995"
	lookup["ibm-865_P100-1995"] = "ibm-865_P100-1995"
	lookup["ibm-865"] = "ibm-865_P100-1995"
	lookup["IBM865"] = "ibm-865_P100-1995"
	lookup["cp865"] = "ibm-865_P100-1995"
	lookup["865"] = "ibm-865_P100-1995"
	lookup["csIBM865"] = "ibm-865_P100-1995"

	lookup["IBM866"] = "IBM866"
	lookup["cp866"] = "IBM866"
	lookup["866"] = "IBM866"
	lookup["csIBM866"] = "IBM866"

	lookup["ibm-867_P100-1998"] = "ibm-867_P100-1998"
	lookup["ibm-867_P100-1998"] = "ibm-867_P100-1998"
	lookup["ibm-867"] = "ibm-867_P100-1998"

	lookup["ibm-868_P100-1995"] = "ibm-868_P100-1995"
	lookup["ibm-868_P100-1995"] = "ibm-868_P100-1995"
	lookup["ibm-868"] = "ibm-868_P100-1995"
	lookup["IBM868"] = "ibm-868_P100-1995"
	lookup["CP868"] = "ibm-868_P100-1995"
	lookup["868"] = "ibm-868_P100-1995"
	lookup["csIBM868"] = "ibm-868_P100-1995"
	lookup["cp-ar"] = "ibm-868_P100-1995"

	lookup["ibm-869_P100-1995"] = "ibm-869_P100-1995"
	lookup["ibm-869_P100-1995"] = "ibm-869_P100-1995"
	lookup["ibm-869"] = "ibm-869_P100-1995"
	lookup["IBM869"] = "ibm-869_P100-1995"
	lookup["cp869"] = "ibm-869_P100-1995"
	lookup["869"] = "ibm-869_P100-1995"
	lookup["cp-gr"] = "ibm-869_P100-1995"
	lookup["csIBM869"] = "ibm-869_P100-1995"
	lookup["windows-869"] = "ibm-869_P100-1995"

	lookup["ibm-870_P100-1995"] = "ibm-870_P100-1995"
	lookup["ibm-870_P100-1995"] = "ibm-870_P100-1995"
	lookup["ibm-870"] = "ibm-870_P100-1995"
	lookup["IBM870"] = "ibm-870_P100-1995"
	lookup["CP870"] = "ibm-870_P100-1995"
	lookup["ebcdic-cp-roece"] = "ibm-870_P100-1995"
	lookup["ebcdic-cp-yu"] = "ibm-870_P100-1995"
	lookup["csIBM870"] = "ibm-870_P100-1995"

	lookup["ibm-871_P100-1995"] = "ibm-871_P100-1995"
	lookup["ibm-871_P100-1995"] = "ibm-871_P100-1995"
	lookup["ibm-871"] = "ibm-871_P100-1995"
	lookup["IBM871"] = "ibm-871_P100-1995"
	lookup["ebcdic-cp-is"] = "ibm-871_P100-1995"
	lookup["csIBM871"] = "ibm-871_P100-1995"
	lookup["CP871"] = "ibm-871_P100-1995"
	lookup["ebcdic-is"] = "ibm-871_P100-1995"
	lookup["871"] = "ibm-871_P100-1995"

	lookup["ibm-874_P100-1995"] = "ibm-874_P100-1995"
	lookup["ibm-874_P100-1995"] = "ibm-874_P100-1995"
	lookup["ibm-874"] = "ibm-874_P100-1995"
	lookup["ibm-9066"] = "ibm-874_P100-1995"
	lookup["cp874"] = "ibm-874_P100-1995"
	lookup["tis620.2533"] = "ibm-874_P100-1995"
	lookup["eucTH"] = "ibm-874_P100-1995"

	lookup["ibm-875_P100-1995"] = "ibm-875_P100-1995"
	lookup["ibm-875_P100-1995"] = "ibm-875_P100-1995"
	lookup["ibm-875"] = "ibm-875_P100-1995"
	lookup["IBM875"] = "ibm-875_P100-1995"
	lookup["cp875"] = "ibm-875_P100-1995"
	lookup["875"] = "ibm-875_P100-1995"

	lookup["ibm-901_P100-1999"] = "ibm-901_P100-1999"
	lookup["ibm-901_P100-1999"] = "ibm-901_P100-1999"
	lookup["ibm-901"] = "ibm-901_P100-1999"

	lookup["ibm-902_P100-1999"] = "ibm-902_P100-1999"
	lookup["ibm-902_P100-1999"] = "ibm-902_P100-1999"
	lookup["ibm-902"] = "ibm-902_P100-1999"

	lookup["ibm-916_P100-1995"] = "ibm-916_P100-1995"
	lookup["ibm-916_P100-1995"] = "ibm-916_P100-1995"
	lookup["ibm-916"] = "ibm-916_P100-1995"
	lookup["cp916"] = "ibm-916_P100-1995"
	lookup["916"] = "ibm-916_P100-1995"

	lookup["ibm-918_P100-1995"] = "ibm-918_P100-1995"
	lookup["ibm-918_P100-1995"] = "ibm-918_P100-1995"
	lookup["ibm-918"] = "ibm-918_P100-1995"
	lookup["IBM918"] = "ibm-918_P100-1995"
	lookup["CP918"] = "ibm-918_P100-1995"
	lookup["ebcdic-cp-ar2"] = "ibm-918_P100-1995"
	lookup["csIBM918"] = "ibm-918_P100-1995"

	lookup["ibm-922_P100-1999"] = "ibm-922_P100-1999"
	lookup["ibm-922_P100-1999"] = "ibm-922_P100-1999"
	lookup["ibm-922"] = "ibm-922_P100-1999"
	lookup["IBM922"] = "ibm-922_P100-1999"
	lookup["cp922"] = "ibm-922_P100-1999"
	lookup["922"] = "ibm-922_P100-1999"

	lookup["ibm-1006_P100-1995"] = "ibm-1006_P100-1995"
	lookup["ibm-1006_P100-1995"] = "ibm-1006_P100-1995"
	lookup["ibm-1006"] = "ibm-1006_P100-1995"
	lookup["IBM1006"] = "ibm-1006_P100-1995"
	lookup["cp1006"] = "ibm-1006_P100-1995"
	lookup["1006"] = "ibm-1006_P100-1995"

	lookup["ibm-1025_P100-1995"] = "ibm-1025_P100-1995"
	lookup["ibm-1025_P100-1995"] = "ibm-1025_P100-1995"
	lookup["ibm-1025"] = "ibm-1025_P100-1995"
	lookup["cp1025"] = "ibm-1025_P100-1995"
	lookup["1025"] = "ibm-1025_P100-1995"

	lookup["ibm-1026_P100-1995"] = "ibm-1026_P100-1995"
	lookup["ibm-1026_P100-1995"] = "ibm-1026_P100-1995"
	lookup["ibm-1026"] = "ibm-1026_P100-1995"
	lookup["IBM1026"] = "ibm-1026_P100-1995"
	lookup["CP1026"] = "ibm-1026_P100-1995"
	lookup["csIBM1026"] = "ibm-1026_P100-1995"
	lookup["1026"] = "ibm-1026_P100-1995"

	lookup["ibm-1047_P100-1995"] = "ibm-1047_P100-1995"
	lookup["ibm-1047_P100-1995"] = "ibm-1047_P100-1995"
	lookup["ibm-1047"] = "ibm-1047_P100-1995"
	lookup["IBM1047"] = "ibm-1047_P100-1995"
	lookup["cp1047"] = "ibm-1047_P100-1995"
	lookup["1047"] = "ibm-1047_P100-1995"

	lookup["ibm-1097_P100-1995"] = "ibm-1097_P100-1995"
	lookup["ibm-1097_P100-1995"] = "ibm-1097_P100-1995"
	lookup["ibm-1097"] = "ibm-1097_P100-1995"
	lookup["cp1097"] = "ibm-1097_P100-1995"
	lookup["1097"] = "ibm-1097_P100-1995"

	lookup["ibm-1098_P100-1995"] = "ibm-1098_P100-1995"
	lookup["ibm-1098_P100-1995"] = "ibm-1098_P100-1995"
	lookup["ibm-1098"] = "ibm-1098_P100-1995"
	lookup["IBM1098"] = "ibm-1098_P100-1995"
	lookup["cp1098"] = "ibm-1098_P100-1995"
	lookup["1098"] = "ibm-1098_P100-1995"

	lookup["ibm-1112_P100-1995"] = "ibm-1112_P100-1995"
	lookup["ibm-1112_P100-1995"] = "ibm-1112_P100-1995"
	lookup["ibm-1112"] = "ibm-1112_P100-1995"
	lookup["cp1112"] = "ibm-1112_P100-1995"
	lookup["1112"] = "ibm-1112_P100-1995"

	lookup["ibm-1122_P100-1999"] = "ibm-1122_P100-1999"
	lookup["ibm-1122_P100-1999"] = "ibm-1122_P100-1999"
	lookup["ibm-1122"] = "ibm-1122_P100-1999"
	lookup["cp1122"] = "ibm-1122_P100-1999"
	lookup["1122"] = "ibm-1122_P100-1999"

	lookup["ibm-1123_P100-1995"] = "ibm-1123_P100-1995"
	lookup["ibm-1123_P100-1995"] = "ibm-1123_P100-1995"
	lookup["ibm-1123"] = "ibm-1123_P100-1995"
	lookup["cp1123"] = "ibm-1123_P100-1995"
	lookup["1123"] = "ibm-1123_P100-1995"

	lookup["ibm-1124_P100-1996"] = "ibm-1124_P100-1996"
	lookup["ibm-1124_P100-1996"] = "ibm-1124_P100-1996"
	lookup["ibm-1124"] = "ibm-1124_P100-1996"
	lookup["cp1124"] = "ibm-1124_P100-1996"
	lookup["1124"] = "ibm-1124_P100-1996"

	lookup["ibm-1125_P100-1997"] = "ibm-1125_P100-1997"
	lookup["ibm-1125_P100-1997"] = "ibm-1125_P100-1997"
	lookup["ibm-1125"] = "ibm-1125_P100-1997"
	lookup["cp1125"] = "ibm-1125_P100-1997"

	lookup["ibm-1129_P100-1997"] = "ibm-1129_P100-1997"
	lookup["ibm-1129_P100-1997"] = "ibm-1129_P100-1997"
	lookup["ibm-1129"] = "ibm-1129_P100-1997"

	lookup["ibm-1130_P100-1997"] = "ibm-1130_P100-1997"
	lookup["ibm-1130_P100-1997"] = "ibm-1130_P100-1997"
	lookup["ibm-1130"] = "ibm-1130_P100-1997"

	lookup["ibm-1131_P100-1997"] = "ibm-1131_P100-1997"
	lookup["ibm-1131_P100-1997"] = "ibm-1131_P100-1997"
	lookup["ibm-1131"] = "ibm-1131_P100-1997"
	lookup["cp1131"] = "ibm-1131_P100-1997"

	lookup["ibm-1132_P100-1998"] = "ibm-1132_P100-1998"
	lookup["ibm-1132_P100-1998"] = "ibm-1132_P100-1998"
	lookup["ibm-1132"] = "ibm-1132_P100-1998"

	lookup["ibm-1133_P100-1997"] = "ibm-1133_P100-1997"
	lookup["ibm-1133_P100-1997"] = "ibm-1133_P100-1997"
	lookup["ibm-1133"] = "ibm-1133_P100-1997"

	lookup["ibm-1137_P100-1999"] = "ibm-1137_P100-1999"
	lookup["ibm-1137_P100-1999"] = "ibm-1137_P100-1999"
	lookup["ibm-1137"] = "ibm-1137_P100-1999"

	lookup["ibm-1140_P100-1997"] = "ibm-1140_P100-1997"
	lookup["ibm-1140_P100-1997"] = "ibm-1140_P100-1997"
	lookup["ibm-1140"] = "ibm-1140_P100-1997"
	lookup["IBM01140"] = "ibm-1140_P100-1997"
	lookup["CCSID01140"] = "ibm-1140_P100-1997"
	lookup["CP01140"] = "ibm-1140_P100-1997"
	lookup["cp1140"] = "ibm-1140_P100-1997"
	lookup["ebcdic-us-37+euro"] = "ibm-1140_P100-1997"

	lookup["ibm-1141_P100-1997"] = "ibm-1141_P100-1997"
	lookup["ibm-1141_P100-1997"] = "ibm-1141_P100-1997"
	lookup["ibm-1141"] = "ibm-1141_P100-1997"
	lookup["IBM01141"] = "ibm-1141_P100-1997"
	lookup["CCSID01141"] = "ibm-1141_P100-1997"
	lookup["CP01141"] = "ibm-1141_P100-1997"
	lookup["cp1141"] = "ibm-1141_P100-1997"
	lookup["ebcdic-de-273+euro"] = "ibm-1141_P100-1997"

	lookup["ibm-1142_P100-1997"] = "ibm-1142_P100-1997"
	lookup["ibm-1142_P100-1997"] = "ibm-1142_P100-1997"
	lookup["ibm-1142"] = "ibm-1142_P100-1997"
	lookup["IBM01142"] = "ibm-1142_P100-1997"
	lookup["CCSID01142"] = "ibm-1142_P100-1997"
	lookup["CP01142"] = "ibm-1142_P100-1997"
	lookup["cp1142"] = "ibm-1142_P100-1997"
	lookup["ebcdic-dk-277+euro"] = "ibm-1142_P100-1997"
	lookup["ebcdic-no-277+euro"] = "ibm-1142_P100-1997"

	lookup["ibm-1143_P100-1997"] = "ibm-1143_P100-1997"
	lookup["ibm-1143_P100-1997"] = "ibm-1143_P100-1997"
	lookup["ibm-1143"] = "ibm-1143_P100-1997"
	lookup["IBM01143"] = "ibm-1143_P100-1997"
	lookup["CCSID01143"] = "ibm-1143_P100-1997"
	lookup["CP01143"] = "ibm-1143_P100-1997"
	lookup["cp1143"] = "ibm-1143_P100-1997"
	lookup["ebcdic-fi-278+euro"] = "ibm-1143_P100-1997"
	lookup["ebcdic-se-278+euro"] = "ibm-1143_P100-1997"

	lookup["ibm-1144_P100-1997"] = "ibm-1144_P100-1997"
	lookup["ibm-1144_P100-1997"] = "ibm-1144_P100-1997"
	lookup["ibm-1144"] = "ibm-1144_P100-1997"
	lookup["IBM01144"] = "ibm-1144_P100-1997"
	lookup["CCSID01144"] = "ibm-1144_P100-1997"
	lookup["CP01144"] = "ibm-1144_P100-1997"
	lookup["cp1144"] = "ibm-1144_P100-1997"
	lookup["ebcdic-it-280+euro"] = "ibm-1144_P100-1997"

	lookup["ibm-1145_P100-1997"] = "ibm-1145_P100-1997"
	lookup["ibm-1145_P100-1997"] = "ibm-1145_P100-1997"
	lookup["ibm-1145"] = "ibm-1145_P100-1997"
	lookup["IBM01145"] = "ibm-1145_P100-1997"
	lookup["CCSID01145"] = "ibm-1145_P100-1997"
	lookup["CP01145"] = "ibm-1145_P100-1997"
	lookup["cp1145"] = "ibm-1145_P100-1997"
	lookup["ebcdic-es-284+euro"] = "ibm-1145_P100-1997"

	lookup["ibm-1146_P100-1997"] = "ibm-1146_P100-1997"
	lookup["ibm-1146_P100-1997"] = "ibm-1146_P100-1997"
	lookup["ibm-1146"] = "ibm-1146_P100-1997"
	lookup["IBM01146"] = "ibm-1146_P100-1997"
	lookup["CCSID01146"] = "ibm-1146_P100-1997"
	lookup["CP01146"] = "ibm-1146_P100-1997"
	lookup["cp1146"] = "ibm-1146_P100-1997"
	lookup["ebcdic-gb-285+euro"] = "ibm-1146_P100-1997"

	lookup["ibm-1147_P100-1997"] = "ibm-1147_P100-1997"
	lookup["ibm-1147_P100-1997"] = "ibm-1147_P100-1997"
	lookup["ibm-1147"] = "ibm-1147_P100-1997"
	lookup["IBM01147"] = "ibm-1147_P100-1997"
	lookup["CCSID01147"] = "ibm-1147_P100-1997"
	lookup["CP01147"] = "ibm-1147_P100-1997"
	lookup["cp1147"] = "ibm-1147_P100-1997"
	lookup["ebcdic-fr-297+euro"] = "ibm-1147_P100-1997"

	lookup["ibm-1148_P100-1997"] = "ibm-1148_P100-1997"
	lookup["ibm-1148_P100-1997"] = "ibm-1148_P100-1997"
	lookup["ibm-1148"] = "ibm-1148_P100-1997"
	lookup["IBM01148"] = "ibm-1148_P100-1997"
	lookup["CCSID01148"] = "ibm-1148_P100-1997"
	lookup["CP01148"] = "ibm-1148_P100-1997"
	lookup["cp1148"] = "ibm-1148_P100-1997"
	lookup["ebcdic-international-500+euro"] = "ibm-1148_P100-1997"

	lookup["ibm-1149_P100-1997"] = "ibm-1149_P100-1997"
	lookup["ibm-1149_P100-1997"] = "ibm-1149_P100-1997"
	lookup["ibm-1149"] = "ibm-1149_P100-1997"
	lookup["IBM01149"] = "ibm-1149_P100-1997"
	lookup["CCSID01149"] = "ibm-1149_P100-1997"
	lookup["CP01149"] = "ibm-1149_P100-1997"
	lookup["cp1149"] = "ibm-1149_P100-1997"
	lookup["ebcdic-is-871+euro"] = "ibm-1149_P100-1997"

	lookup["ibm-1153_P100-1999"] = "ibm-1153_P100-1999"
	lookup["ibm-1153_P100-1999"] = "ibm-1153_P100-1999"
	lookup["ibm-1153"] = "ibm-1153_P100-1999"

	lookup["ibm-1154_P100-1999"] = "ibm-1154_P100-1999"
	lookup["ibm-1154_P100-1999"] = "ibm-1154_P100-1999"
	lookup["ibm-1154"] = "ibm-1154_P100-1999"

	lookup["ibm-1155_P100-1999"] = "ibm-1155_P100-1999"
	lookup["ibm-1155_P100-1999"] = "ibm-1155_P100-1999"
	lookup["ibm-1155"] = "ibm-1155_P100-1999"

	lookup["ibm-1156_P100-1999"] = "ibm-1156_P100-1999"
	lookup["ibm-1156_P100-1999"] = "ibm-1156_P100-1999"
	lookup["ibm-1156"] = "ibm-1156_P100-1999"

	lookup["ibm-1157_P100-1999"] = "ibm-1157_P100-1999"
	lookup["ibm-1157_P100-1999"] = "ibm-1157_P100-1999"
	lookup["ibm-1157"] = "ibm-1157_P100-1999"

	lookup["ibm-1158_P100-1999"] = "ibm-1158_P100-1999"
	lookup["ibm-1158_P100-1999"] = "ibm-1158_P100-1999"
	lookup["ibm-1158"] = "ibm-1158_P100-1999"

	lookup["ibm-1160_P100-1999"] = "ibm-1160_P100-1999"
	lookup["ibm-1160_P100-1999"] = "ibm-1160_P100-1999"
	lookup["ibm-1160"] = "ibm-1160_P100-1999"

	lookup["ibm-1162_P100-1999"] = "ibm-1162_P100-1999"
	lookup["ibm-1162_P100-1999"] = "ibm-1162_P100-1999"
	lookup["ibm-1162"] = "ibm-1162_P100-1999"

	lookup["ibm-1164_P100-1999"] = "ibm-1164_P100-1999"
	lookup["ibm-1164_P100-1999"] = "ibm-1164_P100-1999"
	lookup["ibm-1164"] = "ibm-1164_P100-1999"

	lookup["ibm-4517_P100-2005"] = "ibm-4517_P100-2005"
	lookup["ibm-4517_P100-2005"] = "ibm-4517_P100-2005"
	lookup["ibm-4517"] = "ibm-4517_P100-2005"

	lookup["ibm-4899_P100-1998"] = "ibm-4899_P100-1998"
	lookup["ibm-4899_P100-1998"] = "ibm-4899_P100-1998"
	lookup["ibm-4899"] = "ibm-4899_P100-1998"

	lookup["ibm-4909_P100-1999"] = "ibm-4909_P100-1999"
	lookup["ibm-4909_P100-1999"] = "ibm-4909_P100-1999"
	lookup["ibm-4909"] = "ibm-4909_P100-1999"

	lookup["ibm-4971_P100-1999"] = "ibm-4971_P100-1999"
	lookup["ibm-4971_P100-1999"] = "ibm-4971_P100-1999"
	lookup["ibm-4971"] = "ibm-4971_P100-1999"

	lookup["ibm-5123_P100-1999"] = "ibm-5123_P100-1999"
	lookup["ibm-5123_P100-1999"] = "ibm-5123_P100-1999"
	lookup["ibm-5123"] = "ibm-5123_P100-1999"

	lookup["ibm-8482_P100-1999"] = "ibm-8482_P100-1999"
	lookup["ibm-8482_P100-1999"] = "ibm-8482_P100-1999"
	lookup["ibm-8482"] = "ibm-8482_P100-1999"

	lookup["ibm-9067_X100-2005"] = "ibm-9067_X100-2005"
	lookup["ibm-9067_X100-2005"] = "ibm-9067_X100-2005"
	lookup["ibm-9067"] = "ibm-9067_X100-2005"

	lookup["ibm-12712_P100-1998"] = "ibm-12712_P100-1998"
	lookup["ibm-12712_P100-1998"] = "ibm-12712_P100-1998"
	lookup["ibm-12712"] = "ibm-12712_P100-1998"
	lookup["ebcdic-he"] = "ibm-12712_P100-1998"

	lookup["ibm-16804_X110-1999"] = "ibm-16804_X110-1999"
	lookup["ibm-16804_X110-1999"] = "ibm-16804_X110-1999"
	lookup["ibm-16804"] = "ibm-16804_X110-1999"
	lookup["ebcdic-ar"] = "ibm-16804_X110-1999"

	lookup["KOI8-R"] = "KOI8-R"
	lookup["csKOI8R"] = "KOI8-R"

	lookup["KOI8-U"] = "KOI8-U"

	lookup["ibm-1051_P100-1995"] = "ibm-1051_P100-1995"
	lookup["ibm-1051_P100-1995"] = "ibm-1051_P100-1995"
	lookup["ibm-1051"] = "ibm-1051_P100-1995"
	lookup["hp-roman8"] = "ibm-1051_P100-1995"
	lookup["roman8"] = "ibm-1051_P100-1995"
	lookup["r8"] = "ibm-1051_P100-1995"
	lookup["csHPRoman8"] = "ibm-1051_P100-1995"

	lookup["ibm-1276_P100-1995"] = "ibm-1276_P100-1995"
	lookup["ibm-1276_P100-1995"] = "ibm-1276_P100-1995"
	lookup["ibm-1276"] = "ibm-1276_P100-1995"
	lookup["Adobe-Standard-Encoding"] = "ibm-1276_P100-1995"
	lookup["csAdobeStandardEncoding"] = "ibm-1276_P100-1995"

	lookup["EUC-JP"] = "EUC-JP"
	lookup["extended_unix_code_packed_format_for_japanese"] = "EUC-JP"
	lookup["cseucpkdfmtjapanese"] = "EUC-JP"

	lookup["Big5"] = "Big5"
	lookup["csBig5"] = "Big5"
	lookup["950"] = "Big5"

	lookup["Shift_JIS"] = "Shift_JIS"
	lookup["MS_Kanji"] = "Shift_JIS"
	lookup["csShiftJIS"] = "Shift_JIS"
	lookup["SJIS"] = "Shift_JIS"
	lookup["932"] = "Shift_JIS"

	// dbase file encodings
	// caution! I am not sure if this mapping are really true.
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

	// cross mapping between dbase file encodings and mahonia library
	encodingTable = make(map[string]byte)
	encodingTable["IBM437"] = 0x1b
	encodingTable["IBM850"] = 0x37
	encodingTable["windows-1252"] = 0x59
	//encodingTable[""] = 0x4
	encodingTable["ibm-865_P100-1995"] = 0x66
	encodingTable["IBM437"] = 0x1b
	encodingTable["IBM850"] = 0x37
	encodingTable["IBM437"] = 0x1b
	encodingTable["IBM437"] = 0x1b
	encodingTable["IBM850"] = 0x37
	encodingTable["IBM437"] = 0x1b
	encodingTable["IBM850"] = 0x37
	encodingTable["IBM437"] = 0x1b
	encodingTable["IBM850"] = 0x37
	encodingTable["Shift_JIS"] = 0x7b
	encodingTable["IBM850"] = 0x37
	encodingTable["IBM437"] = 0x1b
	encodingTable["IBM850"] = 0x37
	encodingTable["ibm-865_P100-1995"] = 0x66
	encodingTable["IBM437"] = 0x1b
	encodingTable["IBM437"] = 0x1b
	encodingTable["IBM850"] = 0x37
	encodingTable["IBM437"] = 0x1b
	encodingTable["ibm-863_P100-1995"] = 0x6c
	encodingTable["IBM850"] = 0x37
	encodingTable["IBM852"] = 0x87
	encodingTable["IBM852"] = 0x87
	encodingTable["IBM852"] = 0x87
	encodingTable["ibm-860_P100-1995"] = 0x24
	encodingTable["IBM850"] = 0x37
	encodingTable["IBM866"] = 0x65
	encodingTable["IBM850"] = 0x37
	encodingTable["IBM852"] = 0x87
	//encodingTable[""] = 0x7a
	//encodingTable[""] = 0x79
	encodingTable["Big5"] = 0x78
	encodingTable["windows-874"] = 0x7c
	encodingTable["windows-1252"] = 0x59
	encodingTable["windows-1252"] = 0x59
	encodingTable["IBM852"] = 0x87
	encodingTable["IBM866"] = 0x65
	encodingTable["ibm-865_P100-1995"] = 0x66
	encodingTable["ibm-861_P100-1995"] = 0x67
	//encodingTable[""] = 0x68
	//encodingTable[""] = 0x69
	//encodingTable[""] = 0x86
	encodingTable["ibm-857_P100-1995"] = 0x88
	encodingTable["ibm-863_P100-1995"] = 0x6c
	encodingTable["Big5"] = 0x78
	//encodingTable[""] = 0x79
	//encodingTable[""] = 0x7a
	encodingTable["Shift_JIS"] = 0x7b
	encodingTable["windows-874"] = 0x7c
	//encodingTable[""] = 0x86
	encodingTable["IBM852"] = 0x87
	encodingTable["ibm-857_P100-1995"] = 0x88
	//encodingTable[""] = 0x96
	//encodingTable[""] = 0x97
	//encodingTable[""] = 0x98
	encodingTable["windows-1250"] = 0xc8
	encodingTable["windows-1251"] = 0xc9
	encodingTable["windows-1254"] = 0xca
	encodingTable["windows-1253"] = 0xcb
	encodingTable["windows-1257"] = 0xcc
}
