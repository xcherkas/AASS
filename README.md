# Application Architectures of Software Systems
## Lukáš Tankovič, Andrei Cherkas
### Elektronický obchod (e shop)
Našim zadaním bude vytvoriť elektronický obchod, ďalej len ako e shop, určený na predaj rozličných tovarov. Eshop sa bude skladať z frontendu a backendu, pričom bude koncipovaný ako aplikácia využívajúca mikroslužby, avšak bude riešená aj pomocou BPM a EDA. Pomocou e shopu si budú môcť zákazníci kúpiť tovar rozličných druhov, pričom objednávky budú spracované správcami systému. V rámci projektu bude zainteresovanou stranou aj investor, ktorý poskytne financie na začatie projektu a jeho realizáciu.

Cieľom bude vytvoriť platformu, ktorá ponúkne zákazníkom rozličný tovar, v ktorom bude jednoduché navigovať sa a zadať objednávku. Pre správcov zasa ponúkne možnosť efektívne reagovať na objednávky, a zároveň aj dopĺňať tovar v sklade pomocou interakcie s dodávateľmi.

## Vetvy
Vetva `base_nginx` obsahuje verziu kodu s implementáciou mikrosluzieb a nginx.
Vetva `base_camunda` obsahuje verziu kódu s implementáciou Camunda BPMN.
Vetva `base_kafka` obsahuje verziu kodu s implementáciou mikrosluzieb a Apache Kafka.

## Sluzby

`postgres` - predtavuje instanciu DB, ktoru ale je potrebne inicializovat pomocou suboru `db.sql`

`kafka` - instancia Apache Kafka, vyuzivajuca standartny konfig

`proxy` - proxy server, ktory sluzi na preposielanie requestov z frontendu do Apache Kafka

`nginx` - proxy server, ktory sluzi na preposielanie requestov z frontendu do mikrosluzieb backendu

`frontend` - frontend aplikacia, ktora sluzi na interakciu so zakaznikom

`details`/`filter`/`searcher`/`products` - backend mikrosluzba, ktora sluzi na interakciu s DB

## Spustenie
Spustit vsetky sluzby je mozne pomocou prikazu:

```bash
docker-compose up
```

## Troubleshooting
Niekedy `nginx` servisa nechce fungovať, vtedy je potrebné ho zastaviť a znova spustiť. To sa dá urobiť pomocou príkazov:

```bash
docker-compose stop nginx
docker-compose up nginx
```

To iste moze stat aj pri pouziti Apache Kafka, vtedy je potrebne zastavit a znova spustit kontajner `kafka` a `proxy`

```bash
docker-compose stop kafka
docker-compose stop proxy
docker-compose up kafka
docker-compose up proxy
```
