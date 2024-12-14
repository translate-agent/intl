const fs = require("fs").promises;
const util = require("util");
const exec = util.promisify(require("child_process").exec);

async function getLocales() {
  const locales = (await fs.readdir(".cldr/common/main"))
    .map((v) => v.slice(0, -4))
    .filter((v) => !["root"].includes(v))
    .map((v) => v.replaceAll("_", "-"));

  // Skip the tests for the following locales - ICU does not support them (it falls back to system LANG).
  const supported = await Promise.all(
    locales.map(async (locale) => {
      const a = await exec(
        `LANG=en_US node -e "console.log(new Intl.DateTimeFormat('${locale}', {month: 'numeric', day: 'numeric'}).format(new Date()))"`
      );

      if (a.stderr) {
        console.error(a.stderr);
        process.exit(1);
      }

      const b = await exec(
        `LANG=en_GB node -e "console.log(new Intl.DateTimeFormat('${locale}', {month: 'numeric', day: 'numeric'}).format(new Date()))"`
      );

      if (b.stderr) {
        console.error(b.stderr);
        process.exit(1);
      }

      return a.stdout === b.stdout;
    })
  );

  return locales.filter((_, i) => supported[i]);
}

function generateTests(locales) {
  const date = new Date("2024-01-02 03:04:05");

  const tests = locales.reduce((r, locale) => {
    const result = [];

    [undefined, "long", "short", "narrow"].forEach((era) => {
      [undefined, "numeric", "2-digit"].forEach((year) => {
        [undefined, "numeric", "2-digit"].forEach((month) => {
          [undefined, "numeric", "2-digit"].forEach((day) => {
            // TODO(jhorsts): skip default formatting for now. It can be resolved when formatting is fully implemented.
            if (
              era == undefined &&
              year === undefined &&
              month === undefined &&
              day == undefined
            ) {
              return;
            }

            const options = { era, year, month, day };

            result.push([
              options,
              new Intl.DateTimeFormat(locale, options).format(date),
            ]);
          });
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
