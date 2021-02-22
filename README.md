# Anonymazing

Deadly simple SQL generation based data anonymization command-line utility

For example, given the following data on a Postgres table:


| Name                          | Email                       |
| ----------------------------- | --------------------------- |
| Abner Fiorelli Lyara          | alan.hickle@ies.com.br      |
| Abner Krystine Tramm Ariadna  | jayda.ohara@ies.com.br      |
| Abraão Cássio Tallis Ruas     | nola.nader@ies.com.br       |
| Abraão Sthefanny Delmondes    | salma.treutel@ies.com.br    |
| Acácia Claudiva Kauane        | julie.heaney@ies.com.br     |
| Acácia Craveiro Cecília Silva | berneice.oconner@ies.com.br |
| Acácia Farmácia Victoria      | nadia.mosciski@ies.com.br   |
| Acácia Mickaelle Maboni       | paolo.parisian@ies.com.br   |
| Acácia Rogéria                | gavin.beahan@ies.com.br     |
| Acácia Vasco Panuce Fraporti  | bennie.medhurst@ies.com.br  |


```sh
$ ./anonymazing --postgres-connection-string=postgresql://postgres:123456@localhost:5432/my_database?sslmode=disable --database-table=pessoa --database-columns=name,email --output=./anonymizer_script.sql
```

```sql
UPDATE pessoa SET nome='Tandara Padilha' WHERE nome='Abner Fiorelli Lyara';
UPDATE pessoa SET email='mandy.kozey@ies.com.br' WHERE email='alan.hickle@ies.com.br';
UPDATE pessoa SET nome='Gaspar Edfisica3 Roberta' WHERE nome='Abner Krystine Tramm Ariadna';
UPDATE pessoa SET email='willie.kirlin@ies.com.br' WHERE email='jayda.o"hara@ies.com.br';
UPDATE pessoa SET nome='Claudinei Luna Tozi' WHERE nome='Abraão Cássio Tallis Ruas';
UPDATE pessoa SET email='samantha.jerde@ies.com.br' WHERE email='nola.nader@ies.com.br';
UPDATE pessoa SET nome='Raiane Aquino Roque' WHERE nome='Abraão Sthefanny Delmondes';
UPDATE pessoa SET email='marvin.herman@ies.com.br' WHERE email='salma.treutel@ies.com.br';
UPDATE pessoa SET email='myah.zieme@ies.com.br' WHERE email='julie.heaney@ies.com.br';
UPDATE pessoa SET nome='Janaína Thais Naiara Fantinato' WHERE nome='Acácia Claudiva Kauane';
UPDATE pessoa SET nome='Flávio Fraga Romulo Nicaelle' WHERE nome='Acácia Craveiro Cecília Silva';
UPDATE pessoa SET email='fay.farrell@ies.com.br' WHERE email='berneice.o"conner@ies.com.br';
UPDATE pessoa SET nome='Hariane Santos' WHERE nome='Acácia Farmácia Victoria';
UPDATE pessoa SET email='grayce.hodkiewicz@ies.com.br' WHERE email='nadia.mosciski@ies.com.br';
UPDATE pessoa SET nome='Marcello Reiter' WHERE nome='Acácia Mickaelle Maboni';
UPDATE pessoa SET email='aron.o"conner@ies.com.br' WHERE email='paolo.parisian@ies.com.br';
UPDATE pessoa SET nome='Kelcia Psico2 Arroyane' WHERE nome='Acácia Rogéria';
UPDATE pessoa SET email='vaughn.morissette@ies.com.br' WHERE email='gavin.beahan@ies.com.br';
UPDATE pessoa SET nome='Yanca Sartori' WHERE nome='Acácia Vasco Panuce Fraporti';
```
