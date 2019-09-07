const generateMockData = require("../generateMocks");
const assert = require("./util.js");

// Based off - https://docs.mongodb.com/manual/core/schema-validation/ (JSON Schema)
// TESTS Start here ...
assert(
  [
    generateMockData(),
    generateMockData({ schemawithNoProperties: true }),
    generateMockData(1),
    generateMockData(true),
    generateMockData('lalalala'),
  ],
  "If argument given doesn't contain `properties` key then fail",
  {
    check: "Please use correct schema"
  }
);

assert(
  [
    generateMockData({
      properties: {
        testField: {
          bsonType: "string"
        },
        secondTest: {
          bsonType: "string"
        }
      }
    }),
  ],
  "If argument given is correct then JSON Object should be returned with mock-data",
  {
    funcCheck: mockList => {
      let testPassed = false;
      const result = Object.keys(mockList);
      const expected = ["testField", "secondTest"];
      expected.map(key => {
        testPassed = result.includes(key);
      })

      return testPassed;
    }
  }
);

const typesToHandle = ["string", "int", "bool", "double"].reduce((accu, type, index) => {
  return Object.assign(accu, {
    [`field-${index}`]: {
      bsonType: type
    }
  });
}, {});

assert(
  [
    generateMockData({
      properties: {
        ...typesToHandle
      }
    }),
  ],
  "If bsonType is string/32bit int/boolean/floating point (double), it should be given correct data type",
  {
    funcCheck: mockList => {
      const toCheck = ["field-0", "string", "field-1", "number", "field-2", "boolean", "field-3", "number"];
      let testPassed = false;

      for (let i = 0; i < toCheck.length - 1; i += 2) {
          testPassed = typeof mockList[toCheck[i]] === toCheck[i + 1];

          if (testPassed === false) {
            return false;
          }
      }

      return testPassed;
    }
  }
);

assert(
  [
    generateMockData({
      properties: {
        enumField: {
          enum: ["no", "bson", "type", "here"]
        }
      }
    }),
  ],
  "If type is enum then random value from defined values should be added",
  {
    funcCheck: mockList => ["no", "bson", "type", "here"].includes(mockList.enumField)
  }
);

// TODO: Replace all data with mocks
// TODO: until decide how to deal with this
assert(
  [
    generateMockData({
      properties: {
        newField: {
          bsonType: "string"
        },
        objectField: {
          bsonType: "ObjectId"
        },
        arrField: {
          bsonType: "array"
        },
        objField: {
          bsonType: "object"
        }
      }
    }),
  ],
  "If type isn't enum or bsonType it should be skipped",
  {
      funcCheck: mockList => {
        return Object.keys(mockList).length === 1
      }
  }
);

assert(
  [
    generateMockData({
      properties: {
        newField: {
          bsonType: "string"
        },
      }
    }, "10"),
  ],
  "If multiplier is specified, then schema data should be increased according to integer given",
  {
      funcCheck: mockList => {
        return mockList.length === 10
      }
  }
);
