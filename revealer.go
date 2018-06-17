// Package revealer will "de-obfuscate" email addresses.
//
// See README.md for more info.
package revealer

import (
	"errors"
	"log"
	"net/mail"
	"regexp"
	"strings"
)

var (
	debugEmail = ""
	logSteps   bool
	reg        *regexp.Regexp
)

func init() {
	// Domain names may be formed from the set of alphanumeric
	// ASCII characters (a-z, A-Z, 0-9), but characters are
	// case-insensitive. In addition the hyphen is permitted
	// if it is surrounded by characters, digits or hyphens,
	// although it is not to start or end a label.
	// http://get.bible/blog/post/a-clear-description-for-a-valid-domain-name

	// make a simple regex to clear out anything that doesn't belong
	// to be used later.
	var err error
	reg, err = regexp.Compile("[^a-z0-9-. ]+")
	if err != nil {
		log.Fatal(err)
	}
}

// Fix "de-obfucates" email addresses
func Fix(email string) (string, error) {

	// check for empty string first
	if email == "" {
		return "", errors.New("email address cannot be empty")
	}

	// debugging?
	if email == debugEmail {
		logSteps = true
		logStep("Original:", email)
	} else {
		logSteps = false
	}

	// general fixes
	fixed := strings.ToLower(email)
	fixed = generalFixes(fixed)

	// find the @
	fixed = findTheAt(fixed)

	// find the dots
	fixed = findTheDots(fixed)

	// find the @ (round 2)
	fixed = findTheAt(fixed)

	// find the dots (round 2)
	fixed = findTheDots(fixed)

	// strip bad characters in domain
	fixed = stripBad(fixed)

	// perform any "special" hardcoded fixes
	fixed = handcraftedFixes(fixed)

	// add dots, or @ as needed
	stripped := addDots(fixed)

	// remove any special characters
	stripped = checkSpecial(stripped)

	// check if valid
	_, err := mail.ParseAddress(stripped)
	if err == nil {
		logStep("Valid:", stripped)
		return stripped, nil
	}

	//
	// Try again!
	//

	// find the @
	stripped = findTheAt(stripped)

	// find the dots
	stripped = findTheDots(stripped)

	// check if valid
	_, err = mail.ParseAddress(stripped)
	if err == nil {
		logStep("Valid:", stripped)
		return stripped, nil
	}

	return "", errors.New("unable to fix email address: " + email + " -> " + stripped)
}

func handcraftedFixes(email string) string {

	// If [gmail] is in front
	if len(email) > 7 {
		if email[0:7] == "[gmail]" {
			email = email[7:] + "@gmail.com"
		}
	}

	// If (gmail) is in front
	if len(email) > 7 {
		if email[0:7] == "(gmail)" {
			email = email[7:] + "@gmail.com"
		}
	}

	// If gmail is in front
	if len(email) > 5 {
		if email[0:5] == "gmail" {
			email = email[5:] + "@gmail.com"
		}
	}

	// if it ends in gmail (without .com) add .com
	if len(email) > 6 {
		if email[len(email)-5:] == "gmail" || email[len(email)-6:] == "gmail." {
			email = email + ".com"
		}
	}

	// if it ends in (gmail)
	if len(email) > 7 {
		if email[len(email)-7:] == "(gmail)" {
			email = email[:len(email)-7] + "@gmail.com"
		}
	}

	// if it ends in [gmail]
	if len(email) > 7 {
		if email[len(email)-7:] == "[gmail]" {
			email = email[:len(email)-7] + "@gmail.com"
		}
	}

	// if it ends in qq (without .com) add .com
	if len(email) > 3 {
		if email[len(email)-2:] == "qq" || email[len(email)-3:] == "qq." {
			email = email + ".com"
		}
	}

	// if it ends in 163 (without .com) add .com
	if len(email) > 4 {
		if email[len(email)-3:] == "163" || email[len(email)-4:] == "163." {
			email = email + ".com"
		}
	}

	// if gmail.com does not have an @ in front
	if len(email) > 10 {
		if email[len(email)-9:] == "gmail.com" && email[len(email)-10:len(email)-9] != "@" {
			email = email[0:len(email)-9] + "@gmail.com"
		}
	}

	// if qq.com does not have an @ in front
	if len(email) > 7 {
		if email[len(email)-6:] == "qq.com" && email[len(email)-7:len(email)-6] != "@" {
			email = email[0:len(email)-6] + "@qq.com"
		}
	}

	// if 163.com does not have an @ in front
	if len(email) > 8 {
		if email[len(email)-7:] == "163.com" && email[len(email)-8:len(email)-7] != "@" {
			email = email[0:len(email)-7] + "@163.com"
		}
	}

	// FIXME an attempt to strip off everything after ".com" but this
	// may be inaccurate - can you have .com.uk for example?
	split := strings.SplitAfter(email, ".com")
	email = split[0]

	split = strings.SplitAfter(email, ".org")
	email = split[0]

	split = strings.SplitAfter(email, ".net")
	email = split[0]

	split = strings.SplitAfter(email, ".edu")
	email = split[0]

	logStep("Hand fixes:", email)
	return email
}

func addDots(email string) string {

	// split string on "@"
	portions := strings.Split(email, "@")

	//
	// name portion
	//

	// trim off spaces in first portion
	trimmed := strings.Trim(portions[0], " ")

	// now split first portion (name) on spaces
	spaces := strings.Split(trimmed, " ")

	// if we have no @ and just one space, try an @
	if len(portions) == 1 && len(spaces) == 2 {
		addAt := ""
		for i, s := range spaces {
			s = strings.Trim(s, " ")
			if s != "" && s != "." {
				if i == 0 {
					addAt = s
				} else {

					addAt = addAt + "@" + s
				}
			}
		}
		return addAt
	}

	// if we have a lot of sections just strip spaces
	if len(spaces) > 3 {
		portions[0] = strings.Replace(portions[0], " ", "", -1)
	} else {
		// otherwise reassemble - placing dots between spaces
		newString := ""
		for i, s := range spaces {
			s = strings.Trim(s, " ")
			if s == "+" { // dump "+" and everything after (+spam, +junk, etc.)
				break
			} else {
				if s != "" && s != "." {
					if i == 0 {
						newString = s
					} else {

						newString = newString + "." + s
					}
				}
			}
		}
		portions[0] = strings.TrimRight(newString, ".")
	}

	//
	// domain portion
	//

	if len(portions) > 1 {

		// trim off spaces
		domain := strings.Trim(portions[len(portions)-1], " ")

		// now split last portion (domain) on spaces
		domainChunks := strings.Split(domain, " ")

		// if we have a lot of sections just strip spaces
		if len(domainChunks) > 3 {
			portions[len(portions)-1] = strings.Replace(portions[len(portions)-1], " ", "", -1)
		} else {
			// otherwise reassemble - placing dots between spaces
			domainString := ""
			for i, s := range domainChunks {
				s = strings.Trim(s, " ")
				s = strings.Trim(s, ".")
				if s != "" {
					if i == 0 {
						domainString = s
					} else {
						domainString = domainString + "." + s
					}
				}
			}
			portions[len(portions)-1] = strings.TrimLeft(domainString, ".")
		}
	}

	//
	// reassemble full string
	//

	for i, s := range portions {
		if i == 0 {
			email = s
		} else {
			email = email + "@" + s
		}
	}

	logStep("Add dots:", email)
	return email
}

func stripBad(email string) string {

	// split string on "@"
	portions := strings.Split(email, "@")

	// reassemble full string
	for i, s := range portions {
		if i == 0 {
			email = s
		} else {
			email = email + "@" + reg.ReplaceAllString(s, "")
		}
	}

	logStep("Strip bad:", email)
	return email
}

func checkSpecial(email string) string {

	// Special characters are allowed with restrictions -
	// They must be between double quotes.

	// They are: Space and "(),:;<>@[\]
	r := strings.NewReplacer(
		"(", " ",
		")", " ",
		",", " ",
		":", " ",
		";", " ",
		"<", " ",
		">", " ",
		// "@", " ",
		"[", " ",
		"]", " ",

		// The restrictions for special characters are that they must
		// only be used when contained between quotation marks, and
		// that 3 of them (The space, backslash \ and quotation mark "
		// must also be preceded by a backslash \ (e.g. "\ \\\"").
		"\\ ", " ",
		"\\", " ",
		"\"", " ",
	)

	// split email on quotes
	chunks := strings.Split(email, "\"")

	// range through chunks
	for i, s := range chunks {
		if i == 0 {
			email = r.Replace(s)
		} else {
			// i is odd (because we start at zero it means we are even!)
			// and we are not at the first chunk
			// and there is an odd number of chunks
			if i%2 != 0 && i > 0 && len(chunks)%2 != 0 {
				email = email + "\"" + s
			} else {
				// i is even (because we start at zero it means we are odd!)
				// and we are not at the first chunk
				// and there is an odd number of chunks
				if i%2 == 0 && i > 0 && len(chunks)%2 != 0 {
					email = email + "\"" + r.Replace(s)
				} else {
					email = email + r.Replace(s)
				}
			}
		}
	}

	trimmed := strings.Trim(email, " ")
	trimmed = strings.Trim(trimmed, ".")

	logStep("Special chars:", trimmed)
	return trimmed
}

func findTheAt(email string) string {

	r := strings.NewReplacer(

		// remove duplicates
		"@@@@", "@",
		"@@@", "@",
		"@@", "@",
		"@ @", "@",

		// some strings may appear inside a name so we will
		// only look for those that have spaces around them
		// NOTE multiples in inverse order
		" atatat ", "@",
		" atat ", "@",
		" at ", "@",
		" ta ", "@",
		" at-sign ", "@",
		" located-at ", "@",
		" atsign ", "@",
		" isat ", "@",
		" atmark ", "@",
		" atmk ", "@",
		" shift2 ", "@",
		" 4t ", "@",

		// misc @
		"＠", "@",
		"ät", "@",
		"æt", "@",
		"ət", "@",
		"åt", "@",
		"a-t", "@",
		"а", "a", // a is a different rune, really...
		".a.t.", "@",
		"u+0040", "@",
		"arroba", "@", // at in spanish

		"[@t]", "@",
		"(@t)", "@",
		"{@t}", "@",
		"<@t>", "@",

		// -------------
		// @ substitutes
		// -------------
		".@.", "@",
		"!@!", "@",
		"#@#", "@",
		"$@$", "@",
		"%@%", "@",
		"&@&", "@",
		"'@'", "@",
		"*@*", "@",
		"+@+", "@",
		"-@-", "@",
		"/@/", "@",
		"=@=", "@",
		"?@?", "@",
		"^@^", "@",
		"_@_", "@",
		"`@`", "@",
		"|@|", "@",
		"{@{", "@",
		"}@}", "@",
		"{@}", "@",
		"~@~", "@",
		"\"@\"", "@",
		"(@(", "@",
		")@)", "@",
		"(@)", "@",
		",@,", "@",
		":@:", "@",
		";@;", "@",
		"<@<", "@",
		">@>", "@",
		"<@>", "@",
		"@@@", "@",
		"\\@\\", "@",
		"[@[", "@",
		"]@]", "@",
		"[@]", "@",
		"*@*", "@",
		"#@#", "@",

		// one side
		// NOTE: other side handled via special domain handling
		".@", "@",
		"!@", "@",
		"#@", "@",
		"$@", "@",
		"%@", "@",
		"&@", "@",
		"'@", "@",
		"*@", "@",
		"+@", "@",
		"-@", "@",
		"/@", "@",
		"=@", "@",
		"?@", "@",
		"^@", "@",
		"_@", "@",
		"`@", "@",
		"|@", "@",
		"{@", "@",
		"}@", "@",
		"~@", "@",
		"\"@", "@",
		"(@", "@",
		")@", "@",
		",@", "@",
		":@", "@",
		";@", "@",
		"<@", "@",
		">@", "@",
		"<@", "@",
		"@@", "@",
		"\\@", "@",
		"[@", "@",
		"]@", "@",
		"*@", "@",

		// ----------------
		// "at" substitutes
		// ----------------
		".at.", "@",
		"!at!", "@",
		"#at#", "@",
		"$at$", "@",
		"%at%", "@",
		"&at&", "@",
		"'at'", "@",
		"*at*", "@",
		"+at+", "@",
		"-at-", "@",
		"/at/", "@",
		"=at=", "@",
		"?at?", "@",
		"^at^", "@",
		"_at_", "@",
		"`at`", "@",
		"|at|", "@",
		"{at{", "@",
		"}at}", "@",
		"{at}", "@",
		"~at~", "@",
		"\"at\"", "@",
		"(at(", "@",
		")at)", "@",
		"(at)", "@",
		",at,", "@",
		":at:", "@",
		";at;", "@",
		"<at<", "@",
		">at>", "@",
		"<at>", "@",
		"@at@", "@",
		"\\at\\", "@",
		"[at[", "@",
		"]at]", "@",
		"[at]", "@",
		"*at*", "@",

		// one side
		".at", "@",
		"$at", "@",
		"(at", "@",
		",at", "@",
		"*at", "@",

		// ----------------
		// "a" substitutes
		// ----------------
		"!a!", "@",
		"#a#", "@",
		"$a$", "@",
		"%a%", "@",
		"&a&", "@",
		"'a'", "@",
		"*a*", "@",
		"+a+", "@",
		"-a-", "@",
		"/a/", "@",
		"=a=", "@",
		"?a?", "@",
		"^a^", "@",
		"_a_", "@",
		"`a`", "@",
		"|a|", "@",
		"{a{", "@",
		"}a}", "@",
		"{a}", "@",
		"~a~", "@",
		"\"a\"", "@",
		"(a(", "@",
		")a)", "@",
		"(a)", "@",
		",a,", "@",
		":a:", "@",
		";a;", "@",
		"<a<", "@",
		">a>", "@",
		"<a>", "@",
		"@a@", "@",
		"\\a\\", "@",
		"[a[", "@",
		"]a]", "@",
		"[a]", "@",
		"*a*", "@",

		// last hurrahs
		"#", "@",
		"*", "@",
	)

	fixed := r.Replace(email)

	logStep("Find at:", fixed)
	return fixed
}

func findTheDots(email string) string {

	r := strings.NewReplacer(

		// remove duplicates
		"....", ".",
		"...", ".",
		"..", ".",
		". .", ".",

		// some strings may appear inside a name so we will
		// only look for those that have spaces around them
		// NOTE multiples in inverse order

		" dotdotdot ", ".",
		" dotdot ", ".",
		" dot ", ".",
		" otd ", ".",
		" d0t ", ".",
		" dat ", ".",
		" dotsym ", ".",
		" point ", ".",
		" period ", ".",
		" dt ", ".",
		" daught ", ".",

		// misc .
		",", ".",
		"d-o-t", ".",

		// -------------
		// "." substitutes
		// -------------
		"!.!", ".",
		"#.#", ".",
		"$.$", ".",
		"%.%", ".",
		"&.&", ".",
		"'.'", ".",
		"*.*", ".",
		"+.+", ".",
		"-.-", ".",
		"/./", ".",
		"=.=", ".",
		"?.?", ".",
		"^.^", ".",
		"_._", ".",
		"`.`", ".",
		"|.|", ".",
		"{.{", ".",
		"}.}", ".",
		"{.}", ".",
		"~.~", ".",
		"\".\"", ".",
		"(.(", ".",
		").)", ".",
		"(.)", ".",
		",.,", ".",
		":.:", ".",
		";.;", ".",
		"<.<", ".",
		">.>", ".",
		"<.>", ".",
		// "@.@", ".",
		"\\.\\", ".",
		"[.[", ".",
		"].]", ".",
		"[.]", ".",
		"*.*", ".",

		// ----------------
		// "dot" substitutes
		// ----------------
		".dot.", ".",
		"!dot!", ".",
		"#dot#", ".",
		"$dot$", ".",
		"%dot%", ".",
		"&dot&", ".",
		"'dot'", ".",
		"*dot*", ".",
		"+dot+", ".",
		"-dot-", ".",
		"/dot/", ".",
		"=dot=", ".",
		"?dot?", ".",
		"^dot^", ".",
		"_dot_", ".",
		"`dot`", ".",
		"|dot|", ".",
		"{dot{", ".",
		"}dot}", ".",
		"{dot}", ".",
		"~dot~", ".",
		"\"dot\"", ".",
		"(dot(", ".",
		")dot)", ".",
		"(dot)", ".",
		",dot,", ".",
		":dot:", ".",
		";dot;", ".",
		"<dot<", ".",
		">dot>", ".",
		"<dot>", ".",
		"@dot@", ".",
		"\\dot\\", ".",
		"[dot[", ".",
		"]dot]", ".",
		"[dot]", ".",
		"*dot*", ".",

		// ----------------
		// "d0t" substitutes
		// ----------------
		".d0t.", ".",
		"!d0t!", ".",
		"#d0t#", ".",
		"$d0t$", ".",
		"%d0t%", ".",
		"&d0t&", ".",
		"'d0t'", ".",
		"*d0t*", ".",
		"+d0t+", ".",
		"-d0t-", ".",
		"/d0t/", ".",
		"=d0t=", ".",
		"?d0t?", ".",
		"^d0t^", ".",
		"_d0t_", ".",
		"`d0t`", ".",
		"|d0t|", ".",
		"{d0t{", ".",
		"}d0t}", ".",
		"{d0t}", ".",
		"~d0t~", ".",
		"\"d0t\"", ".",
		"(d0t(", ".",
		")d0t)", ".",
		"(d0t)", ".",
		",d0t,", ".",
		":d0t:", ".",
		";d0t;", ".",
		"<d0t<", ".",
		">d0t>", ".",
		"<d0t>", ".",
		"@d0t@", ".",
		"\\d0t\\", ".",
		"[d0t[", ".",
		"]d0t]", ".",
		"[d0t]", ".",
		"*d0t*", ".",

		// ----------------
		// "dt" substitutes
		// ----------------
		".dt.", ".",
		"!dt!", ".",
		"#dt#", ".",
		"$dt$", ".",
		"%dt%", ".",
		"&dt&", ".",
		"'dt'", ".",
		"*dt*", ".",
		"+dt+", ".",
		"-dt-", ".",
		"/dt/", ".",
		"=dt=", ".",
		"?dt?", ".",
		"^dt^", ".",
		"_dt_", ".",
		"`dt`", ".",
		"|dt|", ".",
		"{dt{", ".",
		"}dt}", ".",
		"{dt}", ".",
		"~dt~", ".",
		"\"dt\"", ".",
		"(dt(", ".",
		")dt)", ".",
		"(dt)", ".",
		",dt,", ".",
		":dt:", ".",
		";dt;", ".",
		"<dt<", ".",
		">dt>", ".",
		"<dt>", ".",
		"@dt@", ".",
		"\\dt\\", ".",
		"[dt[", ".",
		"]dt]", ".",
		"[dt]", ".",
		"*dt*", ".",

		// ----------------
		// "d" substitutes
		// ----------------
		".d.", ".",
		"!d!", ".",
		"#d#", ".",
		"$d$", ".",
		"%d%", ".",
		"&d&", ".",
		"'d'", ".",
		"*d*", ".",
		"+d+", ".",
		"-d-", ".",
		"/d/", ".",
		"=d=", ".",
		"?d?", ".",
		"^d^", ".",
		"_d_", ".",
		"`d`", ".",
		"|d|", ".",
		"{d{", ".",
		"}d}", ".",
		"{d}", ".",
		"~d~", ".",
		"\"d\"", ".",
		"(d(", ".",
		")d)", ".",
		"(d)", ".",
		",d,", ".",
		":d:", ".",
		";d;", ".",
		"<d<", ".",
		">d>", ".",
		"<d>", ".",
		"@d@", ".",
		"\\d\\", ".",
		"[d[", ".",
		"]d]", ".",
		"[d]", ".",
		"*d*", ".",
	)

	fixed := r.Replace(email)

	logStep("Find dots:", fixed)
	return fixed
}

func generalFixes(email string) string {

	r := strings.NewReplacer(

		"{{{{", "{",
		"{{{", "{",
		"{{", "{",
		"}}}}", "}",
		"}}}", "}",
		"}}", "}",

		"((((", "(",
		"(((", "(",
		"((", "(",
		"))))", ")",
		")))", ")",
		"))", ")",

		"[[[[", "[",
		"[[[", "[",
		"[[", "[",
		"]]]]", "]",
		"]]]", "]",
		"]]", "]",

		"++++", "+",
		"+++", "+",
		"++", "+",

		// bad endings
		"@com", ".com",
		"@org", ".org",
		"@net", ".net",
		"@edu", ".edu",

		// misc
		".gmail.com", "@gmail.com",
		"g.mail.com", "gmail.com",
		".@gmail.com", "@gmail.com",
		"gmail.com.com", "gmail.com",
		"gmail com", "gmail.com",
		"_gmail_com", "@gmail.com",
		"gmail_com", "gmail.com",
		"atgmail.com", "@gmail.com",
		"atgmaildotcom", "@gmail.com",
		"google email", "google.com",
		"gee mail", "gmail",
		"ge mail", "gmail",
		"g mail", "gmail",

		// hotmail
		"_hotmail.com", "@hotmail.com",
		".hotmail.com", "@hotmail.com",
		"_hotmail_com", "@hotmail.com",
		"athotmail.com", "@hotmail.com",

		"dotcalm", ".com",
		"dat.com", ".com",
		"dot calm", ".com",
		"dat com", ".com",
		"dat calm", ".com",
		" calm", ".com",
		".calm", ".com",
		"dotcom", ".com",
		"@.com", ".com",

		"atnospam", "@",
		"nospam", "",
		"n0spam", "",
		"n0_spam", "",
		"no_spam", "",
		"n0-spam", "",
		"no-spam", "",
		"n*o*s*p*a*m", "",
		"n*0*s*p*a*m", "",

		"qq@.com", "@qq.com",
		"@.qq.com", "@qq.com",
		".qq.com", "@qq.com",

		"163@.com", "@163.com",
		"@.163.com", "@163.com",
		".163.com", "@163.com",
	)

	fixed := r.Replace(email)

	logStep("General fixes:", fixed)
	return fixed
}

func logStep(step, output string) {
	if logSteps {
		paddedStep, err := padding(step, 14, " ")
		if err != nil {
			log.Println(err)
		}
		log.Println(paddedStep + " " + output)
	}
}

// padding will pad a string out to a defined length using specified character
func padding(s string, length int, padChar string) (string, error) {

	// Check for errors
	if len(s) > length {
		return "", errors.New("string is too long")
	}
	if len(padChar) > 1 {
		return "", errors.New("padding must be only one character")
	}

	// use builder (5x speedup)
	var b strings.Builder
	b.Grow(length)
	b.WriteString(s)

	paddingNeeded := length - len(s)
	for i := 0; i < paddingNeeded; i++ {
		b.WriteString(padChar)
	}

	return b.String(), nil
}
