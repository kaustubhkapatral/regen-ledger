CREATE TABLE geo (
  hash bytea NOT NULL PRIMARY KEY,
  tx bytea NOT NULL REFERENCES tx,
  -- Both the Postgis geography and geometry representations are stored
  geog geography NOT NULL,
  geom geometry NOT NULL
);

CREATE INDEX geo_geom_gist ON geo USING GIST ( geom );
