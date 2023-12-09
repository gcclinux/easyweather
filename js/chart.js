var url = "/json";
var labels = [];
var tempData = [];
var dewpt = [];
var windData = [];
var windchill = [];

fetch(url)
    .then(response => response.json())
    .then(data => {
        for (var i = 0; i < data.length; i++) {
            labels.push(data[i].Obstimelocal);
            tempData.push(data[i].Temp);
            dewpt.push(data[i].Dewpt);
            windData.push(data[i].Windspeed);
            windchill.push(data[i].Windchill);
        }

        var ctx = document.getElementById('mainChart').getContext('2d');
        var mainChart = new Chart(ctx, {
            type: 'line',
            data: {
                labels: labels,
                datasets: [{
                    label: 'Temperature',
                    borderWidth: 1,
                    data: tempData,
                    borderColor: 'rgba(255, 99, 132, 1)',
                    backgroundColor: 'rgba(0, 204, 0, 0.2)',
                    fill: true,
                    tension: 0.1,
                    pointRadius: 0
                }, {
                    label: 'Dew Point',
                    borderWidth: 1,
                    data: dewpt,
                    borderColor: 'rgba(54, 162, 235, 1)',
                    backgroundColor: 'rgba(54, 162, 235, 0.2)',
                    fill: true,
                    tension: 0.1,
                    pointRadius: 0
                    
                }]
            },
            options: {
                scales: {
                    yAxes: [{
                        ticks: {
                            beginAtZero: false
                        }
                    }]
                },
                elements: {
                    line: {
                        borderWidth: 1
                    }
                },
                plugins: {
                    zoom: {
                      zoom: {
                        wheel: {
                          enabled: true,
                        },
                        pinch: {
                          enabled: true
                        },
                        mode: 'xy',
                      }
                    }
                  }
            }
        });

        var ctx2 = document.getElementById('myChart2').getContext('2d');
        var myChart2 = new Chart(ctx2, {
            type: 'line',
            data: {
                labels: labels,
                datasets: [{
                    label: 'Windchill',
                    borderWidth: 1,
                    data: windchill,
                    borderColor: 'rgba(0, 153, 0, 1)',
                    backgroundColor: 'rgba(0, 204, 0, 0.2)',
                    fill: true,
                    tension: 0.1,
                    pointRadius: 0
                },{
                    label: 'Wind Speed',
                    borderWidth: 1,
                    data: windData,
                    borderColor: 'rgba(255, 119, 52, 1)',
                    backgroundColor: 'rgba(255, 153, 102, 0.2)',
                    fill: true,
                    tension: 0.1,
                    pointRadius: 0
                }]
            },
            options: {
                scales: {
                    yAxes: [{
                        ticks: {
                            beginAtZero: false
                        }
                    }]
                },
                elements: {
                    line: {
                        borderWidth: 1
                    }
                },
                plugins: {
                    zoom: {
                      zoom: {
                        wheel: {
                          enabled: true,
                        },
                        pinch: {
                          enabled: true
                        },
                        mode: 'xy',
                      }
                    }
                  }
            }
        });

    });