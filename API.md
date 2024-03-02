POST example.com/api/location?user_id=<user_id>

{
    "latitude": <latitude>,
    "longitude": <longitude>,
    "timestamp": <timestamp>
}

GET example.com/api/location?user_id=<user_id>

{
    "latitude": <latitude>,
    "longitude": <longitude>,
    "last_updated": <timestamp>
}

GET example.com/api/location/nearest?user_id=<user_id>&range=<range>&count=<count>

[
    {
        "latitude": <latitude>,
        "longitude": <longitude>,
        "last_updated": <timestamp>,
        "user_id": <user_id>
    }, ...
]