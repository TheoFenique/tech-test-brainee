create table brainees
(
    id     int auto_increment
        primary key,
    author varchar(127)  null,
    text   varchar(1000) null,
    brand  varchar(31)   null
);

INSERT INTO techTest.brainees (id, author, text, brand) VALUES (1, 'Theo', 'Je veux du coca a la truffe', 'Coca');
INSERT INTO techTest.brainees (id, author, text, brand) VALUES (2, 'Robin', 'Nutella sans huile de palme', 'Nutella');
INSERT INTO techTest.brainees (id, author, text, brand) VALUES (3, 'Alexandre', 'Du café au chocolat', 'Nespresso');
INSERT INTO techTest.brainees (id, author, text, brand) VALUES (10, 'Theo Fenique', 'Le sens de la vie', 'Le guide du voyageurgGalactique');