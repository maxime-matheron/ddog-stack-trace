FROM python:3.8

WORKDIR /code

ADD ./src ./
RUN pip install -r requirements.txt

CMD [ "python", "./main.py" ]