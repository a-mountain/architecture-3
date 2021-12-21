CREATE TABLE MENU_ITEMS
(
    ID    SERIAL PRIMARY KEY,
    NAME  VARCHAR(50) NOT NULL,
    PRICE REAL        NOT NULL
);

CREATE TABLE ORDERS
(
    ID           SERIAL PRIMARY KEY,
    TABLE_NUMBER SMALLINT NOT NULL,
    MENU_ITEM_ID BIGINT   NOT NULL
);

ALTER TABLE ORDERS
    ADD CONSTRAINT MENU_ITEMS_FK FOREIGN KEY (MENU_ITEM_ID) REFERENCES MENU_ITEMS (ID);

CREATE INDEX TABLE_NUMBER_INDEX ON ORDERS (TABLE_NUMBER);

INSERT INTO MENU_ITEMS(NAME, PRICE)
VALUES ('Beef carpaccio', 790),
       ('Burrata', 950),
       ('salmon', 870),
       ('Herring with new potatoes and green onions', 2100),
       ('Antipasto', 550),
       ('Oyster', 390),
       ('Spicy beef tartare with truffle oil', 790);