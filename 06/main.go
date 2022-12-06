package main

import (
	"fmt"
	"container/ring"
)

var (
	use_samples = true
	sample1 = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	sample2 = "bvwbjplbgvbhsrlpgdmjqwftvncz"
	sample3 = "nppdvjthqldpwncqszvftbrmjlhg"
	sample4 = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	sample5 = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	puzzle = "htsslsmstsrrhlrrllqppfnpnzzqtqtjqttslttmvmsmbbnpbbznzjjmrmsrrnjjzczcfcqchcnnhrnrzrnrzzddtrrpjrprbpbwbswsqswqswsqsqffgngtgwtwbtbbhhslsrsmrmffgcgtcgcppnbpnpbnpptcptpltthjhwwlttrrlldzldzzzlssccrfrqfrfdfldfldfdzfdzzcnznwznzqqzpqzqdzzbbslbljlqjjvzjzgzqgqqqvsvffzjfzfmzzrjrhjhrhcrhhtgtmggchggcsggvtttwmmsspmpffzpfpjjnwwnpwwdttcfcmmlblwlvvqrqddcdwcwnnfqnqdnnncggflfdftddtftqthqhrrmsrmsrsffbccjnnjjgddppjmmldllhttqvqzvzrvrsslnlplrprtprtttnvnsvsvzzdndrnnnlznzcnnzwwzjjsnjnwjwqqczqcqwqppqnqqllgblbhhbbzbjzzwjzzrqzqggdppgcgwcclglhhwwgcccgcvvpbbrzzwszsqqjqrrczzvmvpprllsgstsnnnrvvsszgzddzttbccmjccwtccrbccnssdqsdsqqfgqfqbfqfjqqzqqmhmwhwffsjjcjffqmmwjwhjjqttpvvgddzjdddwndwdjdldgdsdmdzmdzmzqqszzcwwwjbbvvlplnpndnddjcjcttrgrggjvjsjvjssnsmshmhrmhhfvvdqqppmbbjnbnzzzwlzzfmmlttvptvpphlhwhzzwffcsfsppdssvsvgsgcscmmdwwpggtcgttmlmglgmmcsccmmgffnvvvvsscstctptzzcmcsclcffnwwndnwddmgmfmvffntfnndgdrrbmrrvprrjnjcnjjsgjgmmjjjqwqhhlfljfljlclttrgttmdmhhcncwncnmccnmnqqmhmwmdwdmdfdttnggcscfsftttcjchhqghhtdthtwtgwgttlffphfhtftpfttlgtltblljmljjsmsszgsghhnnmvnmnvmvmhvvtmmdqdppclplqljllzppwgpgvpvjpvvlwldlcdcmvfrbcqfdrgqjdmpvczgmtrbhtfgbctjghwfmpqtprlnvzvbmhnmgsnbcqpdqhtstncmtrvwmjzzmsbprfvmszgvdwjwchzcvmzncfblffwvwwrdqpfnvrtlftnvhcqpwfgzdnmpqthblmqzhbdbdhmqvscfrpbjbgjgdndhlnnpcrrmrlghcwvmmmhbvdsrztbpdvhwhsphmgblzfpflhnbrvjsgsbvgcfwdlbmbsmbqfwczhhbbwntsczhljgvdwpsfmjppmdwcfchhvtrwlfgqrzbtzndqwbwqflpthhgtzrfmspfnlrrvrjttwnzzfsbzlvfwlntvcqhfntllsrdgqjdbrbvnmjgfsqshmcgtsmlrzjmmqbpfdfhlghqjnjrcnzvvmbwhwfwpbbsmpdmvvznstqpjwmthncjsvfmbhwtbfqgfmvhhwrwzrzccbmmrngfbcqgsdjlfwrlzqwrbhvhmjzqqtgqvjwgllhhgphjthwnscndjfhrdzjctthwfszztvtzcnvsnfdwnvrrbjzplgnghfsvnflhctjntrpzfgjnfzlmwzrvqcfdqqscrffjshhgmvlgdtqpbllfdspdhcccqbmpqcgctljwzdszbfzsltfmdrcvmjqmpvzzhnczlfdbrmsvfffmfzjvnthghqgtnvjtflzrrnrhqshtzslbmbghvtqcmbdgzcmmrpbtqlzbzjhhgfcbrbcztlhdcfqccqsgdqdfrvlsprbhwhtmjfbzlzfpcslcddzdhbhstvqjgdfwdwzfhrzpthnzgmzrrbrszvcwhbdczzcnqsqrwdpmgdmcdmccmfdwdfmzpbbsdtjlmlvcnvhhgcltwmpzqmjlpbsbfrrntrbmrrmgqrdsbfcgnqvnmwwgnmvvpjzpwbvwfzzfcfrbgdshrvbpdlbpzdsvjnrsmfttqgbngjbqcnhhhmrnlgwjnlfcsdwccqqlzcdrcqnjlmsfqldgdmwlctstcvqbvvwvhvrnhvtgpcnsstnhvlttdnstlndnpdglqlbgggqlwqpfztgzvhqqwbctvgtrmrbpmvwlztsfbrmhhpfdtnsjcjmngsqtzvhqjnwmgcrjbhvghpnrjlrhjfrfrmbzvprpzlcshbqtlnlmqmfhsmbznjmpzwljccngbsvqqmghqsqwwmtqhgddnpcmmlfvmplgptzmfsmpntprnwlpjmsdsntpmpnhwlqgfwslrnjlmvhvqsnlcssvrqvtvfhhghbhnvmpngtdmrcmftjmwgnptbmjcwvrqqbczmtdprlltcjcggvffvjwdthbhvjvhbsmfnstqvczmvsgsqfjmddszzrwnbrbvlhgltjddczwpzpsgfgbtmgmssbszhjfbprpbgsqpvgwfnmpcdwfzwbfbtfwhjbjgrctcwqgfhpvjpjssmcpppwpbnlfsrmbzqdpmpmqjtzqmqswpfhfvfltwnfbltvdmfgmpbhzzdrbtnmmjfqgtrgmrbsgplfrmlnjrggtslngnphcvcwvqsdftlhddhpsdwlhzcrfwgwbndplwjmtrltmswwdqjpgmmcdchllzdgpfpfrvqrmvppvzcrvbswzdclnqnqvmfvjtpzpvzlmbngtrrjwpgqjgrrqbqwcvjlgqzddtcfmmrqtnjppqvztfpdrdmjpsqqvzrhwjgvdwwdtmrzrvhwgsqvjjrtsdtwfwhrbcsqdblhgwzghrqrqtfbzszrmjqbhrtsrcfjlzbcjmdnpthrtbhtzmhgpfcqndwpmtlvzschzzcqdnzdrfhczsrscvttlpslrzgvwprwppmpfjqwhtnhzcdjjwmsnvqzmtlzsdpbdtdtsmbmjszsglrldcsgtnmbgsjfmnrftnmvnjtswrmdthvstmdlsnsmrbqhdlpnmnhjhcccgnzrrsljvwswmhjqrrfbwwhgrcttsdcjsdtlgmgblfnvwmgztbbzlbrdnspfnvwvhzlztdhlzhwwppgwwvrhfbjvrmjpjflqzbdnlvptmrsggqmzgmzlsdzbfqnqzzdnzfgmhncvwmgrcfrmlnwzcwdsvtwbvcqpmgcczbnnfdrlsgnfgnpbdstdmwhlprnzvhjpznsgwhfhzfwmjsvcbwccpsnfgbzrltcbwczcmzwljcswgwjnhwtjrjvgrggsrqszscrjcghnwdzpzhttrjcgwvmtcwmrhvwptlgfjpdjpnhphmzdgfdvsncswbdvjdtsgdtlsgjljlczgrwbswtbwrcfpmhgfjfnhpsmfrrtcjmpvvscmgftpprdmbjwcvhzbrmpbvhgzwftwtvvhrjmljgjhbpnbnfntmjhvjrzwlqbtqlrfbblgmfsbgzhhscgwgzrwvflnfctngtwbvdbbgspntclnbwgpppjgbrlqvtfznpdttrsplvjfsgbjwjprbvdfvjtffbhsjwgbsfschfnmlqdfzmdwzvfctjjvzncdvrdttlpvpgvsssflpfnrzgfznzlldrbgnnztngtlbbrmrmlfnnspsvvsfzfbmsblmzdwqcrftvnbvdlhgdqjtglclpzrchtlffrfwslqjvbfpvnhmdgqrcjtbhmjqwqzpfndqbwldbzmwqsptdccczmhwmdqqzcqvmbbnmqndtspmbtggcghdhsfgrjgvwjdbwmltbhdvdtgqnpqhwmhzpzbqjtnlftqrjbsvtwtgpmvpnfwdtsjgdtfnlntpgrmwphphrrzhbdrbzhtqddwvptdllntjzldzrndhfjwdfnmtjmjtfmndbvlwgmlcmwwlpfdjwbfznbllbmqrlmvljngvpmdlqmvdvwvpqwlbqssqcbnrmhdvrzwljstghmhntwsbqmnlpgthbwmrznbbggthtjhnndnqbzmcrtftcbnpctqjzghdfcvmmvpqwtnntstlspfsgwfdbrlsbwgbhhbfcvwrjclsgmmbqmmjwtdjqppjvcbnbvfwczlqzbtnlhzhssglgnlm"
)

func contains(items []byte, item byte) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}

	return false
}

func getMarkerPosition(packet string, numberOfUniqueItems int) int {
	var checked_markers = ring.New(numberOfUniqueItems)

	for idx := 0; idx < numberOfUniqueItems; idx++ {
		checked_markers.Value = packet[idx]
		checked_markers = checked_markers.Next()
	}

	for idx := numberOfUniqueItems; idx < len(packet); idx ++ {
		is_unique := true
		current_marker := packet[idx]
		items := []byte{}
		checked_markers.Do(func(p any) {
			is_unique = is_unique && !contains(items, p.(byte))
			if (is_unique) {
				items = append(items, p.(byte))
			}
		})

		if (is_unique) {
			return idx
		}

		checked_markers.Value = current_marker
		checked_markers = checked_markers.Next()
	}

	return -1
}

func getPacketMarkerPosition(packet string) int {
	return getMarkerPosition(packet, 4)
}

func getMessageMarkerPosition(packet string) int {
	return getMarkerPosition(packet, 14)
}

func main() {
	var markerPosition = getPacketMarkerPosition(sample1)
	if (use_samples) {
		fmt.Printf("%s -> %d\n", sample1, markerPosition)
		markerPosition = getPacketMarkerPosition(sample2)
		fmt.Printf("%s -> %d\n", sample2, markerPosition)
		markerPosition = getPacketMarkerPosition(sample3)
		fmt.Printf("%s -> %d\n", sample3, markerPosition)
		markerPosition = getPacketMarkerPosition(sample4)
		fmt.Printf("%s -> %d\n", sample4, markerPosition)
		markerPosition = getPacketMarkerPosition(sample5)
		fmt.Printf("%s -> %d\n", sample5, markerPosition)
	}
	markerPosition = getPacketMarkerPosition(puzzle)
	fmt.Printf("%s -> %d\n", "puzzle", markerPosition)

	fmt.Println("------------------------------------------")

	markerPosition = getMessageMarkerPosition(sample1)
	if (use_samples) {
		fmt.Printf("%s -> %d\n", sample1, markerPosition)
		markerPosition = getMessageMarkerPosition(sample2)
		fmt.Printf("%s -> %d\n", sample2, markerPosition)
		markerPosition = getMessageMarkerPosition(sample3)
		fmt.Printf("%s -> %d\n", sample3, markerPosition)
		markerPosition = getMessageMarkerPosition(sample4)
		fmt.Printf("%s -> %d\n", sample4, markerPosition)
		markerPosition = getMessageMarkerPosition(sample5)
		fmt.Printf("%s -> %d\n", sample5, markerPosition)
	}
	markerPosition = getMessageMarkerPosition(puzzle)
	fmt.Printf("%s -> %d\n", "puzzle", markerPosition)
}