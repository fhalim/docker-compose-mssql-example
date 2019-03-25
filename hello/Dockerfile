FROM python:3-alpine3.9

COPY . .
RUN pip install -r requirements.txt
ENV FLASK_APP=hello.py
ENTRYPOINT ["flask", "run", "--host=0.0.0.0"]
