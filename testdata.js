const fs = require("fs").promises;

async function getLocales() {
  return (
    (await fs.readdir(".cldr/common/main"))
      .map((v) => v.slice(0, -4))
      .filter((v) => !["root"].includes(v))
      .map((v) => v.replaceAll("_", "-"))
      // Skip the tests for the following locales - ICU does not support them (it falls back to system LANG).
      .filter(
        (locale) =>
          ![
            "aa",
            "ab",
            "an",
            "ann",
            "apc",
            "arn",
            "ba",
            "bal",
            "bew",
            "bgn",
            "blt",
            "bss",
            "byn",
            "cad",
            "cch",
            "cho",
            "cic",
            "co",
            "cu",
            "dv",
            "frr",
            "gaa",
            "gez",
            "gn",
            "hnj",
            "ife",
            "io",
            "iu",
            "jbo",
            "kaj",
            "kcg",
            "ken",
            "kpe",
            "la",
            "mdf",
            "mic",
            "moh",
            "mus",
            "myv",
            "nr",
            "nso",
            "nv",
            "ny",
            "osa",
            "pap",
            "pis",
            "quc",
            "rhg",
            "rif",
            "scn",
            "sdh",
            "shn",
            "sid",
            "skr",
            "sma",
            "smj",
            "sms",
            "ss",
            "ssy",
            "st",
            "tig",
            "tn",
            "tpi",
            "trv",
            "trw",
            "ts",
            "tyv",
            "ve",
            "vo",
            "wa",
            "wal",
            "wbp",
          ].some((v) => locale === v || locale.startsWith(v + "-"))
      )
      .sort()
  );
}

function generateTests(locales) {
  const date = new Date("2024-01-02 03:04:05");

  const tests = locales.reduce((r, locale) => {
    const result = [];

    [undefined, "numeric", "2-digit"].forEach((year) => {
      [undefined, "numeric", "2-digit"].forEach((month) => {
        [undefined, "numeric", "2-digit"].forEach((day) => {
          // TODO(jhorsts): skip default formatting for now. It can be resolved when formatting is fully implemented.
          if (
            (year === undefined && month === undefined && day == undefined) ||
            (year !== undefined && month === undefined && day != undefined)
          ) {
            return;
          }

          const options = { year, month, day };

          result.push([
            options,
            new Intl.DateTimeFormat(locale, options).format(date),
          ]);
        });
      });
    });

    r[locale] = result;

    return r;
  }, {});

  return { date, tests };
}

// main

(async function () {
  const locales = await getLocales();
  const tests = generateTests(locales);
  const data = JSON.stringify(tests);

  await fs.writeFile("tests.json", data, (err) => {
    if (err) {
      console.log(err);
    }
  });
})();
