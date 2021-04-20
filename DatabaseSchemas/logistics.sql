DROP DATABASE IF EXISTS logistics;
CREATE DATABASE logistics;
USE logistics;

CREATE TABLE Customer(
    id_user       INTEGER NOT NULL,
    gender        VARCHAR(50),
    city          VARCHAR(100) NOT NULL,
    country       VARCHAR(70) NOT NULL,
    amount_orders SMALLINT,
    PRIMARY KEY (id_user)
)ENGINE=INNODB;

CREATE TABLE Container(
    id_container   INTEGER NOT NULL,
    container_name VARCHAR(150) NOT NULL,
    amount_bills   SMALLINT,
    to_country     VARCHAR(70) NOT NULL,
    create_date    DATE NOT NULL,
    close_date     DATE,
    status         BOOLEAN,
    curr_value     DECIMAL(6,2) NOT NULL,
    curr_weight    SMALLINT NOT NULL,
    max_weight     SMALLINT, 
    curr_volume    SMALLINT NOT NULL,
    max_volume     SMALLINT,
    PRIMARY KEY(id_container)
)ENGINE=INNODB;

CREATE TABLE Bill(
    id_bill      INTEGER NOT NULL,
    id_container INTEGER NOT NULL,
    id_user      INTEGER NOT NULL,
    bill_date    DATE,
    bill_city    VARCHAR(100),
    bill_country VARCHAR(70),
    total_value  DECIMAL(6,2) NOT NULL,
    total_weight DECIMAL(4,2) NOT NULL,
    total_volume DECIMAL(4,2) NOT NULL,
    PRIMARY KEY (id_bill),
    FOREIGN KEY (id_container) REFERENCES Container(id_container)
    ON UPDATE CASCADE,
    FOREIGN KEY (id_user) REFERENCES Customer(id_user)
    ON UPDATE CASCADE
)ENGINE=INNODB;

CREATE TABLE Product(
    id_product    INTEGER NOT NULL,
    id_bill       INTEGER NOT NULL,
    product_name  VARCHAR(255),
    product_value DECIMAL(6,2) NOT NULL,
    amount        SMALLINT NOT NULL,
    color         VARCHAR(40),
    brand         VARCHAR(100),
    category      VARCHAR(100) NOT NULL,
    weight        DECIMAL(4,2) NOT NULL,
    volume        DECIMAL(4,2) NOT NULL,
    PRIMARY KEY(id_product),
    FOREIGN KEY(id_bill) REFERENCES Bill(id_bill)
    ON UPDATE CASCADE
)ENGINE=INNODB;