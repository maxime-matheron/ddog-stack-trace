from ddtrace import tracer
from ddtrace.opentracer import Tracer, set_global_tracer

def init_tracer(service_name):
    config = {
      'agent_hostname': 'localhost',
      'agent_port': 8126,
    }
    tracer = Tracer(service_name, config=config)
    set_global_tracer(tracer)
    return tracer

init_tracer('stacktrace')

def trigger_exception():
    try:
        with tracer.trace('call.fetch'):
            v = {}['a']
    except KeyError as e:
        raise ValueError('failed in trigger_exception')

def second_function():
    try:
        with tracer.trace('call.trigger'):
            trigger_exception()
    except ValueError as e:
        raise ValueError('failed in second_function') from e

def first_function():
    with tracer.trace('call.second'):
        second_function()

with tracer.trace('call.root'):
    first_function()
