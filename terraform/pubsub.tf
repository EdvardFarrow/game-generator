resource "google_pubsub_topic" "game_events" {
  name = "game-events-topic"

  message_retention_duration = "86400s" 
}

resource "google_pubsub_subscription" "game_events_sub" {
  name  = "casual-game-events-sub"
  topic = google_pubsub_topic.game_events.name

  ack_deadline_seconds = 20 
  
  # TODO: Dead Letter Policy (DLQ) 
  
  expiration_policy {
    ttl = "2592000s" 
  }
}


resource "google_pubsub_subscription" "game_events_to_gcs" {
  name  = "casual-game-events-to-gcs"
  topic = google_pubsub_topic.game_events.name

  cloud_storage_config {
    bucket = google_storage_bucket.bronze_layer.name

    max_duration = "60s"   
    max_bytes    = 10000000 
    
  }

  depends_on = [
    google_storage_bucket_iam_member.pubsub_bucket_reader,
    google_storage_bucket_iam_member.pubsub_object_creator
  ]
}