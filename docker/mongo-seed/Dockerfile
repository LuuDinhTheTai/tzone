FROM python:3.12-alpine

WORKDIR /seed

RUN pip install --no-cache-dir pymongo

COPY docker/mongo-seed/seed_phone.py /seed/seed_phone.py

CMD ["python", "/seed/seed_phone.py"]

