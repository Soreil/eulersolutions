package main

import (
	"fmt"
	"math/big"
)

func main() {
	// The Scan function is rarely used directly;
	// the fmt package recognizes it as an implementation of fmt.Scanner.
	numbers := make([]big.Int, 100)
	fmt.Sscan("37107287533902102798797998220837590246510135740250", &numbers[0])
	fmt.Sscan("46376937677490009712648124896970078050417018260538", &numbers[1])
	fmt.Sscan("74324986199524741059474233309513058123726617309629", &numbers[2])
	fmt.Sscan("91942213363574161572522430563301811072406154908250", &numbers[3])
	fmt.Sscan("23067588207539346171171980310421047513778063246676", &numbers[4])
	fmt.Sscan("89261670696623633820136378418383684178734361726757", &numbers[5])
	fmt.Sscan("28112879812849979408065481931592621691275889832738", &numbers[6])
	fmt.Sscan("44274228917432520321923589422876796487670272189318", &numbers[7])
	fmt.Sscan("47451445736001306439091167216856844588711603153276", &numbers[8])
	fmt.Sscan("70386486105843025439939619828917593665686757934951", &numbers[9])
	fmt.Sscan("62176457141856560629502157223196586755079324193331", &numbers[10])
	fmt.Sscan("64906352462741904929101432445813822663347944758178", &numbers[11])
	fmt.Sscan("92575867718337217661963751590579239728245598838407", &numbers[12])
	fmt.Sscan("58203565325359399008402633568948830189458628227828", &numbers[13])
	fmt.Sscan("80181199384826282014278194139940567587151170094390", &numbers[14])
	fmt.Sscan("35398664372827112653829987240784473053190104293586", &numbers[15])
	fmt.Sscan("86515506006295864861532075273371959191420517255829", &numbers[16])
	fmt.Sscan("71693888707715466499115593487603532921714970056938", &numbers[17])
	fmt.Sscan("54370070576826684624621495650076471787294438377604", &numbers[18])
	fmt.Sscan("53282654108756828443191190634694037855217779295145", &numbers[19])
	fmt.Sscan("36123272525000296071075082563815656710885258350721", &numbers[20])
	fmt.Sscan("45876576172410976447339110607218265236877223636045", &numbers[21])
	fmt.Sscan("17423706905851860660448207621209813287860733969412", &numbers[22])
	fmt.Sscan("81142660418086830619328460811191061556940512689692", &numbers[23])
	fmt.Sscan("51934325451728388641918047049293215058642563049483", &numbers[24])
	fmt.Sscan("62467221648435076201727918039944693004732956340691", &numbers[25])
	fmt.Sscan("15732444386908125794514089057706229429197107928209", &numbers[26])
	fmt.Sscan("55037687525678773091862540744969844508330393682126", &numbers[27])
	fmt.Sscan("18336384825330154686196124348767681297534375946515", &numbers[28])
	fmt.Sscan("80386287592878490201521685554828717201219257766954", &numbers[29])
	fmt.Sscan("78182833757993103614740356856449095527097864797581", &numbers[30])
	fmt.Sscan("16726320100436897842553539920931837441497806860984", &numbers[31])
	fmt.Sscan("48403098129077791799088218795327364475675590848030", &numbers[32])
	fmt.Sscan("87086987551392711854517078544161852424320693150332", &numbers[33])
	fmt.Sscan("59959406895756536782107074926966537676326235447210", &numbers[34])
	fmt.Sscan("69793950679652694742597709739166693763042633987085", &numbers[35])
	fmt.Sscan("41052684708299085211399427365734116182760315001271", &numbers[36])
	fmt.Sscan("65378607361501080857009149939512557028198746004375", &numbers[37])
	fmt.Sscan("35829035317434717326932123578154982629742552737307", &numbers[38])
	fmt.Sscan("94953759765105305946966067683156574377167401875275", &numbers[39])
	fmt.Sscan("88902802571733229619176668713819931811048770190271", &numbers[40])
	fmt.Sscan("25267680276078003013678680992525463401061632866526", &numbers[41])
	fmt.Sscan("36270218540497705585629946580636237993140746255962", &numbers[42])
	fmt.Sscan("24074486908231174977792365466257246923322810917141", &numbers[43])
	fmt.Sscan("91430288197103288597806669760892938638285025333403", &numbers[44])
	fmt.Sscan("34413065578016127815921815005561868836468420090470", &numbers[45])
	fmt.Sscan("23053081172816430487623791969842487255036638784583", &numbers[46])
	fmt.Sscan("11487696932154902810424020138335124462181441773470", &numbers[47])
	fmt.Sscan("63783299490636259666498587618221225225512486764533", &numbers[48])
	fmt.Sscan("67720186971698544312419572409913959008952310058822", &numbers[49])
	fmt.Sscan("95548255300263520781532296796249481641953868218774", &numbers[50])
	fmt.Sscan("76085327132285723110424803456124867697064507995236", &numbers[51])
	fmt.Sscan("37774242535411291684276865538926205024910326572967", &numbers[52])
	fmt.Sscan("23701913275725675285653248258265463092207058596522", &numbers[53])
	fmt.Sscan("29798860272258331913126375147341994889534765745501", &numbers[54])
	fmt.Sscan("18495701454879288984856827726077713721403798879715", &numbers[55])
	fmt.Sscan("38298203783031473527721580348144513491373226651381", &numbers[56])
	fmt.Sscan("34829543829199918180278916522431027392251122869539", &numbers[57])
	fmt.Sscan("40957953066405232632538044100059654939159879593635", &numbers[58])
	fmt.Sscan("29746152185502371307642255121183693803580388584903", &numbers[59])
	fmt.Sscan("41698116222072977186158236678424689157993532961922", &numbers[60])
	fmt.Sscan("62467957194401269043877107275048102390895523597457", &numbers[61])
	fmt.Sscan("23189706772547915061505504953922979530901129967519", &numbers[62])
	fmt.Sscan("86188088225875314529584099251203829009407770775672", &numbers[63])
	fmt.Sscan("11306739708304724483816533873502340845647058077308", &numbers[64])
	fmt.Sscan("82959174767140363198008187129011875491310547126581", &numbers[65])
	fmt.Sscan("97623331044818386269515456334926366572897563400500", &numbers[66])
	fmt.Sscan("42846280183517070527831839425882145521227251250327", &numbers[67])
	fmt.Sscan("55121603546981200581762165212827652751691296897789", &numbers[68])
	fmt.Sscan("32238195734329339946437501907836945765883352399886", &numbers[69])
	fmt.Sscan("75506164965184775180738168837861091527357929701337", &numbers[70])
	fmt.Sscan("62177842752192623401942399639168044983993173312731", &numbers[71])
	fmt.Sscan("32924185707147349566916674687634660915035914677504", &numbers[72])
	fmt.Sscan("99518671430235219628894890102423325116913619626622", &numbers[73])
	fmt.Sscan("73267460800591547471830798392868535206946944540724", &numbers[74])
	fmt.Sscan("76841822524674417161514036427982273348055556214818", &numbers[75])
	fmt.Sscan("97142617910342598647204516893989422179826088076852", &numbers[76])
	fmt.Sscan("87783646182799346313767754307809363333018982642090", &numbers[77])
	fmt.Sscan("10848802521674670883215120185883543223812876952786", &numbers[78])
	fmt.Sscan("71329612474782464538636993009049310363619763878039", &numbers[79])
	fmt.Sscan("62184073572399794223406235393808339651327408011116", &numbers[80])
	fmt.Sscan("66627891981488087797941876876144230030984490851411", &numbers[81])
	fmt.Sscan("60661826293682836764744779239180335110989069790714", &numbers[82])
	fmt.Sscan("85786944089552990653640447425576083659976645795096", &numbers[83])
	fmt.Sscan("66024396409905389607120198219976047599490197230297", &numbers[84])
	fmt.Sscan("64913982680032973156037120041377903785566085089252", &numbers[85])
	fmt.Sscan("16730939319872750275468906903707539413042652315011", &numbers[86])
	fmt.Sscan("94809377245048795150954100921645863754710598436791", &numbers[87])
	fmt.Sscan("78639167021187492431995700641917969777599028300699", &numbers[88])
	fmt.Sscan("15368713711936614952811305876380278410754449733078", &numbers[89])
	fmt.Sscan("40789923115535562561142322423255033685442488917353", &numbers[90])
	fmt.Sscan("44889911501440648020369068063960672322193204149535", &numbers[91])
	fmt.Sscan("41503128880339536053299340368006977710650566631954", &numbers[92])
	fmt.Sscan("81234880673210146739058568557934581403627822703280", &numbers[93])
	fmt.Sscan("82616570773948327592232845941706525094512325230608", &numbers[94])
	fmt.Sscan("22918802058777319719839450180888072429661980811197", &numbers[95])
	fmt.Sscan("77158542502016545090413245809786882778948721859617", &numbers[96])
	fmt.Sscan("72107838435069186155435662884062257473692284509516", &numbers[97])
	fmt.Sscan("20849603980134001723930671666823555245252804609722", &numbers[98])
	fmt.Sscan("53503534226472524250874054075591789781264330331690", &numbers[99])
	result := new(big.Int)
	for i := range numbers {
		result.Add(result, &numbers[i])
	}
	fmt.Println(result)
}
