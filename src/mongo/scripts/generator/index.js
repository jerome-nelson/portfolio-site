// #!/usr/bin/env node
const fs = require("fs");
const generateMockData = require("./generateMocks");

require('yargs')
    .command(
        'generate [file]',
        'Generate mock json data for MongoDB import from JSON Schema Validator',
        yargs => {
            yargs.positional('file', {
                describe: 'filename or file glob to specify',
                default: '',
            });
        },
        argv => {
            init(argv.file, {
                multiplier: argv.multiplier,
                output: argv.output,
            });
        }
    )
    .option('multiplier', {
        alias: 'm',
        default: 1,
    })
    .option('output', {
        alias: 'o',
        default: '',
    }).argv;


function init(file, options) {
    if (!file || !options || options && !options.output) {
        return 'Please set attributes correctly'
    }

    // Sync since if this doesn't return then no point
    try {
        let input = fs.readFileSync (file, {encoding: 'utf8', flag: 'r'});
        input = JSON.parse(input);
        const data = generateMockData(input, options.multiplier);

        if (options.dryrun) {
            return data;
        }
        _saveFile(data, options.output);
    } catch (e) {
        if (e.errno === -4058) {
            return 'File not found'
        }
        return e;
    }
}

function _prepareForMongo(data) {
    return JSON.stringify(data);
    // return stringObj.replace(/\[/, "").replace(/\]/, "").replace(/\}\,/g, "}")
}

function _saveFile (data, path) {
    const mongoStr = _prepareForMongo(data);
    buffer = Buffer.from (mongoStr);

    fs.open (path, 'w', function (err, fd) {
      if (err) {
        throw 'error opening file: ' + err;
      }

      fs.write (fd, buffer, 0, buffer.length, null, function (err) {
        if (err) throw 'error writing file: ' + err;
        fs.close (fd, function () {
          return;
        });
      });
    });
  }

module.exports = init;