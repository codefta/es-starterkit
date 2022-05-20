# ElasticSearch Mapping

Dokumen ini berisi penjelasan tentang mapping yang digunakan pada project ini.

Mapping pada elasticsearch kurang lebih mirip dengan skema pada table di SQL. Pada intinya di mapping ini kita mendefinisikan field-field apa saja yang searchable.

## Mapping Index `foods`

```json
PUT /foods
{
  "mappings": {
    "properties": {
      "name": {"type": "text"}, // kita membuat field `name` menjadi searchable
      "description": {"type": "text"} // kita membuat field `description` menjadi searchable
    },
    "dynamic": false // kita tidak mengizinkan field lain menjadi searchable
  }
}
```