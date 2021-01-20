package solutions

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var reorganizeStringSuccessCaseTestCases = []struct {
	input string
}{
	{"aab"},
	{"baaba"},
	{"cxmwmmm"},
	{"abbabbaaab"},
	{"tndsewnllhrtwsvxenkscbivijfqnysamckzoyfnapuotmdexzkkrpmppttficzerdndssuveompqkemtbwbodrhwsfpbmkafpwyedpcowruntvymxtyyejqtajkcjakghtdwmuygecjncxzcxezgecrxonnszmqmecgvqqkdagvaaucewelchsmebikscciegzoiamovdojrmmwgbxeygibxxltemfgpogjkhobmhwquizuwvhfaiavsxhiknysdghcawcrphaykyashchyomklvghkyabxatmrkmrfsppfhgrwywtlxebgzmevefcqquvhvgounldxkdzndwybxhtycmlybhaaqvodntsvfhwcuhvuccwcsxelafyzushjhfyklvghpfvknprfouevsxmcuhiiiewcluehpmzrjzffnrptwbuhnyahrbzqvirvmffbxvrmynfcnupnukayjghpusewdwrbkhvjnveuiionefmnfxao"},
}

var reorganizeStringFailCaseTestCases = []struct {
	input string
}{
	{"bbbbbbb"},
}

func Test_reorganizeString_SuccessCase(t *testing.T) {
	for _, testCase := range reorganizeStringSuccessCaseTestCases {
		t.Run("reorganizeStringSuccessCase", func(t *testing.T) {
			Convey("Given some random string", t, func() {
				input := testCase.input

				Convey("When pass the random string to reorganizeString function", func() {
					result := reorganizeString(input)

					Convey("Then the result string shouldn't have same Letter adjacent to each other", func() {
						So(len(result), ShouldEqual, len(input))

						for i := 0; i < len(result)-1; i++ {
							So(result[i], ShouldNotEqual, result[i+1])
						}
					})
				})
			})
		})
	}
}

func Test_reorganizeString_FailCase(t *testing.T) {
	for _, testCase := range reorganizeStringFailCaseTestCases {
		t.Run("reorganizeStringFailCase", func(t *testing.T) {
			Convey("Given some random string", t, func() {
				input := testCase.input

				Convey("When pass the random string which cannot reorganize to match the rule to reorganizeString function", func() {
					result := reorganizeString(input)

					Convey("Then the result should be empty string", func() {
						So(result, ShouldEqual, "")
					})
				})
			})
		})
	}
}

func Test_categorizeNumberOfEachRuneIn(t *testing.T) {
	Convey("Given string", t, func() {
		someString := "cxmwmmm"
		runes := []rune(someString)

		Convey("When categorize number of each rune", func() {
			categorize := categorizeNumberOfEachRuneIn(runes)

			Convey("Then the categorize should ", func() {
				So(categorize['c'], ShouldEqual, 1)
				So(categorize['x'], ShouldEqual, 1)
				So(categorize['w'], ShouldEqual, 1)
				So(categorize['m'], ShouldEqual, 4)
			})
		})
	})
}

func Test_createLetterInStringInfosFrom(t *testing.T) {
	Convey("Given map[rune]int", t, func() {
		someString := "cxmwwmmmwc"
		runes := []rune(someString)

		categorize := categorizeNumberOfEachRuneIn(runes)

		Convey("When create letterInStringInfos", func() {

			result := createLetterInStringInfosFrom(categorize)

			Convey("Then it should create slice with LetterInfo", func() {

				for _, letterInfo := range result {
					switch letterInfo.Letter {
					case 'x':
						So(letterInfo.NumberOfLetter, ShouldEqual, 1)
					case 'w':
						So(letterInfo.NumberOfLetter, ShouldEqual, 3)
					case 'm':
						So(letterInfo.NumberOfLetter, ShouldEqual, 4)
					case 'c':
						So(letterInfo.NumberOfLetter, ShouldEqual, 2)
					}
				}
			})
		})
	})
}

func Test_descSort(t *testing.T) {
	Convey("Given letterInStringInfos", t, func() {
		someString := "cxmwwmmmwc"
		runes := []rune(someString)

		categorize := categorizeNumberOfEachRuneIn(runes)
		letterInfos := createLetterInStringInfosFrom(categorize)

		Convey("When descSort letterInfos", func() {

			sortedLetterInfos := descSort(letterInfos)

			Convey("Then it should create slice with LetterInfo", func() {

				So(sortedLetterInfos[0].Letter, ShouldEqual, 'm')
				So(sortedLetterInfos[1].Letter, ShouldEqual, 'w')
				So(sortedLetterInfos[2].Letter, ShouldEqual, 'c')
				So(sortedLetterInfos[3].Letter, ShouldEqual, 'x')
			})
		})
	})
}
