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
  case "ahom":
    return [10]rune{'𑜰','𑜱','𑜲','𑜳','𑜴','𑜵','𑜶','𑜷','𑜸','𑜹',}
  case "arab":
    return [10]rune{'٠','١','٢','٣','٤','٥','٦','٧','٨','٩',}
  case "arabext":
    return [10]rune{'۰','۱','۲','۳','۴','۵','۶','۷','۸','۹',}
  case "bali":
    return [10]rune{'᭐','᭑','᭒','᭓','᭔','᭕','᭖','᭗','᭘','᭙',}
  case "beng":
    return [10]rune{'০','১','২','৩','৪','৫','৬','৭','৮','৯',}
  case "bhks":
    return [10]rune{'𑱐','𑱑','𑱒','𑱓','𑱔','𑱕','𑱖','𑱗','𑱘','𑱙',}
  case "brah":
    return [10]rune{'𑁦','𑁧','𑁨','𑁩','𑁪','𑁫','𑁬','𑁭','𑁮','𑁯',}
  case "cakm":
    return [10]rune{'𑄶','𑄷','𑄸','𑄹','𑄺','𑄻','𑄼','𑄽','𑄾','𑄿',}
  case "cham":
    return [10]rune{'꩐','꩑','꩒','꩓','꩔','꩕','꩖','꩗','꩘','꩙',}
  case "deva":
    return [10]rune{'०','१','२','३','४','५','६','७','८','९',}
  case "diak":
    return [10]rune{'𑥐','𑥑','𑥒','𑥓','𑥔','𑥕','𑥖','𑥗','𑥘','𑥙',}
  case "fullwide":
    return [10]rune{'０','１','２','３','４','５','６','７','８','９',}
  case "gong":
    return [10]rune{'𑶠','𑶡','𑶢','𑶣','𑶤','𑶥','𑶦','𑶧','𑶨','𑶩',}
  case "gonm":
    return [10]rune{'𑵐','𑵑','𑵒','𑵓','𑵔','𑵕','𑵖','𑵗','𑵘','𑵙',}
  case "gujr":
    return [10]rune{'૦','૧','૨','૩','૪','૫','૬','૭','૮','૯',}
  case "guru":
    return [10]rune{'੦','੧','੨','੩','੪','੫','੬','੭','੮','੯',}
  case "hanidec":
    return [10]rune{'〇','一','二','三','四','五','六','七','八','九',}
  case "hmng":
    return [10]rune{'𖭐','𖭑','𖭒','𖭓','𖭔','𖭕','𖭖','𖭗','𖭘','𖭙',}
  case "hmnp":
    return [10]rune{'𞅀','𞅁','𞅂','𞅃','𞅄','𞅅','𞅆','𞅇','𞅈','𞅉',}
  case "java":
    return [10]rune{'꧐','꧑','꧒','꧓','꧔','꧕','꧖','꧗','꧘','꧙',}
  case "kali":
    return [10]rune{'꤀','꤁','꤂','꤃','꤄','꤅','꤆','꤇','꤈','꤉',}
  case "kawi":
    return [10]rune{'𑽐','𑽑','𑽒','𑽓','𑽔','𑽕','𑽖','𑽗','𑽘','𑽙',}
  case "khmr":
    return [10]rune{'០','១','២','៣','៤','៥','៦','៧','៨','៩',}
  case "knda":
    return [10]rune{'೦','೧','೨','೩','೪','೫','೬','೭','೮','೯',}
  case "lana":
    return [10]rune{'᪀','᪁','᪂','᪃','᪄','᪅','᪆','᪇','᪈','᪉',}
  case "lanatham":
    return [10]rune{'᪐','᪑','᪒','᪓','᪔','᪕','᪖','᪗','᪘','᪙',}
  case "laoo":
    return [10]rune{'໐','໑','໒','໓','໔','໕','໖','໗','໘','໙',}
  case "latn":
    return [10]rune{'0','1','2','3','4','5','6','7','8','9',}
  case "lepc":
    return [10]rune{'᱀','᱁','᱂','᱃','᱄','᱅','᱆','᱇','᱈','᱉',}
  case "limb":
    return [10]rune{'᥆','᥇','᥈','᥉','᥊','᥋','᥌','᥍','᥎','᥏',}
  case "mathbold":
    return [10]rune{'𝟎','𝟏','𝟐','𝟑','𝟒','𝟓','𝟔','𝟕','𝟖','𝟗',}
  case "mathdbl":
    return [10]rune{'𝟘','𝟙','𝟚','𝟛','𝟜','𝟝','𝟞','𝟟','𝟠','𝟡',}
  case "mathmono":
    return [10]rune{'𝟶','𝟷','𝟸','𝟹','𝟺','𝟻','𝟼','𝟽','𝟾','𝟿',}
  case "mathsanb":
    return [10]rune{'𝟬','𝟭','𝟮','𝟯','𝟰','𝟱','𝟲','𝟳','𝟴','𝟵',}
  case "mathsans":
    return [10]rune{'𝟢','𝟣','𝟤','𝟥','𝟦','𝟧','𝟨','𝟩','𝟪','𝟫',}
  case "mlym":
    return [10]rune{'൦','൧','൨','൩','൪','൫','൬','൭','൮','൯',}
  case "modi":
    return [10]rune{'𑙐','𑙑','𑙒','𑙓','𑙔','𑙕','𑙖','𑙗','𑙘','𑙙',}
  case "mong":
    return [10]rune{'᠐','᠑','᠒','᠓','᠔','᠕','᠖','᠗','᠘','᠙',}
  case "mroo":
    return [10]rune{'𖩠','𖩡','𖩢','𖩣','𖩤','𖩥','𖩦','𖩧','𖩨','𖩩',}
  case "mtei":
    return [10]rune{'꯰','꯱','꯲','꯳','꯴','꯵','꯶','꯷','꯸','꯹',}
  case "mymr":
    return [10]rune{'၀','၁','၂','၃','၄','၅','၆','၇','၈','၉',}
  case "mymrshan":
    return [10]rune{'႐','႑','႒','႓','႔','႕','႖','႗','႘','႙',}
  case "mymrtlng":
    return [10]rune{'꧰','꧱','꧲','꧳','꧴','꧵','꧶','꧷','꧸','꧹',}
  case "nagm":
    return [10]rune{'𞓰','𞓱','𞓲','𞓳','𞓴','𞓵','𞓶','𞓷','𞓸','𞓹',}
  case "newa":
    return [10]rune{'𑑐','𑑑','𑑒','𑑓','𑑔','𑑕','𑑖','𑑗','𑑘','𑑙',}
  case "nkoo":
    return [10]rune{'߀','߁','߂','߃','߄','߅','߆','߇','߈','߉',}
  case "olck":
    return [10]rune{'᱐','᱑','᱒','᱓','᱔','᱕','᱖','᱗','᱘','᱙',}
  case "orya":
    return [10]rune{'୦','୧','୨','୩','୪','୫','୬','୭','୮','୯',}
  case "osma":
    return [10]rune{'𐒠','𐒡','𐒢','𐒣','𐒤','𐒥','𐒦','𐒧','𐒨','𐒩',}
  case "rohg":
    return [10]rune{'𐴰','𐴱','𐴲','𐴳','𐴴','𐴵','𐴶','𐴷','𐴸','𐴹',}
  case "saur":
    return [10]rune{'꣐','꣑','꣒','꣓','꣔','꣕','꣖','꣗','꣘','꣙',}
  case "segment":
    return [10]rune{'🯰','🯱','🯲','🯳','🯴','🯵','🯶','🯷','🯸','🯹',}
  case "shrd":
    return [10]rune{'𑇐','𑇑','𑇒','𑇓','𑇔','𑇕','𑇖','𑇗','𑇘','𑇙',}
  case "sind":
    return [10]rune{'𑋰','𑋱','𑋲','𑋳','𑋴','𑋵','𑋶','𑋷','𑋸','𑋹',}
  case "sinh":
    return [10]rune{'෦','෧','෨','෩','෪','෫','෬','෭','෮','෯',}
  case "sora":
    return [10]rune{'𑃰','𑃱','𑃲','𑃳','𑃴','𑃵','𑃶','𑃷','𑃸','𑃹',}
  case "sund":
    return [10]rune{'᮰','᮱','᮲','᮳','᮴','᮵','᮶','᮷','᮸','᮹',}
  case "takr":
    return [10]rune{'𑛀','𑛁','𑛂','𑛃','𑛄','𑛅','𑛆','𑛇','𑛈','𑛉',}
  case "talu":
    return [10]rune{'᧐','᧑','᧒','᧓','᧔','᧕','᧖','᧗','᧘','᧙',}
  case "tamldec":
    return [10]rune{'௦','௧','௨','௩','௪','௫','௬','௭','௮','௯',}
  case "tnsa":
    return [10]rune{'𖫀','𖫁','𖫂','𖫃','𖫄','𖫅','𖫆','𖫇','𖫈','𖫉',}
  case "telu":
    return [10]rune{'౦','౧','౨','౩','౪','౫','౬','౭','౮','౯',}
  case "thai":
    return [10]rune{'๐','๑','๒','๓','๔','๕','๖','๗','๘','๙',}
  case "tibt":
    return [10]rune{'༠','༡','༢','༣','༤','༥','༦','༧','༨','༩',}
  case "tirh":
    return [10]rune{'𑓐','𑓑','𑓒','𑓓','𑓔','𑓕','𑓖','𑓗','𑓘','𑓙',}
  case "vaii":
    return [10]rune{'꘠','꘡','꘢','꘣','꘤','꘥','꘦','꘧','꘨','꘩',}
  case "wara":
    return [10]rune{'𑣠','𑣡','𑣢','𑣣','𑣤','𑣥','𑣦','𑣧','𑣨','𑣩',}
  case "wcho":
    return [10]rune{'𞋰','𞋱','𞋲','𞋳','𞋴','𞋵','𞋶','𞋷','𞋸','𞋹',}
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