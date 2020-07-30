db.createUser(
    {
        user: "test",
        pwd: "12345678",
        roles: [
            {
                role: "readWrite",
                db: "vietlott"
            }
        ]
    }
);