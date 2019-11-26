mongo admin --username $MONGO_INITDB_ROOT_USERNAME --password $MONGO_INITDB_ROOT_PASSWORD --eval "var site_user='$DB_USER'; var site_pass='$DB_PASS';
 db.getSiblingDB('website').createUser({
     user: site_user,
     pwd: site_pass,
     roles: [
         { role: 'readWrite', db: 'website' },
     ],
});
 db.createCollection('pages', {
    validator: {
        \$jsonSchema: {
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
"
mongoimport -d website -u website_admin -p website_pass -c pages --type json --jsonArray import.dev.json
