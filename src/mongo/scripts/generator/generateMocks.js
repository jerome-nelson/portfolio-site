/*
 * Function accepts  MongoDB JSON Schema
 * and then will create mockdata based on what it gets
 */
function generateMock(schema, multiplier = 1) {
    if (!schema || !schema.properties) {
        return 'Please use correct schema';
    }
    let result = [];
    for (let i = multiplier; i >= 1; i -=1) {
        result.push(_createJSON(schema.properties));
    }

    return multiplier > 1 ? result : _createJSON(schema.properties);
}

function _createJSON(fieldList) {
    const list = Object.keys(fieldList).reduce((accu, elem) => {
        let value = _generateMockValue(fieldList[elem]);
        if (!value) {
            return accu;
        }

        return {
            ...accu,
            [elem]: _generateMockValue(fieldList[elem])
        }
    }, {});
    return list;
}

function _generateMockValue(value) {
        const key = Object.keys(value)
            .filter(elem =>
                elem.toLowerCase() === 'bsontype' ||
                elem.toLowerCase() === 'enum'
            )[0];

        const type = value[key];

        if (key === "bsonType") {
            switch (type) {
                case "string":
                    return "string-data";
                case "double":
                    return Math.random()
                case "int":
                    return Math.floor(Math.random() * Math.pow(2, 32) - 1)
                case "bool":
                    return true
            }
        }

        if (key === "enum") {
            const randKey = Math.floor(Math.random() * type.length);
            return type[randKey]
        }

        return;
}

module.exports = generateMock