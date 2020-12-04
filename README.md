# Goquotes

goquotes is a simple Go scraper aimed at scraping super-ultra-duper lit quotes, built on [colly](https://github.com/gocolly/colly).

## Usage

```
./goquotes <theme> [amount] [output]
```

## Example

```
bash-3.2$ ./goquotes bird 30 birdquotes.json
Launching Scraper !
........................................Scrapped 40 quotes.

bash-3.2$ head birdquotes.json
[
  "“There is no greater agony than bearing an untold story inside you.”
 ― Maya AngelouI Know Why the Caged Bird Sings\n",
  "“I am no bird; and no net ensnares me: I am a free human being with
an independent will.” ― Charlotte BrontëJane Eyre\n",
  "“Hold fast to dreams,For if dreams dieLife is a broken-winged bird,T
hat cannot fly.” ― Langston Hughes\n",
  "“Clouds come floating into my life, no longer to carry rain or usher
 storm, but to add color to my sunset sky.” ― Rabindranath TagoreStray
Birds\n",
  "“You see, cuckoos are parasites. They lay their eggs in other birds'
 nests. When the egg hatches, the baby cuckoo pushes the other baby bir
ds out of the nest. The poor parent birds work themselves to death tryi
ng to find enough food to feed the enormous cuckoo child who has murder
ed their babies and taken their places.\"\"Enormous?\" said Jace. \"Did
 you just call me fat?\"\"It was an analogy.\"\"I am not fat.” ― Cassan
dra ClareCity of Ashes\n",
```

## Secret

The quotes are scraped from [goodreads](https://www.goodreads.com/quotes).

## LICENSE

MIT
