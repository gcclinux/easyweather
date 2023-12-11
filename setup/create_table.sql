-- Table: public.easyweather

-- DROP TABLE IF EXISTS public.easyweather;

CREATE TABLE IF NOT EXISTS public.easyweather
(
    id integer NOT NULL DEFAULT nextval('easyweather_id_seq'::regclass),
    obstimeutc timestamp without time zone,
    obstimelocal timestamp without time zone,
    neighborhood character varying(50) COLLATE pg_catalog."default",
    country character(2) COLLATE pg_catalog."default",
    solarradiation double precision,
    lon double precision,
    realtimefrequency integer,
    epoch bigint,
    lat double precision,
    uv double precision,
    winddir integer,
    humidity integer,
    qcstatus integer,
    temp double precision,
    heatindex double precision,
    dewpt double precision,
    windchill double precision,
    windspeed double precision,
    windgust double precision,
    pressure double precision,
    preciprate double precision,
    preciptotal double precision,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    freetext character varying COLLATE pg_catalog."default",
    CONSTRAINT easyweather_pkey_1 PRIMARY KEY (id)
)

TABLESPACE pg_default;