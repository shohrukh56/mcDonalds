package burgers

const removeBurger = "UPDATE burgers SET removed = true where id = $1;"
const saveInBurgers = "INSERT INTO burgers(name, price) VALUES ($1, $2);"
const getFalseBurgers = "SELECT id, name, price FROM burgers WHERE removed = FALSE"