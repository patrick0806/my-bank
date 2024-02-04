CREATE TABLE IF NOT EXISTS clients(
    id varchar(40) PRIMARY KEY,
    name varchar(50),
    birthdate date,
    cpf varchar(11),
    email varchar(50),
    phone_number varchar(20)
);

CREATE TABLE IF NOT EXISTS accounts(
    id varchar(40) PRIMARY KEY,
    balance decimal(10,2),
    clientId varchar(40),
    FOREIGN KEY (clientId) REFERENCES clients(id)
);