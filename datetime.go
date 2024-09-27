// DO NOT EDIT!
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

var numberingSystems = [][10]rune{
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
  case s == "root" || strings.HasPrefix(s, "root-"):
    return numberingSystemLatn
  case s == "ar" || strings.HasPrefix(s, "ar-"):
    return numberingSystemArab
  case s == "ar-AE" || strings.HasPrefix(s, "ar-AE-"):
    return numberingSystemLatn
  case s == "ar-BH" || strings.HasPrefix(s, "ar-BH-"):
    return numberingSystemArab
  case s == "ar-DJ" || strings.HasPrefix(s, "ar-DJ-"):
    return numberingSystemArab
  case s == "ar-DZ" || strings.HasPrefix(s, "ar-DZ-"):
    return numberingSystemLatn
  case s == "ar-EG" || strings.HasPrefix(s, "ar-EG-"):
    return numberingSystemArab
  case s == "ar-EH" || strings.HasPrefix(s, "ar-EH-"):
    return numberingSystemLatn
  case s == "ar-ER" || strings.HasPrefix(s, "ar-ER-"):
    return numberingSystemArab
  case s == "ar-IL" || strings.HasPrefix(s, "ar-IL-"):
    return numberingSystemArab
  case s == "ar-IQ" || strings.HasPrefix(s, "ar-IQ-"):
    return numberingSystemArab
  case s == "ar-JO" || strings.HasPrefix(s, "ar-JO-"):
    return numberingSystemArab
  case s == "ar-KM" || strings.HasPrefix(s, "ar-KM-"):
    return numberingSystemArab
  case s == "ar-KW" || strings.HasPrefix(s, "ar-KW-"):
    return numberingSystemArab
  case s == "ar-LB" || strings.HasPrefix(s, "ar-LB-"):
    return numberingSystemArab
  case s == "ar-LY" || strings.HasPrefix(s, "ar-LY-"):
    return numberingSystemLatn
  case s == "ar-MA" || strings.HasPrefix(s, "ar-MA-"):
    return numberingSystemLatn
  case s == "ar-MR" || strings.HasPrefix(s, "ar-MR-"):
    return numberingSystemArab
  case s == "ar-OM" || strings.HasPrefix(s, "ar-OM-"):
    return numberingSystemArab
  case s == "ar-PS" || strings.HasPrefix(s, "ar-PS-"):
    return numberingSystemArab
  case s == "ar-QA" || strings.HasPrefix(s, "ar-QA-"):
    return numberingSystemArab
  case s == "ar-SA" || strings.HasPrefix(s, "ar-SA-"):
    return numberingSystemArab
  case s == "ar-SD" || strings.HasPrefix(s, "ar-SD-"):
    return numberingSystemArab
  case s == "ar-SO" || strings.HasPrefix(s, "ar-SO-"):
    return numberingSystemArab
  case s == "ar-SS" || strings.HasPrefix(s, "ar-SS-"):
    return numberingSystemArab
  case s == "ar-SY" || strings.HasPrefix(s, "ar-SY-"):
    return numberingSystemArab
  case s == "ar-TD" || strings.HasPrefix(s, "ar-TD-"):
    return numberingSystemArab
  case s == "ar-TN" || strings.HasPrefix(s, "ar-TN-"):
    return numberingSystemLatn
  case s == "ar-YE" || strings.HasPrefix(s, "ar-YE-"):
    return numberingSystemArab
  case s == "as" || strings.HasPrefix(s, "as-"):
    return numberingSystemBeng
  case s == "bgc" || strings.HasPrefix(s, "bgc-"):
    return numberingSystemDeva
  case s == "bho" || strings.HasPrefix(s, "bho-"):
    return numberingSystemDeva
  case s == "bn" || strings.HasPrefix(s, "bn-"):
    return numberingSystemBeng
  case s == "ccp" || strings.HasPrefix(s, "ccp-"):
    return numberingSystemCakm
  case s == "ckb" || strings.HasPrefix(s, "ckb-"):
    return numberingSystemArab
  case s == "dz" || strings.HasPrefix(s, "dz-"):
    return numberingSystemTibt
  case s == "fa" || strings.HasPrefix(s, "fa-"):
    return numberingSystemArabext
  case s == "ff-Adlm" || strings.HasPrefix(s, "ff-Adlm-"):
    return numberingSystemAdlm
  case s == "hnj" || strings.HasPrefix(s, "hnj-"):
    return numberingSystemHmnp
  case s == "ks" || strings.HasPrefix(s, "ks-"):
    return numberingSystemArabext
  case s == "lrc" || strings.HasPrefix(s, "lrc-"):
    return numberingSystemArabext
  case s == "mni" || strings.HasPrefix(s, "mni-"):
    return numberingSystemBeng
  case s == "mni-Mtei" || strings.HasPrefix(s, "mni-Mtei-"):
    return numberingSystemMtei
  case s == "mr" || strings.HasPrefix(s, "mr-"):
    return numberingSystemDeva
  case s == "my" || strings.HasPrefix(s, "my-"):
    return numberingSystemMymr
  case s == "ne" || strings.HasPrefix(s, "ne-"):
    return numberingSystemDeva
  case s == "nqo" || strings.HasPrefix(s, "nqo-"):
    return numberingSystemNkoo
  case s == "pa-Arab" || strings.HasPrefix(s, "pa-Arab-"):
    return numberingSystemArabext
  case s == "ps" || strings.HasPrefix(s, "ps-"):
    return numberingSystemArabext
  case s == "raj" || strings.HasPrefix(s, "raj-"):
    return numberingSystemDeva
  case s == "sa" || strings.HasPrefix(s, "sa-"):
    return numberingSystemDeva
  case s == "sat" || strings.HasPrefix(s, "sat-"):
    return numberingSystemOlck
  case s == "sat-Deva" || strings.HasPrefix(s, "sat-Deva-"):
    return numberingSystemDeva
  case s == "sd" || strings.HasPrefix(s, "sd-"):
    return numberingSystemArab
  case s == "sdh" || strings.HasPrefix(s, "sdh-"):
    return numberingSystemArab
  case s == "ur-IN" || strings.HasPrefix(s, "ur-IN-"):
    return numberingSystemArabext
  case s == "uz-Arab" || strings.HasPrefix(s, "uz-Arab-"):
    return numberingSystemArabext
  }
}


func calendarPreferences(locale language.Tag) []string {
  switch v, _ := locale.Region(); v.String() {
  default:
    return nil
  case "001":
    return []string{"gregorian"}
  case "BD", "DJ", "DZ", "EH", "ER", "ID", "IQ", "JO", "KM", "LB", "LY", "MA", "MR", "MY", "NE", "OM", "PK", "PS", "SD", "SY", "TD", "TN", "YE":
    return []string{"gregorian", "islamic", "islamic-civil", "islamic-tbla"}
  case "AL", "AZ", "MV", "TJ", "TM", "TR", "UZ", "XK":
    return []string{"gregorian", "islamic-civil", "islamic-tbla"}
  case "AE", "BH", "KW", "QA":
    return []string{"gregorian", "islamic-umalqura", "islamic", "islamic-civil", "islamic-tbla"}
  case "AF", "IR":
    return []string{"persian", "gregorian", "islamic", "islamic-civil", "islamic-tbla"}
  case "CN", "CX", "HK", "MO", "SG":
    return []string{"gregorian", "chinese"}
  case "EG":
    return []string{"gregorian", "coptic", "islamic", "islamic-civil", "islamic-tbla"}
  case "ET":
    return []string{"gregorian", "ethiopic"}
  case "IL":
    return []string{"gregorian", "hebrew", "islamic", "islamic-civil", "islamic-tbla"}
  case "IN":
    return []string{"gregorian", "indian"}
  case "JP":
    return []string{"gregorian", "japanese"}
  case "KR":
    return []string{"gregorian", "dangi"}
  case "SA":
    return []string{"islamic-umalqura", "gregorian", "islamic", "islamic-rgsa"}
  case "TH":
    return []string{"buddhist", "gregorian"}
  case "TW":
    return []string{"gregorian", "roc", "chinese"}
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