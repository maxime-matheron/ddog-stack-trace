from ddtrace import patch_all, tracer
patch_all()

from flask import Flask
from pkg.handler import first_function
from ddtrace.opentracer import Tracer, set_global_tracer

config = {'agent_hostname': 'dd-agent','agent_port': 8126}
tracer = Tracer('', config=config)
set_global_tracer(tracer)

app = Flask(__name__)

@app.route("/generate-stack")
def generate_stack():
    try:
        first_function()
    except: 
        pass
    return "<p>Python stack trace generated!</p>"

if __name__ == "__main__":
   app.run()

