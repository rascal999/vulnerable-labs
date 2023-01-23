#!/usr/bin/env python
import os
import re

from flask import Flask, redirect, request

app = Flask(__name__)

@app.route('/')
def hello():
    return "Hello!\n"

# Solution: aHR0cHM6Ly9iYWQuY29t
@app.route('/redirect0')
def redirect0():
    url = request.args.get("url")
    return redirect(url, code=302)

# Solution: Ly90ZXN0LmNvbS5iYWQuY29t
@app.route('/redirect1')
def redirect1():
    url = request.args.get("url")
    validate = re.search("^//test.com|^http[s]?://test.com|^/(?!/)", url)
    if validate:
      return redirect(url, code=302)
    return redirect('/', code=302)

#Solution: L1xiYWQuY29tCg==
@app.route('/redirect2')
def redirect2():
    url = request.args.get("url")
    validate = re.search("^//test.com/|^http[s]?://test.com/|^/(?!/)", url)
    if validate:
      return redirect(url, code=302)
    return redirect('/', code=302)

@app.route('/redirect3')
def redirect3():
    url = request.args.get("url")
    validate = re.search("^//test.com/|^http[s]?://test.com/|^/(?![/|\\\\])", url)
    if validate:
      return redirect(url, code=302)
    return redirect('/', code=302)

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=os.environ.get("FLASK_SERVER_PORT", 9090), debug=True)
