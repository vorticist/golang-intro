payload='{"users" :["a31e70aa-77bd-4227-8c34-0b55c8b6c949","1f12f754-ef0c-4b10-878c-05c8144ca61a","a31e70aa-77bd-4227-8c34-0b55c8b6c949"], "whentolaunch" : "2022-01-18T14:42:05-08:00", "description" : "Hora del daily"  }'

curl http://127.0.0.1:8002/schedule/ -d "$payload"
