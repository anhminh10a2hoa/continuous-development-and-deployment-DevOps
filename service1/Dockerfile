FROM python:3.9-slim

WORKDIR /app

COPY service1.py .

RUN apt-get update && apt-get install -y procps  # Install 'procps' for ps command
RUN pip install flask requests  # Install Python dependencies

EXPOSE 8199

CMD ["python", "service1.py"]
