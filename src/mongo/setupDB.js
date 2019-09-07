
// Add ip restriction for container only development machine
db.getSiblingDB("website").createUser({
    user: process.env.DB_USER,
    pwd: process.env.DB_PASS,
    roles: [
      { role: "readWrite", db: "website" },
    ],
  })
