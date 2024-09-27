// DO NOT EDIT!
package intl

import (
  "strings"

  "golang.org/x/text/language"
)

func numberingSystem(id string) [10]rune {
  switch id {
  default:
    return [10]rune{}
  case "adlm":
    return [10]rune{'𞥐','𞥑','𞥒','𞥓','𞥔','𞥕','𞥖','𞥗','𞥘','𞥙',}
  case "arab":
    return [10]rune{'٠','١','٢','٣','٤','٥','٦','٧','٨','٩',}
  case "arabext":
    return [10]rune{'۰','۱','۲','۳','۴','۵','۶','۷','۸','۹',}
  case "beng":
    return [10]rune{'০','১','২','৩','৪','৫','৬','৭','৮','৯',}
  case "cakm":
    return [10]rune{'𑄶','𑄷','𑄸','𑄹','𑄺','𑄻','𑄼','𑄽','𑄾','𑄿',}
  case "deva":
    return [10]rune{'०','१','२','३','४','५','६','७','८','९',}
  case "hmnp":
    return [10]rune{'𞅀','𞅁','𞅂','𞅃','𞅄','𞅅','𞅆','𞅇','𞅈','𞅉',}
  case "latn":
    return [10]rune{'0','1','2','3','4','5','6','7','8','9',}
  case "mtei":
    return [10]rune{'꯰','꯱','꯲','꯳','꯴','꯵','꯶','꯷','꯸','꯹',}
  case "mymr":
    return [10]rune{'၀','၁','၂','၃','၄','၅','၆','၇','၈','၉',}
  case "nkoo":
    return [10]rune{'߀','߁','߂','߃','߄','߅','߆','߇','߈','߉',}
  case "olck":
    return [10]rune{'᱐','᱑','᱒','᱓','᱔','᱕','᱖','᱗','᱘','᱙',}
  case "tibt":
    return [10]rune{'༠','༡','༢','༣','༤','༥','༦','༧','༨','༩',}
  }
}

func defaultNumberingSystem(locale language.Tag) string {
  s := locale.String()

  switch {
  default:
    return "latn"
  case s == "root" || strings.HasPrefix(s, "root-"):
    return "latn"
  case s == "ar" || strings.HasPrefix(s, "ar-"):
    return "arab"
  case s == "ar-AE" || strings.HasPrefix(s, "ar-AE-"):
    return "latn"
  case s == "ar-BH" || strings.HasPrefix(s, "ar-BH-"):
    return "arab"
  case s == "ar-DJ" || strings.HasPrefix(s, "ar-DJ-"):
    return "arab"
  case s == "ar-DZ" || strings.HasPrefix(s, "ar-DZ-"):
    return "latn"
  case s == "ar-EG" || strings.HasPrefix(s, "ar-EG-"):
    return "arab"
  case s == "ar-EH" || strings.HasPrefix(s, "ar-EH-"):
    return "latn"
  case s == "ar-ER" || strings.HasPrefix(s, "ar-ER-"):
    return "arab"
  case s == "ar-IL" || strings.HasPrefix(s, "ar-IL-"):
    return "arab"
  case s == "ar-IQ" || strings.HasPrefix(s, "ar-IQ-"):
    return "arab"
  case s == "ar-JO" || strings.HasPrefix(s, "ar-JO-"):
    return "arab"
  case s == "ar-KM" || strings.HasPrefix(s, "ar-KM-"):
    return "arab"
  case s == "ar-KW" || strings.HasPrefix(s, "ar-KW-"):
    return "arab"
  case s == "ar-LB" || strings.HasPrefix(s, "ar-LB-"):
    return "arab"
  case s == "ar-LY" || strings.HasPrefix(s, "ar-LY-"):
    return "latn"
  case s == "ar-MA" || strings.HasPrefix(s, "ar-MA-"):
    return "latn"
  case s == "ar-MR" || strings.HasPrefix(s, "ar-MR-"):
    return "arab"
  case s == "ar-OM" || strings.HasPrefix(s, "ar-OM-"):
    return "arab"
  case s == "ar-PS" || strings.HasPrefix(s, "ar-PS-"):
    return "arab"
  case s == "ar-QA" || strings.HasPrefix(s, "ar-QA-"):
    return "arab"
  case s == "ar-SA" || strings.HasPrefix(s, "ar-SA-"):
    return "arab"
  case s == "ar-SD" || strings.HasPrefix(s, "ar-SD-"):
    return "arab"
  case s == "ar-SO" || strings.HasPrefix(s, "ar-SO-"):
    return "arab"
  case s == "ar-SS" || strings.HasPrefix(s, "ar-SS-"):
    return "arab"
  case s == "ar-SY" || strings.HasPrefix(s, "ar-SY-"):
    return "arab"
  case s == "ar-TD" || strings.HasPrefix(s, "ar-TD-"):
    return "arab"
  case s == "ar-TN" || strings.HasPrefix(s, "ar-TN-"):
    return "latn"
  case s == "ar-YE" || strings.HasPrefix(s, "ar-YE-"):
    return "arab"
  case s == "as" || strings.HasPrefix(s, "as-"):
    return "beng"
  case s == "bgc" || strings.HasPrefix(s, "bgc-"):
    return "deva"
  case s == "bho" || strings.HasPrefix(s, "bho-"):
    return "deva"
  case s == "bn" || strings.HasPrefix(s, "bn-"):
    return "beng"
  case s == "ccp" || strings.HasPrefix(s, "ccp-"):
    return "cakm"
  case s == "ckb" || strings.HasPrefix(s, "ckb-"):
    return "arab"
  case s == "dz" || strings.HasPrefix(s, "dz-"):
    return "tibt"
  case s == "fa" || strings.HasPrefix(s, "fa-"):
    return "arabext"
  case s == "ff-Adlm" || strings.HasPrefix(s, "ff-Adlm-"):
    return "adlm"
  case s == "hnj" || strings.HasPrefix(s, "hnj-"):
    return "hmnp"
  case s == "ks" || strings.HasPrefix(s, "ks-"):
    return "arabext"
  case s == "lrc" || strings.HasPrefix(s, "lrc-"):
    return "arabext"
  case s == "mni" || strings.HasPrefix(s, "mni-"):
    return "beng"
  case s == "mni-Mtei" || strings.HasPrefix(s, "mni-Mtei-"):
    return "mtei"
  case s == "mr" || strings.HasPrefix(s, "mr-"):
    return "deva"
  case s == "my" || strings.HasPrefix(s, "my-"):
    return "mymr"
  case s == "ne" || strings.HasPrefix(s, "ne-"):
    return "deva"
  case s == "nqo" || strings.HasPrefix(s, "nqo-"):
    return "nkoo"
  case s == "pa-Arab" || strings.HasPrefix(s, "pa-Arab-"):
    return "arabext"
  case s == "ps" || strings.HasPrefix(s, "ps-"):
    return "arabext"
  case s == "raj" || strings.HasPrefix(s, "raj-"):
    return "deva"
  case s == "sa" || strings.HasPrefix(s, "sa-"):
    return "deva"
  case s == "sat" || strings.HasPrefix(s, "sat-"):
    return "olck"
  case s == "sat-Deva" || strings.HasPrefix(s, "sat-Deva-"):
    return "deva"
  case s == "sd" || strings.HasPrefix(s, "sd-"):
    return "arab"
  case s == "sdh" || strings.HasPrefix(s, "sdh-"):
    return "arab"
  case s == "ur-IN" || strings.HasPrefix(s, "ur-IN-"):
    return "arabext"
  case s == "uz-Arab" || strings.HasPrefix(s, "uz-Arab-"):
    return "arabext"
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