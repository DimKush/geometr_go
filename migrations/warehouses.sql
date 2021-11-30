create table warehouses
(
    id                   bigserial
        constraint points_pkey
            primary key,
    created_at           timestamp,
    updated_at           timestamp,
    deleted_at           timestamp,
    name                 varchar,
    description          varchar,
    latitude             double precision,
    longitude            double precision,
    geom                 geometry(Point, 4326),
    poly                 geometry(Polygon, 4326),
    multi_poly           geometry(MultiPolygon, 4326)
);