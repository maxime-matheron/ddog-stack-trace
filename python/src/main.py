from ddtrace import patch_all, tracer
patch_all()

from flask import Flask
from pkg.handler import first_function
from ddtrace.opentracer import Tracer, set_global_tracer

app = Flask(__name__)

def init_tracer():
    config = {
      'agent_hostname': 'dd-agent',
      'agent_port': 8126,
    }
    tracer = Tracer('', config=config)
    set_global_tracer(tracer)
    return tracer

init_tracer()

@app.route("/generate-stack")
def generate_stack():
    with tracer.trace('call.first_function'):
        try:
            first_function()
        except: 
            pass
    return "<p>Python stack trace generated!</p>"

if __name__ == "__main__":
   app.run(host='0.0.0.0', port=5000)

