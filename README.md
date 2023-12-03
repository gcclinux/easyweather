# weatherAPI

**New Weather collector that get weather data from weather.com and openweathermap.org**

This weatherCollect application will collect the weather data from either or both above data source and store it in a PostgreSQL database.

Requires: APIKEY from either or both Weather Sites

**If you don't have your own weather station.**<BR>
*https://openweathermap.org/api*

**If you do have your own weather station that can upload data.**<BR>
*https://www.wunderground.com/member/api-keys*

**Config example if you upload data to weather.com then StationValid = true**
```
{
    "DB_HOST": ["host_name"],
    "DB_PORT": ["5432"],
    "DB_USER": ["user_name"],
    "DB_PASS": ["Passw0rd"],
    "DB_NAME": ["weather"],
    "TB_NAME": ["weather_table"],
    "OpenWeatherApi": ["12345678909876543212345678909876"],
    "StationValid": ["true"],
    "WundergroupApi": ["12345678909876543212345678909876"],
	"StationId": ["Station_Name"],
    "AdminPort": ["8081"]
}
```
**PostgreSQL database structure example and syntax**

```
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
    CONSTRAINT easyweather_pkey_1 PRIMARY KEY (id)
)
```