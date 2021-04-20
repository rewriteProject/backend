# Backend

## Pfade für Website

Für die Website-Pfade wird die Methode GET verwendet

### Liste von  Ländern

```
localhost:8081/onLoad/countries
```

Response-JSON:

```
{
    "countries": [
        "Canada",
        "China",
        "Finland",
        "Brazil"
    ]
}
```

### Listen von Produkteigenschaften

```
localhost:8081/onLoad/properties
```

Response-JSON:

```
{
    "brand": [
        "Edgepulse",
        "Yakidoo",
        "Fliptune",
    ],
    "category": [
        "Beauty",
        "Garden",
        "Automotive",
    ],
    "color": [
        "Green",
        "Red",
        "Blue"
    ]
}
```





## Pfade für Analyse-Komponente

Für diese Pfade wird die Methode POST verwendet

### Information 

localhost:8081/analytics/information/{Case}/{countries}

#### Fall I1

Für alle Länder:

```
localhost:8081/analytics/information/i1/all
```

Für eine Auswahl an Ländern:

Der Name muss großgeschrieben werden und durch ein Comma ohne Abstand von anderen getrennt sein

```
localhost:8081/analytics/information/i1/Austria,Russia
```



#### Fall I2

Dieser Fall funktioniert komplett gleich; es muss nur der Pfad auf diesen Fall geändert werden

```
localhost:8081/analytics/information/i2/all

localhost:8081/analytics/information/i2/Austria,Russia
```



### Statistik

localhost:8081/analytics/statistics/{country}/{attributes}

Parameter für Statistik:

- minDate ... Startdatum (des Zeitraums)
- maxDate ... Enddatum (des Zeitraums)

Für alle Merkmalsarten:

```
localhost:8081/analytics/statistics/France/all
```

Für eine Auswahl an Merkmalsarten:

```
localhost:8081/analytics/statistics/France/color,brand
```

Eine Merkmalsart muss kleingeschrieben werden und durch ein Comma ohne Abstand von anderen getrennt sein



### Prognose

localhost:8081/analytics/forecast/{Case}/{country}

#### Fall P1-1

Parameter für P1-1

- minDate ... Format yyyy-mm-dd

```
localhost:8081/analytics/forecast/p1-1/France?minDate=2020-05-01
```



#### Fall P1-2

Dieser Fall hat keine Parameter

```
localhost:8081/analytics/forecast/p1-2/France
```



#### Fall P2

Parameter für P2

- minDate ... Format yyyy-mm-dd

  oder

- year ... Startjahr

- typ ... Merkmalsart (color, brand oder category)

- feature ... Merkmal 

```
localhost:8081/analytics/forecast/p2/France
```

