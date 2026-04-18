resource "google_bigquery_dataset" "analytics_dataset" {
  dataset_id                  = "game_analytics"
  friendly_name               = "Casual Game Analytics"
  description                 = "Хранилище данных для игровой аналитики"
  
  location                    = "EU" 
}

resource "google_bigquery_table" "bronze_external_events" {
  dataset_id = google_bigquery_dataset.analytics_dataset.dataset_id
  table_id   = "events_raw_ext"

  external_data_configuration {
    autodetect    = true 
    source_format = "NEWLINE_DELIMITED_JSON" 

    source_uris = [
      "gs://${google_storage_bucket.bronze_layer.name}/*"
    ]
  }
}