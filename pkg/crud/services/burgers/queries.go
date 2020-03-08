package burgers

const removeBurger = "UPDATE burgers SET removed = true where id = $1;"
const saveInBurgers = "INSERT INTO burgers(name, price) VALUES ($1, $2);"
const getFalseBurgers = "SELECT id, name, price FROM burgers WHERE removed = FALSE"
const createDB = `CREATE TABLE  if not exists  burgers (
   id BIGSERIAL PRIMARY KEY,
   name TEXT NOT NULL,
   price INT NOT NULL CHECK ( price > 0 ),
   removed BOOLEAN DEFAULT FALSE
);`