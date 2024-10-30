# This example will buy/sell apple randomly every 5 minutes
from CUTE import CUTE
cute = CUTE(step_interval_seconds=5*60, #Each "step" should move the virtual timer forward 5 minutes. There is a minimum time of 1 minute (60 seconds), and each step gets rounded to the nearest minute, so multiples of 60 will work best.
    start_time_unix=1673560800 #This is the unix timestamp for Jan 12 22:00:00 2023 UTC. Check out https://www.unixtimestamp.com/ for more info
)
