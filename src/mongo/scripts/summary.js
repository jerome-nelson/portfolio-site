 db.createCollection("summaries", {
     validator: {
         $jsonSchema: {
            bsonType: "object",
            required: ["updated", "text"],
            properties: {
                content: {
                    bsonType: "object",
                    properties: {
                        text: {
                            bsonType: "array",
                            description: "array of markdown. required"
                        },
                        tags: {
                            bsonType: "array",
                            description: "array of objects"
                        },
                        updated: {
                            bsonType: "array",
                            description: "array of timestamps. required"
                        }
                    }
                }
            }
         }
     }
 })