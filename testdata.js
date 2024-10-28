const fs = require("fs").promises;

async function getLocales() {
  const localesAttributeValue = await fs
    .readFile(".cldr/common/supplemental/supplementalMetadata.xml")
    .then((v) => v.toString())
    .then(
      (v) =>
        v
          .matchAll(/<defaultContent locales="([\sa-zA-Z0-9_]*)/gm)
          .toArray()[0][1]
    );

  return (
    localesAttributeValue
      .split(/\s/)
      .filter((v) => v.trim() !== "")
      .map((v) => v.replace("_", "-"))
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
      .filter((v) => {
        try {
          new Intl.DateTimeFormat(v);
          return true;
        } catch (e) {
          return false;
        }
      })
      .sort()
  );
}

function generateTests(locales) {
  const date = new Date("2024-01-02 03:04:05");

  const tests = locales.reduce((r, locale) => {
    const result = [];

    [undefined, "numeric", "2-digit"].forEach((year) => {
      [undefined, "numeric", "2-digit"].forEach((month) => {
        // TODO(jhorsts): skip default formatting for now. It can be resolved when formatting is fully implemented.
        if (year === undefined && month === undefined) {
          return;
        }

        const options = { year, month };

        result.push([
          options,
          new Intl.DateTimeFormat(locale, options).format(date),
        ]);
      });
    });

    ["numeric", "2-digit"].forEach((day) => {
      const options = { day };

      result.push([
        options,
        new Intl.DateTimeFormat(locale, options).format(date),
      ]);
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
