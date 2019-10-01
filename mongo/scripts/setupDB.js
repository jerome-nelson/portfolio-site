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

//  db.pages.insertMany([
//     {
//         'title': 'home',
//         'type': 'main',
//         'status': 'enabled'
//     },
//     {
//         'title': 'about',
//         'slug': 'about-me',
//         'type': 'main',
//         'status': 'enabled'
//     },
//     {
//         'title': 'portfolio',
//         'slug': 'previous-work',
//         'type': 'main',
//         'status': 'enabled'
//     },
//     {
//         'title': 'contact',
//         'slug': 'contact-me',
//         'type': 'main',
//         'status': 'enabled'
//     },
//     {
//         'title': 'progen orthopadiecs',
//         'slug': 'progen-ortho',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'caciques',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'andrew k hairdresser',
//         'slug': 'andrew-k',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'the anime empire',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'hd virtual art',
//         'slug': 'hdva',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'visioncare opticians',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'vertigo42',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'trafalgar place',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'swishboxes',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'shampangroup',
//         'type': 'portfolio',
//         'slug': 'the-shampan',
//         'status': 'disabled'
//     },
//     {
//         'title': 'punjab58',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'orchard orthodontics',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'cnetso education',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'oceana drycleaners',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'n16/n4 drycleaners',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'mercado cantina',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'trafalgar place',
//         'type': 'portfolio',
//         'status': 'disabled'
//     },
//     {
//         'title': 'htaccessible',
//         'type': 'code',
//         'status': 'disabled'
//     }
//  ])
  db.createCollection('summaries', {
    validator: {
        $jsonSchema: {
           bsonType: 'object',
           required: ['updated', 'text'],
           properties: {
               content: {
                   bsonType: 'object',
                   properties: {
                       text: {
                           bsonType: 'array',
                           description: 'array of markdown. required'
                       },
                       tags: {
                           bsonType: 'array',
                           description: 'array of objects'
                       },
                       updated: {
                           bsonType: 'array',
                           description: 'array of timestamps. required'
                       }
                   }
               }
           }
        }
    }
})

