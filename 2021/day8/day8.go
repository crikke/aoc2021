package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	sum1, sum2 := day8(input)
	fmt.Printf("part1: %d \n", sum1)
	fmt.Printf("part2: %d \n", sum2)
}

func day8(input string) (sum1, sum2 int) {
	letters := make(map[rune]int)

	bit := 0
	for i := 'a'; i <= 'g'; i++ {
		letters[i] = int(math.Pow(2, float64(bit)))
		bit++
	}
	for _, line := range strings.Split(input, "\n") {
		set := make(map[int]int)

		pair := strings.Split(line, "| ")
		output := [4]int{}
		for _, w := range strings.Split(pair[0], " ") {

			if w == " " {
				continue
			}

			num := 0
			for _, r := range w {
				num ^= letters[r]
			}

			set[num] = 0

		}
		for i, w := range strings.Split(pair[1], " ") {

			if w == "" {
				continue
			}

			num := 0
			for _, r := range w {
				num ^= letters[r]
			}

			output[i] = num
			set[num]++
		}

		nums := [10]int{}
		idx := 0
		for bits := range set {
			if bits == 0 {
				continue
			}
			nums[idx] = bits
			idx++
		}
		numbers := decodeSegment(nums)
		res := make(map[int]int)

		for k, v := range set {

			res[numbers[k]] = v
		}
		sum1 += res[1] + res[4] + res[7] + res[8]

		val := 0
		for i := 0; i < len(output); i++ {

			val += int(numbers[output[i]] * int(math.Pow10(len(output)-1-i)))
		}
		sum2 += val
	}
	return
}

// convert letters to binary numbers
// a: 001
// b: 010
// c: 100
func decodeSegment(numbers [10]int) map[int]int {

	one := getWithSize(numbers, 2)[0]
	four := getWithSize(numbers, 4)[0]
	seven := getWithSize(numbers, 3)[0]
	eight := getWithSize(numbers, 7)[0]

	result := make(map[int]int)

	result[one] = 1
	result[four] = 4
	result[seven] = 7
	result[eight] = 8

	var segments [7]int

	segments[0] = seven ^ one

	// find 3
	three := 0
	for _, num := range getWithSize(numbers, 5) {

		remaining := num ^ eight

		// the num is 3 if it does not contain one

		if remaining&one == 0 {
			three = num
		}
	}
	result[three] = 3

	// now segment 1,3,4 can now be calculated
	segments[3] = (three ^ one) & four
	segments[1] = (three ^ eight) & four
	segments[4] = (three ^ eight) ^ segments[1]

	// get 2 and 5
	for _, num := range getWithSize(numbers, 5) {
		if num == three {
			continue
		}

		if num&segments[1] == segments[1] {
			result[num] = 5
		} else {
			result[num] = 2
		}
	}

	for _, num := range getWithSize(numbers, 6) {

		if num&segments[3] == 0 {
			result[num] = 0
		} else if num&segments[4] == 0 {
			result[num] = 9
		} else {
			result[num] = 6
		}
	}

	return result
}

func getWithSize(numbers [10]int, size int) []int {
	var res []int

	for _, num := range numbers {

		n := num
		count := 0
		for n > 0 {
			count += n & 1
			n >>= 1
		}
		if size == count {
			res = append(res, num)
		}
	}
	return res
}

var input = `fdeba beagfd gbafe dagb dbf ecfad bd dgcaefb fbecgd abfecg | dgba dfb ecadf bdf
dgfc egdbaf afgcbe eafdbcg bcgad bdg fdbacg gd ecabd bafgc | gd gbd dg febgac
dfebc bceadf ebacf gdceb fd dfea cdgfba afbgec dfb fdaecbg | cgafbde efbagc faed gecdb
dbaef fabgd fcaedb cgadf aefbgd cabefg gfb bceagdf bg ebdg | gdbfae geafdbc bafgdce dbgfa
aebgd adg dg dgcabef facedg cfaebd dfbega begca bfgd dbafe | abdecf eafbd fdgb ebdfgca
bae febac abcgf eabfgd cbagfd gefcba abedfcg efcda eb becg | fcgab gafbed aedcfbg eb
edbcg ba afcdeg gbaedf gcabdf eabf bagde bga deabgcf edfga | aebf ab bdfecag abfe
dbae dgfec egfbd fbgaed feb fcabdg be bfgda facebgd baefcg | bagecf efb eb fbcgad
cdagbef cd bdgcf efcbg adbgf cdga egdabf acbdfg cdf caebfd | fbdga fcd fcd adcg
abceg dbgca acfd ebagfdc dbecfg gdfbea dc dcg abfdcg gfbad | cdg dbcagf acdf gbcdef
bcdag acdgbef ebcdag facd dfgbe cfg bfagdc fc fbcgd bgfaec | fc dafbcg fc cbdgae
ca feacgb cedfgb bgaed gbaec acefdgb bfcge eafgcd acg fbac | ca cag gac becga
gbdcfa afdcb cebd fadce cbgfdea aec cdebfa ce caegfb daegf | aedfbgc eca eca cebd
cb dcb fdbegc bgaed fcba dacgfe efcadb dceaf eacgfbd dbaec | cbd cb bagde dcb
debgf gaecdfb gcbef cbgaf fbecag bce bgfcda ce aefc ebdgca | ebcfgad ec gefcba bec
cdbfega ae cedaf bedgcf ebac efadbg afcedb afe fgadc cdfbe | fabedcg cafbed efa badgcef
aceg bcafgde deacgb fdcgab dacgb eafbcd dgcbe bdfge bce ce | dbfge bdcafg fbdcga ebdfg
efcbda fcgbe efdgc efgdcab cagbe cbf bagf ecagdb bacefg fb | fgab aecgb acfgbe agecfbd
bcgd egafb adgbcf cafged cegadfb gd dgf dbfag dfeacb dcabf | adbcf dg gd dgf
bgeca fgcbe fcgd egfbd bcgdfe fc agbdef ecf dafgecb bdafec | fc febdag cgfd dfcbgae
agb dbfa gbafcde afgcbe edfgb bagde cfdegb gbdefa ba caged | bdfa bafd cfdgeba agb
bdf baefcdg df cafbgd bdcgef bcdeag bcgad cadf gfbad fbeag | ebacfgd facd cdgfbe df
edagc egbcfad bcfdge agbdcf fbdage dcabg bga bgfdc fbca ba | cebgfda gdcfabe gfcdba bag
bgaf af bcdfge ecfagd fgcdb beacd gdecbfa bcafd fac bfdgac | bagcedf fbedgc ebfgdc gbfa
fagbdc fagdeb cebgf cbd ecagdb cade dagcebf dc becdg geabd | gfbadec cdae dbc bcefg
dgfbeac abfge dcefag cga cgbd baced gc adfebc ecbag aedgbc | bfcaged gac gebaf egdafc
cfdbe afge fa gdfbeca agfcbd abedgc fac cafbe ecfgab acgeb | adcfgeb bcadfg ebfac cfa
gebdcf cdbga eag ae dagecbf bgdfe ebgfad gabed dagfec feab | fadgecb fgbced dgcab ega
gf gcbaf gfdc agf eafcb cadebfg cgfabd dbagc agdcbe egfbda | gf beafc fga fcgd
cbed ebgad gadbfe bc bcfgea gebacd gceadfb cgfad cbg abcgd | bacdg afbegdc bgc bcagdfe
fdg dg acgfdbe gafdec cfbdg ebfgc gdeb bdcgef fbegca fbdca | fdg edbg edbg cgfbe
gfcbaed gfceda cadbg gafdb ca adc fbca dfcabg cedbg aegfdb | ca becadfg dcefga ca
agfcbd edcabg cdf cdgfabe debgcf fdgca bfac gdcab cf gedaf | fc cdf cf bcagdef
cdagb bae ea cdfebg febacd egdfb gdacfbe feag aedfbg gdabe | ea gefa agbde afge
dab cgaedb cafbg faced bdfe bd gfdeca dafgbec fcbda cadbef | db edfb bd dcafe
bde cbaed dcbf ecbag fcbdae acgfde dgafebc db fbdgea dacfe | cadbe bcdf afebdg dcgfea
ag fgaecd egcbdaf ebacd gad dagbfe cfga fcegd dcbgfe gaced | dag ag gda gacf
fdaegb dcgefa bfdec geca ge fge cdbafg decbgaf ecdfg cagdf | gdefc cgea dfacgb gedbaf
ecgfab cgbdea dgabf dcfb bd dgcabf fgdeacb dba efagd cfabg | dab adb bfdc fcdbag
dfcae cfegbd ade cdeabg egacdf gcfde da becfadg agfd efcba | ade aed dfga ade
fdgeab adcbgfe edfca cebad gdbace dgcb beafgc bac cb gdeba | bdcea bgdc dagcfeb abc
cabdge fdgae fbgc fedgcb gfebacd dfbeg bf bdf ecgbd adecfb | fb aegfbcd bdf fbacged
dfbgec feacgdb ca adgec cda ecba gdaef degbc abfdgc gdceab | ca dabfegc ca acdeg
abd fbagec gfdbea dagc ad fcdbega bafcdg facbg ecbdf fbcad | edfagb bcedf dba fbdcag
abf ebfdcg fgcdbae faegc bdgacf bfdec bade ceafdb ba eacbf | fab fbdaecg dcabfg ba
ebcgf aedcb gbaecfd gd cgdeb cadfeb edga bdfgca cbgead cdg | adge fgbec cgd dgbcfa
gdcaef dcgbe fbcdage bacge aec ea bfea bcafg aegbcf cfdgab | bfea gfecadb gbafc gedafcb
facg aebfg ecfba bfaecd fbdge badcfeg cfgeba ega cadebg ag | fgca bdgef gafc acgf
dacefg eb bdecf acbfge ecb bfcda decgf cgdebaf cbfgde gbde | bdge ecbfgd feacbgd efdabcg
bd fbd bfcgde egbaf dafgceb cefgda fcbdae dfgeb gcdb cgfde | dfbceag bd fgbdec gdbfec
gbdacf dacebf adgfbe aedfb acgbedf eacf ca dac aedcb bedgc | ca cad fcbdaeg bdagcf
cbd egdba ebfdga efadc gcdfabe bc decbga dbaec gfbecd bcag | cdb egdba acbg acgb
eacb fdeab fgcade cbdaf dea cfebgda ae bdegf cebdaf fgdcab | ae facbdg bdfecga baec
cfd adbce cf bcdafg fcedb cbedfag fdgbe faegdb fceg gdcbef | fegc cf gfdabe adbcgf
gfbdeac efdgac gcfb edbagf cebda cafeb abf cgfeab bf gacfe | abfce cadefgb abf abfce
dcfge fabgec beg abdg gbaedcf gcdbe bg ebcad cfdaeb ebagcd | bgdaec gbe gb gb
aecb daefbg ebdfc aedfcg efbad cdgabfe cfgdb fdbeca ce cfe | fce ce aceb ec
gadebf fg fga acbgfe agbcd deabf bdagf cedbfa dbafceg dfeg | fg afg gaf fg
afed cdbge dag gabfe fagbde gadbcf adebg ad agfecb gedfabc | da geabf fgcbad gad
de gcdeaf agbcedf efacg gdcea efgacb edc fegd adcgb daecbf | edc ecd de dgceaf
becadgf edagfc fcg edbagf cg adfgcb dfega acgfe gecd bfeca | cgf cg gcde cg
cfaed agcedf db dfcbe bcd dbefca dcbgfa gbecf dcaegfb bdea | dagbcef bcd fegcb eabd
fcdagb ca aedfcgb fca gafbd abegfd afbcg fbgce aefcbd acdg | acf adcg efbgc ca
da dcefgba cdgef egcabf baed fagbcd fadbge dag faebg agedf | fcegab abde gad dbgfae
fbdec cdaeb dac dbage egca ecbdgaf debcga ca abdgfe dagcbf | efbcd adgefbc gbefacd ca
bfdecg dbaecg fcd bedaf fc dcbafg cegf fcdeb cbdfage bcdeg | egcf fbade cfd dbafgec
cg cadegb badfce gfdc agfbc cadbf egafb dbecfag cbg dagcbf | bagdcf fgcbda cbg dgcf
acgbf afbged fcadgb dcefg gebcaf ad adcb acdefbg cdgfa daf | dgcefab egbfacd gabefc cdab
dgb ecfbd fbgcd aegdbcf adcfg gfba gbaced acfgde fcabgd bg | dbg gb efdcb dbgcae
febg bcfaed efagdc fba ebafgc aecbgdf fegca bcgad bf acfbg | febgacd bfeg afb fb
adc ebdfacg ad badec efbdcg fade bcgdaf faedbc baecg cbfde | da edgbcf da edfa
afbgd agdcef dc cda bedcag cgaefb fdec cgfae fgadc deabgfc | cd cfage cedf fecd
facbe edcbf ebdgac fa dfgeac gaebc acf cfaegb dcfegba bfag | ebfac gcebad fac fa
cfebdag daecfb bg gecbdf ebagc ecagd gbe abfgec gabf ebfac | gbe afcbe gecdabf gb
bgefdca fecbgd ec bce cedab ecaf bdeafg cgbda dbaef bcfdae | bfecgd cefa fdeabc ce
gedbac cbade edfgbac gabfe cfaeb caf fc dcbf cdfbea cdefga | acf dafcbe cdfb dfabceg
cfgaed cbda abgfde cea abcdfeg adcgbe ac fcbge egdab eagbc | bcda ac ac ac
cbfaged afdgec ed fdea edcfg dce fdgac bfdagc fecbg ebcdga | fcabgd bfegc fcedag gfdabc
gafec afegb bgafed bafd fb egbda edbgac cedgfb ebf bedgfca | cgefbd bef edcgfb fb
cf gfc fecgd cebf dgecbaf faedgb afbgcd dcaeg efbdg dbcefg | fecb edfgcb ebgafd bfadgc
ea gafecd fbgac fcedb defcab fbeca bfdcge baed dabfceg fae | gbefcd ebda beda fae
afdgb acg aefdgc cebgd ac begdfac dacbfg gbcad abgfde cfab | abgdc abcf bacf ca
cdbga acedbg gefcd abdf fac af bfgaec ecfagdb gfdac bcdfag | caf bfad fa caf
edcagf feadg dab adgeb gceab dgfabc fdbe befcadg fagbed bd | dab dgabef dgebcaf bd
bga baedg abdc ab gefdb gadceb gcdae aedgcbf gcebaf cfdega | gab cdab dagcfeb dafebgc
afb eabdgf eafgcb gcdabe ecfa edbgfac fbdgc fa bacfg ceagb | cefa fab baf aebcgfd
da gdbcef facdeg dafbc efdcb adc edba gbcfa acdbef deafcbg | da ad defbgc ad
acbg ecbdg ecbgfda gdabec bg daefbc cefgd dceba dgb afebdg | bg edfcg gb efcgdab
dgbeca befcd gcfaed edgac agefcbd ceadb ba abe gcba daebgf | aeb acbg efgadc becgda
cgbafde ed bdce gde bdaegc eagbfd cgbad eadgc facbgd cgfae | gfdacbe fdegacb de deg
deacfbg afc dfcea cgaed gdbcea cf dcfg fbeacg gdfcae baedf | dbcaefg fdcg acgdebf gdfc
gca fagdbe cedbgaf dgbc gdaec badeg edafc gc abedcg afbgec | gac dgbae gdcb bgdc
cdfeab dbcfa bdg bg fgbc adcgb gdfeab ecgbdfa gdeca cfadgb | bedgacf acbdfg gdb bg
gbfec bgfde gacedb fcgea cb cgeafd bcg aebgcf gdaebcf afbc | cafbge abfc adbgfec bcg
edgf bcgef edcgfb gfcabd fbdcae ef efc acebg gcbdfea fgcdb | fdge fe dabcfg aefdcb
cdbafge agdebc bcg dfbgac cged ebagd degafb bfcea gc caegb | bcg edbga fcgdab debga
ea dagfbe fae fcdgbe gfaecbd decfb dacbfe abec gdafc caedf | defcba ebgdafc ceab gbcafde
dg cefga bfaed fdecgb bdag bdaecfg fgedab edg adfecb efdag | dgab gfcae bedgfc cgfdeab
degcfb bdcfg edgb cdfab bcadgef fgcdea gdf becfg dg egfbca | gd ceabgfd efcbagd bedg
dfbg bcgeda faecdg dab bfaec dbegacf geafd eabfd agbfde db | dab adfegbc db bd
ecfagdb db gadecb dagfb caefgb gbd fbeag facgd bdfe efadgb | bfedagc db dgb cagfd
edbcga fcdgbe ebcdg bcfage fb bfcgd bgf gcdaf bgfcdea fbde | dgcbf fbg gbf fdgcb
efagd eb dafcgb ebfad eabdgcf cbed aecfbd egbcfa feb adcbf | be eb eb bfe
bdge fdabce begfcd gcbdf bd ebfgc fadcg dagecbf dfb bcafge | gfacd dbf bgcaedf dgfca
geacf cfb gbfcde dfbga bc fbadgce fgdcba cabgf fdaegb dacb | dfagcb cafdebg acfbdg fcdgeba
fecdbg gbcaf ebgacfd acgfed gbad gb fbaec afcgd bafdcg fgb | gcbefd egfdbc fgb fgb
fdce ebagf fdcegb fgacbd dcbeg geadbc cf gebcf gabefdc cbf | edfc cfdgab decbg fc
adcfeg cfgad efgcdba fecda fec gabcdf ef bacde bfdceg gefa | ef edafc ef fgcdab
afeb fbegd daefg agdbcfe agdcf ega adbegc dbgafe efcdbg ae | abedcg efbgcd fgcad cdegba
fdegbca gefbda egbf dbagc gfd bgdaf adgfec fg bfead dfceab | fdg abcdegf gadebf decfgba
dbcaef debfgc gfebdca bcaed ad baceg dca deaf edfbc dagcfb | acd dafe fbdecg adef
beadc fc faedc dbafeg gdfea gafc fegcad cgebadf dfgbec fcd | cdf deacgf gefda cf
fdaceg edgacb dbcg cda cd ebfad acbeg gdfbace dbaec bcfeag | deabf bcgd cdgb defgca
fgcebd dgbcae bfcgad ga bga bfcgd gbfda acfg dbafe fbcgead | bedcgf dbfgc gfac acfbgd
cfaebdg dcg gedbfc gefdba egafcd dc cbde bedfg bcgdf cabgf | dbacfeg feadgb dcbgf fbgdae
fagcbd fedbg fcabge ebgad cbfdg cfedgab feb fdce dfcgeb ef | cfde fgedb bdgecf ef
becgdf efcbda cgbd gdcfe gc gaedf efbcd cfg fbegcad gbfcae | fcg cdefbg gafed efadg
da adbcgef dacf eadfcb bfcage eafbc ceabd begfda bdceg bda | da da acbedf bagcfe
agebf adfb fabedg cgafde dga beadfcg cdbeg egbad cfabeg da | dfegba dafgbce egdbc fecbdga
eac fcdabeg gdabc ebadc bcge gfecad deafb bdgfac dgaecb ce | cea ec cegabdf gadcb
gfedbc gcdfb ecdfga df cfegba defb dbcga gdf cgfbead ecgbf | facbged dfg bfed cfbedag
agbde bgdfca deacgb agbef aebcgdf afg fbecg aedbfg afed af | daef fga fga gdefab
gfbac gaecdb baefc fdebgca fgbadc ec adbef efgc fabegc cea | cae eac gfdcbea fceg
ec dagbc cagfbd adbecg bfdge gebcdaf ced gdebc fagced bcea | gdbeca dec agcdb aebc
geab ebgdfa eafbd dafbc efdgb aed ae bgcaefd dcegfb aegdcf | ebga ade ebfdg dgafceb
bdaeg aebcd cgebad dbc cb fecbgd dafce bcgfdea gcab gfdbae | cbga cbag cdb bagc
dfb egfbdc ceabdfg cfbae fgbade fdbec bd gedfc deafgc cbdg | ebfgdca fdgbec caebdgf db
dfcaeg bagcefd fceab ad dae faedb fcabed ebfgd fcabeg adcb | dea ad ebgfd da
afgcd acgebf da agbdfc cad gcdfe dagb dafbec dagbefc abcgf | fgcde beagcfd cebafg da
cabd bce cb cfbdea adbef agcfbe aebcdfg fabdeg dgcfe fbdec | ebc gfcabe egcfd fedab
efcbga cgfbaed bcfga ecgbf egdbc deabfc fdgcba fage fce ef | ecgbdaf gdecabf bdceg febcga
gbcda ad gacfb dba gfacbd fbgdeca aecbfg fdac gfabde bdgce | dab bdgafce bcegd abegdfc
fcaegb baegdf gfaced gedfa dcebg dbgea agedcbf abdf bga ba | bag fabd bgfcea bga
ebafg gdebafc gedacb gcdefb agbdfc cf gebcd gcf cfde cgbfe | fdgaecb dfbcega cfg fc
fce cbaedf cgbaedf ec bfega dcea agbdfc bfdegc dacbf cefab | ce efabg faebcd gdecbaf
dcfbg fd egcdba efcgb bfad cefagd abegdfc gcafdb cabgd fdg | bafd egfbc cdgabf gfebc
abfegc gecdabf cedb bdg dgcbf gaebdf db fgacd ecfgb egcfdb | bdegaf gfaebd gcefb cdbe
bafeg febdac ged bfcedg gd cbafedg edcaf gafdce cagd fadge | fgcdae daebgcf dagc fadeg
cfdegb efdc gcadbe gdbaf ebadgfc fc gdebc dfcgb fgc egcbaf | fc cf cfed cf
aecfgd efgab bdagfc cedf cge agfce afcdg cdeabg dgbaefc ec | gce gce ec cegfadb
eagdbc ga fecbda ecdaf fdgbc abgdcef afdcg fega cag eafgdc | cga ag ag ag
fabedg cedg cbefga dga gacedf gd fdcga afecg bfdac gabfedc | fgeca gdec fdacg gda
begadc gcfeab dgcba defbac gc fbdag agdebcf degc cag abecd | gcde gca cdge cg
fd dcgba gdbfca afedcgb gbadef cagfe fad dfcb acgfd dacgeb | cbgdea fd cfdb afegbdc
fb fceab fdeb bgfaced acdfbe gbcafd cdabe abf agcef baecdg | bf fb gdebac bf
df fdb egcbd gfed bcaegd abcfg ceadbf egcdbf aegfdcb dbgfc | begdac fdb acefbd dgbec
bgedfa dg bdcgeaf cbfad decg efcdga aecgfb aegcf cgdfa gad | egfac dg gedc dga
gadfe cfgead cdgbae cfage de dbceagf ecgafb cfed dea fgbad | ade fecag aefgc fcagbed
dfab gaecbd degfc bcegfa dgebfa edb ebdfg dbceagf gafeb bd | gbaef bafd agcfbe bfcadeg
edcbg dgbfc bdfecag cfegbd eg gebfda dcaeb afcdgb cgef ebg | fceabdg gdfbec beg fdbgc
ba gfbdec gacfd bfcdae caeb egbfacd cbedf edabgf bad cadbf | becfdg aecb bad efgcdb
gba dabgfec ga fbadc dafgcb fdabg dbecaf agfceb cdag bfegd | dfeabc gab gfabd abg
abgdf gcabfd cgabde agedfcb ab cegfda dbfeg afcgd acfb gba | bcfa dgefca dfcagb ba
bdecf dgebfc bgcd fdgae gec cg gbacef bacefd cgedf afbecgd | abegfc egfcd cg gbedafc
gdecbfa cd bfcagd fagec dfcae dcf adgcfe eabfd cdge gcbeaf | cd dfc dcf gcfbae
fdecb adfce bc ecagbd bdc facgedb bfac fbged egdacf fcadbe | daefc efcgda bcd dagcfeb
fgdba fbae bcafedg dfgebc fag cbgad gedacf gdfeb fbdage fa | dbfag fdbeacg bgedaf bfae
dafbg abfed fe eaf degbca fegabc bafdec dcef aedbc afdcegb | fe fe ebfcag dacgefb
adcegb cf dceab caef ebgadfc dcabfe cbdfe fbc degbf gbdfca | cf afce aecgdb cf
cebdfag acfgbd gcbef bcf dfegc bf egacfd bfde gbfdec abceg | dbef cbf cfgde cbf
bgfaedc gcdfb cadbge fdac bgfdac cf eabfgc bfc acdbg fgbed | cafd fbdagc bdgac cf
fgdc fgceab edf feadg beadfc bdgae egdcfba efagc fd cdagef | fed gfdcaeb cgaef fedgca
bfg ecabg cabdgfe fdgc fgbac dafbge gf fcdbag fadcbe fdcab | bdagecf fgbdaec fdcgeab fg
aefdgb agcfe dfeca dgecfa feg deabgfc ge gacfb defabc dceg | efg edfca egf aefcdg
afbeg fagd adgfbe abefcdg bdefca ga gab ecdgab abefd bgfce | fgad fdag dgcbeaf ga
fbgea ebfgdca badegc bac cb gceba eadcgf daegc gcdafb bedc | cgbdefa bc bc dabecgf
cfbega dgba gb fbegd gdeabf bfcadeg egdcf fgb badef fdbaec | gadb facbdge cfgdbae eadbf
gdecaf fabed bdgcae dcgfba ce bdfegca dcaeb gbcda cae ebgc | afdgce dcbegfa egdcfba bgfcad
bcgdfa cda agecfdb cedba cafdeb caefbg abfec edcf debag cd | cda cda dc ecabd
fdbagce abfc cageb cagfe eba gadefc dgcbe adebgf geafbc ab | ab edbagf afbc facb
gcdfa cadeg dcaeb adcegf gfcdba fcabgde ge gfed gea gfeacb | cadgf gae dagfbc gfed
bfegdac gcabf fdg adgfb gefdcb gd abfed gdae dfgbae befacd | dega bfgdec dfg daefbg
acebgd ceg gcabd afcbgde ebcd afbdcg cbgea ec gbfae gdfeac | gce afbeg adcgb abedcgf
degfab fbade ac acdf feacb dgcbea cfbge abc dgfaecb bdecfa | adfc cdaebg ca acb
adegcfb gf acfdbe ecfdb dgceaf dgfbe dgbae feg fcbg dfcbge | gf fg fg gf
fbacde gfacbe eg dcge fdbga beg bedgcf dcbef fcbaedg bgfde | cged edcg bgedf ge
dgafc dfbegc fagbed cfba adfbcg cfg cf dabfg gafcdbe dgaec | dgbface fcgad cfg fadgc
defcga gdfc afc cadebf bceag degaf fc gaedfb dabcgfe eagfc | adfge dgfc edcfab fgcd
dcefb feb gcdeb efdg fcagbe becgad bfcda gabfdec fe ebfdgc | ecbagf ef dfge decbf
fgebac faegb gcebfad fb fecga bgf ecbf abged gcefad fbgacd | fb fbeag fbega bfg
cbdgfae fbe gdfbce egbda abedf cdfea badcge agfb edbafg fb | ebfad afbg edafbg efb
gbdc gb cfdegba bdgae gab gcafbe fadeb ecfgda gbdace cgaed | abg dbgc dcbg gb
db dgefb eafdbgc ecfbg bcegdf bdg fbcd gefad ecfbag adbgce | db cabedgf dbg fgead
ab bcfa bacfed dba befad fadeg bcdaeg abcgdfe bdecfg bcfde | ba debacfg fdcbe abd
baecfg edgba cdgae fced cd egfdca cgd acedfgb dbfgca feagc | adceg dcaeg efdc dc
cdb bgfeda acdeb becfda fedba dfca gbeca dbgecf dc aedcfbg | bdc egabfd cd eadbc
efgcadb dgebf baedf gcedfa bfg gb eacgbf dgefc bdcg cgbfde | dfgbe gb gbdefca dbcg
afebdgc gbfdc gbca defgb cdb bc bcfagd bcdfae edfgca fdacg | cbd cgab dfbgc cb
bgfade cdgabef cfebd ce febad dgafce dce abce cgbdf edcafb | bfeadg gdfbea cabe ce
gefbcd de dgfce gfdaecb gafec bfdcg bdge abdfgc edf dfbcea | ed fcdgab dfe edgb
cgafeb egadbc bagce gedcafb afcbd fgae bfeac fdegbc feb fe | egcba bfe bef bcdaf
bcdea dgbfc gdae dabcg ecdfab gca ag dcebgaf gadecb beafcg | eagd eadg bfgcd gac
gafcb eba acbfeg cagbfd ecaf fcbedag bdcge ea cagbe fdeabg | bea aeb ae eab
cdegfb cdagef bgf dfgcaeb eadgb fegac febga bf bfcaeg cfba | caefdgb gfb fb ebafgcd
bgedfa dfbca ebdgcf fbgcd dfcega cgf gfbcead degfb ecbg gc | gecb cfg cg aegdfb
agebd bcegda aeg fedabc bdgafce acgd fbgde ag caegbf cbdae | agcd egabfc dacg dgca
cbdfe feacd geacfd dbaefg cbgfd aedfbc ceba eb bde egacfdb | bed ceba adbefgc ebd
egd ge afcbedg bfgeda bdgefc gbce bcgdf fdacbg aecfd dcgfe | cefgbd afcbgd ge cebg
bg gacbe cefgad cbgf cbead gba fcedagb ebdafg gcfbae gacfe | gabec cedfgab acgefb dgecabf
afedg edabg fcdg bfcega ebfdca decaf gf gfa acegfbd afedcg | dgcf efacgdb dcgf ebdfcag`
