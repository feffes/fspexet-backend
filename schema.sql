CREATE TABLE news (
    ID      text             NOT NULL,
    Title   text            NOT NULL,
    Author  text            NOT NULL,
    Content    text            NOT NULL,
    Time    timestamptz     NOT NULL,
    PRIMARY KEY (ID)


);
CREATE TABLE users (
    ID      text             NOT NULL,
    Name   text            NOT NULL,
    Password    text        NOT NULL,
    LastLogin    timestamptz     NOT NULL,
    CreationTime timestamptz     NOT NULL,
    PRIMARY KEY (ID)
);

INSERT INTO news (ID, Title, Author, Content, Time)
    VALUES ('first-test-news', 'This is the Title :)', 'feffe', 'Lorem ipsum dolar sit amet or whatever', '2016-06-22 19:10:25-01');

INSERT INTO users (ID, Name, Password, LastLogin, CreationTime)
    VALUES ('feffe', 'Fredrik Åberg', 'losenord', '2018-06-22 19:10:25-01', '2016-02-22 19:10:25-01');
