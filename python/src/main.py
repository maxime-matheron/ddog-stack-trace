from flask import Flask
from ddtrace import tracer
from pkg.handler import first_function
from ddtrace.opentracer import Tracer, set_global_tracer

server = Flask(__name__)

def init_tracer():
    config = {
      'agent_hostname': 'dd-agent',
      'agent_port': 8126,
    }
    tracer = Tracer('', config=config)
    set_global_tracer(tracer)
    return tracer

init_tracer()

@server.route("/generate-stack")
def generate_stack():
    with tracer.trace('call.first_function'):
        try:
            first_function()
        except: 
            pass
    return "<p>Python stack trace generated!</p>"

if __name__ == "__main__":
   server.run(host='0.0.0.0', port=5000)

