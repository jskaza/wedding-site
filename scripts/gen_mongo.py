import json
import pandas as pd
import pymongo
import os

client = pymongo.MongoClient(os.getenv("MONGO_WEDDING"))
db = client.wedding



data = pd.read_csv("guest_list.csv",dtype={"displayName": str, "nGuests": int}).to_dict(orient="records")

def genGuestsObj(entry):
    guests = [{"name":"", "meal":""}]*entry["nGuests"]
    return {"displayName": entry["displayName"], "guests":guests}



to_upload = list(map(genGuestsObj, data))

db.guests.delete_many({})
db.guests.insert_many(to_upload)
    

