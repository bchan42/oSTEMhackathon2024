FROM python:3.10

COPY signupverif.py requirements.txt ./

RUN pip install -r requirements.txt

CMD ["python", "./signupverif.py"]