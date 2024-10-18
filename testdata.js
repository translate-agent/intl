const fs = require("fs");

// ICU does not support locales (ICU falls back to system LANG):
// - aa ab an ann apc arn
// - ba bal bew bgn blt bss byn
// - cad cch cho cic co cu
// - dv
// - frr
// - gaa gez gn
// - hnj
// - ife io iu
// - jbo
// - kaj kcg ken kpe
// - la
// - mdf mic moh mus myv
// - nr nso nv ny
// - osa
// - pap pis
// - quc
// - rhg rif
// - scn sdh shn sid skr sma smj sms ss ssy st
// - tig tn tpi trv trw ts tyv
// - ve vo
// - wa wal wbp
const locales = `
			af_ZA agq_CM ak_GH am_ET ar_001 as_IN
			asa_TZ ast_ES az_Arab_IR az_Cyrl_AZ az_Latn az_Latn_AZ
			bas_CM be_BY bem_ZM bez_TZ bg_BG
			bgc_IN bho_IN blo_BJ bm_ML bm_Nkoo_ML bn_BD bo_CN br_FR brx_IN
			bs_Cyrl_BA bs_Latn bs_Latn_BA
			ca_ES ccp_BD ce_RU ceb_PH cgg_UG chr_US ckb_IQ
			cs_CZ csw_CA cv_RU cy_GB
			da_DK dav_KE de_DE dje_NE doi_IN dsb_DE dua_CM dyo_SN dz_BT
			ebu_KE ee_GH el_GR en_Dsrt_US en_Shaw_GB en_US eo_001 es_ES et_EE eu_ES ewo_CM
			fa_IR ff_Adlm_GN ff_Latn ff_Latn_SN fi_FI fil_PH fo_FO fr_FR fur_IT fy_NL
			ga_IE gd_GB gl_ES gsw_CH gu_IN guz_KE gv_IM
			ha_Arab_NG ha_NG haw_US he_IL hi_IN hi_Latn_IN hr_HR hsb_DE hu_HU hy_AM
			ia_001 id_ID ie_EE ig_NG ii_CN is_IS it_IT
			ja_JP jgo_CM jmc_TZ jv_ID
			ka_GE kab_DZ kam_KE kde_TZ kea_CV kgp_BR khq_ML ki_KE
			kk_KZ kkj_CM kl_GL kln_KE km_KH kn_IN ko_KR kok_IN ks_Arab ks_Arab_IN
			ks_Deva_IN ksb_TZ ksf_CM ksh_DE ku_TR kw_GB kxv_Deva_IN kxv_Latn kxv_Latn_IN
			kxv_Orya_IN kxv_Telu_IN ky_KG
			lag_TZ lb_LU lg_UG lij_IT lkt_US lmo_IT ln_CD lo_LA lrc_IR lt_LT lu_CD
			luo_KE luy_KE lv_LV
			mai_IN mas_KE mer_KE mfe_MU mg_MG mgh_MZ mgo_CM mi_NZ mk_MK ml_IN
			mn_MN mn_Mong_CN mni_Beng mni_Beng_IN mni_Mtei_IN mr_IN ms_Arab_MY ms_MY
			mt_MT mua_CM my_MM mzn_IR
			naq_NA nb nb_NO nd_ZW nds_DE ne_NP nl_NL nmg_CM nn_NO nnh_CM nqo_GN
			nus_SS nyn_UG
			oc_FR om_ET or_IN os_GE
			pa_Arab_PK pa_Guru pa_Guru_IN pcm_NG pl_PL prg_PL ps_AF pt_BR
			qu_PE
			raj_IN rm_CH rn_BI ro_RO rof_TZ ru_RU rw_RW rwk_TZ
			sa_IN sah_RU saq_KE sat_Deva_IN sat_Olck sat_Olck_IN sbp_TZ sc_IT
			sd_Arab sd_Arab_PK sd_Deva_IN se_NO seh_MZ ses_ML sg_CF shi_Latn_MA
			shi_Tfng shi_Tfng_MA si_LK sk_SK sl_SI
			smn_FI sn_ZW so_SO sq_AL sr_Cyrl sr_Cyrl_RS sr_Latn_RS
			su_Latn su_Latn_ID sv_SE sw_TZ syr_IQ szl_PL
			ta_IN te_IN teo_UG tg_TJ th_TH ti_ET tk_TM to_TO tok_001
			tr_TR tt_RU twq_NE tzm_MA
			ug_CN uk_UA ur_PK uz_Arab_AF uz_Cyrl_UZ uz_Latn uz_Latn_UZ
			vai_Latn_LR vai_Vaii vai_Vaii_LR vec_IT vi_VN vmw_MZ vun_TZ
			wae_CH wo_SN
			xh_ZA xnr_IN xog_UG
			yav_CM yi_UA yo_NG yrl_BR yue_Hans_CN yue_Hant yue_Hant_HK
			za_CN zgh_MA zh_Hans zh_Hans_CN zh_Hant_TW zu_ZA`
  .split(/\s/)
  .filter((v) => v.trim() !== "")
  .map((v) => v.replace("_", "-"))
  .filter((v) => {
    try {
      new Intl.DateTimeFormat(v);
      return true;
    } catch (e) {
      return false;
    }
  })
  .sort();

const date = new Date("2024-01-02 03:04:05");

const tests = locales.reduce((r, locale) => {
  const result = [];

  [undefined, "numeric", "2-digit"].forEach((year) => {
    [undefined, "numeric", "2-digit"].forEach((month) => {
      if (year === undefined && month === undefined) {
        return;
      }

      const options = { year, month };

      result.push([
        options,
        new Intl.DateTimeFormat(locale, options).format(date),
      ]);

      r[locale] = result;
    });
  });

  ["numeric", "2-digit"].forEach((day) => {
    const options = { day };

    result.push([
      options,
      new Intl.DateTimeFormat(locale, options).format(date),
    ]);

    r[locale] = result;
  });

  return r;
}, {});

const data = JSON.stringify({ date, tests });

fs.writeFile("tests.json", data, (err) => {
  if (err) {
    console.log(err);
  }
});
