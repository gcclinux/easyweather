# EASYWEATHER

This application / code is currently still working progress several know facts still to be corrected and completed.


**New Weather collector that get weather data from weather.com and openweathermap.org**

This easyweather application will collect the weather data from either or both above data source and store it in a PostgreSQL database.

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
solarRadiation		solarRadiation(double)		float64		nil				8.6	
lon					lon(double)					float64		lon				-3.456
realtimeFrequency	realtimeFrequency(double)	int			nil				0
epoch				epoch(double)				int			dt				1701684204
lat					lat(double)					float64		lat				55.932
uv					uv(double)					float64		nil				0	
winddir				winddir(integer)			float64		nil				118	
humidity			humidity(integer)			float64		humidity		93
qcStatus			qcStatus(integer)			float64		nil				1	
temp				temp(double)				float64		temp			3.6
heatIndex			heatIndex(double)			float64		feels_like		3.6
dewpt				dewpt(double)				float64		Convert			2.5
windChill			windChill(double)			float64		nil				3.7
windSpeed			windSpeed(double)			float64		speed			1.1
windGust			windGust(double)			float64		gust			1.8
pressure			pressure(double)			float64		pressure		1020.79
precipRate			precipRate(double)			float64		nil				0	
precipTotal			precipTotal(double)			float64		nil				0.71	
nil					freetext(character)			string		description		broken clouds
```
**Config example if you upload data to weather.com**
```
{
    "DB_HOST": ["host_name"],									(Database hostname or ip address)
    "DB_PORT": ["5432"],										(PostgreSQL default port)
    "DB_USER": ["user_name"],									(PostgreSQL database user)
    "DB_PASS": ["Passw0rd"],									(PostgreSQL user password)
    "DB_NAME": ["weather"],										(PostgreSQL database name)
    "TB_NAME": ["easyweather"],									(PostgreSQL database table name)
    "OpenWeatherApi": ["12345678909876543212345678909876"],		(OpenWeatherApi key)
    "WundergroundApi": [""],									(WundergroundApi key)
	"StationId": [""],											(Wunderground StationId)
    "WebPort": ["8081"],										(Web port for Admin & Chart)
    "Language": ["en"],											(OpenWeatherApi Description language)
    "DefaultCity": ["Edinburgh"],								(Default city if not varble passed)
	"Interval": [10],											(Freequency data will be collected in minutes)
	"Retry": [10],												(Freequency of data capture if failed)
	"PrivKeyPATH": [""],										(String with Full PATH to your /path/privkey.pem)
    "CertPemPATH": [""]											(String with Full PATH to you /path/cert.pem)
}
```
**PostgreSQL database structure example and syntax**<BR>
*There is a script under sertup flder called setup/create_table.sql*

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

When everything is created and configured like your PostgreSQL database and your config and API accounts! The first thing yu should do is run an integrity check to make sure everything is working and configured using the --integrity flag<BR>

The compile script creates easyweather.exe but you can compile the program manually if you want:<BR>
*go build -o easyweather.exe *.go*
*go build -o easyweather *.go*
*go build -o easyweather-$(uname)-$(uname -m) *.go*

**Function to run through the setup and check the integrity of the application setup**<BR>
*./compile.sh && ./easyweather.exe --integrity*<BR>

**Quick example how to compile and lanch the application for data collection**<BR>
*./compile.sh && ./easyweather.exe --collect*<BR>

**Quick example how to compile and lanch the application to launch the web interface!**<BR>
*./compile.sh && ./easyweather.exe --web*<BR>

To run the program you will need to set it in the background to collect and to serve the web!<BR><BR>
Linux Example:<BR>
*cd /server/easyweather && /usr/bin/screen -dmS WEATHER-GUI /server/easyweather/easyweather-Linux-aarch64 --web*<BR>
*cd /server/easyweather && /usr/bin/screen -dmS WEATHER-COL /server/easyweather/easyweather-Linux-aarch64 --collect*<BR><BR>
Windows Example (CMD):<BR>
*C:\Users\ricar\easyweather> start /b easyweather.exe --web*<BR>
*C:\Users\ricar\easyweather> start /b easyweather.exe --collect*<BR><BR>
Windows Example (PS):<BR>
*PS C:\Users\ricar\easyweather> Start-Process -FilePath ".\easyweather.exe" -ArgumentList "--web" -NoNewWindow -PassThru | Out-Null*<BR>
*PS C:\Users\ricar\easyweather> Start-Process -FilePath ".\easyweather.exe" -ArgumentList "--collect" -NoNewWindow -PassThru | Out-Null*<BR>