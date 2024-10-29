import requests
class Cute:
    def __init__(self, url: str = 'localhost:3000', step_interval_seconds: int = 60*5, start_time_unix: int = 1673560800):
        assert step_interval_seconds >= 60

        self.url = url
        self.step_interval_seconds = step_interval_seconds
        self.start_time_unix = start_time_unix
