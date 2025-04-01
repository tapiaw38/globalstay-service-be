// init-mongo.js

// Read environment variables
const dbName = process.env.DB_NAME || "reservation-db";
const rootUser = process.env.DB_USER || "admin";
const rootPassword = process.env.DB_PASSWORD || "admin";

// Connect to the admin database
const adminDB = db.getSiblingDB("admin");

// Authenticate with the root user
adminDB.auth({
  user: rootUser,
  pwd: rootPassword,
});

// Use dbName variable for the new database
const newDB = db.getSiblingDB(dbName);

// Create a collection or perform other operations if needed
newDB.createCollection("migrations");
newDB.createCollection("business");
newDB.createCollection("services");
newDB.createCollection("reservations");

// Create a user for the new database using root user and password
newDB.createUser({
  user: rootUser,
  pwd: rootPassword,
  roles: [{ role: "readWrite", db: dbName }],
});
