import json
import os
import sys
import time
from pathlib import Path

from bson import json_util
from pymongo import MongoClient
from pymongo.errors import PyMongoError

MONGO_URI = os.getenv("MONGO_URI", "mongodb://mongo:27017")
MONGO_DB = os.getenv("MONGO_DB", "Cluster0")
MONGO_COLLECTION = os.getenv("MONGO_COLLECTION", "brands")
SEED_FILE = Path(os.getenv("SEED_FILE", "/data/phone.json"))
MAX_ATTEMPTS = int(os.getenv("MONGO_SEED_MAX_ATTEMPTS", "30"))
SLEEP_SECONDS = float(os.getenv("MONGO_SEED_SLEEP_SECONDS", "2"))


def connect() -> MongoClient:
    last_error: Exception | None = None
    for attempt in range(1, MAX_ATTEMPTS + 1):
        try:
            client = MongoClient(MONGO_URI, serverSelectionTimeoutMS=2000)
            client.admin.command("ping")
            print(f"Connected to MongoDB on attempt {attempt}")
            return client
        except Exception as exc:  # noqa: BLE001
            last_error = exc
            print(f"MongoDB not ready yet (attempt {attempt}/{MAX_ATTEMPTS}): {exc}")
            time.sleep(SLEEP_SECONDS)

    raise RuntimeError(f"failed to connect to MongoDB after {MAX_ATTEMPTS} attempts") from last_error


def load_seed_data() -> list[dict]:
    if not SEED_FILE.exists():
        raise FileNotFoundError(f"seed file not found: {SEED_FILE}")

    with SEED_FILE.open("r", encoding="utf-8") as handle:
        data = json_util.loads(handle.read())

    if not isinstance(data, list):
        raise ValueError("phone.json must contain a JSON array of brand documents")

    return data


def main() -> int:
    try:
        data = load_seed_data()
        client = connect()
        collection = client[MONGO_DB][MONGO_COLLECTION]

        collection.delete_many({})
        if data:
            collection.insert_many(data, ordered=True)

        print(f"Seeded {len(data)} brand document(s) into {MONGO_DB}.{MONGO_COLLECTION}")
        return 0
    except (OSError, ValueError, PyMongoError, RuntimeError) as exc:
        print(f"Mongo seed failed: {exc}", file=sys.stderr)
        return 1


if __name__ == "__main__":
    raise SystemExit(main())

