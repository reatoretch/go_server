from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/')
def index():
    return jsonify({
        "version":["1.0.0","1.0.1","1.0.2"]
    })

if __name__ == "__main__":
    app.run()
