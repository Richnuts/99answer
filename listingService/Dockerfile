FROM python:3.10.12

WORKDIR /app

COPY listing_service.py .
COPY python-libs.txt .

RUN pip install --no-cache-dir -r python-libs.txt

EXPOSE 6000

CMD ["python", "listing_service.py"]