FROM python:3.8

ARG UID=1001
RUN useradd -m -u ${UID} docker

USER ${UID}

RUN mkdir /home/docker/src

ARG dir=/home/docker/src

WORKDIR $dir

RUN pip install flask

ENV FLASK_APP "app.py"

ADD . $dir

CMD ["python", "-m", "flask", "run", "--host", "0.0.0.0", "--port", "8080"]

