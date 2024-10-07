from flask import Flask, jsonify
import os, subprocess, requests

app = Flask(__name__)

def get_system_info():
    ip = subprocess.check_output("hostname -I", shell=True).decode().strip()
    processes = subprocess.check_output("ps -ax", shell=True).decode().strip()
    disk_space = subprocess.check_output("df -h /", shell=True).decode().strip()
    uptime = subprocess.check_output("uptime -p", shell=True).decode().strip()
    return {'ip': ip, 'processes': processes, 'disk_space': disk_space, 'uptime': uptime}

@app.route('/')
def system_info():
    service2_info = requests.get('http://service2:5000').json()
    service1_info = get_system_info()
    return jsonify({'Service1': service1_info, 'Service2': service2_info})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8199)
