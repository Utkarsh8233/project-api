GET /events  -> Get a list of available events 
GET /events/<id> -> Get detail of specific event using its id
POST /events -> Create a new bookable event (Auth required)
PUT /events/<id> -> Update an event (Auth required)
DELETE /events/<id> -> Delete an event (Auth required)
POST /signup -> Create a new user
POST /login -> Login (Authenticate user)
POST /events/<id>/register -> Register user for Event (Auth required)
DELETE /events/<id>/register -> Cancel registration (Auth required)
