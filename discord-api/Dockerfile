FROM python:3.9

ENV LANG=C.UTF-8
ENV PYTHONPATH /usr/src/api

WORKDIR /usr/src/api

COPY requirements.txt requirements.txt
RUN pip install --upgrade pip
RUN pip install --no-cache-dir -r requirements.txt
