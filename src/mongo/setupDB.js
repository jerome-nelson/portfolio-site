// Add ip restriction for container only development machine
db.getSiblingDB("website").createUser({
    user: "website_admin",
    pwd: "website_password",
    roles: [
      { role: "readWrite", db: "website" },
    ],
  })
