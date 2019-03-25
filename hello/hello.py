#!/usr/bin/env python3
import sys
import os

who = "world"
greeting = "hello"

def getmessage(who):
        
    if "GREETING" in os.environ:
        greeting = os.environ["GREETING"]

    return f"{greeting}, {who}"


from flask import Flask
app = Flask(__name__)

@app.route("/hello/<who>")
def hello(who):
    return getmessage(who)