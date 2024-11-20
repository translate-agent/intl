package intl

import "golang.org/x/text/language"

func defaultNumberingSystem(locale language.Tag) numberingSystem {
	switch locale.String() {
	default:
		return numberingSystemLatn
	case "ar", "ar-001", "ar-BH", "ar-DJ", "ar-EG", "ar-ER", "ar-IL", "ar-IQ", "ar-JO", "ar-KM", "ar-KW", "ar-LB",
		"ar-MR", "ar-OM", "ar-PS", "ar-QA", "ar-SA", "ar-SD", "ar-SO", "ar-SS", "ar-SY", "ar-TD", "ar-YE", "ckb", "ckb-IR",
		"ckb-IQ", "sd", "sd-Arab", "sdh":
		return numberingSystemArab
	case "ar-AE", "ar-DZ", "ar-EH", "ar-LY", "ar-MA", "ar-TN":
		return numberingSystemLatn
	case "as", "as-IN", "bn", "bn-BD", "bn-IN", "mni", "mni-Beng", "mni-Mtei":
		return numberingSystemBeng
	case "bgn", "fa", "fa-AF", "fa-IR", "ks", "ks-Arab", "lrc", "lrc-IQ", "lrc-IR", "mzn", "mzn-IR",
		"pa-Arab", "ps", "ps-AF", "ps-PK", "ur-IN", "uz-Arab":
		return numberingSystemArabext
	case "bgc", "bgc-IN", "bho", "bho-IN", "mr", "mr-IN", "ne", "ne-IN", "ne-NP", "raj", "raj-IN", "sa", "sa-IN":
		return numberingSystemDeva
	case "ccp", "ccp-BD", "ccp-IN":
		return numberingSystemCakm
	case "dz", "dz-BT":
		return numberingSystemTibt
	case "ff-Adlm":
		return numberingSystemAdlm
	case "hnj":
		return numberingSystemHmnp
	case "my", "my-MM":
		return numberingSystemMymr
	case "nqo", "nqo-GN":
		return numberingSystemNkoo
	case "sat", "sat-Deva", "sat-Olck":
		return numberingSystemOlck
	}
}
