package revealer

import (
	"testing"
)

func TestFixer(t *testing.T) {

	var tests = []struct {
		email          string
		expectedResult string
	}{
		{"dijuek[at]googlemail[dot]com", "dijuek@googlemail.com"},
		{"y.imai at ocaml.jp", "y.imai@ocaml.jp"},
		{"nap-at-zerosum-dot-org", "nap@zerosum.org"},
		{"zxytim[at]gmail[dot]com", "zxytim@gmail.com"},
		{"naranjo dot manuel at gmail dot com", "naranjo.manuel@gmail.com"},
		{"170503326@qq,com", "170503326@qq.com"},
		{"defagos (@) gmail (.) com", "defagos@gmail.com"},
		{"delesley atsign google dotsym com", "delesley@google.com"},
		{"digicyc@gmail .com", "digicyc@gmail.com"},
		{"dima @ secretsauce . net", "dima@secretsauce.net"},
		{"demirozali (@) gmail.com", "demirozali@gmail.com"},
		{"debackerl gmail com", "debackerl@gmail.com"},
		{"smihica_gmail_com", "smihica@gmail.com"},
		{"dhananjay.patkar.gmail.com", "dhananjay.patkar@gmail.com"},
		{"borismoore@gmail.com;", "borismoore@gmail.com"},
		{"syusui.s[a]gmail.com", "syusui.s@gmail.com"},
		{"alex |at| sorokine.info", "alex@sorokine.info"},
		{"shinichiro.hamaji _at_ gmail.com", "shinichiro.hamaji@gmail.com"},
		{"stevendealatgmail.com", "stevendeal@gmail.com"},
		{"digicyc at gmail (dot) com", "digicyc@gmail.com"},
		{"dexgecko (gmail)", "dexgecko@gmail.com"},
		{"dima -at- secretsauce -dot- net", "dima@secretsauce.net"},
		{"mattthieu point dubuget at gmail point com", "mattthieu.dubuget@gmail.com"},
		{"trisk at-sign forkgnu.org", "trisk@forkgnu.org"},
		{"diego -at- vazqueznanini (dot) com", "diego@vazqueznanini.com"},
		{"gabe(a)gundy.org", "gabe@gundy.org"},
		{"crazyjvm #at# gmail.com", "crazyjvm@gmail.com"},
		{"fjfnaranjo 4t gmail d0t com", "fjfnaranjo@gmail.com"},
		{"thomaspollet++@++gmail.com", "thomaspollet@gmail.com"},
		{"felix [dot] mulder [at] gmail [dotcom]", "felix.mulder@gmail.com"},
		{"felix021 # gmail.com", "felix021@gmail.com"},
		{"falko ät briefhansa dot de", "falko@briefhansa.de"},
		{"ev45ive + github @gmail.com", "ev45ive@gmail.com"},
		{"ericwillisson atmark gmail period com", "ericwillisson@gmail.com"},
		{"richard d0t maynard isat gmail d0t com", "richard.maynard@gmail.com"},
		{"elias ((at)) showk ((dot)) me", "elias@showk.me"},
		{"twoerner .at. gmail .dot. com", "twoerner@gmail.com"},
		{"[[tfujiwar at redhat dot com]]", "tfujiwar@redhat.com"},
		{"gentimouton (@gmail.com)", "gentimouton@gmail.com"},
		{"dpcx@dpcloudx@.com", "dpcx@dpcloudx.com"},
		{"costantino.giuliodori at google email", "costantino.giuliodori@google.com"},
		{"ideaplexus [ at ] gmail", "ideaplexus@gmail.com"},
		{"frode andre petterson (at) gmail com", "frode.andre.petterson@gmail.com"},
		{"jach shift2 thejach d0t com", "jach@thejach.com"},
		{"cyprian -at- ironin -dot- pl", "cyprian@ironin.pl"},
		{"ossug.hychen -at- gmail.com", "ossug.hychen@gmail.com"},
		{"hughpearse 'at' gmx \\dot\\ co $dot$ uk", "hughpearse@gmx.co.uk"},
		{"kolar.radim gmail.com", "kolar.radim@gmail.com"},
		{"hirafoo atmk gmail_com", "hirafoo@gmail.com"},
		{"1417262058qq@.com", "1417262058@qq.com"},
		{"1078360711@.qq.com", "1078360711@qq.com"},
		{"13384033080@.163.com", "13384033080@163.com"},
		{"jonathan . david . smith -at- gmail.com", "jonathan.david.smith@gmail.com"},
		{"viszu84atgmaildotcom", "viszu84@gmail.com"},
		{"reddaly at gee mail dot com", "reddaly@gmail.com"},
		{"peter _at_ bloodaxe _dt_ com", "peter@bloodaxe.com"},
		{"kay.cichini[atnospam]gmail.com", "kay.cichini@gmail.com"},
		{"ximen@ximen,ru", "ximen@ximen.ru"},
		{"[thecrazymancan@gmail.com", "thecrazymancan@gmail.com"},
		{"stevendealatgmail.com", "stevendeal@gmail.com"},
		{"harald æt hauknes dot org", "harald@hauknes.org"},
		{"johnish { at } gmail", "johnish@gmail.com"},
		{"jsantiago [[[at]]] fastmail.us", "jsantiago@fastmail.us"},
		{"jeffrey.knight@gmail[.com", "jeffrey.knight@gmail.com"},
		{"aleonardobecerra@gmail,.com", "aleonardobecerra@gmail.com"},
		{"j.wb2007#@163.com", "j.wb2007@163.com"},
		{"zach @ g mail dot com", "zach@gmail.com"},
		{"liam [atatat] w3 [dotdotdot] org", "liam@w3.org"},
		{"levilucas93\"gmail.com", "levilucas93@gmail.com"},
		{"leejaygo.163.com", "leejaygo@163.com"},
		{"sencer.hamarat(at`gmail.com", "sencer.hamarat@gmail.com"},
		{"kurt[ at ]kurtrose.com", "kurt@kurtrose.com"},
		{"kiran[dot]puttur[dot]gmail[dot]com", "kiran.puttur@gmail.com"},
		{"knusul(a-t)gmail.com", "knusul@gmail.com"},
		{"kentoj *at* gmail *dot* com", "kentoj@gmail.com"},
		{"yde001atgmaildotcom", "yde001@gmail.com"},
		{"m.mastrodonato(a-t)gmail(d-o-t)com", "m.mastrodonato@gmail.com"},
		// {"macopy123[attttt]gmai.com", "macopy123@gmail.com"},
		{"togos at gee mail daught com", "togos@gmail.com"},
		// {"lucas [arroba] dillmann.com.br", "lucas@dillmann.com.br"},
		{"uniqueluolong##gmail.com", "uniqueluolong@gmail.com"},
		{"vit at ribachenko dt com", "vit@ribachenko.com"},
		{"joe [@t] webdrake.net", "joe@webdrake.net"},
		{"markus.dahm (_at_) akquinet.de", "markus.dahm@akquinet.de"},
		{"mitsuhiro.okuno＠gmail.com", "mitsuhiro.okuno@gmail.com"},
		{"magnus __at_ hagander.net", "magnus@hagander.net"},
		{"doug holmes gmail", "doug.holmes@gmail.com"},
		{"test1 {at} example {dot} com", "test1@example.com"},
		// {"test2ATexampleDOTcom", "test2@example.com"},
		{"test3 at example dot edu", "test3@example.edu"},
		{"test4 \"at\" example =d0t= biz", "test4@example.biz"},
		{"test5N0_SPAM@example.com", "test5@example.com"},
		{"test6@example[[N*0*S*P*A*M]].com", "test6@example.com"},
		{"test7/@example.com.invalid", "test7@example.com"},
		{"t e s t 7 @ f o o . c o m", "test7@foo.com"},
		{"test9@@example.com", "test9@example.com"},
		{"ren_kai {$at} live.com", "ren_kai@live.com"},
		{"pvandenberk [(at)] mac [(dot)] com", "pvandenberk@mac.com"},
		{"boromil ta gmail otd com", "boromil@gmail.com"},
		{"priikone [ət] iki.fi", "priikone@iki.fi"},
		{"bjardon97gmail.com", "bjardon97@gmail.com"},
		{"pgrabows 'at' mtools.com", "pgrabows@mtools.com"},
		{"\"types\" at \"ccs.neu.edu\"", "types@ccs.neu.edu"},
		{"adam.a.szymczakatgmail.com", "adam.a.szymczak@gmail.com"},
		{"aleksandar topuzovic at gmail dot com", "aleksandar.topuzovic@gmail.com"},
		{"ozkan {.at.) portakal.net", "ozkan@portakal.net"},
		{"ancosen dot gmail dot com", "ancosen@gmail.com"},
		{"747325123qq.com", "747325123@qq.com"},
		{"codemonkey at forsters freehold dat calm", "codemonkey@forsters.freehold.com"},
		{"neutra.at.qq.dot.com", "neutra@qq.com"},
		{"randy ~dot~ secrist ~at~ gmail.com", "randy.secrist@gmail.com"},
		{"luisfmh@gmail.com or luis.hernandez@ryerson.ca", "luisfmh@gmail.com"},
		{"info[-at-]codecraft.de", "info@codecraft.de"},
		{"rpdladps åt gmail.com", "rpdladps@gmail.com"},
		{"voropaev.roma u+0040 gmail.com", "voropaev.roma@gmail.com"},
		{"jian.luo.cn(at_)gmail.com", "jian.luo.cn@gmail.com"},
		{"hongsige1989@163.", "hongsige1989@163.com"},
		{"hongsige1989@qq.", "hongsige1989@qq.com"},
		{"hongsige1989163.com", "hongsige1989@163.com"},
		{"a.a.m.macdonald[at]gmail.com", "a.a.m.macdonald@gmail.com"},
		// {"abc.\"defghi\".xyz@example.com", "abc.\"defghi\".xyz@example.com"},
		// {"\"abcdefghixyz\"@example.com", "\"abcdefghixyz\"@example.com"},
		{"joe located-at intrusion.org", "joe@intrusion.org"},
		{"jmpessoa_hotmail.com", "jmpessoa@hotmail.com"},
		{"bernhard zq1.de", "bernhard@zq1.de"}, // First space =@, last space =.
		{"gmail: test.last", "test.last@gmail.com"},
		{"keunho.yoo __at__ gmail.com", "keunho.yoo@gmail.com"},
		{"kenneth.wong8(.a.t.)gmail.com", "kenneth.wong8@gmail.com"},
		{"[gmail]: kieranrcampbell", "kieranrcampbell@gmail.com"},
		{"(gmail): kieranrcampbell", "kieranrcampbell@gmail.com"},
		{"kieranrcampbell [gmail]", "kieranrcampbell@gmail.com"},
		{"crawford@saao.ac.za.", "crawford@saao.ac.za"},
		{"denisowpavel [аt] yandex.ru", "denisowpavel@yandex.ru"},
		// {"danielsh apache org", "danielsh apache org"},
	}

	for _, test := range tests {
		result, err := Fix(test.email)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if result != test.expectedResult {
			t.Errorf("Expected: %s, Actual: %s", test.expectedResult, result)
		}
	}

	// TODO make errors constants?

	// test errors
	var testErrs = []struct {
		email       string
		expectedErr string
	}{
		{"", "email address cannot be empty"},
		{"broken", "unable to fix email address: broken -> broken"},
		{"0", "unable to fix email address: 0 -> 0"},
		{"01", "unable to fix email address: 01 -> 01"},
		{"012", "unable to fix email address: 012 -> 012"},
		{"0123", "unable to fix email address: 0123 -> 0123"},
		{"01234", "unable to fix email address: 01234 -> 01234"},
		{"012345", "unable to fix email address: 012345 -> 012345"},
		{"0123456", "unable to fix email address: 0123456 -> 0123456"},
		{"01234567", "unable to fix email address: 01234567 -> 01234567"},
		{"012345678", "unable to fix email address: 012345678 -> 012345678"},
		{"0123456789", "unable to fix email address: 0123456789 -> 0123456789"},
	}

	for _, testErr := range testErrs {
		_, err := Fix(testErr.email)
		if err == nil {
			t.Errorf("Should have errored!")
		}
		if err.Error() != testErr.expectedErr {
			t.Errorf("Expected Err: %s, Actual Err: %s", testErr.expectedErr, err)
		}
	}
}

func TestAddDots(t *testing.T) {

	var tests = []struct {
		email          string
		expectedResult string
	}{
		{" frode andre petterson ", "frode.andre.petterson"},
		{"doug holmes gmail", "doug.holmes.gmail"},
		{"jonathan . david . smith ", "jonathan.david.smith"},
		{"test@first second .com", "test@first.second.com"},
		{"doug.holmes gmail.com", "doug.holmes@gmail.com"},
	}

	for _, test := range tests {
		result := addDots(test.email)
		if result != test.expectedResult {
			t.Errorf("Expected: %s, Actual: %s", test.expectedResult, result)
		}
	}
}

func TestCheckSpecial(t *testing.T) {

	var tests = []struct {
		email          string
		expectedResult string
	}{
		{`test`, `test`},
		{`te\ st`, `te st`},
		{`te\"st`, `te st`},
		{`"test"`, `"test"`},
		{`"te()st"`, `"te()st"`},
		{`test"te()st"test`, `test"te()st"test`},
		{`te()st"te()st"test`, `te  st"te()st"test`},
		{`te()st"te()st"te[]st`, `te  st"te()st"te  st`},
	}

	for _, test := range tests {
		result := checkSpecial(test.email)
		if result != test.expectedResult {
			t.Errorf("Expected: %s, Actual: %s", test.expectedResult, result)
		}
	}
}

func TestStripBad(t *testing.T) {

	var tests = []struct {
		email          string
		expectedResult string
	}{
		{"test@example.^com", "test@example.com"},
		{"test@`example.com", "test@example.com"},
		{"test@example_.com", "test@example.com"},
		{"test@first second .com", "test@first second .com"},
	}

	for _, test := range tests {
		result := stripBad(test.email)
		if result != test.expectedResult {
			t.Errorf("Expected: %s, Actual: %s", test.expectedResult, result)
		}
	}
}

func TestPadding(t *testing.T) {
	// regular tests
	var tests = []struct {
		test           string
		length         int
		char           string
		expectedResult string
	}{
		{"test", 20, " ", "test                "},
		{"", 20, "-", "--------------------"},
	}

	for _, test := range tests {
		result, _ := padding(test.test, test.length, test.char)
		if result != test.expectedResult {
			t.Errorf("Expected: %s, Actual: %s", test.expectedResult, result)
		}
	}

	// test errors
	var testErrs = []struct {
		test        string
		length      int
		char        string
		expectedErr string
	}{
		{"test", 2, " ", "string is too long"},
		{"", 20, "--", "padding must be only one character"},
		{"thisisareallylongstringtotestwith", 20, " ", "string is too long"},
	}

	for _, testErr := range testErrs {
		_, err := padding(testErr.test, testErr.length, testErr.char)
		if err == nil {
			t.Errorf("Should have errored!")
		}
		if err.Error() != testErr.expectedErr {
			t.Errorf("Expected Err: %s, Actual Err: %s", testErr.expectedErr, err)
		}
	}
}

//
// Benchmarks
//

// base: BenchmarkPadding-4   	 5000000	       294 ns/op
func BenchmarkPadding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = padding("test", 40, " ")
	}
}

// base: BenchmarkStripBad-4   	 1000000	      1160 ns/op
func BenchmarkStripBad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = stripBad("test@`exa^mp>le.co_m")
	}
}

// base: BenchmarkStripBad-4   	 1000000	      1160 ns/op
func BenchmarkHandcraftedFixes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = handcraftedFixes("test (gmail)")
	}
}
