CREATE TABLE IF NOT EXISTS clients(
    id varchar(40) PRIMARY KEY NOT NULL,
    name varchar(50) NOT NULL,
    birthdate date NOT NULL,
    cpf varchar(11) NOT NULL,
    email varchar(50) NOT NULL,
    phone_number varchar(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS accounts(
    id varchar(40) PRIMARY KEY NOT NULL,
    balance decimal(10,2) NOT NULL,
    clientId varchar(40) NOT NULL,
    FOREIGN KEY (clientId) REFERENCES clients(id)
);