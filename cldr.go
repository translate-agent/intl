// Code generated by "earthly +cldr". DO NOT EDIT.
package intl

import (
  "strings"

  "golang.org/x/text/language"
)

type numberingSystem int

const (
  numberingSystemLatn numberingSystem = iota
  numberingSystemAdlm
  numberingSystemArab
  numberingSystemArabext
  numberingSystemBeng
  numberingSystemCakm
  numberingSystemDeva
  numberingSystemHmnp
  numberingSystemMtei
  numberingSystemMymr
  numberingSystemNkoo
  numberingSystemOlck
  numberingSystemTibt
  numberingSystemLast
)

var numberingSystems = []digits{
  numberingSystemAdlm: {'𞥐','𞥑','𞥒','𞥓','𞥔','𞥕','𞥖','𞥗','𞥘','𞥙',},
  numberingSystemArab: {'٠','١','٢','٣','٤','٥','٦','٧','٨','٩',},
  numberingSystemArabext: {'۰','۱','۲','۳','۴','۵','۶','۷','۸','۹',},
  numberingSystemBeng: {'০','১','২','৩','৪','৫','৬','৭','৮','৯',},
  numberingSystemCakm: {'𑄶','𑄷','𑄸','𑄹','𑄺','𑄻','𑄼','𑄽','𑄾','𑄿',},
  numberingSystemDeva: {'०','१','२','३','४','५','६','७','८','९',},
  numberingSystemHmnp: {'𞅀','𞅁','𞅂','𞅃','𞅄','𞅅','𞅆','𞅇','𞅈','𞅉',},
  numberingSystemMtei: {'꯰','꯱','꯲','꯳','꯴','꯵','꯶','꯷','꯸','꯹',},
  numberingSystemMymr: {'၀','၁','၂','၃','၄','၅','၆','၇','၈','၉',},
  numberingSystemNkoo: {'߀','߁','߂','߃','߄','߅','߆','߇','߈','߉',},
  numberingSystemOlck: {'᱐','᱑','᱒','᱓','᱔','᱕','᱖','᱗','᱘','᱙',},
  numberingSystemTibt: {'༠','༡','༢','༣','༤','༥','༦','༧','༨','༩',},
}

func defaultNumberingSystem(locale language.Tag) numberingSystem {
  s := locale.String()

  switch {
  default:
    return numberingSystemLatn
  case s == "ff-Adlm", strings.HasPrefix(s, "ff-Adlm-"):
    return numberingSystemAdlm
  case s == "ar", strings.HasPrefix(s, "ar-"), s == "ar-BH", strings.HasPrefix(s, "ar-BH-"), s == "ar-DJ", strings.HasPrefix(s, "ar-DJ-"), s == "ar-EG", strings.HasPrefix(s, "ar-EG-"), s == "ar-ER", strings.HasPrefix(s, "ar-ER-"), s == "ar-IL", strings.HasPrefix(s, "ar-IL-"), s == "ar-IQ", strings.HasPrefix(s, "ar-IQ-"), s == "ar-JO", strings.HasPrefix(s, "ar-JO-"), s == "ar-KM", strings.HasPrefix(s, "ar-KM-"), s == "ar-KW", strings.HasPrefix(s, "ar-KW-"), s == "ar-LB", strings.HasPrefix(s, "ar-LB-"), s == "ar-MR", strings.HasPrefix(s, "ar-MR-"), s == "ar-OM", strings.HasPrefix(s, "ar-OM-"), s == "ar-PS", strings.HasPrefix(s, "ar-PS-"), s == "ar-QA", strings.HasPrefix(s, "ar-QA-"), s == "ar-SA", strings.HasPrefix(s, "ar-SA-"), s == "ar-SD", strings.HasPrefix(s, "ar-SD-"), s == "ar-SO", strings.HasPrefix(s, "ar-SO-"), s == "ar-SS", strings.HasPrefix(s, "ar-SS-"), s == "ar-SY", strings.HasPrefix(s, "ar-SY-"), s == "ar-TD", strings.HasPrefix(s, "ar-TD-"), s == "ar-YE", strings.HasPrefix(s, "ar-YE-"), s == "ckb", strings.HasPrefix(s, "ckb-"), s == "sd", strings.HasPrefix(s, "sd-"), s == "sdh", strings.HasPrefix(s, "sdh-"):
    return numberingSystemArab
  case s == "fa", strings.HasPrefix(s, "fa-"), s == "ks", strings.HasPrefix(s, "ks-"), s == "lrc", strings.HasPrefix(s, "lrc-"), s == "pa-Arab", strings.HasPrefix(s, "pa-Arab-"), s == "ps", strings.HasPrefix(s, "ps-"), s == "ur-IN", strings.HasPrefix(s, "ur-IN-"), s == "uz-Arab", strings.HasPrefix(s, "uz-Arab-"):
    return numberingSystemArabext
  case s == "as", strings.HasPrefix(s, "as-"), s == "bn", strings.HasPrefix(s, "bn-"), s == "mni", strings.HasPrefix(s, "mni-"):
    return numberingSystemBeng
  case s == "ccp", strings.HasPrefix(s, "ccp-"):
    return numberingSystemCakm
  case s == "bgc", strings.HasPrefix(s, "bgc-"), s == "bho", strings.HasPrefix(s, "bho-"), s == "mr", strings.HasPrefix(s, "mr-"), s == "ne", strings.HasPrefix(s, "ne-"), s == "raj", strings.HasPrefix(s, "raj-"), s == "sa", strings.HasPrefix(s, "sa-"), s == "sat-Deva", strings.HasPrefix(s, "sat-Deva-"):
    return numberingSystemDeva
  case s == "hnj", strings.HasPrefix(s, "hnj-"):
    return numberingSystemHmnp
  case s == "root", strings.HasPrefix(s, "root-"), s == "ar-AE", strings.HasPrefix(s, "ar-AE-"), s == "ar-DZ", strings.HasPrefix(s, "ar-DZ-"), s == "ar-EH", strings.HasPrefix(s, "ar-EH-"), s == "ar-LY", strings.HasPrefix(s, "ar-LY-"), s == "ar-MA", strings.HasPrefix(s, "ar-MA-"), s == "ar-TN", strings.HasPrefix(s, "ar-TN-"):
    return numberingSystemLatn
  case s == "mni-Mtei", strings.HasPrefix(s, "mni-Mtei-"):
    return numberingSystemMtei
  case s == "my", strings.HasPrefix(s, "my-"):
    return numberingSystemMymr
  case s == "nqo", strings.HasPrefix(s, "nqo-"):
    return numberingSystemNkoo
  case s == "sat", strings.HasPrefix(s, "sat-"):
    return numberingSystemOlck
  case s == "dz", strings.HasPrefix(s, "dz-"):
    return numberingSystemTibt
  }
}

func defaultCalendar(locale language.Tag) string {
  switch v, _ := locale.Region(); v.String() {
  default:
    return "gregorian"
  case "AF", "IR":
    return "persian"
  case "SA":
    return "islamic-umalqura"
  case "TH":
    return "buddhist"
  }
}

func fmtYear(y string, locale language.Tag) string {
  lang, _ := locale.Base()

  switch lang.String() {
  default:
    return y
  case "prg":
    return y+" m."
  case "lv":
    return y+". g."
  case "bs", "hr", "hu", "sr":
    return y+"."
  case "bg":
    return y+" г."
  case "ja", "yue", "zh":
    return y+"年"
  case "ko":
    return y+"년"
  }
}