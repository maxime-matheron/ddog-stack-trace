FROM python:3.9

WORKDIR /code

ADD ./src ./
RUN pip install -r requirements.txt

CMD [ "python", "./main.py" ]