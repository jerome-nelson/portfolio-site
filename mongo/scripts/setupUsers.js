db.createUser(
    {
        user: username,
        pwd: password,
        roles: [{ role: "userAdminAnyDatabase", db: "admin" }],
    }
)