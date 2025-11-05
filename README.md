## Build And Run

```bash
go build -o bin/gowatcher .
```

```bash
./bin/gowatcher check
```

```bash
# -i : fichier d'entrée, -o fichier de sortie
./bin/gowatcher check -i=data/urls.json -o=reports/results.json
```

```bash
# Ajouter une entrée dans le fichier urls.json
./bin/gowatcher add -f=data/urls.json -n=Google -u=https://google.com -o=toto
```

## Exemple de données dans urls.json

```json
{
  "name": "Google",
  "url": "https://google.com",
  "owner": "test"
}
```
