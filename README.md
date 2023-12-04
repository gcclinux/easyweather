# weatherAPI

**New Weather collector that get weather data from weather.com and openweathermap.org**

This weatherCollect application will collect the weather data from either or both above data source and store it in a PostgreSQL database.

Requires: APIKEY from either or both Weather Sites

**If you don't have your own weather station.**<BR>
*https://openweathermap.org/api*

**If you do have your own weather station that can upload data.**<BR>
*https://www.wunderground.com/member/api-keys*

**Table definiton and differences captured between weather.com and openweathermap.org**

```
Wunderground		SQL							Struct		OpenWeather	    Example
					id(integer)			
obsTimeUtc			obsTimeUtc(timestamp)		time.Time	Convert			2023-12-04T10:03:24Z	
obsTimeLocal		obsTimeLocal(timestamp)		string		Convert			2023-12-04 10:03:24
neighborhood		neighborhood(character)		string		name			Edinburgh
country				country(character)			string		country			GB
solarRadiation		solarRadiation(double)		float64						8.6	
lon					lon(double)					float64		lon				-3.456
realtimeFrequency	realtimeFrequency(double)	int				
epoch				epoch(double)				int			dt				1701684204
lat					lat(double)					float64		lat				55.932
uv					uv(double)					float64						0	
winddir				winddir(integer)			float64						118	
humidity			humidity(integer)			float64		humidity		93
qcStatus			qcStatus(integer)			float64						1	
temp				temp(double)				float64		temp			3.6
heatIndex			heatIndex(double)			float64		temp_max		3.6
dewpt				dewpt(double)				float64		feels_like		2.5
windChill			windChill(double)			float64						3.7
windSpeed			windSpeed(double)			float64		speed			1.1
windGust			windGust(double)			float64		gust			1.8
pressure			pressure(double)			float64		pressure		1020.79
precipRate			precipRate(double)			float64						0	
precipTotal			precipTotal(double)			float64						0.71	
					freetext(character)			string		description		broken clouds
```
**Config example if you upload data to weather.com then StationValid = true**
```
{
    "DB_HOST": ["host_name"],									(Database hostname or ip address)
    "DB_PORT": ["5432"],										(PostgreSQL default port)
    "DB_USER": ["user_name"],									(PostgreSQL database user)
    "DB_PASS": ["Passw0rd"],									(PostgreSQL user password)
    "DB_NAME": ["weather"],										(PostgreSQL database name)
    "TB_NAME": ["easyweather"],									(PostgreSQL database table name)
    "OpenWeatherApi": ["12345678909876543212345678909876"],		(OpenWeatherApi key)
    "StationValid": [false],									(Using your own weather station)
    "WundergroundApi": [""],									(WundergroundApi key)
	"StationId": [""],											(Wunderground StationId)
    "WebPort": ["8081"],										(Web port for Admin & Chart)
    "Language": ["en"],											(OpenWeatherApi Description language)
    "DefaultCity": ["Edinburgh"],								(Default city if not varble passed)
	"EcowittKey": [""],											(Ecowitt Application key)
    "EcowittApi": [""],											(Ecowitt API key)
    "EcowittMac": [""]											(Ecowitt Device MacAddress)
}
```
**PostgreSQL database structure example and syntax**

```
CREATE TABLE public.easyweather (
	id int4 NOT NULL DEFAULT nextval('easyweather_id_seq'::regclass),
	obstimeutc timestamp NULL,
	obstimelocal timestamp NULL,
	neighborhood varchar(50) NULL,
	country bpchar(2) NULL,
	solarradiation float8 NULL,
	lon float8 NULL,
	realtimefrequency int4 NULL,
	epoch int8 NULL,
	lat float8 NULL,
	uv float8 NULL,
	winddir int4 NULL,
	humidity int4 NULL,
	qcstatus int4 NULL,
	"temp" float8 NULL,
	heatindex float8 NULL,
	dewpt float8 NULL,
	windchill float8 NULL,
	windspeed float8 NULL,
	windgust float8 NULL,
	pressure float8 NULL,
	preciprate float8 NULL,
	preciptotal float8 NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	freetext varchar NULL,
	CONSTRAINT easyweather_pkey_1 PRIMARY KEY (id)
);
```