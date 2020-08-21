FROM python:3.8

ARG UID=1001
RUN useradd -m -u ${UID} docker

USER ${UID}

RUN mkdir /home/docker/src

ARG dir=/home/docker/src

WORKDIR $dir

RUN pip install flask 

ENV PATH $PATH:/home/docker/.local/bin

ENV FLASK_APP "app.py"

RUN pip install uwsgi


ADD . $dir

CMD ["uwsgi",  "--http", "0.0.0.0:8080", "--wsgi-file", "app.py", "--callable", "app"]

