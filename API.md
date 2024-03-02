POST example.com/api/location?user_id=<user_id>

{
    "latitude": <latitude>,
    "longitude": <longitude>
}

GET example.com/api/location?user_id=<user_id>

{
    "latitude": <latitude>,
    "longitude": <longitude>
}

GET example.com/api/location/nearest?user_id=<user_id>&range=<range>&count=<count>

{
    count: <count>,
    results: [
        {
            "latitude": <latitude>,
            "longitude": <longitude>,
            "user_id": <user_id>
        }
    ]
}