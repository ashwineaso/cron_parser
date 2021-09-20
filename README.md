## Cron Parser

cron_parser is a command line application which parses a cron string and expands each 
field to show the times at which it will run. The application current supports the standard
cron format with five time fields (minute, hour, day, month, and day of week), and a command.

The output will be in the format

```
minute          0 15 30 45
hour            0
day of month    1 15
month           1 2 3 4 5 6 7 8 9 10 11 12
day of week     1 2 3 4 5
command         /usr/bin/find
```

### Running the application

1. Build the project and run

```bash
> go build .
```

2. Running the parser by providing the appropriate command line arguments.

```bash
> ./cron_parser "*/15 0 1,15 * 1-5 /usr/bin/find"
```
