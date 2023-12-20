 SELECT ecowitt_weather.obstimelocal,
    ecowitt_weather.neighborhood,
	ecowitt_weather.country,
	ecowitt_weather.lon,
	ecowitt_weather.lat,
    ecowitt_weather.winddir,
    ecowitt_weather.humidity,
    ecowitt_weather.temp,
	ecowitt_weather.dewpt,
    ecowitt_weather.windchill,
    ecowitt_weather.windspeed,
    ecowitt_weather.windgust,	
    ecowitt_weather.pressure,
	ecowitt_weather.freetext
   FROM ecowitt_weather
  WHERE ecowitt_weather.obstimelocal >= CURRENT_DATE AND ecowitt_weather.obstimelocal < (CURRENT_DATE + '1 day'::interval)
  ORDER BY ecowitt_weather.id DESC;