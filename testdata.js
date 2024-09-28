const fs = require("fs");

const locales_ = `
			aa_ET ab_GE af_ZA agq_CM ak_GH am_ET an_ES ann_NG apc_SY ar_001 arn_CL as_IN
			asa_TZ ast_ES az_Arab_IR az_Cyrl_AZ az_Latn az_Latn_AZ
			ba_RU bal_Arab bal_Arab_PK bal_Latn_PK bas_CM be_BY bem_ZM bew_ID bez_TZ bg_BG
			bgc_IN bgn_PK bho_IN blo_BJ blt_VN bm_ML bm_Nkoo_ML bn_BD bo_CN br_FR brx_IN
			bs_Cyrl_BA bs_Latn bs_Latn_BA bss_CM byn_ER
			ca_ES cad_US cch_NG ccp_BD ce_RU ceb_PH cgg_UG cho_US chr_US cic_US ckb_IQ
			co_FR cs_CZ csw_CA cu_RU cv_RU cy_GB
			da_DK dav_KE de_DE dje_NE doi_IN dsb_DE dua_CM dv_MV dyo_SN dz_BT
			ebu_KE ee_GH el_GR en_Dsrt_US en_Shaw_GB en_US eo_001 es_ES et_EE eu_ES ewo_CM
			fa_IR ff_Adlm_GN ff_Latn ff_Latn_SN fi_FI fil_PH fo_FO fr_FR frr_DE fur_IT
			fy_NL
			ga_IE gaa_GH gd_GB gez_ET gl_ES gn_PY gsw_CH gu_IN guz_KE gv_IM
			ha_Arab_NG ha_NG haw_US he_IL hi_IN hi_Latn_IN hnj_Hmnp hnj_Hmnp_US hr_HR
			hsb_DE hu_HU hy_AM
			ia_001 id_ID ie_EE ife_TG ig_NG ii_CN io_001 is_IS it_IT iu_CA iu_Latn_CA
			ja_JP jbo_001 jgo_CM jmc_TZ jv_ID
			ka_GE kab_DZ kaj_NG kam_KE kcg_NG kde_TZ kea_CV ken_CM kgp_BR khq_ML ki_KE
			kk_KZ kkj_CM kl_GL kln_KE km_KH kn_IN ko_KR kok_IN kpe_LR ks_Arab ks_Arab_IN
			ks_Deva_IN ksb_TZ ksf_CM ksh_DE ku_TR kw_GB kxv_Deva_IN kxv_Latn kxv_Latn_IN
			kxv_Orya_IN kxv_Telu_IN ky_KG
			la_VA lag_TZ lb_LU lg_UG lij_IT lkt_US lmo_IT ln_CD lo_LA lrc_IR lt_LT lu_CD
			luo_KE luy_KE lv_LV
			mai_IN mas_KE mdf_RU mer_KE mfe_MU mg_MG mgh_MZ mgo_CM mi_NZ mic_CA mk_MK ml_IN
			mn_MN mn_Mong_CN mni_Beng mni_Beng_IN mni_Mtei_IN moh_CA mr_IN ms_Arab_MY ms_MY
			mt_MT mua_CM mus_US my_MM myv_RU mzn_IR
			naq_NA nb nb_NO nd_ZW nds_DE ne_NP nl_NL nmg_CM nn_NO nnh_CM nqo_GN nr_ZA
			nso_ZA nus_SS nv_US ny_MW nyn_UG
			oc_FR om_ET or_IN os_GE osa_US
			pa_Arab_PK pa_Guru pa_Guru_IN pap_CW pcm_NG pis_SB pl_PL prg_PL ps_AF pt_BR
			qu_PE quc_GT
			raj_IN rhg_Rohg rhg_Rohg_MM rif_MA rm_CH rn_BI ro_RO rof_TZ ru_RU rw_RW rwk_TZ
			sa_IN sah_RU saq_KE sat_Deva_IN sat_Olck sat_Olck_IN sbp_TZ sc_IT scn_IT
			sd_Arab sd_Arab_PK sd_Deva_IN sdh_IR se_NO seh_MZ ses_ML sg_CF shi_Latn_MA
			shi_Tfng shi_Tfng_MA shn_MM si_LK sid_ET sk_SK skr_PK sl_SI sma_SE smj_SE
			smn_FI sms_FI sn_ZW so_SO sq_AL sr_Cyrl sr_Cyrl_RS sr_Latn_RS ss_ZA ssy_ER
			st_ZA su_Latn su_Latn_ID sv_SE sw_TZ syr_IQ szl_PL
			ta_IN te_IN teo_UG tg_TJ th_TH ti_ET tig_ER tk_TM tn_ZA to_TO tok_001 tpi_PG
			tr_TR trv_TW trw_PK ts_ZA tt_RU twq_NE tyv_RU tzm_MA
			ug_CN uk_UA ur_PK uz_Arab_AF uz_Cyrl_UZ uz_Latn uz_Latn_UZ
			vai_Latn_LR vai_Vaii vai_Vaii_LR ve_ZA vec_IT vi_VN vmw_MZ vo_001 vun_TZ
			wa_BE wae_CH wal_ET wbp_AU wo_SN
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

const tests = locales_.reduce((r, locale) => {
  const result = [];

  ["numeric", "2-digit"].forEach((year) => {
    const options = { year };

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
