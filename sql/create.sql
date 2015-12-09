
CREATE TABLE LocValMap (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  loc varchar(256) not null,
  lat double not null,
  lng double not null,
  extrakey varchar(129),
  url varchar(256) not null,
  createdat date
);


