const path = require("path");
const fs = require("fs");
const assert = require('./util.js');
const init = require("../index");

// If file cannot be found
// If output file is not set
// If data is undefined
// If input is not valid JSON

assert(
    [
        init(),
        init("set"),
        init(null, {
            output: "set"
        })
    ],
    "If file or output is not set then return error",
    {
        check: "Please set attributes correctly"
    }
)

assert(
    init("doesntexist.json", {
        output: "not-generated.json"
    }),
    "If file cannot be found then return error",
    {
        check: "File not found"
    }
)

assert(
    init("./__tests__/mocks/schema.json", {
        dryrun: true,
        output: "not-generated.json"
    }),
    "It should generate output instead of saving to file on dry run",
    {
        funcCheck: result => ["title", "type", "status", "slug"].length === Object.keys(result).length
    }
)

// Test to generate output
const output = "./__tests__/mocks/generated.json";
assert(
    init("./__tests__/mocks/schema.json", {
        output,
        multiplier: 10
    }),
    "It should generate output to file",
    {
        funcCheck: () => {
            try {
                fs.readFileSync (path.resolve('./', output), {encoding: 'utf8', flag: 'r'});
                fs.unlinkSync(output);
                return true;
            } catch (e) {
                console.log(e);
                return false;
            }
        }
    }
)
