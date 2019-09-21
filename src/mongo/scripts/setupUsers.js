db.createUser(
    {
        user: username,
        pwd: password,
        roles: [{ role: "userAdminAnyDatabase", db: "admin" }],
    }
)

// db.getSiblingDB("website").createUser({
//     user: site_user,
//     pwd: site_pass,
//     roles: [
//         { role: "readWrite", db: "website" },
//     ],
// })