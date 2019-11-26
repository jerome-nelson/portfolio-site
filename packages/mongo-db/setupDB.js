  db.createCollection('pages', {
    validator: {
        $jsonSchema: {
            bsonType: 'object',
            required: ['type', 'title', 'status', 'slug'],
            properties: {
                title: {
                    bsonType: 'string',
                    description: 'is a string and is required'
                },
                type: {
                    description: 'can only be one of these values and is required',
                    enum: ['main', 'portfolio', 'code']
                },
                status: {
                    description: 'can only be one of these values and is required',
                    enum: ['active', 'disabled']
                },
                content: {
                    bsonType: 'ObjectId'
                },
                text: {
                    bsonType: 'array',
                    description: 'array of markdown. required'
                },
                tags: {
                    bsonType: 'array',
                    description: 'array of objects'
                },
                slug: {
                    bsonType: 'string',
                },
                updated: {
                    bsonType: 'array',
                    description: 'array of timestamps.'
                }
            },
        }
    }
 });


